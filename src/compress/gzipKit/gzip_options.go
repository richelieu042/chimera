package gzipKit

import (
	"compress/gzip"

	"github.com/gogf/gf/v2/encoding/gcompress"
)

type (
	gzipOptions struct {
		// level 压缩级别
		level int

		// compressThreshold 压缩阈值
		/*
			单位: byte
			<= 0: ALL压缩
		*/
		compressThreshold int
	}

	GzipOption func(opts *gzipOptions)
)

func loadOptions(options ...GzipOption) *gzipOptions {
	opts := &gzipOptions{
		level:             gzip.BestSpeed,
		compressThreshold: -1, /* 都压缩 */
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

func WithLevel(level int) GzipOption {
	return func(opts *gzipOptions) {
		opts.level = level
	}
}

// WithCompressThreshold 设置: 最小压缩长度阈值
func WithCompressThreshold(compressThreshold int) GzipOption {
	return func(opts *gzipOptions) {
		opts.compressThreshold = compressThreshold
	}
}

func (opts *gzipOptions) Compress(data []byte) ([]byte, error) {
	if err := AssertValidLevel(opts.level); err != nil {
		return nil, err
	}

	if len(data) < opts.compressThreshold {
		// (1) 不压缩
		return data, nil
	}
	// (2) 压缩
	return gcompress.Gzip(data, opts.level)
}
