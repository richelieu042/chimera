package randomKit

import "github.com/duke-git/lancet/v2/random"

var (
	// Float64 生成随机float64数字，可以指定范围和精度.（参考: random.RandFloat）
	/*
		@param precision 	(1) 精度（小数点后保留几位）
							(2) 真正返回值的小数位，可能会 小于 传参precision
		@return [min, max)

		e.g. 返回值的小数位，可能会 小于 传参precision
			randomKit.RandFloat(1, 2, 3) => 1.938
			randomKit.RandFloat(1, 2, 3) => 1.36
			randomKit.RandFloat(1, 2, 3) => 1.41
			randomKit.RandFloat(1, 2, 3) => 1.184
	*/
	Float64 func(min, max float64, precision int) float64 = random.RandFloat

	// Float64Slice 生成随机float64数字切片，指定长度，范围和精度.
	/*
		@param precision 精度（小数点后保留几位）
		@return (1) 切片内的元素范围: [min, max)
				(2) 切片内的元素不会重复
	*/
	Float64Slice func(n int, min, max float64, precision int) []float64 = random.RandFloats
)
