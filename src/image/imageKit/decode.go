package imageKit

import (
	"bytes"
	"image"
	"io"

	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

var (
	// Decode 解码图片（部分特殊格式不支持; path => image.Image）.
	/*
	   @param r 类型可以是: *os.File（用完记得调用Close()）
	   @return 第1个: image.Image实例
	   		第2个: 表示图像的格式名称，例如 "png"、"jpeg" 等（不带"." && 转为小写）
	   		第3个: error（可能为nil）
	*/
	Decode func(r io.Reader) (img image.Image, format string, err error) = image.Decode
)

// DecodeWithPath 解码图片.
/*
@return 第2个: 图片的格式名称，例如 "png"、"jpeg" 等（不带"." && 转为小写）
*/
func DecodeWithPath(path string) (img image.Image, format string, err error) {
	path = strKit.TrimSpace(path)
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return nil, "", err
	}

	f, err := fileKit.OpenReadOnly(path)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	return Decode(f)
}

// DecodeWithBytes []byte => image.Image
func DecodeWithBytes(imgData []byte) (img image.Image, format string, err error) {
	if err := sliceKit.AssertNotEmpty(imgData, "imgData"); err != nil {
		return nil, "", err
	}

	// 将 []byte 数据转换为 io.Reader
	imgReader := bytes.NewReader(imgData)

	return Decode(imgReader)
}
