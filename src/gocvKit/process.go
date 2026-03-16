package gocvKit

import (
	"gocv.io/x/gocv"
)

// ToGrayscale 转灰度.
/*
!!!: 如果返回的 err == nil，使用完后需要：手动关闭返回的 mat（建议在调用的下一行加上defer close）.
*/
func ToGrayscale(img gocv.Mat) (dst gocv.Mat, err error) {
	dst = gocv.NewMat()
	defer func() {
		if err != nil {
			dst.Close()
		}
	}()

	err = gocv.CvtColor(img, &dst, gocv.ColorBGRToGray)
	return
}
