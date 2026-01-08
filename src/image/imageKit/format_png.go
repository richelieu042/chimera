package imageKit

import (
	"image/png"
	"os"

	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

// ToPng 将图片格式转换为".png".
func ToPng(src, dest string) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	srcImage, _, err := DecodeWithPath(src)
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	return png.Encode(destFile, srcImage)
}
