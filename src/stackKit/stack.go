package stackKit

import (
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/listKit"
)

// Stack 堆栈（后进先出）
type Stack[V any] interface {
	// Push 放在最后面
	Push(ele V)

	// Pop 移除并返回最后面的元素
	Pop() (V, bool)

	Size() int
}

func NewStack[V any](safe ...bool) Stack[V] {
	return &stackImpl[V]{
		list: listKit.NewDoubleLinkedList(safe...),
	}
}

// NewStackFrom
/*
Deprecated: 有点耗性能，看 GoFrame 后续会不会支持泛型.
*/
func NewStackFrom[V any](s []V, safe ...bool) Stack[V] {
	var s1 []interface{} = sliceKit.ConvertElementType[V, interface{}](s, func(item V, index int) interface{} {
		return item
	})

	return &stackImpl[V]{
		list: listKit.NewDoubleLinkedListFrom(s1, safe...),
	}
}
