package goccyYamlKit

import (
	"context"

	"github.com/goccy/go-yaml"
)

var (
	Marshal func(v any) ([]byte, error) = yaml.Marshal

	MarshalContext func(ctx context.Context, v any, opts ...yaml.EncodeOption) ([]byte, error) = yaml.MarshalContext
)
