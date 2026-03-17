package imageKit

import "image"

var (
	Pt func(X, Y int) image.Point = image.Pt

	Rect func(x0, y0, x1, y1 int) image.Rectangle = image.Rect
)
