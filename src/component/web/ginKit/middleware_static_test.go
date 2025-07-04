package ginKit

import (
	"embed"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNewStaticMiddleware(t *testing.T) {
	engine := gin.Default()

	{
		/*
			e.g.
			http://127.0.0.1/a.txt
		*/
		m, err := NewStaticMiddleware("/", "_test", false)
		if err != nil {
			panic(err)
		}
		engine.Use(m)
	}
	{
		/*
			e.g.
			http://127.0.0.1/c/a.txt
		*/
		m, err := NewStaticMiddleware("/c", "_test", false)
		if err != nil {
			panic(err)
		}
		engine.Use(m)
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "OK")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}

//go:embed _test
var staticFs embed.FS

func TestNewStaticMiddlewareWithEmbedFolder(t *testing.T) {
	engine := gin.Default()

	{
		/*
			e.g.
			http://127.0.0.1/a.txt
		*/
		m, err := NewStaticMiddlewareWithEmbedFolder("/", staticFs, "_test")
		if err != nil {
			panic(err)
		}
		engine.Use(m)
	}
	{
		// Richelieu: 此处有问题!!!
		/*
			e.g.
			http://127.0.0.1/c/a.txt
		*/
		m, err := NewStaticMiddlewareWithEmbedFolder("/c", staticFs, "_test")
		if err != nil {
			panic(err)
		}
		engine.Use(m)
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "OK")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
