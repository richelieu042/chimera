package browserKit

import (
	"io"

	"github.com/pkg/browser"
)

var (
	// OpenURL 让默认浏览器打开网址.
	/*
		Open a new browser window pointing to url.
	*/
	OpenURL func(url string) error = browser.OpenURL

	// OpenFileOrDirectory 用对应工具打开文件（或目录）.
	/*
		Open new browser window for the file path.
	*/
	OpenFileOrDirectory func(path string) error = browser.OpenFile

	// OpenReader
	/*
		Consume the contents of r and presents the results in a new browser window.
	*/
	OpenReader func(r io.Reader) error = browser.OpenReader
)

// SetOutputs
/*
@param stdout 可以为nil（默认: os.Stdout）
@param stderr 可以为nil（默认: os.Stderr）
*/
func SetOutputs(stdout, stderr io.Writer) {
	if stdout != nil {
		browser.Stdout = stdout
	}
	if stderr != nil {
		browser.Stderr = stderr
	}
}
