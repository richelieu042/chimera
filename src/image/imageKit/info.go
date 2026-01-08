package imageKit

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

// GetSize 获取图片的宽和高.
/*
PS: 单位为像素（px）.
*/
func GetSize(path string) (width int, height int, err error) {
	if err = fileKit.AssertExistAndIsFile(path); err != nil {
		return
	}

	imgFile, err := os.Open(path)
	if err != nil {
		err = errorKit.Wrapf(err, "fail to open")
		return
	}
	defer imgFile.Close()

	img, _, err := Decode(imgFile)
	if err != nil {
		err = errorKit.Wrapf(err, "fail decode")
		return
	}
	bounds := img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()
	return
}

// GetInfo 获取图片的信息（宽、高、后缀名）.
/*
【图像处理】Golang 获取常用图像的宽高总结
	https://www.cnblogs.com/voipman/p/16108320.html
*/
func GetInfo(path string) (*Info, error) {
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	extName := fileKit.GetExtName(path)
	if err := strKit.AssertNotEmpty(extName, "extName"); err != nil {
		return nil, err
	}
	extName = strKit.ToLower(extName)

	var imgConf image.Config
	switch extName {
	case "jpg", "jpeg":
		imgConf, err = jpeg.DecodeConfig(bytes.NewReader(data))
	case "webp":
		imgConf, err = webp.DecodeConfig(bytes.NewReader(data))
	case "png":
		imgConf, err = png.DecodeConfig(bytes.NewReader(data))
	case "tif", "tiff":
		imgConf, err = tiff.DecodeConfig(bytes.NewReader(data))
	case "gif":
		imgConf, err = gif.DecodeConfig(bytes.NewReader(data))
	case "bmp":
		imgConf, err = bmp.DecodeConfig(bytes.NewReader(data))
	default:
		return nil, errorKit.Newf("invalid extName(%s)", extName)
	}
	if err != nil {
		return nil, err
	}
	return &Info{
		ExtName:    extName,
		ColorModel: imgConf.ColorModel,
		Width:      imgConf.Width,
		Height:     imgConf.Height,
	}, nil
}
