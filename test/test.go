package main

import (
	"fmt"
	"os"

	"github.com/richelieu042/chimera/v3/src/log/zapKit"
)

func main() {
	w := os.Stdout
	ws := zapKit.NewLockedWriteSyncer(w)
	fmt.Println(ws == os.Stdout)

	{
		w := os.Stderr
		ws := zapKit.NewLockedWriteSyncer(w)
		fmt.Println(ws == os.Stderr)
	}
}
