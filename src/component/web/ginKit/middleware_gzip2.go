package ginKit

import (
	"bytes"
	"compress/gzip"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/compress/gzipKit"
)

// NewGzipMiddleware2
/*
@param minLength 等效于 Nginx 的 gzip_min_length 配置项
*/
func NewGzipMiddleware2(level, minLength int) (middleware gin.HandlerFunc, err error) {
	if err = gzipKit.AssertValidLevel(level); err != nil {
		return
	}

	middleware = func(ctx *gin.Context) {
		req := ctx.Request
		if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") ||
			strings.Contains(req.Header.Get("Connection"), "Upgrade") ||
			strings.Contains(req.Header.Get("Accept"), "text/event-stream") {
			ctx.Next()
			return
		}

		ctx.Header("Content-Encoding", "gzip")
		ctx.Header("Vary", "Accept-Encoding")

		// 创建一个 ResponseWriter 接口的包装器
		grw := &gzipResponseWriter{
			ResponseWriter: ctx.Writer,

			level:     level,
			minLength: minLength,

			buffer: bytes.NewBuffer(nil),
		}
		// 替换原始的 ResponseWriter 接口
		ctx.Writer = grw

		// 继续处理请求
		ctx.Next()

		grw.WriteBody(ctx)
	}
	return
}

// 自定义的 ResponseWriter 接口的包装器
type gzipResponseWriter struct {
	gin.ResponseWriter

	level     int
	minLength int

	buffer *bytes.Buffer
}

// 重写 Write 方法，将响应体写入缓冲区
func (grw *gzipResponseWriter) Write(data []byte) (int, error) {
	return grw.buffer.Write(data)
}

func (grw *gzipResponseWriter) WriteString(s string) (int, error) {
	return grw.buffer.WriteString(s)
}

// WriteBody 在请求结束时进行 gzip 压缩
func (grw *gzipResponseWriter) WriteBody(ctx *gin.Context) {
	length := grw.buffer.Len()

	// (1) 不进行 gzip 压缩
	if length < grw.minLength {
		ctx.Writer.Header().Del("Content-Encoding")

		if length > 0 {
			_, _ = grw.buffer.WriteTo(grw.ResponseWriter)
		}
		return
	}

	// (2)进行 gzip 压缩
	gzipWriter, _ := gzip.NewWriterLevel(grw.ResponseWriter, grw.level)
	defer gzipWriter.Close()
	_, _ = grw.buffer.WriteTo(gzipWriter)
	//_ = gzipWriter.Flush()
}
