package strKit

import (
	"fmt"
	"hash"
	"hash/fnv"
	"sync"
)

// 针对：高并发场景
var fnvPool = sync.Pool{
	New: func() any {
		//return fnv.New32a()
		return fnv.New64a()
	},
}

// StringToDigits 将 字符串 映射为 指定长度的数字字符串.
/*
流程：字符串  →  哈希值(uint64)  →  取模  →  固定长度数字字符串

@param length	返回字符串的长度，建议 6~10
@param salt		[可选] 用于降低碰撞风险，好处：
				（1）碰撞 = 不同的输入，得到相同的输出
				（2）降低跨系统碰撞
				（3）防止被反推规律
				（4）不同业务场景可用不同salt
*/
func StringToDigits(s string, length int, salt string) string {
	if length <= 0 {
		length = 6
	}

	h := fnvPool.Get().(hash.Hash64)
	h.Reset()
	_, _ = h.Write([]byte(salt + s))

	sum := h.Sum64()
	fnvPool.Put(h)

	mod := uint64(1)
	for i := 0; i < length; i++ {
		mod *= 10
	}

	num := sum % mod
	return fmt.Sprintf("%0*d", length, num)
}
