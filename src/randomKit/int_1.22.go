//go:build go1.22

package randomKit

import (
	"math/rand/v2"
)

// Int 生成随机int.（参考: random.RandInt）
/*
	TODO: 看后续 duke-git/lancet(目前v2.3.1) 会不会加条件编译.

	PS:
	(1) 如果min == max，将返回 min;
	(2) 如果min > max，将交换两者的值.

	@param min 可以 < 0
	@param max 可以 < 0
	@return 范围: [min, max)
*/
func Int(min, max int) int {
	//return random.RandInt(min ,max)

	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return rand.IntN(max-min) + min
}
