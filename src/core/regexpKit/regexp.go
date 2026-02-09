package regexpKit

import (
	"regexp"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// StringToRegexp
/*
PS: 主要是处理 传参str 中的"*".

@param str 			(1) 可以为""，此时：第1个返回值将匹配任意字符串，第2个返回值为nil;
					(2) 不会进行trim操作;
					(3) 不同于原生的 regexp.Compile().
@param borderArgs 	边界匹配
*/
func StringToRegexp(str string, borderArgs ...bool) (*regexp.Regexp, error) {
	startBorder := false
	endBorder := false

	switch len(borderArgs) {
	case 0:
		// do nothing
	case 1:
		startBorder = borderArgs[0]
	case 2:
		fallthrough
	default:
		startBorder = borderArgs[0]
		endBorder = borderArgs[1]
	}

	/* (1) 处理"." */
	str = strKit.ReplaceAll(str, ".", "\\.")

	/* (2) 处理"*" */
	if strKit.Index(str, "*") != -1 {
		// 额外兼容情况：连续"*"的情况
		starRegexp := regexp.MustCompile("\\*+")
		str = starRegexp.ReplaceAllString(str, ".+")
	}
	if startBorder {
		str = strKit.PrependIfMissing(str, "^")
	}
	if endBorder {
		str = strKit.AppendIfMissing(str, "$")
	}

	// 此时如果 str == ""：第一个返回值将匹配任意字符串，第二个返回值为nil
	return regexp.Compile(str)
}
