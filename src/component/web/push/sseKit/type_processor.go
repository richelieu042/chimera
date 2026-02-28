package sseKit

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

type sseProcessor struct {
	pushKit.Processor

	idGenerator  func() (string, error)
	listeners    pushKit.Listeners
	msgType      *messageType
	pongInterval time.Duration
}

func (p *sseProcessor) ProcessWithGin(ctx *gin.Context) {
	p.Process(ctx.Writer, ctx.Request)
}

func (p *sseProcessor) Process(w http.ResponseWriter, r *http.Request) {
	if err := IsSseSupported(w, r); err != nil {
		p.listeners.OnFailure(w, r, err.Error())
		return
	}

	// 设置 response header
	SetHeaders(w)

	channel, err := p.newChannel(w, r, make(chan string, 1))
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

	/*
		!!!: gin.Context.Done() 和 r.Context().Done() 不同，因为 gin.Context.Done() 返回nil（普通Gin Server情况下）.
		(1) case为 w.(http.CloseNotifier).CloseNotify() 和 gin.Context.Done()，前端断开会走到 w.(http.CloseNotifier).CloseNotify()
		(2) case为 w.(http.CloseNotifier).CloseNotify() 和 r.Context().Done()，前端断开会走到 r.Context().Done()
	*/
	select {
	//case <-w.(http.CloseNotifier).CloseNotify():
	//	p.listeners.OnClose(channel, "Connection closed")
	case <-r.Context().Done():
		if channel.SetClosed() {
			p.listeners.OnClose(channel, "Context of request is done.")
		}
	case closeInfo := <-channel.GetCloseCh():
		p.listeners.OnClose(channel, closeInfo)
	}
}

func (p *sseProcessor) newChannel(w http.ResponseWriter, r *http.Request, closeCh chan string) (pushKit.Channel, error) {
	id, err := p.idGenerator()
	if err != nil {
		return nil, errorKit.Wrapf(err, "fail to generate id")
	}
	if err := strKit.AssertNotEmpty(id, "id"); err != nil {
		return nil, err
	}

	ip := httpKit.GetClientIP(r)
	channel := &SseChannel{
		BaseChannel: pushKit.BaseChannel{
			CloseCh:      closeCh,
			ClientIP:     ip,
			Type:         "SSE",
			Id:           id,
			Bsid:         "",
			User:         "",
			Group:        "",
			Data:         nil,
			Closed:       false,
			Listeners:    p.listeners,
			PongInterval: p.pongInterval,
		},
		w:       w,
		r:       r,
		msgType: p.msgType,
	}
	return channel, nil
}
