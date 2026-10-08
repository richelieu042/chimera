package interfaceKit

import (
	"fmt"
	"testing"
)

func TestIsNil(t *testing.T) {
	var obj any
	fmt.Println(IsNil(obj)) // true

	type bean struct {
	}
	var b *bean = nil
	fmt.Println(IsNil(b)) // true

	var obj1 any = b
	fmt.Println(IsNil(obj1)) // true
}
