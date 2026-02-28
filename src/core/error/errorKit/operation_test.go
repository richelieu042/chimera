package errorKit

import (
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestIs(t *testing.T) {
	err := redis.Nil
	err1 := Wrapf(err, "1")
	err2 := Wrapf(err1, "2")

	fmt.Printf("%+v\n", err2)

	fmt.Println(Is(err2, err)) // true
	fmt.Println(Is(err1, err)) // true
	fmt.Println(Is(err, err))  // true

	fmt.Println(Is(err2, err1)) // true
	fmt.Println(Is(err1, err2)) // false
}

func TestAs(t *testing.T) {
	var err error = &net.DNSError{
		IsTimeout: true,
	}

	var c net.Error
	if errors.As(err, &c) {
		fmt.Println(c.Timeout())
	}
	fmt.Println("===")
	/*
		output:
		true
		===
	*/
}

type myError struct {
	Text string
}

func (err myError) Error() string {
	return err.Text
}

// TestAs1 receiver为"值类型"
func TestAs1(t *testing.T) {
	err := myError{
		Text: "cyy",
	}
	err1 := Wrapf(err, "1")

	target := &myError{}
	if ok := As(err1, target); !ok {
		panic("ok == false")
	}
	fmt.Println(target.Text) // cyy
	if err.Text != target.Text {
		panic("not equal")
	}
}

type myError1 struct {
	Text string
}

func (err *myError1) Error() string {
	return err.Text
}

// TestAs2 receiver为"指针类型"
func TestAs2(t *testing.T) {
	err := &myError1{
		Text: "cyy",
	}
	err1 := Wrapf(err, "1")

	target := &myError1{}
	// 相较于 TestAs，多了个 &
	if ok := As(err1, &target); !ok {
		panic("ok == false")
	}
	fmt.Println(target.Text) // cyy
	if err.Text != target.Text {
		panic("not equal")
	}
}
