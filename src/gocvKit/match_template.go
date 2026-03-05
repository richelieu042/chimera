package gocvKit

import (
	"image"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"gocv.io/x/gocv"
)

// MatchTemplate 模板匹配 - 在源图中查找模板图的“最佳”匹配位置（在大图中找小图）
/*
PS:
（1）如果需要再大图中截取小图再进行模板匹配，务必使用 imageKit.ClipWithPath 而非通过桌面工具进行截图，
	否则会导致：使用 gocv.TmCcoeffNormed 的情况下，maxVal 值会很小（即使肉眼可见的匹配）.

@param srcPath      源图像文件路径（大图）
@param templatePath 模板图像文件路径（小图）
@param matchMode    匹配算法模式（推荐使用: gocv.TmCcoeffNormed）
                    	- TmCcoeff/TmCcoeffNormed: 相关系数匹配（值越大越好）
						- TmSqdiff/TmSqdiffNormed: 平方差匹配（值越小越好）
                    	- TmCcorr/TmCcorrNormed: 相关性匹配（值越大越好）
@param grayArgs     可选参数，是否转换为灰度图处理（默认 false）
                    设为 true 可提高匹配速度并降低颜色干扰

@return maxVal      最大匹配值（相似度分数）
                    - 归一化模式（Normed）：范围 [-1, 1] 或 [0, 1]
                    - 非归一化模式：取值范围取决于图像内容
@return maxLoc      最佳匹配位置的左上角坐标点
@return err         错误信息
*/
func MatchTemplate(srcPath, templatePath string, matchMode gocv.TemplateMatchMode, grayArgs ...bool) (matchVal float32, matchLoc image.Point, err error) {
	// （1）读取源图像
	srcImg, err := DecodeFromPath(srcPath)
	defer srcImg.Close()
	if err != nil {
		return 0, image.Point{}, errKit.Wrapf(err, "failed to read source image")
	}

	// （2）读取模板图像
	tmplImg, err := DecodeFromPath(templatePath)
	defer tmplImg.Close()
	if err != nil {
		return 0, image.Point{}, errKit.Wrapf(err, "failed to read template image")
	}

	// (3)验证模板尺寸不能大于源图
	if tmplImg.Cols() > srcImg.Cols() || tmplImg.Rows() > srcImg.Rows() {
		err := errKit.Newf("template size (%dx%d) cannot be larger than source size (%dx%d)",
			tmplImg.Cols(), tmplImg.Rows(), srcImg.Cols(), srcImg.Rows())
		return 0, image.Point{}, err
	}

	// （4）确定是否需要灰度转换
	var grayFlag bool
	if len(grayArgs) > 0 {
		grayFlag = grayArgs[0]
	}
	// 准备用于匹配的图像
	var img, tmpl gocv.Mat
	if grayFlag {
		// 转换源图为灰度图
		img = gocv.NewMat()
		defer img.Close()
		if err := gocv.CvtColor(srcImg, &img, gocv.ColorBGRToGray); err != nil {
			return 0, image.Point{}, errKit.Wrapf(err, "fail to convert source image to grayscale")
		}

		// 转换模板图为灰度图
		tmpl = gocv.NewMat()
		defer tmpl.Close()
		if err := gocv.CvtColor(tmplImg, &tmpl, gocv.ColorBGRToGray); err != nil {
			return 0, image.Point{}, errKit.Wrapf(err, "fail to convert template image to grayscale")
		}
	} else {
		// 直接使用彩色图像
		img = srcImg
		tmpl = tmplImg
	}

	// 创建结果矩阵
	result := gocv.NewMat()
	defer result.Close()

	// 执行模板匹配
	if err := gocv.MatchTemplate(img, tmpl, &result, matchMode, gocv.NewMat()); err != nil {
		return 0, image.Point{}, errKit.Wrapf(err, "failed to perform template matching")
	}

	// 查找最佳匹配位置（其他所有匹配位置，无论分数是 0.99 还是 0.80，只要不是全局最大值，都会被直接丢掉）
	// 对于 TmSqdiff 和 TmSqdiffNormed，最小值是最佳匹配
	// 对于其他模式，最大值是最佳匹配
	minVal, maxVal, minLoc, maxLoc := gocv.MinMaxLoc(result)

	// 根据匹配模式返回对应的最佳值和位置
	switch matchMode {
	case gocv.TmSqdiff, gocv.TmSqdiffNormed:
		// 平方差模式：值越小越好
		return minVal, minLoc, nil
	default:
		// 其他模式：值越大越好
		return maxVal, maxLoc, nil
	}
}

// MatchTemplateWithThreshold 带阈值的模板匹配
/*
在 MatchTemplate 基础上增加阈值判断，只有匹配度超过阈值才认为匹配成功

@param srcPath      源图像文件路径
@param templatePath 模板图像文件路径
@param matchMode    匹配算法模式
@param threshold    匹配阈值（建议使用归一化模式，阈值范围 0-1）
                    - TmCcoeffNormed: 推荐阈值 0.8-0.9
                    - TmSqdiffNormed: 推荐阈值 0.1-0.2（越小越好）
@param grayArgs     是否转换为灰度图

@return matched     是否匹配成功
@return matchVal    匹配值
@return matchLoc    匹配位置
@return err         错误信息
*/
func MatchTemplateWithThreshold(srcPath, templatePath string, matchMode gocv.TemplateMatchMode, threshold float32, grayArgs ...bool) (matched bool, matchVal float32, matchLoc image.Point, err error) {
	matchVal, matchLoc, err = MatchTemplate(srcPath, templatePath, matchMode, grayArgs...)
	if err != nil {
		return false, 0, image.Point{}, err
	}

	// 根据匹配模式判断是否超过阈值
	switch matchMode {
	case gocv.TmSqdiff, gocv.TmSqdiffNormed:
		// 平方差模式：值越小越好，需要小于阈值
		matched = matchVal < threshold
	default:
		// 其他模式：值越大越好，需要大于阈值
		matched = matchVal > threshold
	}

	return matched, matchVal, matchLoc, nil
}

// GetMatchRect 获取匹配区域的矩形范围.
/*
根据匹配位置和模板尺寸，计算出完整的匹配矩形区域

@param matchLoc     匹配位置（左上角坐标）
@param templatePath 模板图像路径（用于获取模板尺寸）
@return rect        匹配区域的矩形（可用于绘制标记）
@return err         错误信息
*/
func GetMatchRect(matchLoc image.Point, templatePath string) (rect image.Rectangle, err error) {
	tmpl := gocv.IMRead(templatePath, gocv.IMReadColor)
	if tmpl.Empty() {
		return image.Rectangle{}, errKit.Newf("failed to read template image: %s", templatePath)
	}
	defer tmpl.Close()

	// 计算矩形区域：从匹配点到 (匹配点 + 模板尺寸)
	rect = image.Rect(
		matchLoc.X,
		matchLoc.Y,
		matchLoc.X+tmpl.Cols(),
		matchLoc.Y+tmpl.Rows(),
	)

	return rect, nil
}
