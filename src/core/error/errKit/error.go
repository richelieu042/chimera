package errKit

import (
	"context"
	goerrors "errors"

	"github.com/cockroachdb/errors"
)

var (
	New           func(msg string) error                                    = errors.New
	Newf          func(format string, args ...interface{}) error            = errors.Newf
	NewWithDepth  func(depth int, msg string) error                         = errors.NewWithDepth
	NewWithDepthf func(depth int, format string, args ...interface{}) error = errors.NewWithDepthf

	Is func(err, target error) bool             = goerrors.Is
	As func(err error, target interface{}) bool = errors.As

	Wrap           func(err error, msg string) error                                    = errors.Wrap
	Wrapf          func(err error, format string, args ...interface{}) error            = errors.Wrapf
	WrapWithDepth  func(depth int, err error, msg string) error                         = errors.WrapWithDepth
	WrapWithDepthf func(depth int, err error, format string, args ...interface{}) error = errors.WrapWithDepthf

	Unwrap     func(err error) error = errors.Unwrap
	UnwrapOnce func(err error) error = errors.UnwrapOnce
	UnwrapAll  func(err error) error = errors.UnwrapAll

	EncodeError func(ctx context.Context, err error) errors.EncodedError = errors.EncodeError
	DecodeError func(ctx context.Context, enc errors.EncodedError) error = errors.DecodeError
)
