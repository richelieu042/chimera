package main

import (
	"fmt"
	"time"

	"github.com/richelieu042/chimera/v3/src/time/timeKit"
)

func main() {
	t, err := timeKit.ParseInLocation(timeKit.FormatCommon, "2025-01-22T15:04:05.000", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
