package errKit

import (
	"context"
	goerrors "errors"
	"fmt"

	"github.com/cockroachdb/errors"
)

var (
	// Simple 创建一个简单的错误（不包含堆栈信息）.
	Simple func(text string) error = goerrors.New

	New          func(msg string) error            = errors.New
	NewWithDepth func(depth int, msg string) error = errors.NewWithDepth

	Wrap          func(err error, msg string) error            = errors.Wrap
	WrapWithDepth func(depth int, err error, msg string) error = errors.WrapWithDepth

	Unwrap     func(err error) error = errors.Unwrap
	UnwrapOnce func(err error) error = errors.UnwrapOnce
	UnwrapAll  func(err error) error = errors.UnwrapAll

	EncodeError func(ctx context.Context, err error) errors.EncodedError = errors.EncodeError
	DecodeError func(ctx context.Context, enc errors.EncodedError) error = errors.DecodeError
)

func Simplef(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func Newf(format string, args ...interface{}) error {
	return errors.NewWithDepthf(1, format, args...)
}

func NewfWithDepth(depth int, format string, args ...interface{}) error {
	return errors.NewWithDepthf(depth+1, format, args...)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.WrapWithDepthf(1, err, format, args...)
}

func WrapfWithDepth(depth int, err error, format string, args ...interface{}) error {
	return errors.WrapWithDepthf(depth+1, err, format, args...)
}
