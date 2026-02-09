package pdfKit

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

var (
	// Split 推荐使用 SplitFile，是对本函数的封装.
	Split func(rs io.ReadSeeker, outDir, fileName string, span int, conf *model.Configuration) error = api.Split

	// SplitByPageNr 推荐使用 SplitByPageNrFile，是对本函数的封装.
	SplitByPageNr func(rs io.ReadSeeker, outDir, fileName string, pageNrs []int, conf *model.Configuration) error = api.SplitByPageNr
)

// SplitFile 拆分pdf文件（可以指定生成pdf文件的页数）.
/*
@param outDir 	拆分后的文件保存目录
@param span		每几页拆分为一个pdf文件？应该>=1
@param conf		可以为nil
*/
func SplitFile(inFile, outDir string, span int, conf *model.Configuration) error {
	/* inFile */
	if err := fileKit.AssertExistAndIsFile(inFile); err != nil {
		return err
	}
	/* outDir */
	if err := fileKit.AssertNotExistOrIsDir(outDir); err != nil {
		return err
	}
	if err := fileKit.MkDirs(outDir); err != nil {
		return err
	}

	return api.SplitFile(inFile, outDir, span, conf)
}

// SplitByPageNrFile
/*
@param pageNrs	拆分的页数s（其中元素应该>=2）
@param conf		可以为nil

e.g. 34页的pdf文件，传参pageNrs为[]int{2}，结果: 拆分为2个pdf文件，第一个pdf文件有 1 页，第二个pdf文件有 33页
	inputPDF := "/Users/richelieu/Desktop/34页(1).pdf"
	outputDir := "_tmp"
	if err := pdfKit.SplitByPageNrFile(inputPDF, outputDir, []int{2}, nil); err != nil {
		panic(err)
	}
*/
func SplitByPageNrFile(inFile, outDir string, pageNrs []int, conf *model.Configuration) error {
	/* inFile */
	if err := fileKit.AssertExistAndIsFile(inFile); err != nil {
		return err
	}
	/* outDir */
	if err := fileKit.AssertNotExistOrIsDir(outDir); err != nil {
		return err
	}
	if err := fileKit.MkDirs(outDir); err != nil {
		return err
	}
	/* pageNrs */
	pageNrs = sliceKit.Uniq(pageNrs)
	if err := sliceKit.AssertNotEmpty(pageNrs, "pageNrs"); err != nil {
		return err
	}

	return api.SplitByPageNrFile(inFile, outDir, pageNrs, conf)
}
