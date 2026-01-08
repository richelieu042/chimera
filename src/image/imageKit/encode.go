package imageKit

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

// Encode 根据目标文件扩展名编码保存（保存为文件）.
/*
@param ext 文件扩展名（带'.'且小写字母化） e.g. ".png"
*/
func Encode(file io.Writer, img image.Image, ext string) (err error) {
	switch ext {
	case ".png":
		err = png.Encode(file, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	case ".gif":
		err = gif.Encode(file, img, nil)
	default:
		err = errorKit.Newf("unsupported ext: %s", ext)
	}
	return
}

func EncodeWithPath(path string, img image.Image) (err error) {
	// 参数检查 && 容错
	if err := strKit.AssertNotBlank(path, "path"); err != nil {
		return err
	}
	if img == nil {
		return errorKit.Newf("img is nil")
	}
	path = strKit.TrimSpace(path)

	ext := fileKit.GetExt(path)
	switch ext {
	case ".png":
	case ".jpg", ".jpeg":
	default:
		if strKit.IsEmpty(ext) {
			return errorKit.Newf("ext is empty")
		}
		return errorKit.Newf("unsupported ext: %s", ext)
	}

	file, err := fileKit.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()

		// 失败的情况下，毁尸灭迹（把生成的目标文件删了）
		if err != nil {
			_ = fileKit.Delete(path)
		}
	}()

	return Encode(file, img, ext)
}
