package cmdKit

import (
	"syscall"
)

type (
	// CmdOptions 参考: exec_linux.go
	CmdOptions struct {
		// Setpgid sets the process group ID of the child to Pgid,
		// or, if Pgid == 0, to the new child's process ID.
		Setpgid bool

		// Pdeathsig, if non-zero, is a signal that the kernel will send to
		// the child process when the creating thread dies. Note that the signal
		// is sent on thread termination, which may happen before process termination.
		// There are more details at https://go.dev/issue/27505.
		Pdeathsig syscall.Signal
	}

	CmdOption func(opts *CmdOptions)
)

func loadOptions(options ...CmdOption) *CmdOptions {
	opts := &CmdOptions{
		Setpgid:   false,
		Pdeathsig: 0,
	}

	for _, option := range options {
		option(opts)
	}
	return opts
}

// WithSetpgid
/*
适用环境: Linux、macOS
*/
func WithSetpgid(setpgid bool) CmdOption {
	return func(opts *CmdOptions) {
		opts.Setpgid = setpgid
	}
}

// WithPdeathsig
/*
适用环境: Linux
*/
func WithPdeathsig(pdeathsig syscall.Signal) CmdOption {
	return func(opts *CmdOptions) {
		opts.Pdeathsig = pdeathsig
	}
}
