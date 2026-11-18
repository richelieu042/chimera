package pdfKit

import (
	"context"
	"fmt"
	"testing"
)

func TestMergeCreateFile(t *testing.T) {
	files := []string{"/Users/richelieu/Desktop/a.pdf", "/Users/richelieu/Desktop/b.pdf"}
	output := "_test-merge-create-file.pdf"

	err := MergeCreateFile(context.Background(), files, output, false, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully.")
}

func TestMergeAppendFile(t *testing.T) {
	files := []string{"/Users/richelieu/Desktop/a.pdf", "/Users/richelieu/Desktop/b.pdf"}
	output := "_test-merge-append-file.pdf"

	err := MergeAppendFile(context.Background(), files, output, false, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully.")
}
