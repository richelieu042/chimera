package timeKit

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

var sources = []string{
	"https://www.google.com",
	"https://github.com",
	"https://www.bilibili.com",
	"https://www.baidu.com",
	"https://www.taobao.com",
	"https://www.360.cn",
	"https://www.kingsoft.com",
	"https://www.yozosoft.com",
}

// GetNetworkTime
/*
PS: 方法体内不要直接使用 reqKit，以防import cycle.

@return time.Time 	获取到的网络时间
		string		网络时间的来源
		error		错误
*/
func GetNetworkTime(ctx context.Context) (time.Time, string, error) {
	type bean struct {
		source string
		time   time.Time
	}

	ch := make(chan *bean, len(sources))
	client := &http.Client{}
	ctx1, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	for _, source := range sources {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			t, err := getNetworkTimeByUrl(ctx1, client, url)
			if err != nil {
				return
			}
			ch <- &bean{
				source: url,
				time:   t,
			}
		}(source)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case b, ok := <-ch:
		if !ok {
			return time.Time{}, "", errKit.New("all network time sources failed")
		}
		return b.time, b.source, nil
	case <-ctx1.Done():
		select {
		case b, ok := <-ch:
			if ok {
				return b.time, b.source, nil
			}
		default:
		}
		return time.Time{}, "", ctx1.Err()
	}
}

func getNetworkTimeByUrl(ctx context.Context, client *http.Client, url string) (t time.Time, err error) {
	// 创建请求，并绑定 context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		err = errKit.Wrapf(err, "fail to new http request")
		return
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		err = errKit.Wrapf(err, "fail to send http request")
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	value := resp.Header.Get(httpKit.HeaderDate)
	if value == "" {
		err = errKit.New("value of header is empty")
		return
	}
	t, err = Parse(http.TimeFormat, value)
	if err != nil {
		err = errKit.Wrap(err, "fail to parse value of header")
		return
	}
	return
}
