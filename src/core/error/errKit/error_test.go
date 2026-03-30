package errKit

import (
	goerrors "errors"
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

func TestIs(t *testing.T) {
	var ErrNotFound = goerrors.New("not found")
	var findUser = func(id int) error {
		return fmt.Errorf("findUser: %w", ErrNotFound) // 用 %w 包装错误
	}

	err := findUser(42)
	if Is(err, ErrNotFound) {
		fmt.Println("用户不存在") // ✅ 即使被包装，也能匹配到
	}
}

func TestIs_1(t *testing.T) {
	var ErrNotFound = goerrors.New("not found")
	var findUser = func(id int) error {
		return Wrap(ErrNotFound, "findUser") // 用 Wrap 包装
	}

	err := findUser(42)
	if Is(err, ErrNotFound) {
		fmt.Println("用户不存在") // ✅ 即使被包装，也能匹配到
	}
}

type validationError struct {
	Field   string
	Message string
}

func (e *validationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

func TestAs(t *testing.T) {
	var err error = &validationError{
		Field:   "name",
		Message: "cannot be empty",
	}
	err = fmt.Errorf("validate: %w", err) // case 1: 用 %w 包装错误
	//err = Wrap(err, "wrap") // case 2: 用 Wrap 包装错误

	var ve *validationError

	//if As(err, &ve) {
	if As1(err, &ve) {
		fmt.Println("字段:", ve.Field)   // 字段: name
		fmt.Println("原因:", ve.Message) // 原因: cannot be empty
	}
}
