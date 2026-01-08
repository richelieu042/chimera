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

func ResizeImage(src image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	/*
		支持的插值算法包括：
			draw.NearestNeighbor - 最快但质量最低
			draw.ApproxBiLinear - 速度和质量平衡
			draw.BiLinear - 双线性插值
			draw.CatmullRom - 高质量，推荐使用
	*/
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}

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
	if width <= 0 {
		return errorKit.Newf("invalid width: %d", width)
	}
	if height <= 0 {
		return errorKit.Newf("invalid height: %d", height)
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
	dstImg := ResizeImage(srcImg, width, height)

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

//// ResizeWithScale 按指定比例缩放图片.
///*
//	@param scale 缩放比例，如 0.5 表示缩小到原来的 50%，2.0 表示放大到 200%
//*/
//func ResizeWithScale(srcPath, dstPath string, scale float64) error {
//	// (0) 参数检查
//	if err := strKit.AssertNotBlank(srcPath, "srcPath"); err != nil {
//		return err
//	}
//	if err := strKit.AssertNotBlank(dstPath, "dstPath"); err != nil {
//		return err
//	}
//	if scale <= 0 {
//		return errorKit.Newf("invalid scale: %f", scale)
//	}
//
//	OpenAndDecode(srcPath)
//
//	bounds := src.Bounds()
//	srcWidth := bounds.Dx()
//	srcHeight := bounds.Dy()
//	targetWidth := int(float64(srcWidth) * scale)
//	targetHeight := int(float64(srcHeight) * scale)
//
//	return nil
//}
