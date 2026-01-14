package main

import (
	"fmt"

	"github.com/richelieu-yang/chimera/v3/src/gocvKit"
)

func main() {
	err := gocvKit.MatchTemplate("/Users/richelieu/Downloads/big.jpeg", "/Users/richelieu/Downloads/small.png", true)
	fmt.Println(err)

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

//func main() {
//	// 读取原图和模板
//	img := gocv.IMRead("/Users/richelieu/Downloads/big.jpeg", gocv.IMReadColor)
//	defer img.Close()
//
//	tmpl := gocv.IMRead("/Users/richelieu/Downloads/small.jpeg", gocv.IMReadColor)
//	defer tmpl.Close()
//
//	//// 转换为灰度图
//	//grayImg := gocv.NewMat()
//	//defer grayImg.Close()
//	//gocv.CvtColor(img, &grayImg, gocv.ColorBGRToGray)
//	//
//	//grayTmpl := gocv.NewMat()
//	//defer grayTmpl.Close()
//	//gocv.CvtColor(tmpl, &grayTmpl, gocv.ColorBGRToGray)
//
//	// 创建结果矩阵
//	result := gocv.NewMat()
//	defer result.Close()
//
//	// 进行模板匹配
//	if err := gocv.MatchTemplate(img, tmpl, &result, gocv.TmCcoeffNormed, gocv.NewMat()); err != nil {
//		fmt.Println("匹配模板错误:", err)
//		return
//	}
//
//	// 查找最佳匹配位置
//	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
//
//	fmt.Printf("最佳匹配位置: %v, 匹配度: %.2f\n", maxLoc, maxVal)
//}
