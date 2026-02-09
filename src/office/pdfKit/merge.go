package pdfKit

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

var (
	// Merge
	/*
		Deprecated: 使用 MergeCreateFile 或 MergeAppendFile，它们实际上是对此函数的封装.
	*/
	Merge func(destFile string, inFiles []string, w io.Writer, conf *model.Configuration, dividerPage bool) error = api.Merge
)

// MergeCreateFile 合并pdf（如果outFile已经存在且是个文件，会"覆盖"内容）.
func MergeCreateFile(inFiles []string, outFile string, dividerPage bool, conf *model.Configuration) error {
	if err := fileKit.MkParentDirs(outFile); err != nil {
		return err
	}

	return api.MergeCreateFile(inFiles, outFile, dividerPage, conf)
}

// MergeAppendFile 合并pdf（如果outFile已经存在且是个文件，会"在最后追加"内容）.
/*
	@param inFiles		要合并的pdf文件（复数）
	@param outFile		合并后的pdf文件
	@param dividerPage	true: 在每个pdf文件之间插入一个空白页
	@param conf 		可以为nil，将使用默认值
*/
func MergeAppendFile(inFiles []string, outFile string, dividerPage bool, conf *model.Configuration) error {
	if err := fileKit.MkParentDirs(outFile); err != nil {
		return err
	}

	return api.MergeAppendFile(inFiles, outFile, dividerPage, conf)
}
