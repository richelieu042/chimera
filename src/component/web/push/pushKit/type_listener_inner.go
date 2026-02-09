package pushKit

import (
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/bytesKit"
)

type innerListener struct {
	Listener
}

func (listener *innerListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
}

func (listener *innerListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
	// 加入管理
	BindId(channel, channel.GetId())
}

// OnMessage 仅针对WebSocket连接
func (listener *innerListener) OnMessage(channel Channel, messageType int, data []byte) {
	// 响应 WebSocket客户端 发来的 "ping"
	if bytesKit.Equals(data, PingData) {
		if err := channel.Push(PongData); err != nil {
			logger.WithError(err).Error("Fail to pong.")
			return
		}
	}
}

func (listener *innerListener) BeforeClosedByBackend(channel Channel, closeInfo string) {
}

func (listener *innerListener) OnClose(channel Channel, closeInfo string, bsid, user, group string) {
	channel.Dispose()

	// 移除管理
	UnbindId(channel)
	UnbindBsid(channel)
	UnbindUser(channel)
	UnbindGroup(channel)
}
