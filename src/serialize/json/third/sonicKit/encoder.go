package sonicKit

import (
	"io"

	"github.com/bytedance/sonic"
)

// NewEncoder 编码器（to json）
func NewEncoder(api sonic.API, writer io.Writer) sonic.Encoder {
	return api.NewEncoder(writer)
}
