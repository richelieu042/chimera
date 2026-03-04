package main

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/ocr/gosseractKit"
)

func main() {
	fmt.Println(gosseractKit.GertText("clipped.png", "chi_sim"))
	fmt.Println(gosseractKit.GertText("ccc.jpg", "chi_sim"))
}
