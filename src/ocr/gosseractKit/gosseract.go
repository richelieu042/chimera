package gosseractKit

import (
	"github.com/otiai10/gosseract/v2"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// GertText
/*
!!!: 使用此函数，必须确保"CGO_ENABLED=1"，否则go run或go build会报错: undefined: gosseract.NewClient
*/
func GertText(imgPath string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(imgPath); err != nil {
		return "", err
	}

	client := gosseract.NewClient()
	defer client.Close()

	// 设置语言 (支持中英文)
	if err := client.SetLanguage("chi_sim", "eng"); err != nil {
		return "", err
	}

	// 设置PSM模式
	// gosseract.PSM_AUTO: 自动检测布局
	if err := client.SetPageSegMode(gosseract.PSM_AUTO); err != nil {
		return "", err
	}

	if err := client.SetImage(imgPath); err != nil {
		return "", err
	}
	return client.Text()
}
