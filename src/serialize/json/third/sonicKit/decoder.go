package sonicKit

import (
	"io"

	"github.com/bytedance/sonic"
)

// NewDecoder 解码器（from json）
func NewDecoder(api sonic.API, reader io.Reader) sonic.Decoder {
	return api.NewDecoder(reader)
}
