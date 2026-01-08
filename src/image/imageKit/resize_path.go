package imageKit

// Resize 缩放图片到指定尺寸（不保证纵横比）.
/*
	@param srcPath	源图片文件路径
	@param dstPath	目标图片文件路径
	@param width	目标宽度（像素）
	@param height	目标高度（像素）
	@return 返回错误信息，如果成功则返回 nil
*/
func Resize(srcPath, dstPath string, width, height int) (err error) {
	srcImg, _, err := DecodeFromPath(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImage(srcImg, width, height)
	if err != nil {
		return
	}

	return EncodeToPath(dstPath, dstImg)
}

// ResizeWithScale 按指定比例缩放图片（保证纵横比）.
func ResizeWithScale(srcPath, dstPath string, scale float64) (err error) {
	srcImg, _, err := DecodeFromPath(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImageWithScale(srcImg, scale)
	if err != nil {
		return
	}

	return EncodeToPath(dstPath, dstImg)
}

// ResizeKeepAspectRatio 按比例调整图片大小（保证纵横比；适应指定尺寸）.
func ResizeKeepAspectRatio(srcPath, dstPath string, maxWidth, maxHeight int) (err error) {
	srcImg, _, err := DecodeFromPath(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImageKeepAspectRatio(srcImg, maxWidth, maxHeight)
	if err != nil {
		return
	}

	return EncodeToPath(dstPath, dstImg)
}

// ResizeByWidth 按宽度等比例缩放图片.
func ResizeByWidth(srcPath, dstPath string, width int) (err error) {
	srcImg, _, err := DecodeFromPath(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImageByWidth(srcImg, width)
	if err != nil {
		return
	}

	return EncodeToPath(dstPath, dstImg)
}

// ResizeByHeight 按高度等比例缩放图片.
func ResizeByHeight(srcPath, dstPath string, height int) (err error) {
	srcImg, _, err := DecodeFromPath(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ResizeImageByHeight(srcImg, height)
	if err != nil {
		return
	}

	return EncodeToPath(dstPath, dstImg)
}
