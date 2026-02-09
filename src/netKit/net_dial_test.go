package netKit

import (
	"testing"
	"time"

	"github.com/richelieu042/chimera/v3/src/log/console"
)

func TestDialContextWithTimeout(t *testing.T) {
	//go func() {
	//	port := 8888
	//	engine := gin.Default()
	//	engine.Any("/test", func(ctx *gin.Context) {
	//		ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
	//	})
	//	if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
	//		panic(err)
	//	}
	//}()
	//time.Sleep(time.Second)

	timeout := time.Second * 2
	console.Info("---")
	conn, err := DialTimeout("tcp", "127.0.0.1:8888", timeout)
	if err != nil {
		console.Error(err.Error())
		return
	}
	defer conn.Close()
	console.Info("===")
}
