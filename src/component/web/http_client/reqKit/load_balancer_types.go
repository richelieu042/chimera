package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu042/chimera/v3/src/randomKit"
)

// LoadBalancerClient http客户端的负载均衡
type LoadBalancerClient struct {
	client *req.Client

	urls []string
}

func (lbc *LoadBalancerClient) Get(queryParams map[string][]string) (*req.Response, error) {
	/* 不能每次都从0开始，否则第一个url压力太大 */
	index := randomKit.Int(0, len(lbc.urls))
	startUrl := lbc.urls[index]
	//console.Infof("start url: %s", startUrl)

	r := lbc.client.Get(startUrl)
	r.SetRetryHook(func(resp *req.Response, err error) {
		index++
		index = index % len(lbc.urls)
		retryUrl := lbc.urls[index]
		//console.Infof("retry url: %s", retryUrl)

		/* 修改请求的url */
		r.SetURL(retryUrl)
	})

	for key, s := range queryParams {
		r.AddQueryParams(key, s...)
	}

	resp := r.Do()
	return resp, resp.Err
}

func (lbc *LoadBalancerClient) Post(contentType string, body interface{}) (*req.Response, error) {
	/* 不能每次都从0开始，否则第一个url压力太大 */
	index := randomKit.Int(0, len(lbc.urls))
	startUrl := lbc.urls[index]
	//console.Infof("start url: %s", startUrl)

	r := lbc.client.Get(startUrl)
	r.SetRetryHook(func(resp *req.Response, err error) {
		index++
		index = index % len(lbc.urls)
		retryUrl := lbc.urls[index]
		//console.Infof("retry url: %s", retryUrl)

		/* 修改请求的url */
		r.SetURL(retryUrl)
	})

	if contentType == "" {
		contentType = ContentTypeJson
	}
	r.SetContentType(contentType)
	r.SetBody(body)

	resp := r.Do()
	return resp, resp.Err
}
