package proxyKit

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
)

// Proxy 代理请求（反向代理，请求转发）.
/*
Deprecated: Use forwardKit instead.

调用此方法就行请求转发前，按照实际场景可以:
(1) POST请求，覆盖 request body;
(2) 变更请求头（request header）.

@param w			e.g.ctx.Writer
@param r 			e.g.ctx.Request
@param host			e.g."127.0.0.1:12345"
@param options 		optional
*/
func Proxy(w http.ResponseWriter, r *http.Request, host string, options ...ProxyOption) error {
	opts := loadOptions(options...)

	return opts.proxy(w, r, host)
}

// ProxyWithGin
/*
Deprecated: Use forwardKit instead.
*/
func ProxyWithGin(ctx *gin.Context, host string, options ...ProxyOption) error {
	opts := loadOptions(options...)
	opts.ctx = ctx

	return opts.proxy(ctx.Writer, ctx.Request, host)
}

// ProxyToUrl
/*
Deprecated: Use forwardKit instead.

项目实战：用 Go 创建一个简易负载均衡器
	https://mp.weixin.qq.com/s/pe0CQa3tdrUmC86OSRBNeg

@param targetUrl 可以是 url.Parse() 的返回值
*/
func ProxyToUrl(w http.ResponseWriter, r *http.Request, targetUrl *url.URL) (err error) {
	/* reset Request.Body */
	if err = httpKit.TryToResetRequestBody(r); err != nil {
		return
	}

	/* Richelieu: 在请求头加个标记，证明此请求被 chimera 代理过 */
	mark(r.Header)

	reverseProxy := httputil.NewSingleHostReverseProxy(targetUrl)
	reverseProxy.ErrorHandler = func(writer http.ResponseWriter, r *http.Request, err1 error) {
		err = errorKit.Wrapf(err1, "fail to proxy")
	}
	// Richelieu: 此处的 recover() 是针对 ReverseProxy.ServeHTTP() 中的 panic(http.ErrAbortHandler)
	defer func() {
		if obj := recover(); obj != nil {
			if err1, ok := obj.(error); ok {
				err = err1
				return
			}
			err = errorKit.Newf("recover from %v", obj)
		}
	}()
	reverseProxy.ServeHTTP(w, r)
	return
}
