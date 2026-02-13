package imageKit

import (
	"image"

	"github.com/richelieu042/chimera/v3/src/core/errorKit"
)

// Clip 裁剪图片.
/*
	@param img				原始图片实例（不能为nil！）
	@param x, y				裁剪区域左上角坐标
	@param width, height	裁剪区域的宽度和高度
*/
func Clip(img image.Image, x, y, width, height int) (image.Image, error) {
	if img == nil {
		return nil, errorKit.Newf("img is nil")
	}

	// 检查裁剪区域是否在图片范围内
	bounds := img.Bounds()
	if x < bounds.Min.X || y < bounds.Min.Y ||
		x+width > bounds.Max.X || y+height > bounds.Max.Y {
		return nil, errorKit.Newf("clip area out of bounds")
	}

	// 检查宽度和高度是否有效
	if width <= 0 || height <= 0 {
		return nil, errorKit.Newf("invalid crop dimensions: width=%d, height=%d", width, height)
	}

	// 定义裁剪区域
	cropRect := image.Rect(x, y, x+width, y+height)

	// 使用类型断言检查是否支持 SubImage
	subImager, ok := img.(interface {
		SubImage(r image.Rectangle) image.Image
	})
	if !ok {
		// 方案1：如果不支持 SubImage，手动复制像素
		return manualCrop(img, cropRect), nil
	}
	// 方案2：通过 SubImage 裁剪图片
	return subImager.SubImage(cropRect), nil
}

// manualCrop 手动裁剪图片（用于不支持 SubImage 的图片类型）
func manualCrop(img image.Image, cropRect image.Rectangle) image.Image {
	croppedImg := image.NewRGBA(image.Rect(0, 0, cropRect.Dx(), cropRect.Dy()))

	for y := cropRect.Min.Y; y < cropRect.Max.Y; y++ {
		for x := cropRect.Min.X; x < cropRect.Max.X; x++ {
			croppedImg.Set(x-cropRect.Min.X, y-cropRect.Min.Y, img.At(x, y))
		}
	}

	return croppedImg
}
