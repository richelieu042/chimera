package excelKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/xuri/excelize/v2"
)

// NewFile 新建一个空白的Excel文件.
/*
PS:
(1) 返回的 *excelize.File实例，不再使用时应当调用 File.Close();
(2) 新建的空白的Excel文件，其内只有1个空白的工作表，表名为"Sheet1".
*/
var NewFile func(opts ...excelize.Options) *excelize.File = excelize.NewFile

// NewFileWithPath 新建一个空白的Excel文件.
/*
PS:
(1) 返回的 *excelize.File实例，不再使用时应当调用 File.Close();
(2) 新建的空白的Excel文件，其内只有1个空白的工作表，表名为"Sheet1".

@param path 文件的路径（如果文件已经存在，会覆盖它）
*/
func NewFileWithPath(filePath string, opts ...excelize.Options) (*excelize.File, error) {
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

	f := NewFile()
	if err := f.SaveAs(filePath, opts...); err != nil {
		_ = f.Close()
		return nil, errorKit.Wrapf(err, "fail to save as")
	}
	return f, nil
}
