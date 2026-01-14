package main

import (
	"fmt"

	"github.com/richelieu-yang/chimera/v3/src/gocvKit"
	"gocv.io/x/gocv"
)

func main() {
	{
		maxVal, maxLoc, err := gocvKit.MatchTemplate("/Users/richelieu/Downloads/big.jpeg",
			"/Users/richelieu/Downloads/small.png",
			gocv.TmCcoeffNormed, true)

		if err != nil {
			panic(err)
		}
		fmt.Println(maxVal, maxLoc.X, maxLoc.Y)
		fmt.Println("---")
	}

	matched, maxVal, maxLoc, err := gocvKit.MatchTemplateWithThreshold("/Users/richelieu/Downloads/big.jpeg",
		"/Users/richelieu/Downloads/small.png",
		gocv.TmCcoeffNormed, 0.9, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(matched, maxVal, maxLoc.X, maxLoc.Y)

	//// 在屏幕截图中查找某个按钮
	//screenshot := gocv.IMRead("/Users/richelieu/Downloads/big.jpeg", gocv.IMReadColor) // 大图
	//buttonImg := gocv.IMRead("/Users/richelieu/Downloads/small.png", gocv.IMReadColor) // 小图（模板）
	//
	//result := gocv.NewMat()
	//defer result.Close()
	//
	//// 在大图中查找小图
	//if err := gocv.MatchTemplate(screenshot, buttonImg, &result, gocv.TmCcoeffNormed, gocv.NewMat()); err != nil {
	//	fmt.Println("匹配模板错误:", err)
	//	return
	//}
	//
	//_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
	//fmt.Println("maxVal:", maxVal)
	//fmt.Println("maxLoc.X:", maxLoc.X)
	//fmt.Println("maxLoc.Y:", maxLoc.Y)
	//
	//if maxVal > 0.8 {
	//	fmt.Printf("在屏幕坐标 (%d, %d) 找到了登录按钮\n", maxLoc.X, maxLoc.Y)
	//}

}
