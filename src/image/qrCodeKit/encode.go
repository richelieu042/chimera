package qrCodeKit

import (
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

// EncodeToFile 生成二维码图片.
/*
@param content 二维码的内容
@param options 可以指定:	(1) logo（放在二维码的中心）
						(2) halftone（比较酷炫），可参考 https://github.com/yeqown/go-qrcode 中: example/with-halftone/README.md
						(3) 尺寸（但单位并非像素）
						(4) 透明
						(5) 前景色、背景色
*/
func EncodeToFile(content, outputImgPath string, options ...standard.ImageOption) error {
	/* content */
	if err := strKit.AssertNotEmpty(content, "content"); err != nil {
		return err
	}
	/* outputImgPath */
	if err := fileKit.AssertNotExistOrIsFile(outputImgPath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(outputImgPath); err != nil {
		return err
	}

	qrCode, err := qrcode.New(content)
	if err != nil {
		return err
	}
	// 设置二维码的输出文件名
	w, err := standard.New(outputImgPath, options...)
	if err != nil {
		return err
	}
	// 将二维码写入文件
	return qrCode.Save(w)
}
