package main

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/ocr/gosseractKit"
)

func main() {
	path := "/Users/richelieu/GolandProjects/chimera/screen.png"

	text, err := gosseractKit.GertText(path)
	if err != nil {
		panic(err)
	}
	console.Info(text)
}
