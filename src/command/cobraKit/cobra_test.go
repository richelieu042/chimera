package cobraKit

import (
	"fmt"
	"os"
	"testing"

	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/spf13/cobra"
)

func TestNewSimpleCommand(t *testing.T) {
	rootCmd := NewSimpleCommand("ccc", "ccc的说明.", "", func(cmd *cobra.Command, args []string) {
		console.Info("logic")
	})

	versionCmd := NewSimpleCommand("version", "Print the version number of newApp.", "", func(cmd *cobra.Command, args []string) {
		fmt.Println("version: 0.0.1")
	})

	// 添加子命令
	rootCmd.AddCommand(versionCmd)

	os.Args = []string{"ccc", "version", "yjs", "ylx", "wmq"}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
