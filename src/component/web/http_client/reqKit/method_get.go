package reqKit

import (
	"context"

	"github.com/imroc/req/v3"
)

func SimpleGet(ctx context.Context, url string) *req.Response {
	return Get(ctx, url, nil)
}

// Get
/*
PS: 要注意 req.Response 结构体的 Err 字段.

@param queryParams 可以为nil
*/
func Get(ctx context.Context, url string, queryParams map[string][]string) *req.Response {
	r := GetGlobalClient().Get(url)

	for key, s := range queryParams {
		r.AddQueryParams(key, s...)
	}
	return r.Do(ctx)
}

func GetAndInto(ctx context.Context, url string, queryParams map[string][]string, v interface{}) error {
	resp := Get(ctx, url, queryParams)
	if resp.Err != nil {
		return resp.Err
	}

	// 反序列化
	return resp.Into(v)
}
