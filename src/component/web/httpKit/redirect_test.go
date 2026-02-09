package httpKit

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/netKit"
)

/*
访问:
(1) http://127.0.0.1/test => https://www.baidu.com/
(2) http://127.0.0.1/a => http://127.0.0.1/b
*/
func TestRedirect(t *testing.T) {
	port := 80

	engine := gin.Default()

	engine.GET("/test", func(ctx *gin.Context) {
		if err := Redirect(ctx.Writer, ctx.Request, "https://www.baidu.com", http.StatusPermanentRedirect); err != nil {
			ctx.String(500, err.Error())
			return
		}
		// do nothing
	})

	engine.GET("/a", func(ctx *gin.Context) {
		if err := Redirect(ctx.Writer, ctx.Request, "/b", http.StatusPermanentRedirect); err != nil {
			ctx.String(500, err.Error())
			return
		}
		// do nothing
	})
	engine.GET("/b", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})

	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
