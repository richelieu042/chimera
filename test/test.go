package main

import (
	"image"

	"github.com/richelieu042/chimera/v3/src/image/imageKit"
)

func main() {
	p0 := image.Pt(1414, 503)
	p1 := image.Pt(1490, 580)

	if err := imageKit.ClipWithPath("/Users/richelieu/Desktop/gift.png", "c.png", p0.X, p0.Y, p1.X-p0.X, p1.Y-p0.Y); err != nil {
		panic(err)
	}
}
