package main

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/atomic/atomicKit"
)

func main() {
	i := atomicKit.NewInt32(3)

	i.Dec()

	fmt.Println(i.Load())
}
