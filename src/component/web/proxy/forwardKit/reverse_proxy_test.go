package forwardKit

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/log/logKit"
	"github.com/richelieu042/chimera/v3/src/netKit"
)

/*
浏览器访问url: 	http://127.0.0.1/test
效果: 			将 http://127.0.0.1/test 转发给 http://127.0.0.1:8000/test
*/
func TestNewSingleHostReverseProxyWithUrl(t *testing.T) {
	/* 被代理服务 */
	go func() {
		port := 8000
		engine := gin.Default()
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	/* 代理服务 */
	port := 80
	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		rp, err := NewSingleHostReverseProxyWithUrl("http://127.0.0.1:8000")
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		rp.ErrorLog = logKit.NewStdoutLogger("")
		// (1) Richelieu: 此处应该为 true
		fmt.Printf("rp.ErrorHandler == nil? [%t]\n", rp.ErrorHandler == nil)

		if err := ForwardByReverseProxy(ctx.Writer, ctx.Request, rp); err != nil {
			// (2.1) Richelieu: 此处应该为 true
			fmt.Printf("rp.ErrorHandler == nil? [%t]\n", rp.ErrorHandler == nil)
			ctx.String(500, err.Error())
			return
		}
		// (2.2) Richelieu: 此处应该为 true
		fmt.Printf("rp.ErrorHandler == nil? [%t]\n", rp.ErrorHandler == nil)
		return
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
