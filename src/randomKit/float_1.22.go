//go:build go1.22

package randomKit

import (
	"math/rand/v2"

	"github.com/duke-git/lancet/v2/mathutil"
)

// RandFloat 生成随机float64数字，可以指定范围和精度.（参考: random.RandFloat）
/*
	TODO: 看后续 duke-git/lancet(目前v2.3.1) 会不会加条件编译.

	@param precision 	(1) 精度（小数点后保留几位）
						(2) 真正返回值的小数位，可能会 小于 传参precision
	@return [min, max)

	e.g. 返回值的小数位，可能会 小于 传参precision
		randomKit.RandFloat(1, 2, 3) => 1.938
		randomKit.RandFloat(1, 2, 3) => 1.36
		randomKit.RandFloat(1, 2, 3) => 1.41
		randomKit.RandFloat(1, 2, 3) => 1.184
*/
func RandFloat(min, max float64, precision int) float64 {
	if min == max {
		return min
	}

	if max < min {
		min, max = max, min
	}

	n := rand.Float64()*(max-min) + min

	return mathutil.RoundToFloat(n, precision)
}

// RandFloatSlice 生成随机float64数字切片，指定长度，范围和精度.（参考: random.RandFloats）
/*
	TODO: 看后续 duke-git/lancet(目前v2.3.1) 会不会加条件编译.

	@param precision 精度（小数点后保留几位）
	@return (1) 切片内的元素范围: [min, max)
			(2) 切片内的元素不会重复
*/
func RandFloatSlice(n int, min, max float64, precision int) []float64 {
	nums := make([]float64, n)
	used := make(map[float64]struct{}, n)
	for i := 0; i < n; {
		r := RandFloat(min, max, precision)
		if _, use := used[r]; use {
			continue
		}
		used[r] = struct{}{}
		nums[i] = r
		i++
	}

	return nums
}
