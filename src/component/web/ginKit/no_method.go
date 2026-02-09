package ginKit

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

func DefaultNoMethod(engine *gin.Engine) {
	engine.NoMethod(func(ctx *gin.Context) {
		var allowed []string
		route := ctx.Request.URL.Path
		routeInfoSlice := engine.Routes()
		for _, routeInfo := range routeInfoSlice {
			if routeInfo.Path == route {
				allowed = append(allowed, routeInfo.Method)
			}
		}
		// http状态码405，需要加上响应头 Allow
		ctx.Header("Allow", sliceKit.Join(allowed, ", "))

		text := fmt.Sprintf("Method(%s) isn't allowed for route(%s), allowed methods is %s.", ctx.Request.Method, ctx.Request.URL.Path, allowed)
		if strKit.IsNotEmpty(serviceInfo) {
			text = fmt.Sprintf("[%s] %s", serviceInfo, text)
		}
		ctx.String(http.StatusMethodNotAllowed, text)
	})
}
