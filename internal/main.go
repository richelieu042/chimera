package main

import (
	"fmt"

	"github.com/richelieu-yang/chimera/v3/src/image/imageKit"
)

func main() {
	////err := gfile.RemoveFile("/Users/richelieu/Desktop/111.png")
	//err := gfile.RemoveAll("/Users/richelieu/Desktop/222.png")
	//
	//fmt.Println(err)

	srcPath := "/Users/richelieu/Desktop/iShot_2026-01-07_15.29.11.PNG"

	err := imageKit.Resize(srcPath, "222.png", 1000, 100)
	fmt.Println(err)
}
