package jsonRespKit

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// Pack 封装成响应结构体实例.
/*
PS: 需要先成功调用 MustSetUp || SetUp.
*/
func Pack(code string, data interface{}, msgArgs ...interface{}) interface{} {
	return PackFully(code, "", data, msgArgs...)
}

// PackFully
/*
PS: 需要先成功调用 MustSetUp || SetUp.
*/
func PackFully(code, msg string, data interface{}, msgArgs ...interface{}) interface{} {
	if strKit.IsEmpty(msg) {
		msg = msgMap[code]
	}
	if strKit.IsNotEmpty(msg) && msgArgs != nil {
		msg = fmt.Sprintf(msg, msgArgs...)
	}
	return provider(code, msg, data)
}
