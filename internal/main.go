package main

import (
	"fmt"

	"github.com/richelieu-yang/chimera/v3/src/command/cobraKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/spf13/cobra"
)

func main() {
	cc := cobraKit.NewSimpleCommand("ccc", "简短描述。", "详细描述", func(cmd *cobra.Command, args []string) {
		console.Info("Run...")
	})

	err := cc.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
