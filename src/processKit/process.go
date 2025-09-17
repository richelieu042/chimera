package processKit

import (
	"os"

	"github.com/shirou/gopsutil/v4/process"
)

var PID int = os.Getpid()

var GetRunningPids func() ([]int32, error) = process.Pids

// PidExists pid是否存在?
var PidExists func(pid int32) (bool, error) = process.PidExists
