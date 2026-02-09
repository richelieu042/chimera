package pdfKit

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// ExtractImagesFile 导出pdf文件中的图片.
/*
!!!: pdfcpu 库仅支持提取原始嵌入的图像（如果图像经过压缩或编码，则可能无法提取）.

@param outDir			输出目录
@param selectedPages	可以为nil，即该pdf文件的所有页
						e.g. 第1页和第2页
							[]string{"1", "2"}
@param conf				可以为nil
*/
func ExtractImagesFile(inFile, outDir string, selectedPages []string, conf *model.Configuration) error {
	if err := fileKit.MkDirs(outDir); err != nil {
		return err
	}

	return api.ExtractImagesFile(inFile, outDir, selectedPages, conf)
}
