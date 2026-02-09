package slbKit

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/netKit"
)

// 访问url: http://127.0.0.1:80/test
func TestNewLoadBalancer(t *testing.T) {
	//go func() {
	//	port := 8000
	//	engine := gin.Default()
	//	engine.Any("/test", func(ctx *gin.Context) {
	//		ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
	//	})
	//	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
	//		panic(err)
	//	}
	//}()
	//go func() {
	//	port := 8001
	//	engine := gin.Default()
	//	engine.Any("/test", func(ctx *gin.Context) {
	//		ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
	//	})
	//	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
	//		panic(err)
	//	}
	//}()
	//go func() {
	//	port := 8002
	//	engine := gin.Default()
	//	engine.Any("/test", func(ctx *gin.Context) {
	//		ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
	//	})
	//	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
	//		panic(err)
	//	}
	//}()

	time.Sleep(time.Second * 3)

	lb := NewLoadBalancer(nil)
	urls := []string{"http://127.0.0.1:8000", "http://127.0.0.1:8001", "http://127.0.0.1:8002"}
	for _, urlStr := range urls {
		backend, err := NewBackend(urlStr)
		if err != nil {
			panic(err)
		}
		if err := lb.AddBackend(backend); err != nil {
			panic(err)
		}
	}
	if err := lb.Start(); err != nil {
		panic(err)
	}

	port := 80
	engine := gin.Default()
	engine.Any("/*path", func(ctx *gin.Context) {
		if err := lb.HandleRequest(ctx.Writer, ctx.Request); err != nil {
			// 代理失败
			ctx.String(http.StatusBadGateway, err.Error())
			return
		}
		// 代理成功
		return
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
