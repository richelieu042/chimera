package gocvKit

import "image"

var (
	NewPoint func(X, Y int) image.Point = image.Pt

	NewRectangle func(x0, y0, x1, y1 int) image.Rectangle = image.Rect
)
