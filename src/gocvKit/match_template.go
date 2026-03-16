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
func MatchTemplate(srcPath, templatePath string, matchMode gocv.TemplateMatchMode, grayArgs ...bool) (matchVal float32, matchRect image.Rectangle, err error) {
	// （1）读取源图像
	srcImg, err := DecodeFromPath(srcPath)
	if err != nil {
		err = errKit.Wrapf(err, "failed to read source image")
		return
	}
	defer srcImg.Close()

	// （2）读取模板图像
	templImg, err := DecodeFromPath(templatePath)
	if err != nil {
		err = errKit.Wrapf(err, "failed to read template image")
		return
	}
	defer templImg.Close()

	// (3)验证模板尺寸不能大于源图
	if templImg.Cols() > srcImg.Cols() || templImg.Rows() > srcImg.Rows() {
		err = errKit.Newf("template size (%dx%d) cannot be larger than source size (%dx%d)",
			templImg.Cols(), templImg.Rows(), srcImg.Cols(), srcImg.Rows())
		return
	}

	// （4）确定是否需要灰度转换
	var grayFlag bool
	if len(grayArgs) > 0 {
		grayFlag = grayArgs[0]
	}
	// 准备用于匹配的图像
	var img, templ gocv.Mat
	if grayFlag {
		// （1）转换源图为灰度图
		img, err = ToGrayscale(srcImg)
		if err != nil {
			err = errKit.Wrapf(err, "fail to convert source image to grayscale")
			return
		}
		defer img.Close()

		// （2）转换模板图为灰度图
		templ, err = ToGrayscale(templImg)
		if err != nil {
			err = errKit.Wrapf(err, "fail to convert template image to grayscale")
			return
		}
		defer templ.Close()
	} else {
		// 直接使用彩色图像
		img = srcImg
		templ = templImg
	}

	result := gocv.NewMat() // 创建结果矩阵
	defer result.Close()
	mask := gocv.NewMat()
	defer mask.Close()
	// 执行模板匹配
	if err = gocv.MatchTemplate(img, templ, &result, matchMode, mask); err != nil {
		err = errKit.Wrapf(err, "failed to perform template matching")
		return
	}

	// 查找最佳匹配位置（其他所有匹配位置，无论分数是 0.99 还是 0.80，只要不是全局最大值，都会被直接丢掉）
	// 对于 TmSqdiff 和 TmSqdiffNormed，最小值是最佳匹配
	// 对于其他模式，最大值是最佳匹配
	minVal, maxVal, minLoc, maxLoc := gocv.MinMaxLoc(result)

	// 根据匹配模式返回对应的最佳值和位置
	switch matchMode {
	case gocv.TmSqdiff, gocv.TmSqdiffNormed:
		// 平方差模式：值越小越好
		loc := minLoc
		matchRect = image.Rect(loc.X, loc.Y, loc.X+templ.Cols(), loc.Y+templ.Rows())
		return minVal, matchRect, nil
	default:
		// 其他模式：值越大越好
		loc := maxLoc
		matchRect = image.Rect(loc.X, loc.Y, loc.X+templ.Cols(), loc.Y+templ.Rows())
		return maxVal, matchRect, nil
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
func MatchTemplateWithThreshold(srcPath, templatePath string, matchMode gocv.TemplateMatchMode, threshold float32, grayArgs ...bool) (matched bool, matchVal float32, rect image.Rectangle, err error) {
	matchVal, rect, err = MatchTemplate(srcPath, templatePath, matchMode, grayArgs...)
	if err != nil {
		return
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
	return matched, matchVal, rect, nil
}
