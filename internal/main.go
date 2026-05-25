package main

import "github.com/richelieu042/chimera/v3/src/image/imageKit"

func main() {
	p0 := imageKit.Pt(1415, 504)
	p1 := imageKit.Pt(1487, 577)

	if err := imageKit.ClipWithPath("/Users/richelieu/Desktop/screenshot.png", "gift1.png", p0.X, p0.Y, p1.X-p0.X, p1.Y-p0.Y); err != nil {
		panic(err)
	}
}
