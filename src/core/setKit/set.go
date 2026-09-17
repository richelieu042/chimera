package setKit

import (
	mapset "github.com/deckarep/golang-set/v3"
)

// NewSet
/*
@param threadSafe 是否并发安全？

e.g.
	set := setKit.NewSet[interface{}](false)
	// Add成功
	fmt.Println(set.Add(1)) // true
	// Add失败
	fmt.Println(set.Add(1)) // false
*/
func NewSet[T comparable](threadSafe bool, args ...T) mapset.Set[T] {
	if !threadSafe {
		return mapset.NewThreadUnsafeSet(args...)
	}
	return mapset.NewSet(args...)
}

func NewSetFromMapKeys[T comparable, V any](threadSafe bool, val map[T]V) mapset.Set[T] {
	if !threadSafe {
		return mapset.NewThreadUnsafeSetFromMapKeys(val)
	}
	return mapset.NewSetFromMapKeys(val)
}

func NewSetWithSize[T comparable](threadSafe bool, cardinality int) mapset.Set[T] {
	if !threadSafe {
		return mapset.NewThreadUnsafeSetWithSize[T](cardinality)
	}
	return mapset.NewSetWithSize[T](cardinality)
}
