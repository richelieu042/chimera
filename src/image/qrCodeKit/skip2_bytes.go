package qrCodeKit

import (
	"image/color"

	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/skip2/go-qrcode"
)

// Generate 生成二维码([]byte类型).
/*
Deprecated: 推荐使用 EncodeToFile.

@param content 	二维码的内容
@param level 	一般使用 qrcode.Medium
@param size		二维码的宽高，单位: px
@return png图片的字节流
*/
func Generate(content string, level qrcode.RecoveryLevel, size int) ([]byte, error) {
	/* content */
	if err := strKit.AssertNotEmpty(content, "content"); err != nil {
		return nil, err
	}

	return qrcode.Encode(content, level, size)
}

// GenerateWithColor 参考了 qrcode.WriteColorFile.
/*
Deprecated: 推荐使用 EncodeToFile.
*/
func GenerateWithColor(content string, level qrcode.RecoveryLevel, size int, background, foreground color.Color) ([]byte, error) {
	/* content */
	if err := strKit.AssertNotEmpty(content, "content"); err != nil {
		return nil, err
	}

	var q *qrcode.QRCode

	q, err := qrcode.New(content, level)

	q.BackgroundColor = background
	q.ForegroundColor = foreground

	if err != nil {
		return nil, err
	}

	return q.PNG(size)
}
