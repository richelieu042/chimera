package ginKit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
)

// NewOptionsMiddleware
/*
Deprecated: NewCorsMiddleware 已经自带对 OPTIONS请求 的处理了.

PS:
(1) 就算 NoRoute || NoMethod，也会走到中间件.
*/
func NewOptionsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Header(httpKit.HeaderAccessControlAllowCredentials, "true")
			ctx.Header(httpKit.HeaderAccessControlAllowMethods, "OPTIONS, GET, POST")

			ctx.Header(httpKit.HeaderAccessControlAllowHeaders, "Accept, Accept-Encoding, Authorization, Cache-Control, Content-Type, Content-Length, Origin, X-CSRF-Token, X-Requested-With")

			ctx.Header(httpKit.HeaderAccessControlAllowOrigin, httpKit.GetOrigin(ctx.Request.Header))
			// 预检请求的返回结果能缓存多久？24h
			ctx.Header(httpKit.HeaderAccessControlMaxAge, "86400")
			ctx.Header(httpKit.HeaderContentType, "text/plain; charset=utf-8")

			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}
