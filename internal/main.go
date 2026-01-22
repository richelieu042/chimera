package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	A bool
	B bool
	C bool
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
			fmt.Println("args", args)
			fmt.Println("A", A)
			fmt.Println("B", B)
			fmt.Println("C", C)
		},
	}

	rootCmd.Flags().BoolVar(&A, "a", false, "usage of a")
	userCmd.Flags().BoolVar(&B, "b", false, "usage of b")
	addCmd.Flags().BoolVar(&C, "c", false, "usage of c")

	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(addCmd)

	os.Args = []string{"app", "--a", "user" /*"--b",*/, "add", "--c"}
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("错误: %v\n", err)
	}
}
