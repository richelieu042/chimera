package gocvKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"gocv.io/x/gocv"
)

func Decode(imgData []byte) (gocv.Mat, error) {
	if err := sliceKit.AssertNotEmpty(imgData, "imgData"); err != nil {
		return gocv.NewMat(), err
	}

	return gocv.IMDecode(imgData, gocv.IMReadColor)
}

// DecodeFromPath
/*
@param path 图片路径
*/
func DecodeFromPath(path string) (gocv.Mat, error) {
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return gocv.NewMat(), err
	}

	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		return gocv.NewMat(), errKit.Simple("read: mat is empty")
	}
	return img, nil
}
