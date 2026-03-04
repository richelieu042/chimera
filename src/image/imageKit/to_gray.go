package imageKit

import (
	"image"
	"image/color"
)

// ToGrayscale 灰度处理：image.Image => 灰度图
/*
	使用标准 Rec.601 亮度公式：Y = 0.299R + 0.587G + 0.114B
*/
func ToGrayscale(src image.Image) image.Image {
	bounds := src.Bounds()
	gray := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray.Set(x, y, color.GrayModel.Convert(src.At(x, y)))
		}
	}
	return gray
}

// ToGrayscaleWithPath 灰度处理：图片 => 灰度图
func ToGrayscaleWithPath(input, output string) error {
	inputImg, _, err := DecodeFromPath(input)
	if err != nil {
		return err
	}

	outputImg := ToGrayscale(inputImg)

	return EncodeToPath(output, outputImg)
}
