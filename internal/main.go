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

	err := imageKit.Resize("/Users/richelieu/Desktop/iShot_2026-01-07_15.27.36.png", "222.png", 100, 100)
	fmt.Println(err)
}
