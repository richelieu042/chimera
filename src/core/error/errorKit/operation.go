package errorKit

import (
	"errors"

	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	// Is 作用: 传参err 的错误链中，是否存在和 传参target 匹配的error实例？(reports whether any error in err's tree matches target.)
	/*
		Deprecated: Use errors.Is() instead.

		PS: 支持第三方依赖 "github.com/gogf/gf/v2/errors/gerror".

		@param target 可以为nil
	*/
	Is func(err, target error) bool = errors.Is

	// As
	/*
		Deprecated: Use errors.As() instead.

		!!!:
		对于传参target，
		(1) 如果 Error() 绑定在结构体上（receiver为"值类型"），需要 1个 '&';
		(2) 如果 Error() 绑定在结构体指针上（receiver为"指针类型"），需要 2个 '&'.

		查找 传参err 的错误链中与 传参target 匹配的第一个错误，(1) 如果找到，则 将 传参target 设置为该错误值 && 返回true;
														(2) 否则 返回false.
		finds the first error in err's tree that matches target, and if one is found, sets
		target to that error value and returns true. Otherwise, it returns false.

		PS: 支持第三方依赖 "github.com/gogf/gf/v2/errors/gerror".

		@param target 	(1) 不能为nil，否则会 panic
						(2) 必须是指针类型，否则会 panic
	*/
	As func(err error, target any) bool = errors.As

	// Equal 错误比较
	Equal func(err, target error) bool = gerror.Equal

	// HasStack 判断错误是否带堆栈
	HasStack func(err error) bool = gerror.HasStack

	// Stack 获取堆栈信息
	Stack func(err error) string = gerror.Stack

	// Current 获取当前error
	Current func(err error) error = gerror.Current

	// Unwrap 获取层级错误的下一级错误error接口对象(当下一层级不存在时，返回nil)
	Unwrap func(err error) error = gerror.Unwrap

	// Cause 获取根错误error
	Cause func(err error) error = gerror.Cause
)
