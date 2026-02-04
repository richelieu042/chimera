package goccyYamlKit

import (
	"context"

	"github.com/goccy/go-yaml"
)

// Marshal 序列化（结构体实例 => yaml文本）
func Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func MarshalWithContextAndOptions(ctx context.Context, v interface{}, opts ...yaml.EncodeOption) ([]byte, error) {
	return yaml.MarshalContext(ctx, v, opts...)
}
