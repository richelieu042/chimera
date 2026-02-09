package forwardKit

import (
	"net/http"
	"net/http/httputil"

	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
)

// ForwardByReverseProxy !!!: 应该使用 此函数 而非直接使用 httputil.ReverseProxy 的 ServeHTTP 方法（因为代理失败时，要从 ErrorHandler字段 中获取error）.
/*
PS: 代理请求失败时，建议返回状态码502(http.StatusBadGateway, 网关错误).

@param reverseProxy 不能为nil
*/
func ForwardByReverseProxy(w http.ResponseWriter, r *http.Request, reverseProxy *httputil.ReverseProxy) (err error) {
	if err = interfaceKit.AssertNotNil(reverseProxy, "reverseProxy"); err != nil {
		return
	}
	// !!!: 下面会修改 httputil.ReverseProxy 的字段，但又不希望影响 传参reverseProxy，因此在此处对指针取值(*)
	rp := *reverseProxy

	/*
		设置 ErrorHandler 字段，以免: 代理请求失败时，返回的err == nil.
		PS: 并不会影响外部的rp，因为不是"指针类型".
	*/
	old := rp.ErrorHandler
	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		err = e
		if rp.ErrorLog != nil {
			rp.ErrorLog.Printf("Fail to forward request, error: %s", err.Error())
		}

		if old != nil {
			old(w, r, e)
		}
	}

	// Richelieu: try to reset http.Request.Body
	if err = httpKit.TryToResetRequestBody(r); err != nil {
		return
	}

	// Richelieu: 请求转发前再检查下，以防请求已经被取消了
	if err = r.Context().Err(); err != nil {
		return
	}

	// 真正转发请求
	rp.ServeHTTP(w, r)
	return
}
