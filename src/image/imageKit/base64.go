package imageKit

import (
	"fmt"
	"os"
	"regexp"

	"github.com/richelieu-yang/chimera/v3/src/component/web/http_client/httpClientKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

// GetImageBase64 (硬盘上的)图片 => base64字符串
/*
参考: golang 将图片生成Base64 https://blog.csdn.net/weixin_40292098/article/details/126029489
*/
func GetImageBase64(imagePath string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(imagePath); err != nil {
		return "", err
	}

	imageData, err := fileKit.ReadFile(imagePath)
	if err != nil {
		return "", err
	}
	return EncodeToBase64String(imageData)
}

// GetWebImageBase64 网络图片 => base64字符串
/*
参考: golang 将图片生成Base64 https://blog.csdn.net/weixin_40292098/article/details/126029489

@param url e.g."https://img.redocn.com/sheying/20150507/pugongying_4267498.jpg"
*/
func GetWebImageBase64(url string) (string, error) {
	// 先获取网络图片的内容
	_, imageData, err := httpClientKit.Get(url)
	if err != nil {
		return "", err
	}

	return EncodeToBase64String(imageData)
}

func EncodeToBase64String(data []byte) (string, error) {
	mimeType := fileKit.DetectContentType(data)
	switch mimeType {
	case "image/jpeg":
		fallthrough
	case "image/png":
		base64Str := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Kit.EncodeToString(data))
		return base64Str, nil
	default:
		return "", errorKit.Newf("mimeType(%s) isn't supported currently", mimeType)
	}
}

// DecodeFromBase64 图片的base64数据 => 图片的数据（可以直接存储到硬盘上）
/*
@param base64 带不带前缀都无所谓(e.g. "data:image/png;base64,"、"data:image/jpeg;base64,"、"data:image/gif;base64,")
*/
func DecodeFromBase64(base64 []byte) ([]byte, error) {
	// 如果有前缀的话，去掉前缀
	re := regexp.MustCompile(`^data:(.+);base64,`)
	tmp := re.Find(base64)
	length := len(tmp)
	if length > 0 {
		base64 = base64[length:]
	}

	return base64Kit.Decode(base64)
}

// DecodeToImageFile
/*
@param target 要生成的图片的路径
*/
func DecodeToImageFile(base64 []byte, dest string, perm os.FileMode) error {
	data, err := DecodeFromBase64(base64)
	if err != nil {
		return err
	}
	return fileKit.WriteToFile(dest, data)
}
