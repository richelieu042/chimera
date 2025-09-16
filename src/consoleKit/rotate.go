package consoleKit

import (
	"fmt"
	"time"

	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/cronKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"github.com/richelieu-yang/chimera/v3/src/time/timeKit"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

// RotateOutput spec触发时，备份并清空传参output对应的文件.
/*
PS:
(1) 缺陷: 可能会丢部分输出，目前没啥好的办法解决.
(2) 建议每天一次.
(3) 不能将nohup.out重命名然后再新建个nohup.out，因为输出还是会到原先那个nohup.out中.

@param output 	控制台输出文件（e.g.nohup.out）的路径
@param backDir	备份文件存放的目录
*/
func RotateOutput(output, backupDir, spec string) (*cron.Cron, error) {
	c, _, err := cronKit.NewCronWithTask(spec, func() {
		if backupPath, err := rotate(output, backupDir); err != nil {
			logrus.WithError(err).Errorf("[%s] Fail to rotate.", strKit.ToUpper(consts.ProjectName))
		} else {
			logrus.WithFields(logrus.Fields{
				"backupPath": backupPath,
			}).Infof("[%s] Succeed to rotate.", strKit.ToUpper(consts.ProjectName))
		}
	})
	if err != nil {
		return nil, err
	}
	c.Start()
	return c, nil
}

func rotate(output, backupDir string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(output); err != nil {
		return "", err
	}
	if err := fileKit.MkDirs(backupDir); err != nil {
		return "", err
	}

	dateStr := timeKit.Format(time.Now().Add(-timeKit.Day), "2006-01-02")
	name := fileKit.GetName(output)
	ext := fileKit.GetExt(output)
	target := pathKit.Join(backupDir, fmt.Sprintf("%s-%s%s", name, dateStr, ext))
	// 复制文件
	if err := fileKit.CopyFile(output, target); err != nil {
		return "", err
	}
	// 清空文件
	if err := fileKit.Truncate(output, 0); err != nil {
		return "", err
	}
	return target, nil
}
