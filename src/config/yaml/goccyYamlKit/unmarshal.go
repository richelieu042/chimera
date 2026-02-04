package goccyYamlKit

import (
	"context"

	"github.com/goccy/go-yaml"
)

// Unmarshal 反序列化（yaml文本 => 结构体实例）
func Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func UnmarshalWithContextAndOptions(ctx context.Context, data []byte, v interface{}, opts ...yaml.DecodeOption) error {
	return yaml.UnmarshalContext(ctx, data, v, opts...)
}
