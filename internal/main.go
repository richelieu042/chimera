package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/image/imageKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/ocr/gosseractKit"
)

func main() {
	x0 := 748
	y0 := 995
	x1 := 962
	y1 := 1040

	path := "3.png"

	err := imageKit.ClipWithPath(path, "aaa.png", x0, y0, x1-x0+1, y1-y0+1)
	if err != nil {
		panic(err)
	}
	text, err := gosseractKit.GertText("aaa.png")
	if err != nil {
		panic(err)
	}
	fmt.Printf("text: [%s]\n", text)

	sailing := strKit.Index(text, "航行中") != -1
	console.Infof("航行中: [%t]", sailing)

	if sailing {
		days, err := getDays(text)
		if err != nil {
			panic(err)
		}
		console.Infof("天数: [%.2f]", days)
	}

	//if text == "航行中" {
	//	x0 := 865
	//	y0 := 995
	//	x1 := 965
	//	y1 := 1040
	//
	//	err := imageKit.ClipWithPath(path, "bbb.png", x0, y0, x1-x0+1, y1-y0+1)
	//	if err != nil {
	//		panic(err)
	//	}
	//	text1, err := gosseractKit.GertText("bbb.png")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("text1: [%s]\n", text1)
	//
	//	days, err := getDays(text1)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("天数: [%.2f]\n", days)
	//}

	fmt.Println("$$$")
}

// 提取字符串中"天"前面的数字（严格匹配）
func getDays(s string) (float64, error) {
	re := regexp.MustCompile(`(\d+\.?\d*)天`)
	match := re.FindStringSubmatch(s)

	if len(match) > 1 {
		return strconv.ParseFloat(match[1], 64)
	}
	return 0, fmt.Errorf("invalid string: %s", s)
}
