package main

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/android/adbKit"
	"github.com/richelieu042/chimera/v3/src/log/zapKit"
)

func main() {
	//c, err := adbKit.NewClient("127.0.0.1:5555", true, zapKit.NewSimpleConsoleLogger())
	c, err := adbKit.NewClient("192.168.60.205:16384", true, zapKit.NewSimpleConsoleLogger())
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
