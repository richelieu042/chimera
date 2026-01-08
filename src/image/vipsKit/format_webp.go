package vipsKit

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

// ToWebpData
/*
@param exportParams 可以为nil
*/
func ToWebpData(src string, exportParams *vips.WebpExportParams) ([]byte, error) {
	if exportParams == nil {
		exportParams = vips.NewWebpExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return nil, err
	}

	data, _, err := imageRef.ExportWebp(exportParams)
	return data, err
}

// ToWebp
/*
@param exportParams 可以为nil
*/
func ToWebp(src, dest string, exportParams *vips.WebpExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	imageData, err := ToWebpData(src, exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, defaultPerm)
}
