package ginKit

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
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
			ctx.String(200, "ok")
			return
		})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(true), WithDefaultNoRouteHtml(true))
}
