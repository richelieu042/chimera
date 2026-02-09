package sliceKit

import (
	"strings"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// PolyfillStringSlice
/*
@return 可能是一个新的slice实例
*/
func PolyfillStringSlice(s []string) []string {
	s = RemoveEmpty(s, true)
	s = Uniq(s)
	return s
}

// Join []string => string
/*
@param sep 分隔符

e.g.
(nil, "-")			=> ""
([]string{}, "-") 	=> ""
e.g.1
([]string{"1"}, ";") 					=> "1"
([]string{"0", "1", "2", "3", ""}, "-") => "0-1-2-3-"
*/
var Join func(s []string, sep string) string = strings.Join

// RemoveEmpty
/*
@param trimArgs 是否 先 对每个元素进行trim操作？默认：false

e.g.
(nil) 			=> nil
([]string{""})	=> []string{}
*/
func RemoveEmpty(s []string, trimArgs ...bool) []string {
	if s == nil {
		return nil
	}

	trimFlag := false
	if len(trimArgs) > 0 {
		trimFlag = trimArgs[0]
	}

	rst := make([]string, 0, len(s))
	if trimFlag {
		for _, str := range s {
			str = strKit.TrimSpace(str)
			if strKit.IsNotEmpty(str) {
				rst = append(rst, str)
			}
		}
	} else {
		for _, str := range s {
			if strKit.IsNotEmpty(str) {
				rst = append(rst, str)
			}
		}
	}
	return rst
}

// ContainsStringIgnoreCase 字符串str是否在切片s中？（不区分大小写）
func ContainsStringIgnoreCase(s []string, str string) bool {
	for _, tmp := range s {
		if strKit.EqualsIgnoreCase(tmp, str) {
			return true
		}
	}
	return false
}
