package errKit

import (
	"context"
	goerrors "errors"

	"github.com/cockroachdb/errors"
)

var (
	// Simple 创建一个简单的错误（不包含堆栈信息）.
	Simple func(text string) error = goerrors.New

	New           func(msg string) error                                    = errors.New
	Newf          func(format string, args ...interface{}) error            = errors.Newf
	NewWithDepth  func(depth int, msg string) error                         = errors.NewWithDepth
	NewWithDepthf func(depth int, format string, args ...interface{}) error = errors.NewWithDepthf

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

var (
	// Is 判断错误是否匹配某个目标值.
	/*
		（1）用于检查错误链中是否包含某个特定的错误值（用 == 语义比较）
		（2）支持错误链（即使被包装，也能匹配到）
	*/
	Is func(err, target error) bool = goerrors.Is

	// As 从错误链中提取某个类型的错误.
	/*
		用于检查错误链中是否存在某个特定类型的错误，并将其提取出来。
	*/
	As func(err error, target interface{}) bool = errors.As

	// As1
	/*
		Deprecated: 推荐使用 As，原因：（分布式系统 / RPC 场景）cockroachdb 有跨节点错误类型恢复能力，是其核心优势
	*/
	As1 func(err error, target any) bool = goerrors.As
)
