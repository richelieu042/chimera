package gocvKit

import (
	"fmt"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"gocv.io/x/gocv"
)

// MatchTemplate 模板匹配
/*
@param src		原图（大图）
@param template	模板（小图）
*/
func MatchTemplate(srcPath, templatePath string, grayFlag bool) error {
	// 在屏幕截图中查找某个按钮
	img := gocv.IMRead(srcPath, gocv.IMReadColor)
	tmpl := gocv.IMRead(templatePath, gocv.IMReadColor)

	// 转换为灰度图
	if grayFlag {
		grayImg := gocv.NewMat()
		defer grayImg.Close()
		if err := gocv.CvtColor(img, &grayImg, gocv.ColorBGRToGray); err != nil {
			return errorKit.Wrapf(err, "fail to convert color for source")
		}

		grayTmpl := gocv.NewMat()
		defer grayTmpl.Close()
		if err := gocv.CvtColor(tmpl, &grayTmpl, gocv.ColorBGRToGray); err != nil {
			return errorKit.Wrapf(err, "fail to convert color for templatePath")
		}

		img = grayImg
		tmpl = grayTmpl
	}

	result := gocv.NewMat()
	defer result.Close()

	// 在大图中查找小图
	if err := gocv.MatchTemplate(img, tmpl, &result, gocv.TmCcoeffNormed, gocv.NewMat()); err != nil {
		return errorKit.Wrapf(err, "fail to match template")
	}

	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
	fmt.Println("maxVal:", maxVal)
	fmt.Println("maxLoc.X:", maxLoc.X)
	fmt.Println("maxLoc.Y:", maxLoc.Y)

	return nil
}
