package httpClientKit

import (
	"io"
	"net/http"

	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/urlKit"
)

// Get Deprecated: Use req111Kit instead.
func Get(url string, options ...Option) (int, []byte, error) {
	resp, err := GetForResponse(url, options...)
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

// GetForResponse Deprecated: Use req111Kit instead.
/*
!!!: 第2个返回值 == nil 的情况下，需要手动调用 resp.Body.Close() 来手动关闭 第1个返回值.
*/
func GetForResponse(url string, options ...Option) (*http.Response, error) {
	opts := loadOptions(options...)

	// url
	if err := httpKit.AssertHttpUrl(url); err != nil {
		return nil, err
	}
	url, err := urlKit.PolyfillUrl(url, opts.queryParams)
	if err != nil {
		return nil, err
	}

	// client
	client := opts.newHttpClient()

	// req
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// send
	return client.Do(req)
}
