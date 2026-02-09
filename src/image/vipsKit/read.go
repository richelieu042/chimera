package vipsKit

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// Read
/*
@param params 可以为nil
*/
func Read(path string, params *vips.ImportParams) (*vips.ImageRef, error) {
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return nil, err
	}

	ref, err := vips.LoadImageFromFile(path, params)
	if err != nil {
		return nil, err
	}

	// Rotate the picture upright and reset EXIF orientation tag
	err = ref.AutoRotate()
	if err != nil {
		return nil, err
	}

	return ref, nil
}
