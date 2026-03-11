package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
)

func main() {
	engine := gin.Default()

	engine.Any("/ccc", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, timeKit.FormatCurrent(timeKit.FormatB))
	})

	if err := engine.Run(":8888"); err != nil {
		panic(err)
	}
}
