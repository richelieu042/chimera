package gocvKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"gocv.io/x/gocv"
)

// Encode gocv.Mat => 图片（二进制数据）
/*
@param format	".jpg" | ".png" | ".bmp" 等
*/
func Encode(mat gocv.Mat, format string) ([]byte, error) {
	if mat.Empty() {
		return nil, errKit.New("mat is empty")
	}

	buf, err := gocv.IMEncode(gocv.FileExt(format), mat)
	if err != nil {
		return nil, errKit.Wrap(err, "encode failed")
	}
	defer buf.Close()

	// copy 是必要的，因为 buf.Close() 后 GetBytes() 返回的底层内存会被释放，直接返回引用会导致数据异常。
	// 未定义行为（UB）的典型特征——不是每次都会崩溃，而是"可能"出问题
	data := make([]byte, len(buf.GetBytes()))
	copy(data, buf.GetBytes())

	return data, nil
}

// EncodeToPath gocv.Mat => 图片（文件路径）
func EncodeToPath(mat gocv.Mat, path string) error {
	if err := fileKit.MkParentDirs(path); err != nil {
		return err
	}

	// 直接保存
	ok := gocv.IMWrite(path, mat)
	if !ok {
		return errKit.New("fail to write")
	}
	return nil
}
