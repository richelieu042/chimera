package errorKit

import (
	"fmt"
	"testing"
)

func TestSimplef(t *testing.T) {
	err := Simplef("111")
	fmt.Println(err.Error())

	fmt.Println("---")

	fmt.Printf("%v\n", err)

	fmt.Println("---")

	fmt.Printf("%+v\n", err)
}
