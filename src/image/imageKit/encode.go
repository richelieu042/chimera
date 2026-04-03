package imageKit

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// Encode image.Image => 文件
/*
@param ext 文件扩展名（带'.'且小写字母化） e.g. ".png"
*/
func Encode(file io.Writer, img image.Image, ext string) (err error) {
	// 参数检查 && 容错
	if file == nil {
		return errorKit.Newf("file is nil")
	}
	if img == nil {
		return errorKit.Newf("img is nil")
	}

	switch ext {
	case ".png":
		err = png.Encode(file, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	case ".gif":
		err = gif.Encode(file, img, nil)
	default:
		err = errorKit.Newf("unsupported ext: [%s]", ext)
	}
	return
}

// EncodeToPath image.Image => 文件
func EncodeToPath(path string, img image.Image) (err error) {
	path = strKit.TrimSpace(path)
	if img == nil {
		return errorKit.Newf("img is nil")
	}

	file, err := fileKit.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()

		// 失败的情况下，毁尸灭迹（把生成的目标文件删了）
		if err != nil {
			_ = fileKit.Remove(path)
		}
	}()

	ext := fileKit.GetExt(path)

	return Encode(file, img, ext)
}
