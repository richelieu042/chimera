package qrCodeKit

import (
	"testing"

	"github.com/richelieu-yang/chimera/v3/src/image/imageKit"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func TestEncodeToFile(t *testing.T) {
	err := EncodeToFile("https://www.example.com", "_test0.png")
	if err != nil {
		panic(err)
	}
}

// 带 logo 的二维码（注意，logo的尺寸不能大于二维码的1/5）.
func TestEncodeToFile1(t *testing.T) {
	img, _, err := imageKit.DecodeWithPath("_logo.png")
	if err != nil {
		panic(err)
	}

	err = EncodeToFile("https://www.example.com", "_test1.png", standard.WithLogoImage(img))
	if err != nil {
		panic(err)
	}
}

// halftone
func TestEncodeToFile2(t *testing.T) {
	halftonePath := "_halftone0.jpg"
	//halftonePath := "_halftone1.jpg"

	err := EncodeToFile("https://github.com/richelieu-yang/chimera", "_test2.png",
		standard.WithHalftone(halftonePath),
		standard.WithQRWidth(18),
	)
	if err != nil {
		panic(err)
	}
}
