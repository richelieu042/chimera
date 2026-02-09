package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
)

// SetHttpStatusCode （响应头）设置http状态码
func SetHttpStatusCode(ctx *gin.Context, statusCode int) {
	ctx.Writer.WriteHeader(statusCode)
}

// SetResponseHeader （响应头）
func SetResponseHeader(ctx *gin.Context, key, value string) {
	ctx.Header(key, value)
	//httpKit.SetHeader(ctx.Writer.Header(), key, value)
}

// AddResponseHeader （响应头）
func AddResponseHeader(ctx *gin.Context, key, value string) {
	httpKit.AddHeader(ctx.Writer.Header(), key, value)
}

// DelResponseHeader （响应头）删除响应头.
func DelResponseHeader(ctx *gin.Context, key string) {
	httpKit.DelHeader(ctx.Writer.Header(), key)
}
