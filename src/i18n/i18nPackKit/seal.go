package i18nPackKit

import "github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"

// Seal
/*
@param langs 可以为空（包括nil）
*/
func Seal(langs []string, code string, data interface{}, msgArgs ...interface{}) (string, error) {
	bean := Pack(langs, code, data, msgArgs...)
	return jsonKit.MarshalToString(bean)
}

// SealFully
/*
@param langs 可以为空（包括nil）
*/
func SealFully(langs []string, code, msg string, data interface{}, msgArgs ...interface{}) (string, error) {
	bean := PackFully(langs, code, msg, data, msgArgs...)
	return jsonKit.MarshalToString(bean)
}
