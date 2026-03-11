package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

// NewClient
/*
!!!:
(1) 不要每次发请求都创建 Client，造成不必要的开销，通常可以复用同一 Client 发所有请求;
(2) 自动重试(retry): 可以在返回值上自行添加.

@param options 不传参的情况下，	(1) 生产模式
								(2) 超时时间: 30s
								(3) 跳过https证书验证
								(4) 日志输出: os.Stdout
								(5) 不重试
*/
func NewClient(options ...ClientOption) (client *req.Client) {
	opts := loadOptions(options...)
	client = req.C()

	/* 开发者模式（甩锅） */
	if opts.Dev {
		client.DevMode()
	}

	/*
		启用自动检测字符集并解码为utf-8
		（imroc/req默认: 启用）
	*/
	client.EnableAutoDecode()

	/*
		启用压缩
		（imroc/req默认: 启用）
	*/
	client.EnableCompression()

	/*
		disable sending GET method requests with body
		（imroc/req默认: 启用，即发送GET请求时附带body）
	*/
	client.DisableAllowGetMethodPayload()

	/* json序列化和反序列化 */
	client.SetJsonMarshal(jsonKit.Marshal).
		SetJsonUnmarshal(jsonKit.Unmarshal)

	/* 伪装成Chrome浏览器发起请求，主要针对: 反爬虫检测 */
	client.ImpersonateChrome()

	/* BaseURL */
	if opts.BaseURL != "" {
		client.SetBaseURL(opts.BaseURL)
	}

	/*
		超时时间（发送请求的整个周期，includes connection time, any redirects, and reading the response body）
		(1) imroc/req默认: 2min
		(2) 0: no timeout
	*/
	client.SetTimeout(opts.Timeout)

	/*
		https证书验证
		（imroc/req默认: 不跳过）
	*/
	if opts.InsecureSkipVerify {
		client.EnableInsecureSkipVerify()
	} else {
		// 推荐正式环境使用，更安全
		client.DisableInsecureSkipVerify()
	}

	/*
		logger
		（imroc/req默认: 输出到 os.Stdout）
	*/
	client.SetLogger(opts.Logger)

	/* 事件 */
	if opts.CommonErrorResult != nil {
		client.SetCommonErrorResult(opts.CommonErrorResult)
	}
	if opts.OnAfterResponse != nil {
		client.OnAfterResponse(opts.OnAfterResponse)
	}

	return
}
