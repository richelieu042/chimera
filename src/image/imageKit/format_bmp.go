package imageKit

import (
	"os"

	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"golang.org/x/image/bmp"
)

// ToBmp 将图片格式转换为".bmp".
func ToBmp(src, dest string) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	srcImage, _, err := DecodeFromPath(src)
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	return bmp.Encode(destFile, srcImage)
}
