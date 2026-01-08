package qrCodeKit

import (
	"os"

	"github.com/tuotoo/qrcode"
)

// Decode 解析二维码（仅支持部分图片）.
/*
PS:
(1) 能解析 yeqown/go-qrcode 生成的二维码.
(2) 不能解析 skip2/go-qrcode 生成的二维码.
*/
func Decode(path string) (content string, err error) {
	// 打开二维码图片文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 解析二维码
	qrCode, err := qrcode.Decode(file)
	if err != nil {
		return "", err
	}

	return qrCode.Content, nil
}
