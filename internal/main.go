package main

import (
	"os"

	"github.com/richelieu-yang/chimera/v3/src/cmd/cmdKit"
)

func main() {
	out, err := cmdKit.Run(nil,
		"adb", "-s", "127.0.0.1:5555",
		"exec-out", "screencap", "-p",
	)
	if err != nil {
		panic(err)
	}

	//cmd := exec.Command(
	//	"adb", "-s", "127.0.0.1:5555",
	//	"exec-out", "screencap", "-p",
	//)
	//out, err := cmd.Output()
	//if err != nil {
	//	panic(err)
	//}

	err = os.WriteFile("screen.png", out, 0644)
	if err != nil {
		panic(err)
	}
}
