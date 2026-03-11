package reqKit

import (
	"context"

	"github.com/imroc/req/v3"
)

// Post TODO: 待验证.
/*
PS: 要注意 req.Response 结构体的 Err 字段.

@param body Set the request Body, accepts string、[]byte、io.Reader、map and struct.
*/
func Post(ctx context.Context, url string, contentType string, body interface{}) *req.Response {
	if contentType == "" {
		contentType = ContentTypeJson
	}

	r := GetGlobalClient().Post(url)
	r.SetContentType(contentType)
	r.SetBody(body)
	return r.Do(ctx)
}

//func PostAndInto(ctx context.Context, url string, body interface{}, v interface{}) error {
//	resp := Post(ctx, url, body)
//	if resp.Err != nil {
//		return resp.Err
//	}
//
//	// 反序列化
//	return resp.Into(v)
//}
