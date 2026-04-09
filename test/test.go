package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(getDays("阿是擦上次   1.5天期望的   2天"))
}

// getDays 提取字符串中"天"前面的数字（严格匹配）
func getDays(s string) (float64, error) {
	re := regexp.MustCompile(`(\d+\.?\d*)天`)
	match := re.FindStringSubmatch(s)

	if len(match) > 1 {
		return strconv.ParseFloat(match[1], 64)
	}
	return 0, fmt.Errorf("invalid string: %s", s)
}
