package imageKit

import (
	"os"

	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"golang.org/x/image/tiff"
)

// ToTiff 将图片格式转换为".tiff".
func ToTiff(src, dest string, opts *tiff.Options) error {
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

	return tiff.Encode(destFile, srcImage, opts)
}
