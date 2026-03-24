package strKit

import (
	"strings"
)

// IsEmpty strings.Compare比自建方法“==”的速度要快
/*
PS: 由于此方法被调用的频率较高，因此方法体内禁止调用自己写的方法，全用原生的.
*/
func IsEmpty(str string) bool {
	return strings.Compare(str, "") == 0
}

func IsNotEmpty(str string) bool {
	return strings.Compare(str, "") != 0
}

func IsEmptyString(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return IsEmpty(str)
	}
	return false
}

func HasEmpty(strings ...string) bool {
	for _, str := range strings {
		if IsEmpty(str) {
			return true
		}
	}
	return false
}

func IsAllEmpty(strings ...string) bool {
	for _, str := range strings {
		if IsNotEmpty(str) {
			return false
		}
	}
	return true
}

func IsAllNotEmpty(strings ...string) bool {
	return !HasEmpty(strings...)
}

// EmptyToDefault
/*
@param trimArgs 是否先对 传参str 进行trim处理？（默认：false，不处理）
*/
func EmptyToDefault(str, def string, trimArgs ...bool) string {
	trimFlag := false
	if len(trimArgs) > 0 {
		trimFlag = trimArgs[0]
	}
	if trimFlag {
		str = TrimSpace(str)
	}

	if IsEmpty(str) {
		return def
	}
	return str
}
