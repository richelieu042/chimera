package ioKit

//import (
//	"github.com/richelieu042/chimera/v3/src/core/errorKit"
//	"github.com/richelieu042/chimera/v3/src/core/strKit"
//	"github.com/zeromicro/go-zero/core/logx"
//	"io"
//)
//
//const (
//	defaultDelimiter = "-"
//)
//
//// NewDailyRotateRuleWriteCloser 调用了go-zero中的方法.
///*
//Deprecated: 调用完 Write() 后立即调用 Close() || 进程退出，Write的内容可能不会立即保存到文件中，导致丢失.
//
//@param filePath	会自动创建父级目录；
//				文件不存在: 会创建；
//				文件已存在: (1)是个文件，会append；(2)是个目录，会返回error(e.g. open /Users/richelieu/Downloads: is a directory)
//*/
//func NewDailyRotateRuleWriteCloser(filePath, delimiter string, days int, compress bool) (io.WriteCloser, error) {
//	if days <= 0 {
//		return nil, errorKit.Sign("invalid days(%d)", days)
//	}
//	delimiter = strKit.EmptyToDefault(delimiter, defaultDelimiter)
//
//	rule := logx.DefaultRotateRule(
//		filePath,
//		delimiter,
//		days,
//		compress,
//	)
//	return logx.NewLogger(filePath, rule, compress)
//}
//
//// NewSizeLimitRotateRuleWriteCloser 调用了go-zero中的方法.
///*
//Deprecated: 调用完 Write() 后立即调用 Close() || 进程退出，Write的内容可能不会立即保存到文件中，导致丢失.
//
//PS: 最多生成文件的数量: maxBackups + 1(filePath)
//
//@param filePath		会自动创建父级目录；
//					文件不存在: 会创建；
//					文件已存在: (1)是个文件，会append；(2)是个目录，会返回error(e.g. open /Users/richelieu/Downloads: is a directory)
//@param delimiter	分隔符
//@param maxSize 		单位: MB
//@param maxBackups	备份数量的上限
//@param compress		如果设置为true，compress前后各会输出一条信息（logx.Infof、logx.Errorf）
//
//e.g. Mac环境
//	("test.log", "-", 1, 10, 3, true)
//	4个文件:
//		test.log
//		test-2023-03-10T16/50/11+08/00.log.gz
//		test-2023-03-10T16/50/10+08/00.log.gz
//		test-2023-03-10T16/50/09+08/00.log.gz
//*/
//func NewSizeLimitRotateRuleWriteCloser(filePath, delimiter string, days, maxSize, maxBackups int, compress bool) (io.WriteCloser, error) {
//	if days <= 0 {
//		return nil, errorKit.Sign("invalid days(%d)", days)
//	}
//	if maxSize <= 0 {
//		return nil, errorKit.Sign("invalid maxSize(%d)", maxSize)
//	}
//	if maxBackups <= 0 {
//		return nil, errorKit.Sign("invalid maxBackups(%d)", maxBackups)
//	}
//	delimiter = strKit.EmptyToDefault(delimiter, defaultDelimiter)
//
//	rule := logx.NewSizeLimitRotateRule(
//		filePath,
//		delimiter,
//		days,
//		maxSize,
//		maxBackups,
//		compress,
//	)
//	return logx.NewLogger(filePath, rule, compress)
//}
