package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"adb", "-s", "127.0.0.1:5555",
		"exec-out", "screencap", "-p",
	)

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("screen.png", out, 0644)
	if err != nil {
		panic(err)
	}
}
