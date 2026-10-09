package randomKit

import "github.com/duke-git/lancet/v2/random"

var (
	// Int 生成随机int
	/*
		PS:
		(1) 如果min == max，将返回 min;
		(2) 如果min > max，将交换两者的值.

		@param min 可以 < 0
		@param max 可以 < 0
		@return 范围: [min, max)
	*/
	Int func(min, max int) int = random.RandInt

	RandIntSlice func(length, min, max int) []int = random.RandIntSlice
)
