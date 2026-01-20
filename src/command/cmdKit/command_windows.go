package cmdKit

import (
	"context"
	"os/exec"
)

func (opts CmdOptions) NewCommand(ctx context.Context, name string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)

	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Setpgid:   opts.Setpgid,
	//	Pdeathsig: opts.Pdeathsig,
	//}
	return cmd
}
