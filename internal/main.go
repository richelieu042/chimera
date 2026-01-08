package main

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gfile"
)

func main() {
	//err := gfile.RemoveFile("/Users/richelieu/Desktop/111.png")
	err := gfile.RemoveAll("/Users/richelieu/Desktop/222.png")

	fmt.Println(err)
}
