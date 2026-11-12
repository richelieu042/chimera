package goccyYamlKit

import (
	"context"

	"github.com/goccy/go-yaml"
)

var (
	// Marshal
	/*
		PS:
		(1) 支持 yaml tag、json tag;
		(2) 同一字段同时有 yaml 和 json tag 时，优先使用非空的 yaml tag.
	*/
	Marshal func(v any) ([]byte, error) = yaml.Marshal

	MarshalContext func(ctx context.Context, v any, opts ...yaml.EncodeOption) ([]byte, error) = yaml.MarshalContext
)
