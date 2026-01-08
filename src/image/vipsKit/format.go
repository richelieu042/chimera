package vipsKit

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

var defaultPerm os.FileMode = 0644

func ToTiff(src, dest string, exportParams *vips.TiffExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	if exportParams == nil {
		exportParams = vips.NewTiffExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportTiff(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, defaultPerm)
}

func ToGif(src, dest string, exportParams *vips.GifExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	if exportParams == nil {
		exportParams = vips.NewGifExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportGIF(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, defaultPerm)
}

func ToAvif(src, dest string, exportParams *vips.AvifExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	if exportParams == nil {
		exportParams = vips.NewAvifExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportAvif(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, defaultPerm)
}

// ToHeif
/*
PS:
(0) 必需先set up;
(1) iphone使用 heif格式 存储照片（占用内存比JPEG格式少）;
(2) 常见的HEIF封装类型文件有 .heic 和 .avif。
*/
func ToHeif(src, dest string, exportParams *vips.HeifExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	// 默认值
	if exportParams == nil {
		exportParams = vips.NewHeifExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportHeif(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, defaultPerm)
}
