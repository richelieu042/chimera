package mailKit

import (
	"crypto/tls"
	"sync"

	"github.com/emersion/go-imap"
	id "github.com/emersion/go-imap-id"
	"github.com/emersion/go-imap/client"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
)

type (
	ImapConfig struct {
		Address   string
		TlsConfig *tls.Config
		UserName  string
		Password  string
	}
)

var imapLock = &sync.RWMutex{}
var imapConfig *ImapConfig

func InitializeImap(config *ImapConfig) error {
	imapLock.Lock()
	defer imapLock.Unlock()

	c, err := newImapClient(config)
	if err != nil {
		return err
	}
	defer func() {
		_ = c.Close()
	}()

	imapConfig = config
	return nil
}

// newImapClient 连接服务器 && 登录
/*
@param addr		e.g. "imap.163.com:993"
@param username 邮箱，e.g. "miro42@163.com"
@param password 邮箱的授权密码（并非登录密码！！！），e.g. "ZRZUBJXZOFIOBNXM"
*/
func newImapClient(config *ImapConfig) (c *client.Client, err error) {
	defer func() {
		// 如果创建客户端失败，为防万一，额外调用下Logout()
		if err != nil && c != nil {
			_ = c.Logout()
		}
	}()
	if config == nil {
		err = errorKit.Newf("config == nil")
		return
	}

	// (1) 连接服务器
	if c, err = client.DialTLS(config.Address, config.TlsConfig); err != nil {
		return
	}

	// (2) 登录
	err = c.Login(config.UserName, config.Password)
	if err != nil {
		return
	}

	// (3) 添加ID信息（以防收取163邮箱报错(error): SELECT Unsafe Login. Please contact kefu@188.com for help）
	idClient := id.NewClient(c)
	_, err = idClient.ID(id.ID{
		id.FieldName:    "IMAPClient",
		id.FieldVersion: "3.1.0",
	})
	return
}

func GetMailboxInfoSlice() ([]*imap.MailboxInfo, error) {
	c, err := newImapClient(imapConfig)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = c.Close()
	}()

	ch := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", ch)
	}()
	// 流程: 先处理ch信道，再处理done信道（不能反过来！）
	s := make([]*imap.MailboxInfo, 0, len(ch))
	for m := range ch {
		s = append(s, m)
	}
	if err := <-done; err != nil {
		return nil, err
	}
	return s, nil
}
