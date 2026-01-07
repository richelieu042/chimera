package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/richelieu-yang/chimera/v3/src/ocr/gosseractKit"
	"gocv.io/x/gocv"
)

func convertToBinary(img image.Image, threshold uint8) *image.Gray {
	bounds := img.Bounds()
	binaryImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor).(color.Gray)

			// 二值化: 高于阈值为白色，低于为黑色
			if grayColor.Y > threshold {
				binaryImg.Set(x, y, color.Gray{255})
			} else {
				binaryImg.Set(x, y, color.Gray{0})
			}
		}
	}

	return binaryImg
}

func preprocessForOCR(inputPath, outputPath string) {
	img := gocv.IMRead(inputPath, gocv.IMReadColor)
	defer img.Close()

	// 1. 转灰度
	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	// 2. 去噪
	denoised := gocv.NewMat()
	defer denoised.Close()
	gocv.GaussianBlur(gray, &denoised, image.Pt(3, 3), 0, 0, gocv.BorderDefault)

	// 3. 二值化
	binary := gocv.NewMat()
	defer binary.Close()
	gocv.Threshold(denoised, &binary, 0, 255, gocv.ThresholdBinary|gocv.ThresholdOtsu)

	// 4. 形态学操作(可选，去除小噪点)
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(2, 2))
	defer kernel.Close()
	gocv.MorphologyEx(binary, &binary, gocv.MorphClose, kernel)

	gocv.IMWrite(outputPath, binary)
}

func main() {
	//file, _ := os.Open("/Users/richelieu/GolandProjects/chimera/screen_帆.png")
	//defer file.Close()
	//
	//img, _, _ := image.Decode(file)
	//
	//// 阈值128，可根据图片调整
	//binaryImg := convertToBinary(img, 128)
	//
	//outFile, _ := os.Create("output_binary.png")
	//defer outFile.Close()
	//
	//png.Encode(outFile, binaryImg)

	//preprocessForOCR("/Users/richelieu/GolandProjects/chimera/screen_帆.png", "output_binary.png")
	//preprocessForOCR("/Users/richelieu/GolandProjects/chimera/screen_帆.png", "output_binary.png")

	fmt.Println(gosseractKit.GertText("/Users/richelieu/GolandProjects/chimera/screen_帆.png"))
}
