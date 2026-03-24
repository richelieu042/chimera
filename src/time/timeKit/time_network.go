package timeKit

import (
	"context"
	"io"
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
!!!: 方法体内不要直接使用 reqKit，以防import cycle.

@param timeout 超时时间
@return time.Time 	获取到的网络时间
		string		网络时间的来源
		error		错误
*/
func GetNetworkTime(timeout time.Duration) (time.Time, string, error) {
	type bean struct {
		source string
		time   time.Time
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client := &http.Client{
		Timeout: timeout,
	}
	var wg sync.WaitGroup
	for _, source := range sources {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			t, err := getNetworkTimeByUrl(ctx, client, url)
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
	case b := <-ch:
		cancel()
		return b.time, b.source, nil
	case <-ctx.Done():
		return time.Time{}, "", ctx.Err()
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
		/*
			为什么不直接用 defer resp.Body.Close() ？
			函数目的是取 Date 响应头，状态码理论上不影响结果，但如果对端返回了 4xx/5xx，resp.Body 里可能有较大的 error body，而代码里 defer resp.Body.Close() 但从未读取 body，可能导致底层 TCP 连接无法复用（Transport 要求 body 被完全读取才能归还连接）
		*/
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
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
