package sliceKit

import (
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/samber/lo"
)
import lop "github.com/samber/lo/parallel"

// ConvertElementType 遍历并修改元素的类型（返回值是一个新的slice实例）.
/*
@param string	可以为nil
@param iteratee 不能为nil，会导致panic: runtime error: invalid memory address or nil pointer dereference
@return			必定不为nil（保底空的slice实例）

e.g. []int => []string
	s := sliceKit.ConvertElementType([]int{0, 1, 2, 3}, func(item int, index int) string {
		return "0x" + strconv.Itoa(item)
	})
	fmt.Println(s) // [0x0 0x1 0x2 0x3]
*/
func ConvertElementType[T any, R any](collection []T, converter func(item T, index int) R) []R {
	return lo.Map(collection, converter)
}

// ConvertElementTypeE
/*
对 lo.Map 就行了拓展.
*/
func ConvertElementTypeE[T any, R any](collection []T, converter func(item T, index int) (R, error)) ([]R, error) {
	result := make([]R, len(collection))

	var err error
	for i, item := range collection {
		result[i], err = converter(item, i)
		if err != nil {
			return nil, errorKit.Wrapf(err, "fail to convert element(index: %d, value: %v)", i, item)
		}
	}

	return result, nil
}

// ConvertElementTypeInParallel 并发地遍历并修改元素的类型（返回值是一个新的slice实例）.
/*
PS: 使用协程更加高效，但要慎用!!!

@param string	可以为nil
@param iteratee 不能为nil，会导致panic: runtime error: invalid memory address or nil pointer dereference
@return			必定不为nil（保底空的slice实例）
*/
func ConvertElementTypeInParallel[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	return lop.Map(collection, iteratee)
}
