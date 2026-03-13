package timeKit

import (
	"context"
	"net/http"
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

@param ctx 	(1) 不能为nil
			(2) 建议附带timeout！！！
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
	for _, source := range sources {
		go func(url string) {
			t, err := getNetworkTimeByUrl(ctx, url)
			if err != nil {
				return
			}

			ch <- &bean{
				source: url,
				time:   t,
			}
		}(source)
	}

	select {
	case b := <-ch:
		return b.time, b.source, nil
	case <-ctx.Done():
		return time.Time{}, "", ctx.Err()
	}
}

func getNetworkTimeByUrl(ctx context.Context, url string) (t time.Time, err error) {
	// 创建请求，并绑定 context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		err = errKit.Wrapf(err, "fail to new http request")
		return
	}

	// 可自定义客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		err = errKit.Wrapf(err, "fail to send http request")
		return
	}
	defer resp.Body.Close()

	value := resp.Header.Get(httpKit.HeaderDate)
	if value == "" {
		err = errKit.New("value of header is empty")
		return
	}
	t, err = Parse(FormatRFC1123, value)
	if err != nil {
		err = errKit.Wrap(err, "fail to parse value of header")
		return
	}
	return
}
