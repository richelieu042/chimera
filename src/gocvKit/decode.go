package gocvKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"gocv.io/x/gocv"
)

// Decode 图片（二进制数据） => gocv.Mat
func Decode(imgData []byte) (mat gocv.Mat, err error) {
	if err = sliceKit.AssertNotEmpty(imgData, "imgData"); err != nil {
		return
	}

	defer func() {
		if err != nil {
			mat.Close()
		}
	}()

	mat, err = gocv.IMDecode(imgData, gocv.IMReadColor)
	if err != nil {
		return
	}
	if mat.Empty() {
		err = errKit.New("mat is empty")
		return
	}
	return mat, nil
}

// DecodeFromPath 图片（路径） => gocv.Mat
/*
@param path 图片路径
*/
func DecodeFromPath(path string) (mat gocv.Mat, err error) {
	if err = fileKit.AssertExistAndIsFile(path); err != nil {
		return
	}

	defer func() {
		if err != nil {
			mat.Close()
		}
	}()

	mat = gocv.IMRead(path, gocv.IMReadColor)
	if mat.Empty() {
		err = errKit.New("mat is empty")
		return
	}
	return mat, nil
}
