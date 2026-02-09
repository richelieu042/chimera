package httpKit

import (
	"net/http"
	"net/url"
	"unicode"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/file/fileKit"
)

// DownloadFile 下载文件（文件路径）.
/*
参考: gin里面的 Context.File() 和 Context.FileAttachment().
*/
func DownloadFile(w http.ResponseWriter, r *http.Request, path, name string) error {
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return err
	}
	if strKit.IsEmpty(name) {
		name = fileKit.GetFileName(path)
	}

	if isFileNameASCII(name) {
		w.Header().Set("Content-Disposition", `attachment; filename="`+name+`"`)
	} else {
		w.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(name))
	}
	http.ServeFile(w, r, path)
	return nil
}

// DownloadFileContent 下载文件（文件内容）.
/*
参考: gin里面的 Context.File() 和 Context.FileAttachment().
支持: office文档、图片...
*/
func DownloadFileContent(w http.ResponseWriter, r *http.Request, content []byte, name string) error {
	if err := strKit.AssertNotEmpty(name, "name"); err != nil {
		return err
	}

	if isFileNameASCII(name) {
		w.Header().Set("Content-Disposition", `attachment; filename="`+name+`"`)
	} else {
		w.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(name))
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err := w.Write(content)
	return err
}

/*
https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
*/
func isFileNameASCII(fileName string) bool {
	for i := 0; i < len(fileName); i++ {
		if fileName[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
