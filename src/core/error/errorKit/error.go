package errorKit

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/richelieu042/chimera/v3/src/funcKit"
)

var (
	// Simplef 使用 err.Error() 或 "%v" 时，不会附带调用此函数的调用方的名称和文件信息.
	Simplef = gerror.Newf
)

// Newf 用于创建一个自定义错误信息的error对象，并包含堆栈信息.
/*
@param format !!!: 类似于fmt.Errorf()，	(1) 不应该首字母大写;
										(2) 不应该以标点符号结尾.
*/
func Newf(format string, args ...interface{}) error {
	skip := 1
	return gerror.NewSkipf(skip, funcKit.AddEntireCaller(skip+1, format), args...)
}

// NewfWithSkip 用于创建一个自定义错误信息的error对象，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。高级功能，一般开发者很少用得到.
/*
	@param skip (1) >=0
				(2) 0: 等价于 Newf || Newf
				(2) 1: 跳过1层（e.g. assert工具类）
*/
func NewfWithSkip(skip int, format string, args ...interface{}) error {
	skip++
	return gerror.NewSkipf(skip, funcKit.AddEntireCaller(skip+1, format), args...)
}

// Wrapf 用于包裹其他错误error对象，构造成多级的错误信息，包含堆栈信息.
/*
@param format !!!: 类似于fmt.Errorf()，	(1) 不应该首字母大写;
										(2) 不应该以标点符号结尾.
*/
func Wrapf(err error, format string, args ...interface{}) error {
	skip := 1
	return gerror.WrapSkipf(skip, err, funcKit.AddEntireCaller(skip+1, format), args...)
}
