package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	output, err := cmd.Output()
	if err != nil {
		// 如果命令执行失败，err 是 *ExitError 类型
		// 可以通过类型断言获取 stderr
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			fmt.Println("stderr:", string(exitErr.Stderr))
		}
	}
	fmt.Println("stdout:", string(output))
}
