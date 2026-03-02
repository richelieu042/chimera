package errKit

import (
	"context"
	goerrors "errors"

	"github.com/cockroachdb/errors"
)

var (
	// Simple 创建一个简单的错误，不包含堆栈信息.
	Simple func(text string) error = goerrors.New

	New           func(msg string) error                                    = errors.New
	Newf          func(format string, args ...interface{}) error            = errors.Newf
	NewWithDepth  func(depth int, msg string) error                         = errors.NewWithDepth
	NewWithDepthf func(depth int, format string, args ...interface{}) error = errors.NewWithDepthf

	Wrap           func(err error, msg string) error                                    = errors.Wrap
	Wrapf          func(err error, format string, args ...interface{}) error            = errors.Wrapf
	WrapWithDepth  func(depth int, err error, msg string) error                         = errors.WrapWithDepth
	WrapWithDepthf func(depth int, err error, format string, args ...interface{}) error = errors.WrapWithDepthf

	As  func(err error, target interface{}) bool = errors.As
	As1 func(err error, target any) bool         = goerrors.As

	// Is 判断错误是否匹配某个目标值.
	/*
		（1）用于检查错误链中是否包含某个特定的错误值（用 == 语义比较）。
		（2）即使被包装，也能匹配到
	*/
	Is func(err, target error) bool = goerrors.Is

	Unwrap     func(err error) error = errors.Unwrap
	UnwrapOnce func(err error) error = errors.UnwrapOnce
	UnwrapAll  func(err error) error = errors.UnwrapAll

	EncodeError func(ctx context.Context, err error) errors.EncodedError = errors.EncodeError
	DecodeError func(ctx context.Context, enc errors.EncodedError) error = errors.DecodeError
)
