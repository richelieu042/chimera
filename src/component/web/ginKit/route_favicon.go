package ginKit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
)

func DefaultFavicon(engine *gin.Engine) {
	httpFileSystem := http.FS(efs)

	engine.GET("/favicon.ico", func(ctx *gin.Context) {
		// 缓存24h
		ctx.Header(httpKit.HeaderCacheControl, "public, max-age=86400, must-revalidate")

		ctx.FileFromFS("_icon/favicon.ico", httpFileSystem)
	})
}
