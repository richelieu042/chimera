package excelKit

import (
	"io"

	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/xuri/excelize/v2"
)

// OpenFile 打开本地文件.
/*
PS: 可能会返回 excelize.ErrWorkbookFileFormat，原因: 文件格式不支持（比如.xls文件）.
*/
func OpenFile(filePath string, opts ...excelize.Options) (*excelize.File, error) {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}

	return excelize.OpenFile(filePath, opts...)
}

// OpenReader 打开数据流.
/*
PS: 可能会返回 excelize.ErrWorkbookFileFormat，原因: 文件格式不支持（比如.xls文件）.

@params r 数据流（包括: 远程文件）
*/
var OpenReader func(r io.Reader, opts ...excelize.Options) (*excelize.File, error) = excelize.OpenReader
