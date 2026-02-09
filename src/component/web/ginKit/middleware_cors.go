package ginKit

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
)

// NewCorsMiddleware 新建一个cors中间件（自带对 OPTIONS请求 的处理）.
/*
「Go框架」 Gin 怎么实现允许前端跨域请求？
	https://mp.weixin.qq.com/s/2eJUJKJ3Xu5jOYXAfknqtA

@param origins 	(0) 可以为nil
				(1) origin白名单
				(2) 支持wildcard（*）
				(3) len(origins) == 0，则通配（请求的origin是什么就允许什么）

e.g.
	传参: []string{"https://*.github.com", "https://api.*", "http://*", "https://facebook.com", "*.golang.org"}
*/
func NewCorsMiddleware(origins []string, allowCredentials bool) gin.HandlerFunc {
	/*
		CORS 中间件提供三个函数，代表三种使用方式，分别是:
			cors.Sign()
			cors.Default()
			cors.DefaultConfig()
	*/
	config := cors.Config{
		// 允许的请求方式，默认值是 GET，POST，PUT，PATCH，DELETE，HEAD，和 OPTIONS
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		// 用在对预请求的响应中，指示实际的请求中可以使用哪些 HTTP 请求头
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "Accept", "Origin", "X-Requested-With"},
		// 可以在响应中显示的请求头
		ExposeHeaders: []string{"Content-Length", "Content-Type", "Last-Modified"},

		// 预检请求（OPTIONS）的结果能被缓存多久
		MaxAge:                    timeKit.Day,
		OptionsResponseStatusCode: http.StatusNoContent,

		// 添加请求源是否允许使用通配符，例如 http://some-domain/*，https://api. 或 http://some.*.subdomain.com
		AllowWildcard: true,
		// 表示请求附带请求凭据时是否响应请求，例如 cookie、HTTP authentication 或客户端 SSL 证书
		AllowCredentials: allowCredentials,
		// 允许使用常用的浏览器的扩展模式
		AllowBrowserExtensions: true,
		// 允许使用 WebSocket 协议
		AllowWebSockets: true,
		// 允许使用 file:// 协议
		AllowFiles: false,
	}
	if len(origins) > 0 {
		/*
			(1) 允许部分origin
			AllowOrigins: 指定允许请求源的列表，如果列表中存在 *，则允许所有请求源，默认值是 []
		*/
		config.AllowOrigins = origins
	} else {
		/*
			(2) 允许全部origin
			AllowOriginFunc: 接收参数 origin，函数体中的验证逻辑返回是否允许跨域请求。该配置项优先级高于 AllowOrigins []string，如果设置该配置项，AllowOrigins []string 配置项的设置被忽略
		*/
		config.AllowOriginFunc = func(origin string) bool {
			return true
		}
	}
	return cors.New(config)
}
