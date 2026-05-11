package cmdKit

import (
	"context"
	"os/exec"
	"strings"
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

// Run 执行命令（会阻塞直到命令结束） - 只有 stdout 在 output 中，标准错误（stderr）会被丢弃或继承父进程的 stderr
/*
!!!:
(1) exec.Cmd结构体执行时，会处理路径中的空格（e.g. java可执行文件的绝对路径、-Djava.ext.dirs=的路径...）
(2) 假如自行处理命令行中的路径，反而会导致: 命令执行失败

能从 err 中获取 stderr 的情况：	只有当命令以非零退出码结束时，Output() 才会返回 *ExitError，此时可以从中获取 stderr
无法从 err 中获取 stderr 的情况：	如果命令成功执行（退出码为 0），即使有 stderr 输出，Output() 也会返回 err = nil，stderr 内容会丢失
*/
func Run(ctx context.Context, name string, args ...string) ([]byte, *exec.Cmd, error) {
	cmd := exec.CommandContext(ctx, name, args...)

	data, err := cmd.Output()
	return data, cmd, err
}

// RunCombinedly 执行命令（会阻塞直到命令结束） - stdout 和 stderr 都在 output 中，两者的内容会合并到一起返回
func RunCombinedly(ctx context.Context, name string, args ...string) ([]byte, *exec.Cmd, error) {
	cmd := exec.CommandContext(ctx, name, args...)

	data, err := cmd.CombinedOutput()
	return data, cmd, err
}

func RunToString(ctx context.Context, name string, args ...string) (string, *exec.Cmd, error) {
	data, cmd, err := Run(ctx, name, args...)

	resp := string(data)
	resp = strings.TrimSpace(resp) // 一般最后会有一个'\n'

	return resp, cmd, err
}

func RunCombinedlyToString(ctx context.Context, name string, args ...string) (string, *exec.Cmd, error) {
	data, cmd, err := RunCombinedly(ctx, name, args...)

	resp := string(data)
	resp = strings.TrimSpace(resp) // 一般最后会有一个'\n'

	return resp, cmd, err
}
