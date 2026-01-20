## Windows环境，执行.exe文件
golang 调用cmd执行EXE  
    http://www.manongjc.com/detail/22-yljaaewtenbdcop.html
go 调用可执行程序并传参(windows 系统exe程序示例)  
    https://blog.csdn.net/vily_lei/article/details/129985497

实际上是通过start命令.（start /B类似于linux的nohup，但目前重定向输出有问题，后续研究下）

#### macOS demo
cmdKit.ExecuteToString(context.TODO(), "sh", "-c", `cmd top -n 1 -b -p 12345 | grep 12345 | awk -F " " '{print $6}'`)

#### Windows demo
exec.Command("cmd.exe", "/c", "start", "logon.exe")
exec.Command("cmd.exe", "/c", "start", "/B", "logon.exe")

#### Windows demo1（缺陷: lib/main.exe的输出丢了）
PS: 如果想要lib/main.exe的输出，可以考虑使用命令: start /b lib/main.exe > output.txt，将输出重定向到文件中，然后再读取文件内容.

```golang
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path := "lib/main.exe"
	cmd := exec.Command("cmd.exe", "/c", "start", path)

	data, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
```
