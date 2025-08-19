package logrusKit

import (
	"github.com/rifflock/lfshook"
)

// NewLfsHook
/*
使用场景:
(1) 指定日志级别(s)，可以也输出到其它输出，不影响原输出.
*/
var NewLfsHook = lfshook.NewHook
