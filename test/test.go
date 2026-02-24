package main

import (
	"fmt"
	"time"

	"github.com/richelieu042/chimera/v3/src/time/timeKit"
)

func main() {
	now := time.Now()
	fmt.Println(timeKit.Format(now, "2006-01-02T15.04.05.000"))
	fmt.Println(timeKit.Format(now, "2006-01-02T15.04"))
	fmt.Println(timeKit.Format(now, "05.000"))

	//fmt.Println(strKit.Replace("abcdcba", "a", "0", 1))  // "0bcdcba"
	//fmt.Println(strKit.Replace("abcdcba", "a", "0", 2))  // "0bcdcb0"
	//fmt.Println(strKit.Replace("abcdcba", "a", "0", -1)) // "0bcdcb0"

}
