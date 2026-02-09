package rocksdbKit

import (
	"github.com/linxGnu/grocksdb"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// OpenDB
/*
@param opts 可以为nil（将采用默认值）
*/
func OpenDB(dirPath string, opts *grocksdb.Options) (*grocksdb.DB, error) {
	if err := fileKit.AssertNotExistOrIsDir(dirPath); err != nil {
		return nil, err
	}
	if err := fileKit.MkDirs(dirPath); err != nil {
		return nil, err
	}

	if opts == nil {
		opts = NewDefaultDBOptions()
	}

	return grocksdb.OpenDb(opts, dirPath)
}
