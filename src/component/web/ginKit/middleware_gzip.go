package ginKit

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/gzip"
	gzip "github.com/richelieu042/gin-gzip-middleware"
)

// NewGzipMiddleware
/*
Richelieu: 官方版（github.com/gin-contrib/gzip v1.2.6）在请求转发forward时有问题，因此使用了自己的修改版（github.com/richelieu042/gin-gzip-middleware）.

PS:
(1) 涉及多个服务（请求转发）的场景下，(a) 最外层的务使用gzip压缩;
								(b) 内层的服务不使用gzip压缩.
(2) Gzip通常不建议用来压缩图片.
*/
func NewGzipMiddleware(level int, options ...gzip.Option) gin.HandlerFunc {
	return gzip.Gzip(level, options...)
}
