package ginKit

import (
	"net/http"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/validateKit"
)

var (
	RequestTooLargeError = errorKit.Newf("HTTP request too large")
)

// NewSizeLimiterMiddleware 参考了echo中的 middleware.BodyLimit()
/*
	@param limit 	(1) 单位: MiB
					(2) 必须 > 0
*/
func NewSizeLimiterMiddleware(limit int64) (gin.HandlerFunc, error) {
	if err := validateKit.Var(limit, "gt=0"); err != nil {
		return nil, errorKit.Wrapf(err, "invalid limit(%d)", limit)
	}

	// bodyLimit 单位: B
	bodyLimit := limit << 20
	return func(ctx *gin.Context) {
		// (1) do nothing
		if ctx.Request.Body == nil || ctx.Request.Body == http.NoBody {
			ctx.Next()
			return
		}

		// (2) Based on content length
		if ctx.Request.ContentLength > bodyLimit {
			// 参考: github.com/gin-contrib/size
			_ = ctx.Error(RequestTooLargeError)
			ctx.Header("connection", "close")
			ctx.String(http.StatusRequestEntityTooLarge, "request too large")
			ctx.AbortWithStatus(http.StatusRequestEntityTooLarge)

			return
		}

		// (3) Based on content read
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, bodyLimit)
		ctx.Next()
	}, nil
}

// NewSizeLimiterMiddleware1
/*
gin-contrib/size
	https://github.com/gin-contrib/size
*/
func NewSizeLimiterMiddleware1(limit int64) gin.HandlerFunc {
	return limits.RequestSizeLimiter(limit)
}
