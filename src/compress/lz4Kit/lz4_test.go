package lz4Kit

import (
	"fmt"
	"testing"

	"github.com/richelieu042/chimera/v3/src/core/bytesKit"
)

func TestCompressAndDecompress(t *testing.T) {
	data := []byte(`===
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
---`)

	compressedData, err := Compress(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(compressedData))
	fmt.Println(len(compressedData))

	fmt.Println("======")

	decompressData, err := Decompress(compressedData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decompressData))
	fmt.Println(len(decompressData))

	if !bytesKit.Equals(data, decompressData) {
		panic("not equal")
	}
}
