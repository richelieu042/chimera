package mailKit

import (
	"testing"

	"github.com/richelieu042/chimera/v3/src/core/intKit"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
)

func TestSendMail(t *testing.T) {
	/* 仅供参考，下面的 Password 已失效 */

	sender := &SmtpConfig{
		Address:  "smtp.163.com:25",
		Host:     "smtp.163.com",
		Account:  "miro42@163.com",
		Password: "ZRZUBJXZOFIOBNXM",
		//NickName: "yjs",
	}
	err := InitializeSmtp(sender, 2)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1; i++ {
		time := timeKit.FormatCurrent(timeKit.FormatEntire)
		mail := NewMail("", []string{"yjs@yozosoft.com", "richelieu042@gmail.com"}, "主题"+time, []byte(intKit.FormatInt(int64(i), 10)), nil, []string{"miro42@163.com"}, nil)
		err = SendMail(mail)
		if err != nil {
			panic(err)
		}
	}
}
