package ginKit

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// LoadHTMLFiles 加载（多个）html文件.
/*
PS: 搭配 gin.Context.HTML() 使用.
*/
func LoadHTMLFiles(engine *gin.Engine, filePaths ...string) {
	engine.LoadHTMLFiles(filePaths...)
}

// LoadHTMLGlob
/*
PS: 搭配 gin.Context.HTML() 使用.
*/
func LoadHTMLGlob(engine *gin.Engine, pattern string) {
	engine.LoadHTMLGlob(pattern)
}

// SetHTMLTemplate
/*
PS:
(1) 搭配 gin.Context.HTML() 使用.
(2) 参考: https://mp.weixin.qq.com/s/07YhlR3fFIbRPrvT6HUqng
*/
func SetHTMLTemplate(engine *gin.Engine, templ *template.Template) {
	engine.SetHTMLTemplate(templ)
}
