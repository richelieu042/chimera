package gocvKit

import (
	"image"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
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

// Resize 将 src 缩放到指定的 (width x height)。
/*
!!!: 如果返回的 err == nil，使用完后需要：手动关闭返回的 mat（建议在调用的下一行加上defer close）.

@param width, height	若 width 或 height <= 0，则按原始宽高比自动推算
@param interp			对应 OpenCV 的插值算法枚举，常用 gocv.InterpolationDefault（通用默认，速度与质量均衡）
*/
func Resize(src gocv.Mat, width, height int, interp gocv.InterpolationFlags) (dst gocv.Mat, err error) {
	if src.Empty() {
		err = errKit.New("source mat is empty")
		return
	}

	origW := src.Cols()
	origH := src.Rows()
	w, h := calcSize(origW, origH, width, height)
	if w <= 0 || h <= 0 {
		err = errKit.Newf("invalid target size(%dx%d)", w, h)
		return
	}

	dst = gocv.NewMat()
	defer func() {
		if err != nil {
			dst.Close()
		}
	}()
	err = gocv.Resize(src, &dst, image.Pt(w, h), 0, 0, interp)
	return
}

// calcSize 根据目标宽高（任意一方可为 0）及原始尺寸，自动推算等比例后的实际宽高。
func calcSize(origW, origH, targetW, targetH int) (int, int) {
	switch {
	case targetW > 0 && targetH > 0:
		return targetW, targetH
	case targetW > 0:
		// 按宽度等比
		ratio := float64(targetW) / float64(origW)
		return targetW, int(float64(origH)*ratio + 0.5)
	case targetH > 0:
		// 按高度等比
		ratio := float64(targetH) / float64(origH)
		return int(float64(origW)*ratio + 0.5), targetH
	default:
		return origW, origH
	}
}
