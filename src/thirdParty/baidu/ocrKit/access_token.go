package ocrKit

import (
	"fmt"
	"sync"
	"time"

	"github.com/richelieu042/chimera/v3/src/component/web/http_client/httpClientKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/serialize/json/jsonKit"
)

var (
	// clientId 即 apiKey
	clientId = ""
	// clientSecret 即 secretKey
	clientSecret = ""
)

var lock = new(sync.Mutex)
var token *accessToken = nil

func SetApiKeyAndSecretKey(apiKey, secretKey string) error {
	lock.Lock()
	defer lock.Unlock()

	apiKey = strKit.TrimSpace(apiKey)
	secretKey = strKit.TrimSpace(secretKey)
	if strKit.IsEmpty(apiKey) {
		return errorKit.Newf("apiKey is empty")
	}
	if strKit.IsEmpty(secretKey) {
		return errorKit.Newf("secretKey is empty")
	}
	token = nil
	clientId = apiKey
	clientSecret = secretKey
	return nil
}

func getAccessToken() (*accessToken, error) {
	lock.Lock()
	defer lock.Unlock()

	// (1) token存在
	if token != nil {
		if !token.isExpired() {
			return token, nil
		}
		// 去除已过期的token
		token = nil
	}
	// (2) token存在，重新生成
	var err error
	token, err = newAccessToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}

// newAccessToken 向百度获取 Access Token
/*
参考: https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu

@return 必定一个为nil，一个非nil
*/
func newAccessToken() (*accessToken, error) {
	if strKit.HasEmpty(clientId, clientSecret) {
		return nil, errorKit.Newf("SetApiKeyAndSecretKey() should be invoked firstly")
	}

	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s",
		grantType,
		clientId,
		clientSecret)
	_, respData, err := httpClientKit.Get(url)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	if err := jsonKit.Unmarshal(respData, &m); err != nil {
		return nil, err
	}

	token, err := parseMapToAccessToken(m)
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errorKit.Newf("failure response(%s)", string(respData))
	}
	// 对token进行简单验证
	// 正常情况下，token的有效期为30天，此处判断是为了防止特殊情况（86400秒 == 1天）
	if token.ExpiresIn <= 86400 {
		return nil, errorKit.Newf("token.ExpiresIn(%d) is invalid", 86400)
	}
	// 提前1h(3600s)认为token失效，以防特殊情况
	token.ExpireTime = time.Now().Add(time.Second * time.Duration(token.ExpiresIn-3600))
	return token, nil
}
