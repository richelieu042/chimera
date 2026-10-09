//go:build go1.22

package randomKit

import (
	"math"
	"math/rand/v2"
)

type floatGrid struct {
	firstIndex float64
	scale      float64
	count      int
}

func newFloatGrid(min, max float64, precision int) (floatGrid, bool) {
	scale := math.Pow10(precision)
	if scale == 0 || math.IsInf(scale, 0) || math.IsNaN(scale) {
		return floatGrid{}, false
	}

	firstIndex := math.Ceil(min * scale)
	endIndex := math.Ceil(max * scale)
	count := endIndex - firstIndex
	if count < 1 || math.IsInf(count, 0) || math.IsNaN(count) {
		return floatGrid{}, false
	}
	if count >= float64(math.MaxInt) {
		return floatGrid{firstIndex: firstIndex, scale: scale, count: math.MaxInt}, true
	}
	return floatGrid{firstIndex: firstIndex, scale: scale, count: int(count)}, true
}

func (grid floatGrid) value(index int) float64 {
	return (grid.firstIndex + float64(index)) / grid.scale
}

// RandFloat 生成随机float64数字，可以指定范围和精度.（参考: random.RandFloat）
/*
	TODO: 看后续 duke-git/lancet(目前v2.3.9) 会不会加条件编译.

	@param precision 	(1) 精度（小数点后保留几位）
						(2) 真正返回值的小数位，可能会 小于 传参precision
						(3) 区间内没有指定精度的可表示值时，返回 min
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

	grid, ok := newFloatGrid(min, max, precision)
	if !ok {
		// 区间内没有指定精度的可表示值时，保留边界值作为唯一结果。
		return min
	}
	return grid.value(rand.IntN(grid.count))
}

// RandFloatSlice 生成随机float64数字切片，指定长度，范围和精度.（参考: random.RandFloats）
/*
	TODO: 看后续 duke-git/lancet(目前v2.3.9) 会不会加条件编译.

	@param n         请求的元素数量；超过指定范围和精度可生成的唯一值数量时，按可生成数量返回
	@param precision 精度（小数点后保留几位）
	@return (1) 切片内的元素范围: [min, max)
			(2) 切片内的元素不会重复
			(3) 区间内没有指定精度的可表示值时，仅返回 min
*/
func RandFloatSlice(n int, min, max float64, precision int) []float64 {
	if n <= 0 {
		return []float64{}
	}
	if max < min {
		min, max = max, min
	}

	grid, ok := newFloatGrid(min, max, precision)
	if !ok {
		return []float64{min}
	}
	if n > grid.count {
		n = grid.count
	}

	selected := make(map[int]struct{}, n)
	for index := grid.count - n; index < grid.count; index++ {
		candidate := rand.IntN(index + 1)
		if _, exists := selected[candidate]; exists {
			candidate = index
		}
		selected[candidate] = struct{}{}
	}

	nums := make([]float64, 0, n)
	for index := range selected {
		nums = append(nums, grid.value(index))
	}

	return nums
}
