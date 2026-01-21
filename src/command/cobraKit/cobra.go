package cobraKit

import (
	"github.com/spf13/cobra"
)

// NewSimpleCommand
/*
PS:
(1) 自动支持 "-h" 或 "--help" 标识.

@param use		命令的名称
@param short	命令的 "简短描述" ，帮助用户快速理解其功能
@param long		命令的 "详细描述" ，进一步解释工具的用途
				（1）如果为空，会使用 short
				（2）可以是多行的文本
@param run		定义了当执行该命令时的行为
*/
func NewSimpleCommand(use, short, long string, run func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,

		Run: run,
	}
}
