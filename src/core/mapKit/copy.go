package mapKit

import "github.com/richelieu042/chimera/v3/src/copyKit"

// Copy 浅拷贝
/*
@param src 可以为nil
@return 保底为空的map实例（不为nil）
*/
func Copy[K comparable, V any](m map[K]V) map[K]V {
	dest := map[K]V{}
	for k, v := range m {
		dest[k] = v
	}
	return dest
}

func DeepCopy[K comparable, V any](m map[K]V) map[K]V {
	return copyKit.DeepCopy(m)
}
