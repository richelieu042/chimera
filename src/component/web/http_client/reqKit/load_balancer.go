package reqKit

import (
	"net/http"
	"time"

	"github.com/imroc/req/v3"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
)

// NewLbClient http客户端的负载均衡.
/*
原理: 请求失败retry时，修改请求的url.

@param baseClient 	(1) 可以为nil，将采用默认值
					(2) !!!: 本函数会修改此传参
@param commonRetryInterval 	重试周期
@param commonRetryCondition	重试条件
*/
func NewLbClient(baseClient *req.Client, urls []string, commonRetryInterval time.Duration, commonRetryCondition req.RetryConditionFunc) (*LoadBalancerClient, error) {
	if err := sliceKit.AssertNotEmpty(urls, "urls"); err != nil {
		return nil, err
	}

	if baseClient == nil {
		baseClient = NewClient()
	}
	if commonRetryInterval <= 0 {
		commonRetryInterval = time.Millisecond * 100
	}
	if commonRetryCondition == nil {
		commonRetryCondition = func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode != http.StatusOK
		}
	}

	baseClient.SetCommonRetryCount(len(urls) - 1).
		SetCommonRetryFixedInterval(commonRetryInterval).
		SetCommonRetryCondition(commonRetryCondition)

	return &LoadBalancerClient{
		client: baseClient,
		urls:   urls,
	}, nil
}
