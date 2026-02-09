package oxyKit

import (
	_ "github.com/richelieu042/chimera/v3/src/log/logrusInitKit"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestNewLoadBalancerHandler(t *testing.T) {
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.DebugLevel)

	lb, err := NewLoadBalancerHandler(nil, []string{
		"http://127.0.0.1:8000",
		//"http://127.0.0.1:8001",
		//"http://127.0.0.1:8002",
	}, logger, true)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	engine.Any("/*path", func(ctx *gin.Context) {
		lb(ctx.Writer, ctx.Request)
	})
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
