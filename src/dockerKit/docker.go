package dockerKit

import "github.com/shirou/gopsutil/v4/docker"

// GetDockerIdList returns a list of DockerID.
var GetDockerIdList func() ([]string, error) = docker.GetDockerIDList
