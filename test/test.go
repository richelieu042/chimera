package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	A()
}

func A() {
	panic("666")
}
