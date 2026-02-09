package ginKit

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

//go:embed _test
var tmpFs embed.FS

/*
url: http://127.0.0.1/a/b/a.txt
*/
func TestStaticFS(t *testing.T) {
	// 通过 fs.Sub函数，获取 一个embed.FS实例 的 子embed.FS实例
	subFs, err := fs.Sub(tmpFs, "_test")
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	group := engine.Group("a")
	if err := StaticFS(group, "b", http.FS(subFs)); err != nil {
		panic(err)
	}

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}

/*
url: http://127.0.0.1/a/b/a.txt
*/
func TestStaticDir(t *testing.T) {
	engine := gin.Default()

	if err := StaticDir(engine, "/a/b", "_test", false); err != nil {
		panic(err)
	}

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
