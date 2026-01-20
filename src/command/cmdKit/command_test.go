package cmdKit

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestLookPath(t *testing.T) {
	path, err := LookPath("java")
	if err != nil {
		panic(err)
	}
	fmt.Println(path) // /usr/bin/java
}

// 命令因超时被取消
func TestNewCommand(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := NewCommand(ctx, "sleep", []string{"5"}).Run(); err != nil {
		fmt.Println("got error:", err) // got error: signal: killed
	}
}

func TestNewCommand1(t *testing.T) {
	cmd := NewCommand(context.TODO(), "echo", nil)
	data, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func TestExecute(t *testing.T) {
	data, err := Run(context.TODO(), "java", "-version")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output:\n%s\n", string(data))
}
