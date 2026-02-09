package httpKit

import (
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/conditionKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// GetScheme 请求使用的Web协议.
/*
@return "http" || "https"
*/
func GetScheme(req *http.Request) string {
	scheme := req.URL.Scheme
	if strKit.IsNotEmpty(scheme) {
		return scheme
	}

	if req.TLS != nil {
		return "https"
	}
	return "http"
}

// GetClientScheme 客户端发送的原始请求 使用的Web协议.
/*
!!!: 返回值不一定准确，除非 代理s 好好配合（有的话）.

@return "http" || "https"
*/
func GetClientScheme(req *http.Request) string {
	tmp := GetHeader(req.Header, "X-Forwarded-Proto")

	if strKit.IsEmpty(tmp) {
		/*
			(1) 没有代理
			(2) 有代理（s），但是它们没有正确设置 X-Forwarded-Proto
		*/
		return GetScheme(req)
	}
	/* 有代理且设置了 X-Forwarded-Proto，，但不一定正确的设置了 */
	return conditionKit.TernaryOperator(tmp == "https", "https", "http")
}

// GetProto 服务器的HTTP版本.
/*
@return (1) "HTTP/1.0"
		(2) "HTTP/1.1"
		(3) "HTTP/2.0"（https）
*/
func GetProto(req *http.Request) string {
	return req.Proto
}
