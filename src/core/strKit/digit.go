package strKit

import (
	"fmt"
	"hash/fnv"
)

// StringToDigits 将 字符串 映射为 指定长度的数字字符串.
/*
@param length	返回字符串的长度，建议 6~10
@param salt		[可选] 用于降低碰撞风险，好处：
				（1）降低跨系统碰撞
				（2）防止被反推规律
				（3）不同业务场景可用不同salt
*/
func StringToDigits(s string, length int, salt string) string {
	if length <= 0 {
		length = 6
	}

	h := fnv.New32a()
	_, _ = h.Write([]byte(salt + s))

	sum := h.Sum32()

	// 计算 10^length
	mod := uint32(1)
	for i := 0; i < length; i++ {
		mod *= 10
	}

	num := sum % mod

	format := fmt.Sprintf("%%0%dd", length)
	return fmt.Sprintf(format, num)
}
