package main

import "time"

func main() {
	start := time.Now()
	time.Sleep(time.Second * 3)

	elapsed := time.Since(start)
	println(elapsed.String())
}
