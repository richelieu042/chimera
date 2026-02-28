package httpKit

import (
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// Redirect 请求重定向.
/*
PS:
301: http.StatusMovedPermanently 	表示资源已被永久移动到新位置，以后对该资源的请求都应使用新的URL。
302: http.StatusFound 				有时称为"Moved Temporarily"，表示资源临时被移动到另一个位置，客户端应继续使用原有URL进行以后的请求。
303: http.StatusSeeOther			指示客户端应使用GET方法访问另一个URL以获取资源，通常用于在POST请求后重定向到一个新的资源。
307: http.StatusTemporaryRedirect	表示资源临时被移动到另一个位置，但客户端应继续使用原有URL进行以后的请求（与302不同的是，307要求重定向请求方法和主体不变）。
308: http.StatusPermanentRedirect	表示资源已被永久移动到新位置，客户端以后应使用新的URL进行请求（与301不同的是，308要求重定向请求方法和主体不变）。

!!!选择使用:
(1) 302：当重定向后的操作不需要保持原请求方法（例如将POST转换为GET是可以接受的）时使用。
(2) 307：当需要保持原请求方法（例如POST必须保持为POST）时使用，以确保请求数据和操作的正确传递。

@param url 	e.g. 完整的网址
				"https://www.baidu.com"
			e.g.1 仅仅是路由
				"/b"
@param code (1) [300, 308] || 201
			(2) Richelieu: 推荐使用 307 或 308
*/
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) error {
	if err := strKit.AssertNotEmpty(url, "url"); err != nil {
		return err
	}

	// Richelieu: 参考了 gin v1.10.0 中的 render.Redirect.Render()，对 code 进行了限制
	if (code < http.StatusMultipleChoices || code > http.StatusPermanentRedirect) && code != http.StatusCreated {
		return errorKit.Newf("can't redirect with status code(%d)", code)
	}

	http.Redirect(w, r, url, code)
	return nil
}
