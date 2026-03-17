package ginKit

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/concurrency/rateLimitKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"golang.org/x/time/rate"
)

// NewGinMiddleware Gin的限流器，用于限流.
/*
PS: 传参说明详见 NewLimiter.

@param r				每秒生成的令牌数，	(1) 可以直接是 某个数值;
										(2) 也可以是 rate.Every() 的返回值;
@param b 				令牌桶的容量
@param forbiddenText 	被限流时，响应给前端的内容（状态码固定为 403）
*/
func NewGinMiddleware(r rate.Limit, b int, forbiddenText string) gin.HandlerFunc {
	limiter := rateLimitKit.NewLimiter(r, b)

	if strKit.IsEmpty(forbiddenText) {
		forbiddenText = fmt.Sprintf("Exceed rate limiter(r: %f, b: %d).", r, b)
	}

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.String(http.StatusForbidden, forbiddenText)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
