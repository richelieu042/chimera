package main

import (
	"image"

	"github.com/richelieu042/chimera/v3/src/image/imageKit"
)

func main() {
	path := "/Users/richelieu/Desktop/screenshot 2.png"
	point0 := &image.Point{
		X: 1657,
		Y: 65,
	}
	point1 := &image.Point{
		X: 1754,
		Y: 163,
	}

	err := imageKit.ClipWithPath(path, "ccc.png", point0.X, point0.Y, point1.X-point0.X, point1.Y-point0.Y)
	if err != nil {
		panic(err)
	}
}
