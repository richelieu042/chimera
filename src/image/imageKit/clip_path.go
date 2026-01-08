package imageKit

// Clip 裁剪图片.
func Clip(srcPath, dstPath string, x, y, width, height int) (err error) {
	srcImg, _, err := DecodeFromPath(srcPath)
	if err != nil {
		return
	}

	dstImg, err := ClipImage(srcImg, x, y, width, height)
	if err != nil {
		return
	}

	return EncodeToPath(dstPath, dstImg)
}
