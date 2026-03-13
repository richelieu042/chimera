package restyKit

import (
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	defClient     *resty.Client
	defClientOnce sync.Once
)

func getDefClient() *resty.Client {
	defClientOnce.Do(func() {
		/*
			SetTimeout(10 * time.Second) — 整个请求的超时时间（含连接+读写），超过 10 秒直接报错
			SetRetryCount(3) — 失败时最多重试 3 次（即总共最多发 4 次请求）
			SetRetryWaitTime(500 * time.Millisecond) — 每次重试前的初始等待时间 500ms
			SetRetryMaxWaitTime(time.Second) — 重试等待时间的上限是 1 秒，防止退避时间过长

			resty 的重试等待时间用的是指数退避策略：
				SetRetryWaitTime(500ms) — 第一次重试前等待的初始值
				SetRetryMaxWaitTime(1s) — 等待时间的上限，再怎么退避也不会超过这个值

			每次重试等待时间会翻倍增长，但不超过上限：
				第1次重试：等 500ms
				第2次重试：等 1000ms → 超过上限，截断为 1s
				第3次重试：等 1s
		*/
		defClient = resty.New().SetTimeout(15 * time.Second).
			SetRetryCount(3).
			SetRetryWaitTime(300 * time.Millisecond).
			SetRetryMaxWaitTime(time.Second)
	})
	return defClient
}

// Get
/*
@param queryParams 可以为 nil
*/
func Get(utl string, queryParams map[string]string) (int, string, error) {
	client := getDefClient()

	resp, err := client.R().
		SetQueryParams(queryParams).
		Get(utl)
	if err != nil {
		return 0, "", err
	}
	return resp.StatusCode(), string(resp.Body()), nil
}
