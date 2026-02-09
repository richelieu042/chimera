package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/netKit"
)

func main() {
	port := 8002

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
