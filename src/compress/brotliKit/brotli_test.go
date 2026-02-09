package brotliKit

import (
	"fmt"
	"testing"

	"github.com/richelieu042/chimera/v3/src/core/bytesKit"
)

func TestEntire(t *testing.T) {
	data := []byte(`===
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
---`)
	fmt.Println(len(data))

	// Compress
	compressed, err := Compress(data, WithLevel(LevelDefaultCompression))
	if err != nil {
		panic(err)
	}
	fmt.Println("compressed:", string(compressed))
	fmt.Println(len(compressed))
	fmt.Println("---")

	// Decompress
	decompressed, err := Decompress(compressed)
	if err != nil {
		panic(err)
	}
	fmt.Println("decompressed:", string(decompressed))
	fmt.Println(len(decompressed))

	if !bytesKit.Equals(data, decompressed) {
		panic("not equals")
	}
}
