package cmdKit

import (
	"context"
	"os/exec"
)

var LookPath func(file string) (string, error) = exec.LookPath

// NewCommand
/*
PS:
(1) Cmd.Start():	不会阻塞;
(2) Cmd.Run(): 		会阻塞.

@param ctx	(1) 如果命令在ctx超时之前没有完成，CommandContext将会杀死该进程;
			(2) !!!: 存在部分杀不死的情况，比如yozo的logon.exe，ctx超时了进程还是卡死在那（通过go代码执行会卡死在那，且无法通过ctx结束；但直接在command里面执行就没问题）.
@param args 可以为nil
*/
func NewCommand(ctx context.Context, name string, args []string, options ...CmdOption) *exec.Cmd {
	opts := loadOptions(options...)
	return opts.NewCommand(ctx, name, args...)
}

// Run 执行命令（会阻塞直到命令结束）
/*
!!!:
(1) exec.Cmd结构体执行时，会处理路径中的空格（e.g. java可执行文件的绝对路径、-Djava.ext.dirs=的路径...）
(2) 假如自行处理命令行中的路径，反而会导致: 命令执行失败
*/
func Run(ctx context.Context, name string, args ...string) ([]byte, error) {
	if ctx == nil {
		ctx = context.TODO()
	}

	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.Output()
}

// RunCombinedly 执行命令（会阻塞直到命令结束）
func RunCombinedly(ctx context.Context, name string, args ...string) ([]byte, error) {
	if ctx == nil {
		ctx = context.TODO()
	}

	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.CombinedOutput()
}
