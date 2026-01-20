package cobraKit

import (
	"github.com/spf13/cobra"
)

func NewSimpleCommand(use, short, long string, run func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,

		Run: run,
	}
}
