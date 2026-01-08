package main

import (
	"fmt"

	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

func main() {
	//////err := gfile.RemoveFile("/Users/richelieu/Desktop/111.png")
	////err := gfile.RemoveAll("/Users/richelieu/Desktop/222.png")
	////
	////fmt.Println(err)
	//
	//srcPath := "/Users/richelieu/Desktop/iShot_2026-01-08_14.07.26.png"
	//
	////err := imageKit.Resize(srcPath, "222.png", 1000, 100)
	//
	////err := imageKit.ResizeWithScale(srcPath, "333.png", 2)
	//
	////err := imageKit.ResizeKeepAspectRatio(srcPath, "444.png", 1000, 1000)
	//
	//err := imageKit.ResizeByHeight(srcPath, "555.png", 1000)
	//
	//fmt.Println(err)

	fmt.Println(fileKit.GetExt("	111.exe "))

}
