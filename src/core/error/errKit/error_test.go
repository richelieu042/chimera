package errKit

import (
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestNew(t *testing.T) {
	err := New("ccc %d")

	fmt.Printf("err: %v\n", err)
	fmt.Println("------")
	fmt.Printf("err: %+v\n", err)
}

func TestNewf(t *testing.T) {
	err := Newf("hello %s", "world")

	fmt.Printf("err: %v\n", err)
	fmt.Println("------")
	fmt.Printf("err: %+v\n", err)
}

func TestWrap(t *testing.T) {
	err := Wrap(redis.Nil, "ccc")

	fmt.Printf("err: %v\n", err)
	fmt.Println("------")
	fmt.Printf("err: %+v\n", err)
}
