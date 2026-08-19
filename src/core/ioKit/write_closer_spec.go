package ioKit

import (
	"fmt"
	"io"
	"math"

	"github.com/richelieu042/chimera/v3/src/consts"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/cronKit"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type DailyWriteCloser struct {
	writeCloser io.WriteCloser
	cron        *cron.Cron
}

func (dwc *DailyWriteCloser) Write(p []byte) (int, error) {
	return dwc.writeCloser.Write(p)
}

func (dwc *DailyWriteCloser) Close() error {
	_ = dwc.cron.Stop()
	return dwc.writeCloser.Close()
}

// NewDailyWriteCloser 每天凌晨0点，执行Rotate().
/*
@param options 可选配置，参考 NewLumberJackWriteCloser()
*/
func NewDailyWriteCloser(filePath string, options ...LumberJackOption) (io.WriteCloser, error) {
	return NewRotatableWriteCloserWithSpec(filePath, "0 0 0 * * *", options...)
}

// NewRotatableWriteCloserWithSpec 满足条件（spec），执行Rotate().
/*
PS:
(1) 可能存在情况，Rotate()后，生成的旧日志文件大小为0B.
*/
func NewRotatableWriteCloserWithSpec(filePath string, cronSpec string, options ...LumberJackOption) (io.WriteCloser, error) {
	// math.MaxInt64: 8.0 EiB
	wc, err := NewLumberJackWriteCloser(filePath, math.MaxInt64, options...)
	if err != nil {
		return nil, err
	}

	c, _, err := cronKit.NewCronWithTask(cronSpec, func() {
		text := fmt.Sprintf("[%s] Rotate by cron.\n", strKit.ToUpper(consts.ProjectName))
		_, _ = wc.Write([]byte(text))
		if err := wc.Rotate(); err != nil {
			text := fmt.Sprintf("[%s] Fail to rotate by cron, error:\n%v\n", strKit.ToUpper(consts.ProjectName), err)
			_, _ = wc.Write([]byte(text))
			logrus.Error(text)
		}
	})
	if err != nil {
		return nil, err
	}
	// Start() 不阻塞
	c.Start()

	return &DailyWriteCloser{
		writeCloser: wc,
		cron:        c,
	}, nil
}
