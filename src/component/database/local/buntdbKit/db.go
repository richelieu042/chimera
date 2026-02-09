package buntdbKit

import (
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/tidwall/buntdb"
)

const (
	MemoryPath = ":memory:"
)

// Open
/*
@param 存储数据的文件路径	(1) 会自动创建父目录
						(2) It will be created if it doesn't exist.
*/
func Open(path string) (*buntdb.DB, error) {
	if path == MemoryPath {
		return OpenInMemory()
	}

	if err := fileKit.AssertNotExistOrIsFile(path); err != nil {
		return nil, err
	}
	if err := fileKit.MkParentDirs(path); err != nil {
		return nil, err
	}

	return buntdb.Open(path)
}

// OpenInMemory 不会持久化到磁盘上.
/*
PS: 读写更快.
*/
func OpenInMemory() (*buntdb.DB, error) {
	return buntdb.Open(MemoryPath)
}
