package ocrKit

import (
	"time"

	"github.com/richelieu042/chimera/v3/src/core/mapKit"
)

type (
	accessToken struct {
		// AccessToken 要获取的Access Token
		AccessToken string `mapstructure:"access_token"`
		// ExpiresIn Access Token的有效期(秒为单位，有效期30天)
		ExpiresIn uint32 `mapstructure:"expires_in"`

		ExpireTime time.Time

		RefreshToken  string `mapstructure:"refresh_token"`
		Scope         string `mapstructure:"scope"`
		SessionKey    string `mapstructure:"session_key"`
		SessionSecret string `mapstructure:"session_secret"`
	}

	Words struct {
		LogId          uint64        `mapstructure:"log_id"`
		WordsResultNum uint32        `mapstructure:"words_result_num"`
		WordsResults   []WordsResult `mapstructure:"words_result"`
	}

	WordsResult struct {
		Words string `mapstructure:"words"`
	}
)

func (token *accessToken) isExpired() bool {
	return time.Now().After(token.ExpireTime)
}

func parseMapToAccessToken(m map[string]interface{}) (*accessToken, error) {
	if !mapKit.ContainKeys(m, "access_token", "expires_in") {
		return nil, nil
	}

	token := &accessToken{}
	err := mapKit.Decode(m, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func parseMapToWords(m map[string]interface{}) (*Words, error) {
	if !mapKit.ContainKeys(m, "log_id", "words_result_num", "words_result") {
		return nil, nil
	}
	words := &Words{}
	err := mapKit.Decode(m, words)
	if err != nil {
		return nil, err
	}
	return words, nil
}
