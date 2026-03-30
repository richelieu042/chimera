package errKit

import (
	goerrors "errors"

	"github.com/cockroachdb/errors"
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
