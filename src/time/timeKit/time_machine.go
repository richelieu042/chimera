package timeKit

import (
	"context"
	"fmt"
	"time"

	"github.com/richelieu-yang/chimera/v3/src/command/cmdKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

// GetMachineTime 获取系统时间（机器时间；本地时间；time.Local）.
var GetMachineTime func() time.Time = time.Now

// SetMachineTime 设置系统时间（机器时间）
/*
PS:
(1) 通过date命令设置root权限，需要root权限.
(2) 只能精确到"秒".

@param password root用户的密码
*/
func SetMachineTime(ctx context.Context, t time.Time, rootPassword string) error {
	// 将时间转换为（date命令认可的）字符串
	format := "010215042006.05"
	timeStr := Format(t, TimeFormat(format))

	var script string
	if strKit.IsEmpty(rootPassword) {
		script = fmt.Sprintf("date %s", timeStr)
	} else {
		script = fmt.Sprintf(`echo "%s" | sudo -S date %s`, rootPassword, timeStr)
	}

	_, err := cmdKit.RunCombinedly(ctx, "sh", "-c", script)
	return err
}

// CorrectMachineTime （根据网络时间）纠正系统时间
func CorrectMachineTime(ctx context.Context, rootPassword string) (t time.Time, err error) {
	t, _, err = GetNetworkTime(ctx)
	if err != nil {
		return
	}
	err = SetMachineTime(ctx, t, rootPassword)
	return
}
