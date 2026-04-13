package gzipKit

import (
	"io"

	"github.com/gogf/gf/v2/encoding/gcompress"
)

//import "github.com/zeromicro/go-zero/core/codec"
//
//var (
//	// Compress 压缩.
//	Compress func(bs []byte) []byte = codec.Gzip
//
//	// Decompress 解压缩.
//	/*
//	   PS: 大小限制: 100MB.
//	*/
//	Decompress func(bs []byte) ([]byte, error) = codec.Gunzip
//)

// Compress
/*
PS: 前端的解压，可以参考 notes/_Vue3Projects/ws-client 中的 GzipKit.
*/
func Compress(data []byte, options ...GzipOption) ([]byte, error) {
	opts := loadOptions(options...)

	return opts.Compress(data)
}

var (
	//Compress func(data []byte, level ...int) ([]byte, error) = gcompress.Gzip

	Decompress func(data []byte) ([]byte, error) = gcompress.UnGzip
)

var (
	GzipFile func(srcFilePath, dstFilePath string, level ...int) (err error) = gcompress.GzipFile

	GzipPathWriter func(filePath string, writer io.Writer, level ...int) error = gcompress.GzipPathWriter
)

var (
	UnGzipFile func(srcFilePath, dstFilePath string) error = gcompress.UnGzipFile
)
