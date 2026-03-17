package main

import (
	"log"
	"os"
)

func main() {
	//p0 := image.Pt(1414, 503)
	//p1 := image.Pt(1490, 580)
	//
	//if err := imageKit.ClipWithPath("/Users/richelieu/Desktop/gift.png", "c.png", p0.X, p0.Y, p1.X-p0.X, p1.Y-p0.Y); err != nil {
	//	panic(err)
	//}

	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	l := log.New(os.Stdout, "", flag)

	l.Println("hello world")
}
