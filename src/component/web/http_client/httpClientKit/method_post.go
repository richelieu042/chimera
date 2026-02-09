package httpClientKit

import (
	"io"
	"net/http"
	"strings"

	httpKit2 "github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/urlKit"
)

// Post Deprecated: Use req111Kit instead.
func Post(url string, options ...Option) (int, []byte, error) {
	resp, err := PostForResponse(url, options...)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, data, nil
}

// PostForResponse Deprecated: Use req111Kit instead.
/*
!!!: 第2个返回值 == nil 的情况下，需要手动调用 resp.Body.Close() 来手动关闭 第1个返回值.
*/
func PostForResponse(url string, options ...Option) (*http.Response, error) {
	opts := loadOptions(options...)

	// url
	if err := httpKit2.AssertHttpUrl(url); err != nil {
		return nil, err
	}
	url, err := urlKit.PolyfillUrl(url, opts.queryParams)
	if err != nil {
		return nil, err
	}

	// payload
	content := httpKit2.ToRequestBodyString(opts.postParams)
	payload := strings.NewReader(content)

	// req
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", "charset=utf-8")

	// client
	client := opts.newHttpClient()

	return client.Do(req)
}
