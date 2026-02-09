package httpKit

import (
	"net/http"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// GetAcceptLanguages 根据请求头"Accept-Language".
func GetAcceptLanguages(req *http.Request) []string {
	langs := make([]string, 0)

	/*
		e.g.
		"zh-CN,zh;q=0.9"
		"zh,zh-TW;q=0.9,en-US;q=0.8,en;q=0.7,zh-CN;q=0.6"
		"fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5"
	*/
	headerValue := GetHeader(req.Header, HeaderAcceptLanguage)
	s := strKit.Split(headerValue, ",")
	for _, ele := range s {
		ele = strKit.TrimSpace(ele)
		ele = strKit.SubBeforeString(ele, ";q=")
		if strKit.IsEmpty(ele) {
			continue
		}
		langs = append(langs, ele)
	}
	return langs
}
