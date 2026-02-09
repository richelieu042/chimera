package httpClientKit

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/richelieu042/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/urlKit"
)

// Upload Deprecated: Use req111Kit instead.
func Upload(url string, fileParams map[string]string, options ...Option) (int, []byte, error) {
	resp, err := UploadForResponse(url, fileParams, options...)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, data, nil
}

// UploadForResponse Deprecated: Use req111Kit instead.
/*
@param fileParams 	(1)可以为nil或空;
					(2)key: 键, value: 要上传文件的路径.
@return !!!: 第一个返回值如果不为nil的话，一般来说需要手动调用 "resp.Body.Close()"
*/
func UploadForResponse(url string, fileParams map[string]string, options ...Option) (*http.Response, error) {
	opts := loadOptions(options...)

	// url
	if err := httpKit.AssertHttpUrl(url); err != nil {
		return nil, err
	}
	url, err := urlKit.PolyfillUrl(url, opts.queryParams)
	if err != nil {
		return nil, err
	}

	// payload
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	// 可能会多次关闭 multipart.Writer实例
	defer writer.Close()
	for k, s := range opts.postParams {
		for _, v := range s {
			// PS: 此处无需对v进行编码操作
			if err := writer.WriteField(k, v); err != nil {
				return nil, err
			}
		}
	}
	for field, path := range fileParams {
		if err := fileKit.AssertExistAndIsFile(path); err != nil {
			return nil, err
		}
		err := func() error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			part, err := writer.CreateFormFile(field, fileKit.GetFileName(path))
			if err != nil {
				return err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	// req
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Content-Type", "charset=utf-8")

	// client
	client := opts.newHttpClient()

	return client.Do(req)
}
