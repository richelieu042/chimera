package httpClientKit

import (
	"crypto/tls"
	"net/http"
	"time"
)

// DefaultHttpClient 默认的 *http.Client 实例
var DefaultHttpClient *http.Client = NewHttpClient(time.Second*3, true)

// NewHttpClient
/*
@param insecureSkipVerify true: 跳过https证书验证（适用于: https证书是自己生成的，非法的）
*/
func NewHttpClient(timeout time.Duration, insecureSkipVerify bool) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecureSkipVerify,
			},
		},
	}
}
