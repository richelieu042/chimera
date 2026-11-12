package goccyYamlKit

import (
	"context"

	"github.com/goccy/go-yaml"
)

var (
	Unmarshal func(data []byte, v any) error = yaml.Unmarshal

	UnmarshalContext func(ctx context.Context, data []byte, v any, opts ...yaml.DecodeOption) error = yaml.UnmarshalContext
)
