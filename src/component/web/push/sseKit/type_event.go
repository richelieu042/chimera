package sseKit

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sse"
	"github.com/richelieu042/chimera/v3/src/urlKit"
)

type Event sse.Event

func (e Event) Push(w http.ResponseWriter) error {
	if err := sse.Encode(w, sse.Event(e)); err != nil {
		return err
	}

	// 没有下面这一代码会导致: 前端无法及时收到推送的消息
	w.(http.Flusher).Flush()

	return nil
}

// MessageEvent 定义SSE事件.
/*
Deprecated: Use sse.Event of github.com/gin-contrib/sse instead.
*/
type MessageEvent struct {
	// Id
	/*
		PS:
		(1) 对应前端的 e.lastEventId.
		(2) 可以为"".
	*/
	Id string

	// Event
	/*
		PS:
		(1) 对应前端的 e.type.
		(2) 可以为""（此时等价于"message"）
		(3) 如果不是 "message" 的话，前端需要自行添加对应的监听.
	*/
	Event string

	// Data
	/*
		PS:
		(1) 对应前端的 e.data.
		(2) 可以为"".
		(3) 建议对内容编码下，以防其中有特殊字符(\n等).
	*/
	Data string
}

// String 实现SSE事件的 String() 方法
func (e MessageEvent) String() string {
	/*
		对 e.Data 编码1次，以防其中有特殊字符（比如'\n'）
		!!!: 前端需要手动解码1次
	*/
	data := urlKit.EncodeURIComponent(e.Data)

	return fmt.Sprintf("id: %s\nevent: %s\ndata: %s\n\n", e.Id, e.Event, data)
}
