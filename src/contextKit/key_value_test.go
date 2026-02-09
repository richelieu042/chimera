package contextKit

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/component/web/proxy/forwardKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/netKit"
	"go.uber.org/zap"
)

func TestAttachKeyValue(t *testing.T) {
	ctx := context.Background()

	ctx = AttachKeyValue(ctx, "key", "a")
	fmt.Println("value:", ctx.Value("key")) // value: a

	ctx = AttachKeyValue(ctx, "key", "b")
	ctx = AttachKeyValue(ctx, "key", "c")
	fmt.Println("value:", ctx.Value("key")) // value: c
}

// 启动后，需要手动访问url: http://127.0.0.1:8001/test
func TestAttachKeyValue1(t *testing.T) {
	go func() {
		port := 8002
		engine := gin.Default()
		engine.Any("/test", func(ctx *gin.Context) {
			/* 拿不到 */
			console.Info("[8002]", zap.Any("value", ctx.Request.Context().Value("key")))

			ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	port := 8001
	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		tmpCtx := AttachKeyValue(ctx.Request.Context(), "key", "a")
		ctx.Request = ctx.Request.WithContext(tmpCtx) // Richelieu: 此处必须覆盖 request!!!
		/* 能拿到 */
		console.Info("[8001]", zap.Any("value", ctx.Request.Context().Value("key")))

		if err := forwardKit.ForwardToHost(ctx.Writer, ctx.Request, "127.0.0.1:8002", nil); err != nil {
			console.Error("Fail to forward.", zap.String("error", err.Error()))
			ctx.String(502, err.Error())
			return
		}
		/* 能拿到 */
		console.Info("[8001]", zap.Any("value", ctx.Request.Context().Value("key")))
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
