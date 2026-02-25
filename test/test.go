package main

import (
	"fmt"
	"hash/fnv"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

func stringTo6Digits(s string) string {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	num := h.Sum32() % 1000000      // 取模得到0~999999
	return fmt.Sprintf("%06d", num) // 补零保证6位
}

func main() {
	fmt.Println(stringTo6Digits("ylx")) // 例如输出: 082345

	fmt.Println(strKit.StringToDigits("ylx", 6, ""))
}
