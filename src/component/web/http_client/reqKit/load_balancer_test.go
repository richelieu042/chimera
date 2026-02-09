package reqKit

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/netKit"
)

func TestNewLbClient(t *testing.T) {
	go func() {
		port := 8002

		engine := gin.Default()
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	// 此时只有 "http://127.0.0.1:8002/test" 是有效的
	urls := []string{
		"http://127.0.0.1:8000/test",
		"http://127.0.0.1:8001/test",
		"http://127.0.0.1:8002/test",
		"http://127.0.0.1:8003/test",
	}
	c := NewClient(WithDev())
	lbc, err := NewLbClient(c, urls, time.Millisecond*100, nil)
	if err != nil {
		panic(err)
	}
	resp, err := lbc.Get(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ToString())
}
