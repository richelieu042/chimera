package forwardKit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/log/logKit"
	"github.com/richelieu042/chimera/v3/src/netKit"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

/*
证明: 代理失败（目标服务不存在）会返回error.

访问url: http://127.0.0.1/test
效果: 	将 http://127.0.0.1/test 转发给 http://127.0.0.1:8000/test
*/
func TestForwardToUrl(t *testing.T) {
	url := "http://127.0.0.1:8000"

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		errLog := logKit.NewStdoutLogger("")
		err := ForwardToSingleHost(ctx.Writer, ctx.Request, url, errLog)
		if err != nil {
			console.Error("Fail to forward.", zap.String("error", err.Error()))
			ctx.String(http.StatusBadGateway /*502*/, err.Error())
			return
		}
		console.Info("Manager to forward.")
		return
	})
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}

func TestForwardToHostComplexly(t *testing.T) {
	/* 目标服务(target) */
	go func() {
		port := 8001

		engine := gin.Default()
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(500, fmt.Sprintf("This is [%d].", port))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	/* 代理服务 */
	errLog := logKit.NewStdoutLogger("")
	engine := gin.Default()
	modifyResponse := func(resp *http.Response) error {
		if resp.StatusCode != 200 {
			return errorKit.Simplef("invalid status(%s)", resp.Status)
		}
		return nil
	}
	engine.Any("/test", func(ctx *gin.Context) {
		err := ForwardToHostComplexly(ctx.Writer, ctx.Request, "127.0.0.1:8001", errLog, nil, modifyResponse)
		if err != nil {
			logrus.WithError(err).Info("Fail to forward.")
			ctx.String(500, err.Error())
			return
		}
		logrus.Info("Manager to forward.")
		return
	})
	if err := engine.Run(":8000"); err != nil {
		panic(err)
	}
}
