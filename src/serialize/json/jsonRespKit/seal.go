package jsonRespKit

import "github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"

// Seal 封装成响应结构体实例，再序列化为json.
/*
PS: 需要先成功调用 MustSetUp || SetUp.
*/
func Seal(code string, data interface{}, msgArgs ...interface{}) (string, error) {
	bean := Pack(code, data, msgArgs...)

	return jsonKit.MarshalToString(bean)
}

// SealFully 封装成响应结构体实例，再序列化为json.
/*
PS: 需要先成功调用 MustSetUp || SetUp.
*/
func SealFully(code, msg string, data interface{}, msgArgs ...interface{}) (string, error) {
	bean := PackFully(code, msg, data, msgArgs...)

	return jsonKit.MarshalToString(bean)
}
