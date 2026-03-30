package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()

	engine.Any("/x", func(ctx *gin.Context) {
		ctx.String(200, "xixi")
	})

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
