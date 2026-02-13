package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/gocvKit"
	"github.com/richelieu042/chimera/v3/src/image/imageKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/ocr/gosseractKit"
	"gocv.io/x/gocv"
)

func main() {
	path := "666.png"

	x0 := 748
	y0 := 995
	x1 := 962
	y1 := 1040

	err := imageKit.ClipWithPath(path, "aaa.png", x0, y0, x1-x0+1, y1-y0+1)
	if err != nil {
		panic(err)
	}
	text, err := gosseractKit.GertText("aaa.png")
	if err != nil {
		panic(err)
	}
	console.Infof("text: [%s]", text)

	sailing := strKit.Index(text, "航行中") != -1
	console.Infof("航行中: [%t]", sailing)

	if sailing {
		days, err := getDays(text)
		if err != nil {
			panic(err)
		}
		console.Infof("天数: [%.2f]", days)

		{
			//x0 := 1094
			//y0 := 637
			//x1 := 1395
			//y1 := 894
			//
			//err := imageKit.ClipWithPath(path, "bbb.png", x0, y0, x1-x0+1, y1-y0+1)
			//if err != nil {
			//	panic(err)
			//}
			//matchVal, maxLoc, err := gocvKit.MatchTemplate("bbb.png", "sail.png", gocv.TmCcoeffNormed, true)
			matchVal, matchLoc, err := gocvKit.MatchTemplate(path, "sail.png", gocv.TmCcoeffNormed, true)
			if err != nil {
				panic(err)
			}
			console.Infof("matchVal: [%.2f], matchLoc: [%v]", matchVal, matchLoc)
		}
	}

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

//package main
//
//import "github.com/richelieu042/chimera/v3/src/image/imageKit"
//
//func main() {
//	x0 := 1112
//	y0 := 650
//	x1 := 1222
//	y1 := 760
//
//	imageKit.ClipWithPath("2.png", "sail.png", x0, y0, x1-x0+1, y1-y0+1)
//}
