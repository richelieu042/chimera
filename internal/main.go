package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // 会在此处卡住
	text := scanner.Text()
	fmt.Printf("input: [%s]\n", text)
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("=== end ===")
}
