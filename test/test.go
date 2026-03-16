package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	var mat gocv.Mat

	fmt.Println("-")
	if err := mat.Close(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("-")
}
