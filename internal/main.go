package main

import (
	"image"
	"image/jpeg"
	"os"
)

func cropImage(inputPath, outputPath string, x, y, width, height int) error {
	// 打开原图
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// 定义裁剪区域
	cropRect := image.Rect(x, y, x+width, y+height)

	// 裁剪图片
	croppedImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(cropRect)

	// 保存裁剪后的图片
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 根据输出格式编码
	return jpeg.Encode(outFile, croppedImg, &jpeg.Options{Quality: 90})
	// 或者使用 png.Encode(outFile, croppedImg)
}

func main() {
	// 从 (100, 100) 位置开始，裁剪 300x200 的区域
	err := cropImage("input.jpg", "output.jpg", 100, 100, 300, 200)
	if err != nil {
		panic(err)
	}
}
