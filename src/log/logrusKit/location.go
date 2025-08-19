package logrusKit

import (
	"fmt"
	"runtime"

	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

func GetLocation(f *runtime.Frame) string {
	var funcName, fileName string

	s := strKit.Split(f.Function, ".")
	funcName = s[len(s)-1]

	s1 := strKit.Split(f.File, "/")
	length := len(s1)
	if length >= 2 {
		fileName = fmt.Sprintf("%s/%s:%d", s1[length-2], s1[length-1], f.Line)
	} else {
		fileName = fmt.Sprintf("%s:%d", f.File, f.Line)
	}

	return fmt.Sprintf("%s(%s)", funcName, fileName)
}
