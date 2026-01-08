package imageKit

import (
	"image"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"golang.org/x/image/draw"
)

// resizeImage 私有函数，调用此函数前须确保传参没有问题
func resizeImage(srcImg image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	/*
		支持的插值算法包括：
			draw.NearestNeighbor - 最快但质量最低
			draw.ApproxBiLinear - 速度和质量平衡
			draw.BiLinear - 双线性插值
			draw.CatmullRom - 高质量，推荐使用
	*/
	draw.CatmullRom.Scale(dst, dst.Bounds(), srcImg, srcImg.Bounds(), draw.Over, nil)
	return dst
}

// ResizeImage 缩放图片到指定尺寸（不保证纵横比）.
func ResizeImage(srcImg image.Image, width, height int) (image.Image, error) {
	if srcImg == nil {
		return nil, errorKit.Newf("srcImg is nil")
	}
	if width <= 0 {
		return nil, errorKit.Newf("invalid width: %d", width)
	}
	if height <= 0 {
		return nil, errorKit.Newf("invalid height: %d", height)
	}

	return resizeImage(srcImg, width, height), nil
}

// ResizeImageWithScale 按指定比例缩放图片（保证纵横比）.
func ResizeImageWithScale(srcImg image.Image, scale float64) (image.Image, error) {
	if srcImg == nil {
		return nil, errorKit.Newf("srcImg is nil")
	}
	if scale <= 0 {
		return nil, errorKit.Newf("invalid scale: %f", scale)
	}

	bounds := srcImg.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()
	targetWidth := int(float64(srcWidth) * scale)
	targetHeight := int(float64(srcHeight) * scale)

	return resizeImage(srcImg, targetWidth, targetHeight), nil
}

// ResizeImageKeepAspectRatio 按比例调整图片大小（保证纵横比；适应指定尺寸）.
/*
	@param	src			源图片对象
	@param	maxWidth	最大宽度
	@param	maxHeight	最大高度
*/
func ResizeImageKeepAspectRatio(srcImg image.Image, maxWidth, maxHeight int) (image.Image, error) {
	if srcImg == nil {
		return nil, errorKit.Newf("srcImg is nil")
	}
	if maxWidth <= 0 {
		return nil, errorKit.Newf("invalid maxWidth: %d", maxWidth)
	}
	if maxHeight <= 0 {
		return nil, errorKit.Newf("invalid maxHeight: %d", maxHeight)
	}

	bounds := srcImg.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()

	// 计算缩放比例
	ratio := float64(srcWidth) / float64(srcHeight)
	targetWidth := maxWidth
	targetHeight := maxHeight
	if float64(maxWidth)/float64(maxHeight) > ratio {
		targetWidth = int(float64(maxHeight) * ratio)
	} else {
		targetHeight = int(float64(maxWidth) / ratio)
	}

	return resizeImage(srcImg, targetWidth, targetHeight), nil
}

// ResizeImageByWidth 按宽度等比例缩放图片.
/*
	@param src		源图片对象
	@param width	目标宽度
*/
func ResizeImageByWidth(srcImg image.Image, width int) (image.Image, error) {
	if srcImg == nil {
		return nil, errorKit.Newf("srcImg is nil")
	}
	if width <= 0 {
		return nil, errorKit.Newf("invalid width: %d", width)
	}

	bounds := srcImg.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()

	scale := float64(width) / float64(srcWidth)
	height := int(float64(srcHeight) * scale)

	return resizeImage(srcImg, width, height), nil
}

// ResizeImageByHeight 按高度等比例缩放图片.
/*
	@param src		源图片对象
	@param height	目标高度
*/
func ResizeImageByHeight(srcImg image.Image, height int) (image.Image, error) {
	if srcImg == nil {
		return nil, errorKit.Newf("srcImg is nil")
	}
	if height <= 0 {
		return nil, errorKit.Newf("invalid height: %d", height)
	}

	bounds := srcImg.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()

	scale := float64(height) / float64(srcHeight)
	width := int(float64(srcWidth) * scale)

	return resizeImage(srcImg, width, height), nil
}
