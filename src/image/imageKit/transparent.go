package imageKit

import (
	"image"
	"image/png"

	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

// TrimTransparentForPng 裁剪透明边缘（适用的图片格式：.png）
/*
@param inputPath	输入图片的路径，要求后缀为.png
@param outputPath	输出图片的路径，为空则覆盖原文件
*/
func TrimTransparentForPng(inputPath, outputPath string) error {
	inputPath = strKit.TrimSpace(inputPath)
	outputPath = strKit.TrimSpace(outputPath)

	if err := fileKit.AssertExistAndIsFile(inputPath); err != nil {
		return err
	}
	if strKit.IsEmpty(outputPath) {
		// 如果输出路径为空，则默认覆盖原文件
		outputPath = inputPath
	}

	// 读取原始图片
	file, err := fileKit.OpenReadOnly(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return err
	}

	// 查找非透明区域的边界
	bounds := findBounds(img)

	// 创建裁剪后的图片
	outputImg := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			outputImg.Set(x-bounds.Min.X, y-bounds.Min.Y, img.At(x, y))
		}
	}

	// 保存裁剪后的图片
	outputFile, err := fileKit.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return png.Encode(outputFile, outputImg)
}

// findBounds 查找非透明像素的边界
func findBounds(img image.Image) image.Rectangle {
	bounds := img.Bounds()
	minX, minY := bounds.Max.X, bounds.Max.Y
	maxX, maxY := 0, 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			// alpha > 0 表示非完全透明
			if a > 0 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// 如果整张图都是透明的，返回原始边界
	if minX > maxX || minY > maxY {
		return bounds
	}

	return image.Rect(minX, minY, maxX+1, maxY+1)
}
