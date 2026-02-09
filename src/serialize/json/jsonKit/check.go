package jsonKit

import "github.com/richelieu042/chimera/v3/src/validateKit"

func IsJson(data []byte) bool {
	return validateKit.Json(data) == nil
}

func IsJsonString(str string) bool {
	return validateKit.Json(str) == nil
}

//// IsJson 简单的检测，并不一定准确.
///*
//PS: 这种方法可以粗略地判断字符串是否是 JSON 格式，但不如反序列化方法那样严谨。它不能处理所有的边缘情况，比如嵌套错误或格式错误的 JSON。为了更准确的判断，仍建议使用反序列化方法。
//*/
//func IsJson(data []byte) (rst bool) {
//	length := len(data)
//	if length >= 2 {
//		if data[0] == '{' {
//			rst = data[length-1] == '}'
//		} else if data[0] == '[' {
//			rst = data[length-1] == ']'
//		}
//	}
//	return
//}
//
//// IsJsonString 简单的检测，并不一定准确.
///*
//PS: 这种方法可以粗略地判断字符串是否是 JSON 格式，但不如反序列化方法那样严谨。它不能处理所有的边缘情况，比如嵌套错误或格式错误的 JSON。为了更准确的判断，仍建议使用反序列化方法。
//*/
//func IsJsonString(str string) (rst bool) {
//	if strKit.StartWith(str, "{") {
//		rst = strKit.EndWith(str, "}")
//	} else if strKit.StartWith(str, "[") {
//		rst = strKit.EndWith(str, "]")
//	}
//	return
//}
