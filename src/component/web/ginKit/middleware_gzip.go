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

@param level 压缩级别，此参数直接映射自 Go 标准库 compress/gzip，共有以下值：
				gzip.HuffmanOnly			-2		仅使用 Huffman 编码，不做 LZ77 匹配，速度极快但压缩率低
				gzip.DefaultCompression		-1		默认压缩级别（推荐，平衡速度与压缩率）（对应的实际级别是 6）
				gzip.NoCompression			0		不压缩（仅封装 gzip 格式，无实际压缩）
				gzip.BestSpeed				1		最快速度，压缩率最低
				2~8							2~8		中间级别，数字越大压缩率越高、速度越慢
				gzip.BestCompression		9		最高压缩率，速度最慢
			常见选择：
				一般 API 响应 → gzip.DefaultCompression（-1）
				延迟敏感场景 → gzip.BestSpeed（1）
				带宽极度受限 → gzip.BestCompression（9）
*/
func NewGzipMiddleware(level int, options ...gzip.Option) gin.HandlerFunc {
	return gzip.Gzip(level, options...)
}
