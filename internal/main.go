package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:              "app",
		TraverseChildren: true,
	}

	userCmd := &cobra.Command{
		Use: "user",
	}

	addCmd := &cobra.Command{
		Use: "add",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("user add")
			fmt.Println(args)
		},
	}

	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(addCmd)

	os.Args = []string{"app", "user", "foo", "add"}
	//os.Args = []string{"app", "user", "add", "foo"}
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("错误: %v\n", err)
	}
}

//func main() {
//	// 根命令
//	rootCmd := &cobra.Command{
//		Use:   "app",
//		Short: "示例应用",
//	}
//
//	//// 父命令 - 启用 TraverseChildren
//	//parentCmd := &cobra.Command{
//	//	Use:   "parent",
//	//	Short: "父命令",
//	//	Run: func(cmd *cobra.Command, args []string) {
//	//		fmt.Println("执行父命令")
//	//	},
//	//	// 关键设置:允许遍历子命令
//	//	TraverseChildren: true,
//	//}
//	//
//	//// 全局标志(父命令级别)
//	//var globalFlag string
//	//parentCmd.PersistentFlags().StringVarP(&globalFlag, "global", "g", "", "全局标志")
//	//
//	//// 子命令
//	//childCmd := &cobra.Command{
//	//	Use:   "child",
//	//	Short: "子命令",
//	//	Run: func(cmd *cobra.Command, args []string) {
//	//		fmt.Printf("执行子命令,全局标志值: %s\n", globalFlag)
//	//	},
//	//}
//	//
//	//// 子命令的本地标志
//	//var localFlag string
//	//childCmd.Flags().StringVarP(&localFlag, "local", "l", "", "本地标志")
//	//
//	//parentCmd.AddCommand(childCmd)
//	//rootCmd.AddCommand(parentCmd)
//	//
//	//// 测试不同的命令行参数
//	//fmt.Println("=== 示例 1: TraverseChildren = true ===")
//	//// 这种情况下,即使 --global 放在 child 后面也能正确解析
//	//os.Args = []string{"app", "parent", "child", "--global", "value1"}
//	//if err := rootCmd.Execute(); err != nil {
//	//	fmt.Println(err)
//	//}
//	//
//	//fmt.Println("\n=== 示例 2: 标志在子命令前 ===")
//	//os.Args = []string{"app", "parent", "--global", "value2", "child"}
//	//if err := rootCmd.Execute(); err != nil {
//	//	fmt.Println(err)
//	//}
//
//	// 对比:TraverseChildren = false 的情况
//	fmt.Println("\n=== 示例 3: TraverseChildren = false ===")
//	parentCmd2 := &cobra.Command{
//		Use:              "parent2",
//		Short:            "父命令(TraverseChildren=false)",
//		TraverseChildren: false, // 禁用遍历
//	}
//
//	var globalFlag2 string
//	parentCmd2.PersistentFlags().StringVarP(&globalFlag2, "global", "g", "", "全局标志")
//
//	childCmd2 := &cobra.Command{
//		Use:   "child",
//		Short: "子命令",
//		Run: func(cmd *cobra.Command, args []string) {
//			fmt.Printf("执行子命令,全局标志值: %s\n", globalFlag2)
//		},
//	}
//
//	parentCmd2.AddCommand(childCmd2)
//	rootCmd.AddCommand(parentCmd2)
//
//	// 这种情况下,标志必须放在正确的位置
//	os.Args = []string{"app", "parent2", "--global", "value3", "child"}
//	if err := rootCmd.Execute(); err != nil {
//		fmt.Println(err)
//	}
//
//	// 如果标志位置不对,会报错
//	fmt.Println("\n=== 示例 4: 错误的标志位置(TraverseChildren=false) ===")
//	os.Args = []string{"app", "parent2", "child", "--global", "value4"}
//	if err := rootCmd.Execute(); err != nil {
//		fmt.Printf("错误: %v\n", err)
//	}
//}
