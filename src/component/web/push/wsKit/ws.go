package wsKit

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/richelieu042/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
)

// DefaultUpgrader 默认的Upgrader.
/*
@return 并发安全的
*/
func DefaultUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		HandshakeTimeout: time.Second * 3,
		CheckOrigin: func(r *http.Request) bool {
			// 允许跨域
			return true
		},
	}
}

// NewProcessor
/*
!!!: 需要先调用 pushKit.MustSetUp 或 pushKit.SetUp.

@param upgrader		可以为nil（将使用默认的）
@param idGenerator	可以为nil（将使用xid）
@param listener		不能为nil
@param msgType		消息类型
@param pongInterval	pong的周期（<=0则不发送pong）
*/
func NewProcessor(upgrader *websocket.Upgrader, idGenerator func() (string, error), listener pushKit.Listener, messageType *MessageType, pongInterval time.Duration) (pushKit.Processor, error) {
	if err := pushKit.CheckSetup(); err != nil {
		return nil, err
	}

	if err := interfaceKit.AssertNotNil(messageType, "MessageType"); err != nil {
		return nil, err
	}

	if upgrader == nil {
		upgrader = DefaultUpgrader()
	}
	if idGenerator == nil {
		idGenerator = pushKit.DefaultIdGenerator()
	}
	listeners, err := pushKit.NewListeners(listener, false)
	if err != nil {
		return nil, err
	}

	return &wsProcessor{
		upgrader:     upgrader,
		idGenerator:  idGenerator,
		listeners:    listeners,
		msgType:      messageType,
		pongInterval: pongInterval,
	}, nil
}
