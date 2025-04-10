package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		console.Info("---")
		s := []string{
			"X-Forwarded-For",
			"X-Real-IP",
			"Forwarded",
			"Via",
			"CF-Connecting-IP",
			"True-Client-IP",
		}
		for _, key := range s {
			value := ctx.Request.Header.Get(key)
			console.Infof("[HEADER] key: %s, value: %s", key, value)
		}

		console.Infof("ClientIP: %s", ctx.ClientIP())
		console.Infof("RemoteIP: %s", ctx.RemoteIP())
		console.Info("===")

		ctx.String(200, "Hello world!")
	})

	if err := engine.Run(":8001"); err != nil {
		panic(err)
	}
}
