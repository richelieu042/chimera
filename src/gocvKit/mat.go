package gocvKit

import "gocv.io/x/gocv"

// ToGrayscale 转灰度.
/*
!!!: 不管返回的 error 是否为nil，都要先手动关闭 mat.
*/
func ToGrayscale(img gocv.Mat) (gocv.Mat, error) {
	dst := gocv.NewMat()

	if err := gocv.CvtColor(img, &dst, gocv.ColorBGRToGray); err != nil {
		return dst, err
	}
	return dst, nil
}
