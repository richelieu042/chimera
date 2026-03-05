package gosseractKit

import (
	"github.com/otiai10/gosseract/v2"
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// newClient
/*
	!!!: err == nil 的情况下，调用完此函数后，应该立即跟上: defer client.Close()
*/
func newClient(languages ...string) (client *gosseract.Client, err error) {
	if len(languages) == 0 {
		languages = []string{"eng"}
	}

	client = gosseract.NewClient()
	defer func() {
		if err != nil {
			client.Close()
		}
	}()

	// （1）设置语言
	if err := client.SetLanguage(languages...); err != nil {
		return client, errKit.Wrap(err, "fail to set language")
	}

	// （2）设置PSM（页面分割模式）
	// gosseract.PSM_AUTO: 自动检测布局
	if err := client.SetPageSegMode(gosseract.PSM_AUTO); err != nil {
		return client, errKit.Wrap(err, "fail to set page seg mode")
	}

	return
}

// GertText
/*
!!!: 使用此函数，必须确保"CGO_ENABLED=1"，否则go run或go build会报错: undefined: gosseract.NewClient

@param languages 	支持指定语言，默认英文（eng）
					e.g.
						eng — 英文
						chi_sim — 简体中文
						chi_tra — 繁体中文
						jpn — 日文
					e.g.1
						如果图片中只有中文和数字的话，建议使用 ("chi_sim") 而非 ("chi_sim", "eng")，否则可能出问题（e.g.把识 "天" 别为 "%"）.
*/
func GertText(imgPath string, languages ...string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(imgPath); err != nil {
		return "", err
	}

	client, err := newClient(languages...)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if err := client.SetImage(imgPath); err != nil {
		return "", errKit.Wrap(err, "fail to set image")
	}
	return client.Text()
}

// GertTextFromBytes 从 "图片的二进制数据" 中获取文字.
/*
@param bytes 图片的二进制数据
*/
func GertTextFromBytes(bytes []byte, languages ...string) (string, error) {
	if err := sliceKit.AssertNotEmpty(bytes, "bytes"); err != nil {
		return "", err
	}

	client, err := newClient(languages...)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if err := client.SetImageFromBytes(bytes); err != nil {
		return "", errKit.Wrap(err, "fail to set image")
	}
	return client.Text()
}
