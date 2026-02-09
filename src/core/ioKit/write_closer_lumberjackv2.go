package ioKit

//import (
//	"github.com/richelieu042/chimera/v3/src/file/fileKit"
//	"gopkg.in/natefinch/lumberjack.v2"
//	"io"
//)
//
//type (
//	lumberjackOptions struct {
//		maxSize    int
//		maxAge     int
//		maxBackups int
//		localTime  bool
//		compress   bool
//		console    bool
//	}
//
//	LumberjackOption func(opts *lumberjackOptions)
//)
//
//func WithMaxSize(maxSize int) LumberjackOption {
//	return func(opts *lumberjackOptions) {
//		opts.maxSize = maxSize
//	}
//}
//
//func WithMaxAge(maxAge int) LumberjackOption {
//	return func(opts *lumberjackOptions) {
//		opts.maxAge = maxAge
//	}
//}
//
//func WithMaxBackups(maxBackups int) LumberjackOption {
//	return func(opts *lumberjackOptions) {
//		opts.maxBackups = maxBackups
//	}
//}
//
//func WithLocalTime(localTime bool) LumberjackOption {
//	return func(opts *lumberjackOptions) {
//		opts.localTime = localTime
//	}
//}
//
//func WithCompress(compress bool) LumberjackOption {
//	return func(opts *lumberjackOptions) {
//		opts.compress = compress
//	}
//}
//
//// WithConsole Deprecated: 暂不使用
//func WithConsole(console bool) LumberjackOption {
//	return func(opts *lumberjackOptions) {
//		opts.console = console
//	}
//}
//
//func loadOptions(options ...LumberjackOption) *lumberjackOptions {
//	opts := &lumberjackOptions{}
//	for _, option := range options {
//		option(opts)
//	}
//	return opts
//}
//
//// NewLumberjackWriteCloser 文件分割（根据文件大小）
///*
//参考: https://github.com/natefinch/lumberjack
//
//PS:
//(1) 传参可以参考下 NewLumberjackWriteCloser1 的注释；
//(2) 仅配置 filePath 的情况: 	(a) 超过100MB就rotate一下；
//						  	(b) 保留所有文件；
//						  	(c) 不压缩；
//						  	(d) 默认使用UTC时间；
//(3) 文件不存在，会自动创建；文件存在，内容会追加在最后.
//
//@param filePath 文件路径
//@param options 	可选配置
//*/
//func NewLumberjackWriteCloser(filePath string, options ...LumberjackOption) (*lumberjack.Logger /*io.WriteCloser*/, error) {
//	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
//		return nil, err
//	}
//	if err := fileKit.MkParentDirs(filePath); err != nil {
//		return nil, err
//	}
//
//	opts := loadOptions(options...)
//	var writeCloser = &lumberjack.Logger{
//		Filename:   filePath,
//		MaxSize:    opts.maxSize,
//		MaxBackups: opts.maxBackups,
//		MaxAge:     opts.maxAge,
//		LocalTime:  opts.localTime,
//		Compress:   opts.compress,
//	}
//	//if opts.console {
//	//	writeCloser = MultiWriteCloser(writeCloser, NopCloserToWriter(os.Stdout))
//	//}
//	return writeCloser, nil
//}
//
//// NewLumberjackWriteCloser1
///*
//Deprecated: Use NewLumberjackWriteCloser instead. 唯一作用: 看如何传参.
//
//PS:
//(1) lumberjack: 日志切割组件.一般情况下，lumberjack配合其他日志库，实现日志的滚动(rolling)记录.
//(2) 根据 maxBackups、maxAge 删除过期文件，根据 compress 决定是否压缩哪些未压缩的旧日志文件。
//
//Golang 语言三方库 lumberjack 日志切割组件怎么使用？
//	https://mp.weixin.qq.com/s/gGnovwzS1ucW3Afxcytp_Q
//go语言的日志滚动(rolling)记录器——lumberjack
//	https://zhuanlan.zhihu.com/p/430224518
//
//@param filePath		日志路径，归档日志也会保存在对应目录下（若该值为空，则日志会保存到os.TempDir()目录下，日志文件名为<processname>-lumberjack.log）
//					(1) 会尝试创建父级目录；
//					(2) 文件不存在，会自动创建；
//					(3) 文件存在，
//						(3.1) 是文件，内容追加在最后；
//						(3.2) 是目录，返回error.
//@param maxSize		[单位: MB，默认: 100MB] 日志大小到达此值，就开始backup
//@param maxBackups	旧日志保存的最大数量，默认保存所有旧日志文件
//@param maxAge		[单位: days] 旧日志保存的最大天数，默认保存所有旧日志文件
//@param localTime	[默认使用UTC时间] 是否使用本地时间戳？
//@param compress		[默认: false] 对backup的日志是否进行压缩？（压缩实际上是打成压缩包，文件名最后加上".gz"）
//
//e.g. rotate && compress == false
//	"aaa.log"（最新的日志文件；当前的指向）
//	"aaa-2023-05-10T08-51-01.320.log"
//	"aaa-2023-05-10T08-50-51.010.log"
//
//e.g.1 rotate && compress == true
//	"aaa.log"（最新的日志文件；当前的指向）
//	"aaa-2023-05-10T08-55-08.504.log.gz"
//	"aaa-2023-05-10T08-54-57.786.log.gz"
//*/
//func NewLumberjackWriteCloser1(filePath string, maxSize, maxBackups, maxAge int, localTime, compress bool) (io.WriteCloser, error) {
//	if err := fileKit.MkParentDirs(filePath); err != nil {
//		return nil, err
//	}
//	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
//		return nil, err
//	}
//
//	return &lumberjack.Logger{
//		Filename:   filePath,
//		MaxSize:    maxSize,
//		MaxBackups: maxBackups,
//		MaxAge:     maxAge,
//		LocalTime:  localTime,
//		Compress:   compress,
//	}, nil
//}
