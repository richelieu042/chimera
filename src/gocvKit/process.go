package gocvKit

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

// Resize 将 src 缩放到指定的 (width x height)。
/*
	@param width, height	若 width 或 height <= 0，则按原始宽高比自动推算
	@param interp			对应 OpenCV 的插值算法枚举，常用 gocv.InterpolationDefault（通用默认，速度与质量均衡）
*/
func Resize(src gocv.Mat, width, height int, interp gocv.InterpolationFlags) (gocv.Mat, error) {
	if src.Empty() {
		return gocv.NewMat(), fmt.Errorf("resize: source mat is empty")
	}

	origW := src.Cols()
	origH := src.Rows()
	w, h := calcSize(origW, origH, width, height)
	if w <= 0 || h <= 0 {
		return gocv.NewMat(), fmt.Errorf("resize: invalid target size (%dx%d)", w, h)
	}

	dst := gocv.NewMat()
	err := gocv.Resize(src, &dst, image.Pt(w, h), 0, 0, interp)
	return dst, err
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
