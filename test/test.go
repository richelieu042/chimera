package main

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

func main() {

	fmt.Println(strKit.Replace("abcdcba", "a", "0", 1))  // "0bcdcba"
	fmt.Println(strKit.Replace("abcdcba", "a", "0", 2))  // "0bcdcb0"
	fmt.Println(strKit.Replace("abcdcba", "a", "0", -1)) // "0bcdcb0"

}
