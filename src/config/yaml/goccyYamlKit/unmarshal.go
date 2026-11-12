package goccyYamlKit

import (
	"context"

	"github.com/goccy/go-yaml"
)

var (
	// Unmarshal
	/*
		PS:
		(1) 支持 yaml tag、json tag;
		(2) 同一字段同时有 yaml 和 json tag 时，优先使用非空的 yaml tag.
	*/
	Unmarshal func(data []byte, v any) error = yaml.Unmarshal

	UnmarshalContext func(ctx context.Context, data []byte, v any, opts ...yaml.DecodeOption) error = yaml.UnmarshalContext
)
