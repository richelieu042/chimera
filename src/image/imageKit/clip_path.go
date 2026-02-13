package imageKit

// ClipWithPath 裁剪图片.
func ClipWithPath(srcPath, dstPath string, x, y, width, height int) error {
	srcImg, _, err := DecodeFromPath(srcPath)
	if err != nil {
		return err
	}

	dstImg, err := Clip(srcImg, x, y, width, height)
	if err != nil {
		return err
	}

	return EncodeToPath(dstPath, dstImg)
}
