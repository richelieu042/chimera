package httpKit

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"
	"time"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

// GetCertificateInfo
/*
获取https过期时间
	https://www.topgoer.cn/docs/gochajian/gofdgjh

@return 仅返回第一个证书信息（有多个的话）
*/
func GetCertificateInfo(rawURL string) (*x509.Certificate, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, errKit.Wrapf(err, "fail to parse url(%s)", rawURL)
	}
	if parsedURL.Scheme != "https" || parsedURL.Hostname() == "" {
		return nil, errKit.Newf("invalid HTTPS url(%s)", rawURL)
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	defer transport.CloseIdleConnections()

	client := &http.Client{
		Transport: transport,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			// 只查询传入 URL 对应的证书，不跟随到其他端点。
			return http.ErrUseLastResponse
		},
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(parsedURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.TLS == nil {
		return nil, errKit.Newf("response from url(%s) has no TLS connection state", rawURL)
	}
	certs := resp.TLS.PeerCertificates
	if len(certs) == 0 {
		return nil, errKit.Newf("length of certs is zero")
	}
	return certs[0], nil
}
