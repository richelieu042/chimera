package sseKit

import (
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/errorKit"
)

// IsSseSupported
/*
@return 为"": 支持SSE
*/
func IsSseSupported(w http.ResponseWriter, r *http.Request) error {
	if _, ok := w.(http.Flusher); !ok {
		// 不支持: 流信息（streaming）
		return errorKit.Newf("http.Flusher(Streaming) isn't supported")
	}
	//if _, ok := w.(http.CloseNotifier); !ok {
	//	// 不支持: 监听关闭
	//	return errorKit.Newf("http.CloseNotifier isn't supported!")
	//}
	return nil
}

// SetHeaders 设置response header.
func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}
