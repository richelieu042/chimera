package main

import (
	"fmt"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

type validationError struct {
	Field   string
	Message string
}

func (e *validationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

func main() {
	var err error = &validationError{
		Field:   "name",
		Message: "cannot be empty",
	}
	//err = fmt.Errorf("validate: %w", err) // case 1: 用 %w 包装错误
	err = errKit.Wrap(err, "wrap") // case 2: 用 Wrap 包装错误

	var ve *validationError
	//if errKit.As(err, &ve) {
	if errKit.As1(err, &ve) {
		fmt.Println("字段:", ve.Field)   // 字段: name
		fmt.Println("原因:", ve.Message) // 原因: cannot be empty
	}
}
