package ginKit

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// NewGzipMiddleware
/*
缺陷: github.com/gin-contrib/gzip v1.2.5 还不支持设置 minContentLength.

PS:
(1) 涉及多个服务（请求转发）的场景下，(a) 最外层的务使用gzip压缩;
								(b) 内层的服务不使用gzip压缩.
(2) Gzip通常不建议用来压缩图片.
*/
func NewGzipMiddleware(level int, options ...gzip.Option) gin.HandlerFunc {
	return gzip.Gzip(level, options...)
}
