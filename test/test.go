package main

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/image/imageKit"
	"github.com/richelieu042/chimera/v3/src/ocr/gosseractKit"
)

func main() {
	fmt.Println(imageKit.ToGrayscaleWithPath("clipped.png", "1.png"))

	fmt.Println(gosseractKit.GertText("clipped.png", "chi_sim", "eng"))
	fmt.Println(gosseractKit.GertText("1.png", "chi_sim", "eng"))
}
