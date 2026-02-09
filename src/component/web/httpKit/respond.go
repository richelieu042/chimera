package httpKit

import (
	"fmt"
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

// Status 设置响应的http状态码
/*
PS:
(1) 不建议多次设置 http状态码；
(2) 如果多次设置的话，感觉 第一次设置的值 会生效.

@param code -1: 不设置http状态码
*/
func Status(w http.ResponseWriter, code int) {
	if code <= 0 {
		return
	}
	w.WriteHeader(code)
}

// RespondString
/*
参考: gin里面的 Context.String() .
*/
func RespondString(w http.ResponseWriter, code int, format string, values ...any) error {
	data := strKit.StringToBytes(fmt.Sprintf(format, values...))
	return RespondData(w, code, PlainContentType, data)
}

func RespondStringData(w http.ResponseWriter, code int, data []byte) error {
	return RespondData(w, code, PlainContentType, data)
}

// RespondJson
/*
参考: gin里面的 Context.JSON() .
*/
func RespondJson(w http.ResponseWriter, code int, obj any) error {
	data, err := jsonKit.Marshal(obj)
	if err != nil {
		return err
	}

	return RespondData(w, code, JsonContentType, data)
}

// RespondData 响应字节流（二进制流）
/*
参考: gin里面的 Context.Content() .

@return 如果不为nil，建议输出到控制台
*/
func RespondData(w http.ResponseWriter, code int, contentType string, data []byte) error {
	if strKit.IsEmpty(contentType) {
		contentType = fileKit.DetectContentType(data)
	}

	Status(w, code)

	setContentType(w, []string{contentType})

	if !bodyAllowedForStatus(code) {
		return nil
	}
	_, err := w.Write(data)
	return err
}

// setContentType
/*
PS: copy from gin/render/render.go
*/
func setContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
/*
PS:
(1) copy from gin/context.go
(2) @return 在对应http状态码的情况下，是否允许写内容？
*/
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}
