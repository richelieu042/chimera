package imageKit

import (
	"image/jpeg"
	"image/png"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

// Resize 缩放图片到指定尺寸（不保证纵横比）
/*
	@param srcPath	源图片文件路径
	@param dstPath	目标图片文件路径
	@param width	目标宽度（像素）
	@param height	目标高度（像素）
	@return 返回错误信息，如果成功则返回 nil
*/
func Resize(srcPath, dstPath string, width, height int) (err error) {
	// (0) 参数检查
	if err := strKit.AssertNotBlank(srcPath, "srcPath"); err != nil {
		return err
	}
	if err := strKit.AssertNotBlank(dstPath, "dstPath"); err != nil {
		return err
	}

	// (1) 打开源图片文件 && 解码图片（自动识别格式）
	srcImg, srcFormat, err := OpenAndDecode(srcPath)
	if err != nil {
		return errorKit.Wrapf(err, "fail to decode source image")
	}

	// (2) 确定目标文件的格式 && 创建文件（会自动创建父目录）
	dstExt := fileKit.GetExt(dstPath)
	if strKit.IsEmpty(dstExt) {
		dstExt = "." + srcFormat
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
		dstFile.Close()

		// 失败的情况下，毁尸灭迹
		if err != nil {
			_ = fileKit.Delete(dstPath)
		}
	}()

	// (3) 创建目标图片对象，使用 CatmullRom 算法进行高质量缩放
	dstImg, err := ResizeImage(srcImg, width, height)
	if err != nil {
		return
	}

	// (4) 根据目标文件扩展名编码保存
	switch dstExt {
	case ".png":
		// PNG 格式
		err = png.Encode(dstFile, dstImg)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(dstFile, dstImg, &jpeg.Options{Quality: 100})
	default:
		err = errorKit.Newf("unsupported dstExt: %s", dstExt)
	}
	return
}
