package ginKit

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/component/web/proxy/forwardKit"
	"github.com/richelieu042/chimera/v3/src/config/viperKit"
	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/pathKit"
)

func TestMustSetUp(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		fmt.Printf("working dir: %s\n", wd)
	}

	type config struct {
		Gin *Config `json:"gin" yaml:"gin"`
	}

	path := "_chimera-lib/config.yaml"
	c := &config{}
	//err := yamlKit.UnmarshalFromFile(path, c)
	_, err := viperKit.UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}

	MustSetUp(c.Gin, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			//ctx.String(200, "ok")

			ctx.JSON(200,
				map[string]string{
					"ok": "111",
				})

			return
		})

		engine.Any("/small", func(ctx *gin.Context) {
			if err := forwardKit.ForwardToSingleHost(ctx.Writer, ctx.Request, "http://127.0.0.1:81", nil); err != nil {
				ctx.JSON(502,
					map[string]string{
						"error": err.Error(),
					})
			}
		})

		engine.Any("/big", func(ctx *gin.Context) {
			if err := forwardKit.ForwardToSingleHost(ctx.Writer, ctx.Request, "http://127.0.0.1:81", nil); err != nil {
				ctx.JSON(502,
					map[string]string{
						"error": err.Error(),
					})
			}
		})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(true), WithDefaultNoRouteHtml(true))
}
