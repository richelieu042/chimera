package httpKit

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"time"

	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// GetCertificateInfo
/*
获取https过期时间
	https://www.topgoer.cn/docs/gochajian/gofdgjh

@return 仅返回第一个证书信息（有多个的话）
*/
func GetCertificateInfo(url string) (*x509.Certificate, error) {
	if !strKit.StartWith(url, "https://") {
		return nil, errorKit.Newf("invalid url(%s)", url)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	certs := resp.TLS.PeerCertificates
	if len(certs) == 0 {
		return nil, errorKit.Newf("length of certs is zero")
	}
	return certs[0], nil
}
