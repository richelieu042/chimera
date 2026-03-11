package httpClientKit

import (
	"net/http"
	"time"
)

const (
	// DefaultTimeout 发送请求的默认超时时间.
	DefaultTimeout = time.Second * 10
)

type (
	options struct {
		// timeout 请求的超时时间
		timeout time.Duration
		// safe 默认false（跳过ssl证书验证，即使url 以https开头 && 证书非法，请求也能成功）
		safe bool

		// queryParams 适用于: POST、GET
		queryParams map[string][]string
		// postParams 适用于: POST
		postParams map[string][]string
	}

	Option func(opts *options)
)

func (opts *options) newHttpClient() *http.Client {
	return NewHttpClient(opts.timeout, !opts.safe)
}

func loadOptions(s ...Option) *options {
	opts := &options{}

	for _, option := range s {
		option(opts)
	}

	if opts.timeout <= 0 {
		opts.timeout = DefaultTimeout
	}

	return opts
}

// WithTimeout
/*
适用于: GET、POST
*/
func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.timeout = timeout
	}
}

// WithSafe
/*
适用于: GET、POST
*/
func WithSafe(safe bool) Option {
	return func(opts *options) {
		opts.safe = safe
	}
}

// WithQueryParams
/*
适用于: GET、POST
*/
func WithQueryParams(queryParams map[string][]string) Option {
	return func(opts *options) {
		opts.queryParams = queryParams
	}
}

// WithPostParams
/*
适用于: POST
*/
func WithPostParams(postParams map[string][]string) Option {
	return func(opts *options) {
		opts.postParams = postParams
	}
}
