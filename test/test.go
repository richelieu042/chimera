package main

import (
	"image"

	"github.com/richelieu042/chimera/v3/src/image/imageKit"
)

func main() {
	p0 := image.Pt(1778, 946)
	p1 := image.Pt(1882, 1050)

	if err := imageKit.ClipWithPath("/Users/richelieu/Desktop/收起.png", "folded.png", p0.X, p0.Y, p1.X-p0.X, p1.Y-p0.Y); err != nil {
		panic(err)
	}

	//flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	//l := log.New(os.Stdout, "", flag)
	//
	//l.Println("hello world")
}
