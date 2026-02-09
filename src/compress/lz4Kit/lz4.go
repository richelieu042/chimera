package lz4Kit

import (
	"bytes"
	"io"
	"os"

	"github.com/pierrec/lz4/v4"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

var (
	NewWriter func(w io.Writer) *lz4.Writer = lz4.NewWriter

	NewReader func(r io.Reader) *lz4.Reader = lz4.NewReader
)

func Compress(data []byte) (compressedData []byte, err error) {
	buf := bytes.NewBuffer(nil)

	lz4Writer := lz4.NewWriter(buf)
	_, err = lz4Writer.Write(data)
	if err != nil {
		return
	}
	// !!!: 需要先 关闭Writer 再调用 buf.Bytes().
	if err = lz4Writer.Close(); err != nil {
		return nil, err
	}

	compressedData = buf.Bytes()
	return
}

func Decompress(compressedData []byte) (decompressData []byte, err error) {
	lz4Reader := lz4.NewReader(bytes.NewReader(compressedData))
	decompressData, err = io.ReadAll(lz4Reader)
	return
}

// CompressFile
/*
@param dest (1) 建议文件后缀为 ".lz4"
			(2) 如果已经存在 && 是个文件，会覆盖.
*/
func CompressFile(src, dest string) error {
	if err := fileKit.AssertExistAndIsFile(src); err != nil {
		return err
	}

	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	// 读取原始文件内容
	inputData, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// 使用lz4.NewWriter创建压缩器，写入到一个新的文件
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	lz4Writer := lz4.NewWriter(destFile)
	defer lz4Writer.Close()

	// 写入压缩数据
	_, err = lz4Writer.Write(inputData)
	return err
}
