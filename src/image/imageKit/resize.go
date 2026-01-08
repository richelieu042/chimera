package imageKit

import (
	"image"
	"image/jpeg"
	"image/png"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"golang.org/x/image/draw"
)

// ResizeImage 缩放图片到指定尺寸
/*
	@param srcPath	源图片文件路径
	@param dstPath	目标图片文件路径
	@param width	目标宽度（像素）
	@param height	目标高度（像素）
	@return 返回错误信息，如果成功则返回 nil
*/
func ResizeImage(srcPath, dstPath string, width, height int) (err error) {
	// （1）打开源图片文件
	srcFile, err := fileKit.OpenReadOnly(srcPath)
	if err != nil {
		return errorKit.Wrapf(err, "fail to open source image")
	}
	defer srcFile.Close()

	// （2）解码图片（自动识别格式）
	srcImg, srcFormat, err := image.Decode(srcFile)
	if err != nil {
		return errorKit.Wrapf(err, "fail to decode source image")
	}

	// （3）确定目标文件的格式 && 创建文件（会自动创建父目录）
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

		if err != nil {
			fileKit.Delete(dstPath)
		}
	}()

	// (4) 创建目标图片对象
	dstImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// (5) 使用 CatmullRom 算法进行高质量缩放
	/*
		支持的插值算法包括：
			draw.NearestNeighbor - 最快但质量最低
			draw.ApproxBiLinear - 速度和质量平衡
			draw.BiLinear - 双线性插值
			draw.CatmullRom - 高质量，推荐使用
	*/
	draw.CatmullRom.Scale(dstImg, dstImg.Bounds(), srcImg, srcImg.Bounds(), draw.Over, nil)

	// (6) 根据目标文件扩展名编码保存
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
