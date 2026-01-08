package imageKit

import (
	"image"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

func processSrc(srcPath string) (image.Image, string, error) {
	// 检查 && 容错
	if err := strKit.AssertNotBlank(srcPath, "srcPath"); err != nil {
		return nil, "", err
	}
	srcPath = strKit.TrimSpace(srcPath)

	srcImg, srcFormat, err := DecodeFromPath(srcPath)
	if err != nil {
		return nil, "", errorKit.Wrapf(err, "fail to decode source image")
	}
	srcExt := "." + srcFormat

	return srcImg, srcExt, nil
}

func processDst(dstPath string, dstImg image.Image, srcExt string) (err error) {
	// 检查 && 容错
	if err := strKit.AssertNotBlank(dstPath, "dstPath"); err != nil {
		return err
	}
	dstPath = strKit.TrimSpace(dstPath)

	// 处理特殊情况: dstPath 不带格式，则使用 srcPath 的格式
	dstExt := fileKit.GetExt(dstPath)
	if strKit.IsEmpty(dstExt) {
		dstExt = srcExt
		dstPath += dstExt
	}

	return EncodeToPath(dstPath, dstImg)
}

// Resize 缩放图片到指定尺寸（不保证纵横比）.
/*
	@param srcPath	源图片文件路径
	@param dstPath	目标图片文件路径
	@param width	目标宽度（像素）
	@param height	目标高度（像素）
	@return 返回错误信息，如果成功则返回 nil
*/
func Resize(srcPath, dstPath string, width, height int) (err error) {
	srcImg, srcExt, err := processSrc(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImage(srcImg, width, height)
	if err != nil {
		return
	}

	return processDst(dstPath, dstImg, srcExt)
}

// ResizeWithScale 按指定比例缩放图片（保证纵横比）.
func ResizeWithScale(srcPath, dstPath string, scale float64) (err error) {
	srcImg, srcExt, err := processSrc(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImageWithScale(srcImg, scale)
	if err != nil {
		return
	}

	return processDst(dstPath, dstImg, srcExt)
}

// ResizeKeepAspectRatio 按比例调整图片大小（保证纵横比；适应指定尺寸）.
func ResizeKeepAspectRatio(srcPath, dstPath string, maxWidth, maxHeight int) (err error) {
	srcImg, srcExt, err := processSrc(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImageKeepAspectRatio(srcImg, maxWidth, maxHeight)
	if err != nil {
		return
	}

	return processDst(dstPath, dstImg, srcExt)
}

// ResizeByWidth 按宽度等比例缩放图片.
func ResizeByWidth(srcPath, dstPath string, width int) (err error) {
	srcImg, srcExt, err := processSrc(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImageByWidth(srcImg, width)
	if err != nil {
		return
	}

	return processDst(dstPath, dstImg, srcExt)
}

// ResizeByHeight 按高度等比例缩放图片.
func ResizeByHeight(srcPath, dstPath string, height int) (err error) {
	srcImg, srcExt, err := processSrc(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImageByHeight(srcImg, height)
	if err != nil {
		return
	}

	return processDst(dstPath, dstImg, srcExt)
}
