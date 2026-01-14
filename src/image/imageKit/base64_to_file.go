package imageKit

import (
	"encoding/base64"
	"strings"

	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

// Base64ToFile 将base64字符串转换为图片文件
// 参数:
//   - base64Str: base64编码的图片字符串，可能包含"data:"前缀
//   - outputPath: 输出文件路径
//
// 返回:
//   - error: 错误信息，成功则返回nil
func Base64ToFile(base64Str, outputPath string) error {
	// 移除可能存在的data:前缀
	// Data URL格式示例: data:image/png;base64,iVBORw0KGgo...
	// 我们只需要逗号后面的纯base64数据部分
	if strings.HasPrefix(base64Str, "data:") {
		// 查找逗号分隔符的位置
		commaIndex := strings.Index(base64Str, ",")
		if commaIndex == -1 {
			return errorKit.Newf("invalid base64 format: missing comma separator")
		}
		// 截取逗号之后的base64数据
		base64Str = base64Str[commaIndex+1:]
	}

	// 将base64字符串解码为原始字节数据
	imageData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return errorKit.Wrapf(err, "fail to decode base64")
	}

	// 将解码后的图片数据写入文件
	// 0644权限: 所有者可读写(6)，组用户可读(4)，其他用户可读(4)
	err = fileKit.WriteToFile(outputPath, imageData, 0644)
	if err != nil {
		return errorKit.Wrapf(err, "fail to write file")
	}

	return nil
}
