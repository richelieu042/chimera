package wsKit

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
)

// PolyfillWebSocketRequest
/*
此函数是为了避免情况: 代理（e.g.Nginx）没有设置WebSocket穿透，导致WebSocket服务收到的WebSocket请求的header有问题.
*/
func PolyfillWebSocketRequest(r *http.Request) {
	httpKit.SetHeader(r.Header, httpKit.HeaderConnection, "Upgrade")
	httpKit.SetHeader(r.Header, httpKit.HeaderUpgrade, "websocket")

	//httpKit.SetHeaderIfMissingIgnoreCase(r.Header, httpKit.HeaderConnection, "Upgrade")
	//httpKit.SetHeaderIfMissingIgnoreCase(r.Header, httpKit.HeaderUpgrade, "websocket")
}

// IsWebSocketUpgrade 是否是WebSocket请求？
var IsWebSocketUpgrade func(r *http.Request) bool = websocket.IsWebSocketUpgrade
