package reqKit

import (
	"time"

	"github.com/imroc/req/v3"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/log/zapKit"
)

const (
	// DefaultTimeout 默认超时时间
	DefaultTimeout = time.Second * 30
)

type (
	clientOptions struct {
		Dev     bool
		BaseURL string
		// Timeout 超时时间（发送请求的整个周期，includes connection time, any redirects, and reading the response body）
		Timeout time.Duration
		// InsecureSkipVerify
		/*
			true:  跳过证书验证
			false: 不跳过证书验证（默认; 更加安全）
		*/
		InsecureSkipVerify bool

		// Logger 日志输出
		/*
			nil: 不输出
		*/
		Logger req.Logger

		/* 下面两者一般搭配起来用，参考: https://req.cool/zh/docs/prologue/quickstart/#%E6%9B%B4%E9%AB%98%E7%BA%A7%E7%9A%84-get-%E8%AF%B7%E6%B1%82 */
		CommonErrorResult interface{}
		OnAfterResponse   req.ResponseMiddleware
	}

	ClientOption func(*clientOptions)
)

func loadOptions(options ...ClientOption) *clientOptions {
	logger := zapKit.NewLogger(nil).Sugar()

	opts := &clientOptions{
		Dev:                false,
		BaseURL:            "",
		Timeout:            0, // 默认值在下面
		InsecureSkipVerify: true,
		// imroc/req默认: 输出到 os.Stdout
		Logger: logger,

		CommonErrorResult: nil,
		OnAfterResponse:   nil, // 默认值在下面
	}

	for _, option := range options {
		option(opts)
	}

	/* 默认值s */
	if opts.Timeout <= 0 {
		opts.Timeout = DefaultTimeout
	}
	if opts.OnAfterResponse == nil {
		opts.OnAfterResponse = func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil { // There is an underlying error, e.g. network error or unmarshal error.
				return nil
			}

			//if errMsg, ok := resp.ErrorResult().(*ErrorMessage); ok {
			//	resp.Err = errMsg // Convert api error into go error
			//	return nil
			//}

			/* 处理不成功的http状态码 */
			if !resp.IsSuccessState() {
				// Neither a success response nor a error response, record details to help troubleshooting
				//resp.Err = fmt.Errorf("bad status: %s\nraw content:\n%s", resp.Status, resp.Dump())
				bodyStr, err := resp.ToString()
				if err != nil {
					resp.Err = errorKit.Newf("bad status: %s, fail to get body string: %s", resp.Status, err.Error())
				} else {
					resp.Err = errorKit.Newf("bad status: %s, body string: %s", resp.Status, bodyStr)
				}
			}
			return nil
		}
	}
	return opts
}

func WithDev() ClientOption {
	return func(options *clientOptions) {
		options.Dev = true
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(options *clientOptions) {
		options.Timeout = timeout
	}
}

func WithInsecureSkipVerify(insecureSkipVerify bool) ClientOption {
	return func(options *clientOptions) {
		options.InsecureSkipVerify = insecureSkipVerify
	}
}

// WithLogger
/*
@param logger	(1) 可以是 *zap.SugaredLogger 实例 || *logrus.Logger 实例
				(2) 可以为nil（disable log, 禁用输出）
*/
func WithLogger(logger req.Logger) ClientOption {
	return func(options *clientOptions) {
		options.Logger = logger
	}
}

func WithCommonErrorResult(result interface{}) ClientOption {
	return func(options *clientOptions) {
		options.CommonErrorResult = result
	}
}

func WithOnAfterResponse(middleware req.ResponseMiddleware) ClientOption {
	return func(options *clientOptions) {
		options.OnAfterResponse = middleware
	}
}
