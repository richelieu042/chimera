package pushKit

import (
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
)

type Listeners []Listener

func (listeners Listeners) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	for _, listener := range listeners {
		listener.OnFailure(w, r, failureInfo)
	}
}

func (listeners Listeners) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
	for _, listener := range listeners {
		listener.OnHandshake(w, r, channel)
	}
}

func (listeners Listeners) OnMessage(channel Channel, messageType int, data []byte) {
	for _, listener := range listeners {
		listener.OnMessage(channel, messageType, data)
	}
}

func (listeners Listeners) BeforeClosedByBackend(channel Channel, closeInfo string) {
	for _, listener := range listeners {
		listener.BeforeClosedByBackend(channel, closeInfo)
	}
}

func (listeners Listeners) OnClose(channel Channel, closeInfo string) {
	// !!!: 此处先把取出来，以防 inner listener 解绑时去掉了相关信息（bsid, user, group），不会去掉id、data，导致轮到 另一个listener 时取不到数据
	bsid := channel.GetBsid()
	user := channel.GetUser()
	group := channel.GetGroup()

	for _, listener := range listeners {
		listener.OnClose(channel, closeInfo, bsid, user, group)
	}
}

// NewListeners
/*
PS: 本方法仅供本项目使用，严禁外部调用.
*/
func NewListeners(listener Listener, sseFlag bool) (Listeners, error) {
	if err := interfaceKit.AssertNotNil(listener, "listener"); err != nil {
		return nil, err
	}

	inner := &innerListener{}
	return []Listener{inner, listener}, nil
}
