package qrCodeKit

import (
	"image/color"
	"testing"

	"github.com/skip2/go-qrcode"
)

func TestGenerateFile(t *testing.T) {
	content := "https://example.org"

	err := GenerateFile(content, qrcode.Medium, 256, "_test-qr0.png")
	if err != nil {
		panic(err)
	}
}

func TestGenerateFileWithColor(t *testing.T) {
	content := "https://example.org"

	err := GenerateFileWithColor(content, qrcode.Medium, 256, color.Transparent, color.Black, "_test-qr1.png")
	if err != nil {
		panic(err)
	}
}

func TestGenerateFileWithBackgroundImage(t *testing.T) {
	content := "https://example.org"
	bgPath := "/Users/richelieu/Desktop/iShot_2024-09-05_08.54.28.png"

	level := qrcode.High

	{
		err := GenerateFileWithBackgroundImage(content, level, -1, bgPath, color.Black, "_test-qr2.png")
		if err != nil {
			panic(err)
		}
	}

	{
		err := GenerateFileWithBackgroundImage(content, level, -1, bgPath, color.Black, "_test-qr2.jpg")
		if err != nil {
			panic(err)
		}
	}
}
