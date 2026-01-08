package imageKit

import (
	"image"
	"image/jpeg"
	"image/png"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

func processSrc(srcPath string) (image.Image, string, error) {
	if err := strKit.AssertNotBlank(srcPath, "srcPath"); err != nil {
		return nil, "", err
	}

	srcImg, srcFormat, err := OpenAndDecode(srcPath)
	if err != nil {
		return nil, "", errorKit.Wrapf(err, "fail to decode source image")
	}
	srcExt := "." + srcFormat

	return srcImg, srcExt, nil
}

func processDst(dstPath string, dstImg image.Image, srcExt string) (err error) {
	if err := strKit.AssertNotBlank(dstPath, "dstPath"); err != nil {
		return err
	}

	// 处理特殊情况: dstPath不带格式
	dstExt := fileKit.GetExt(dstPath)
	if strKit.IsEmpty(dstExt) {
		dstExt = srcExt
		dstPath += dstExt
	}
	switch dstExt {
	case ".png":
	case ".jpg", ".jpeg":
	default:
		return errorKit.Newf("unsupported dstExt: %s", dstExt)
	}

	dstFile, err := fileKit.Create(dstPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = dstFile.Close()

		// 失败的情况下，毁尸灭迹（把生成的目标文件删了）
		if err != nil {
			_ = fileKit.Delete(dstPath)
		}
	}()

	// (4) 根据目标文件扩展名编码保存
	switch dstExt {
	case ".png":
		err = png.Encode(dstFile, dstImg)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(dstFile, dstImg, &jpeg.Options{Quality: 100})
	default:
		err = errorKit.Newf("unsupported dstExt: %s", dstExt)
	}
	return
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
