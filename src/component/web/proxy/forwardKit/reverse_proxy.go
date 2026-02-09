package forwardKit

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// NewSingleHostReverseProxyWithUrl
/*
@param urlStr 		目标url
					e.g. 	"http://127.0.0.1:8000": 将请求转发给"http://127.0.0.1:8000"，路由不变
					e.g.1 	"http://127.0.0.1:8000/a": 将请求转发给"http://127.0.0.1:8000"，路由的最前面加上"/a"
@return (1) Director字段: != nil;
		(2) Transport、ModifyResponse、ErrorLog、ErrorHandler等字段: == nil.
*/
func NewSingleHostReverseProxyWithUrl(urlStr string) (*httputil.ReverseProxy, error) {
	if err := strKit.AssertNotEmpty(urlStr, "urlStr"); err != nil {
		return nil, err
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, errorKit.Newf("invalid urlStr(%s)", urlStr)
	}
	return NewSingleHostReverseProxy(u)
}

// NewSingleHostReverseProxy
/*
@return (1) Director字段: != nil;
		(2) Transport、ModifyResponse、ErrorLog、ErrorHandler等字段: == nil.
*/
func NewSingleHostReverseProxy(u *url.URL) (rp *httputil.ReverseProxy, err error) {
	if err = interfaceKit.AssertNotNil(u, "u"); err != nil {
		return
	}

	rp = httputil.NewSingleHostReverseProxy(u)
	return
}

// NewCustomReverseProxy 新建自定义的 *httputil.ReverseProxy 实例.
/*
PS: 对于 httputil.ReverseProxy 结构体，Rewrite 和 Director 只能有一个非nil.

@param director			不能为nil!!!
@param transport		可以为nil
@param modifyResponse	可以为nil
@param errLog			可以为nil（即无输出，但不推荐这么干）
@param errHandler		可以为nil
*/
func NewCustomReverseProxy(director func(*http.Request), transport http.RoundTripper, modifyResponse func(*http.Response) error, errLog *log.Logger, errHandler func(http.ResponseWriter, *http.Request, error)) (*httputil.ReverseProxy, error) {
	if err := interfaceKit.AssertNotNil(director, "director"); err != nil {
		return nil, err
	}

	rp := &httputil.ReverseProxy{
		Director: director,

		Transport:      transport,
		ModifyResponse: modifyResponse,
		//BufferPool:     nil,
		//FlushInterval:  0,
		ErrorLog:     errLog,
		ErrorHandler: errHandler,
	}
	return rp, nil
}
