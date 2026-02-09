package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
)

func DefaultEngine() *gin.Engine {
	engine := gin.New()

	/*
		默认: false
		true: enable fallback Context.Deadline(), Context.Done(), Context.Err() and Context.Value() when Context.Request.Context() is not nil
	*/
	engine.ContextWithFallback = true

	// 默认: true
	engine.ForwardedByClientIP = true
	// 默认: []string{"X-Forwarded-For", "X-Real-IP"}
	engine.RemoteIPHeaders = httpKit.RemoteIPHeaders

	/*
		(1) 默认: true
		(2) RedirectTrailingSlash 字段用于处理请求 URL 末尾是否包含斜杠 (/) 的情况。如果设置为 true，Gin 会自动将没有末尾斜杠的 URL 重定向到带有斜杠的 URL，反之亦然。
		(3) 这种重定向有助于统一 URL 格式，避免由于末尾斜杠不同而导致的路由匹配失败。

		e.g.	设置为true
		如果请求的 URL 是 /foo，而实际路由是 /foo/，则会将请求重定向到 /foo/。
		如果请求的 URL 是 /bar/，而实际路由是 /bar，则会将请求重定向到 /bar。
	*/
	engine.RedirectTrailingSlash = true
	/*
		(1) 默认: false
		(2) RedirectFixedPath 字段用于处理 URL 中多余或错误的斜杠和大小写问题。如果设置为 true，Gin 会尝试修正 URL 并进行重定向。
		(3) 这种重定向可以帮助用户访问错误格式的 URL 时，自动调整到正确的路径，提高路由匹配的容错性和用户体验。

		e.g.	设置为true
		如果请求的 URL 是 /FOO///bar，而实际路由是 /foo/bar，则会将请求重定向到 /foo/bar。
		如果请求的 URL 是 /foo/BAR，而实际路由是 /foo/bar，则会将请求重定向到 /foo/bar。
	*/
	engine.RedirectFixedPath = true

	/*
		MaxMultipartMemory只是限制内存，不是针对文件上传文件大小，即使文件大小比这个大，也会写入临时文件。
		默认32MiB，并不涉及"限制上传文件的大小"，原因：上传的文件s按顺序存入内存中，累加大小不得超出 32Mb ，最后累加超出的文件就存入系统的临时文件中。非文件字段部分不计入累加。所以这种情况，文件上传是没有任何限制的。
		参考: https://studygolang.com/articles/22643
	*/
	engine.MaxMultipartMemory = 64 << 20

	return engine
}
