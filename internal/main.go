package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "Physical size: 1920x1080."

	re := regexp.MustCompile(`(\d+)x(\d+)$`)
	matches := re.FindStringSubmatch(input)
	width := matches[1]  // "1920"
	height := matches[2] // "1080"

	fmt.Println(width, height)
}
