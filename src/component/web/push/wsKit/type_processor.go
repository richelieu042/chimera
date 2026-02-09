package wsKit

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/sirupsen/logrus"
)

type wsProcessor struct {
	pushKit.Processor

	// upgrader 是并发安全的
	upgrader *websocket.Upgrader

	idGenerator  func() (string, error)
	listeners    pushKit.Listeners
	msgType      *MessageType
	pongInterval time.Duration
}

func (p *wsProcessor) ProcessWithGin(ctx *gin.Context) {
	p.Process(ctx.Writer, ctx.Request)
}

func (p *wsProcessor) Process(w http.ResponseWriter, r *http.Request) {
	PolyfillWebSocketRequest(r)

	// 先判断是不是websocket请求
	if !IsWebSocketUpgrade(r) {
		failureInfo := "Not a websocket upgrade request"
		p.listeners.OnFailure(w, r, failureInfo)
		return
	}

	// Upgrade（升级为WebSocket协议）
	conn, err := p.upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		err = errorKit.Wrapf(err, "fail to upgrade")
		p.listeners.OnFailure(w, r, err.Error())
		return
	}
	// PS: 对于 Conn.Close() ，可以多次调用，不会panic，但从第二次关闭开始，返回非nil的error（可以直接忽略）.
	defer conn.Close()

	channel, err := p.newChannel(r, conn, make(chan string, 1))
	if err != nil {
		err = errorKit.Wrapf(err, "fail to new channel")
		p.listeners.OnFailure(w, r, err.Error())
		return
	}
	if err := channel.Initialize(); err != nil {
		err = errorKit.Wrapf(err, "fail to initialize channel")
		p.listeners.OnFailure(w, r, err.Error())
		return
	}

	p.listeners.OnHandshake(w, r, channel)

	conn.SetCloseHandler(func(code int, text string) error {
		closeInfo := fmt.Sprintf("code: %d, text: %s", code, text)
		if channel.SetClosed() {
			channel.GetCloseCh() <- closeInfo
		}

		// 默认的close handler
		message := websocket.FormatCloseMessage(code, text)
		_ = conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
		return nil
	})

	go func() {
		// test
		defer func() {
			logrus.Info("[TEST] receive goroutine ends...")
		}()

		/* 接收WebSocket客户端发来的消息 */
		for {
			messageType, data, err := conn.ReadMessage()
			if err != nil {
				var closeInfo string
				var closeErr *websocket.CloseError
				if errors.As(err, &closeErr) {
					closeInfo = fmt.Sprintf("code: %d, text: %s", closeErr.Code, closeErr.Text)
				} else {
					closeInfo = fmt.Sprintf("Fail to read message because of error(%s)", err.Error())
				}

				if channel.SetClosed() {
					channel.GetCloseCh() <- closeInfo
				}
				break
			}
			p.listeners.OnMessage(channel, messageType, data)
		}
	}()

	select {
	case <-r.Context().Done():
		if channel.SetClosed() {
			p.listeners.OnClose(channel, "Context of request is done.")
		}
	case closeInfo := <-channel.GetCloseCh():
		p.listeners.OnClose(channel, closeInfo)
	}
}

func (p *wsProcessor) newChannel(r *http.Request, conn *websocket.Conn, closeCh chan string) (pushKit.Channel, error) {
	id, err := p.idGenerator()
	if err != nil {
		return nil, errorKit.Wrapf(err, "fail to generate id")
	}
	if err := strKit.AssertNotEmpty(id, "id"); err != nil {
		return nil, err
	}

	ip := httpKit.GetClientIP(r)
	channel := &WsChannel{
		BaseChannel: pushKit.BaseChannel{
			CloseCh:      closeCh,
			ClientIP:     ip,
			Type:         "WebSocket",
			Id:           id,
			Bsid:         "",
			User:         "",
			Group:        "",
			Data:         nil,
			Closed:       false,
			Listeners:    p.listeners,
			PongInterval: p.pongInterval,
		},
		msgType: p.msgType,
		conn:    conn,
	}
	return channel, nil
}
