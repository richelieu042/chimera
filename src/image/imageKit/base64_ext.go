package imageKit

import (
	"encoding/base64"
	"strings"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

// GetBase64ImageExt 从 base64 字符串中识别图片格式并返回文件后缀
/*
	@return 文件扩展名（小写 && 带"."），e.g. ".jpg" ".png"
*/
func GetBase64ImageExt(base64Str string) (string, error) {
	if err := strKit.AssertNotBlank(base64Str, "base64Str"); err != nil {
		return "", err
	}

	// 去除可能的空格和换行
	base64Str = strings.TrimSpace(base64Str)

	// 处理 data URI 格式
	if strings.HasPrefix(base64Str, "data:") {
		return getExtensionFromDataURI(base64Str)
	}

	// 直接解析 base64 数据
	return getExtensionFromBase64Data(base64Str)
}

// getExtensionFromDataURI 从 data URI 中提取扩展名
func getExtensionFromDataURI(dataURI string) (string, error) {
	// 格式: data:image/jpeg;base64,/9j/4AAQ...
	parts := strings.SplitN(dataURI, ",", 2)
	if len(parts) != 2 {
		return "", errorKit.Newf("invalid dataURI")
	}

	// 提取 MIME 类型
	header := parts[0]
	if strings.Contains(header, "image/") {
		mimeType := strings.TrimPrefix(header, "data:")
		mimeType = strings.Split(mimeType, ";")[0]

		// MIME 类型到扩展名的映射
		mimeToExt := map[string]string{
			"image/jpeg":    ".jpg",
			"image/jpg":     ".jpg",
			"image/png":     ".png",
			"image/gif":     ".gif",
			"image/webp":    ".webp",
			"image/bmp":     ".bmp",
			"image/svg+xml": ".svg",
			"image/tiff":    ".tiff",
			"image/x-icon":  ".ico",
		}

		if ext, ok := mimeToExt[mimeType]; ok {
			return ext, nil
		}
	}

	// 如果 MIME 类型不明确，尝试检测实际数据
	return getExtensionFromBase64Data(parts[1])
}

// getExtensionFromBase64Data 通过检测文件头魔数来识别格式
func getExtensionFromBase64Data(base64Data string) (string, error) {
	// 解码 base64
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", errorKit.Wrapf(err, "fail to decode base64")
	}

	if len(decoded) < 4 {
		return "", errorKit.Newf("数据太短，无法识别格式")
	}

	// 检测文件魔数（文件头特征字节）
	magicNumbers := []struct {
		signature []byte
		ext       string
	}{
		{[]byte{0xFF, 0xD8, 0xFF}, ".jpg"},        // JPEG
		{[]byte{0x89, 0x50, 0x4E, 0x47}, ".png"},  // PNG
		{[]byte{0x47, 0x49, 0x46, 0x38}, ".gif"},  // GIF
		{[]byte{0x52, 0x49, 0x46, 0x46}, ".webp"}, // WEBP (需要进一步检查)
		{[]byte{0x42, 0x4D}, ".bmp"},              // BMP
		{[]byte{0x49, 0x49, 0x2A, 0x00}, ".tiff"}, // TIFF (little-endian)
		{[]byte{0x4D, 0x4D, 0x00, 0x2A}, ".tiff"}, // TIFF (big-endian)
		{[]byte{0x00, 0x00, 0x01, 0x00}, ".ico"},  // ICO
	}

	for _, magic := range magicNumbers {
		if len(decoded) >= len(magic.signature) {
			match := true
			for i, b := range magic.signature {
				if decoded[i] != b {
					match = false
					break
				}
			}
			if match {
				// WEBP 需要额外验证
				if magic.ext == ".webp" && len(decoded) >= 12 {
					if string(decoded[8:12]) == "WEBP" {
						return ".webp", nil
					}
					continue
				}
				return magic.ext, nil
			}
		}
	}

	return "", errorKit.Newf("无法识别的图片格式")
}
