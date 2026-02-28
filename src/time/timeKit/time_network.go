package timeKit

import (
	"context"
	"sync"
	"time"

	"github.com/imroc/req/v3"
	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
)

/*
PS: 腾讯的时间不准.
*/
var sources = []string{
	"https://www.google.com",
	//"https://www.tencent.com",
	"https://github.com",
	"https://www.bilibili.com",
	"https://www.baidu.com",
	"https://www.taobao.com",
	"https://www.360.cn",
	"https://www.kingsoft.com",
	"https://www.yozosoft.com",
}

var (
	defaultReqClientOnce sync.Once

	// Deprecated: 仅供内部使用 && 不要直接使用，想用就调用 getDefaultReqClient()
	defaultReqClient *req.Client
)

func getDefaultReqClient() *req.Client {
	defaultReqClientOnce.Do(func() {
		defaultReqClient = req.C().
			ImpersonateChrome().
			SetTimeout(time.Second * 30).
			EnableInsecureSkipVerify()
	})

	return defaultReqClient
}

// GetNetworkTime
/*
!!!: 方法体内不要直接使用 reqKit，以防import cycle.

@param ctx 	(1) 不能为nil
			(2) 建议附带timeout
*/
func GetNetworkTime(ctx context.Context) (t time.Time, source string, err error) {
	type bean struct {
		source string
		time   time.Time
	}

	ch := make(chan *bean, len(sources))
	for _, source := range sources {
		go func(url string) {
			t, err = GetNetworkTimeByUrl(ctx, url)
			if err != nil {
				//fmt.Printf("fail to get network time from url(%s), error: %s\n", url, err)
				return
			}

			//fmt.Printf("get network time, source: %s, time: %s\n", url, t)
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
		err = ctx.Err()
		return
	}
}

// GetNetworkTimeByUrl
/*
!!!: 方法体内不要直接使用 reqKit，以防import cycle.
*/
func GetNetworkTimeByUrl(ctx context.Context, url string) (t time.Time, err error) {
	resp := getDefaultReqClient().Get(url).Do(ctx)
	if resp.Err != nil {
		err = resp.Err
		return
	}

	str := resp.GetHeader(httpKit.HeaderDate)
	if str == "" {
		err = errorKit.Newf("empty header: %s", httpKit.HeaderDate)
		return
	}
	t, err = Parse(FormatRFC1123, str)
	if err != nil {
		err = errorKit.Wrapf(err, "fail to prase with FormatRFC1123")
		return
	}
	return
}
