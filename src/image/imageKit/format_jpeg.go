package imageKit

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// ToJpeg 将图片格式转换为".jpg"(||".jpeg").
/*
@param qualityArgs （默认: 100; 取值范围: [1, 100]）生成jpeg图片的质量
*/
func ToJpeg(src, dest string, qualityArgs ...int8) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	var quality int8 = 100
	if len(qualityArgs) > 0 {
		quality = qualityArgs[0]
	}

	srcImage, _, err := DecodeFromPath(src)
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	//return jpeg.Encode(destFile, srcImage, &jpeg.Options{Quality: 100})

	// 先画一层白色的底色，再画src对应图片的内容，以避免: 将透明背景的PNG转换为JPG（或JPEG），默认背景色为黑色 https://www.zongscan.com/demo333/95729.html
	destImage := image.NewRGBA(srcImage.Bounds())
	draw.Draw(destImage, destImage.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(destImage, destImage.Bounds(), srcImage, srcImage.Bounds().Min, draw.Over)
	return jpeg.Encode(destFile, destImage, &jpeg.Options{
		Quality: int(quality),
	})
}
