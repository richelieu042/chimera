package ginKit

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultNoRouteHtml 使用自带的404页面.
func DefaultNoRouteHtml(engine *gin.Engine) error {
	templ, err := template.New("").ParseFS(efs, "_html/*.html")
	if err != nil {
		return err
	}
	engine.SetHTMLTemplate(templ)

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.min.html", gin.H{
			//"route": ctx.Request.URL.Path,
		})
	})
	return nil
}
