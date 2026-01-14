package fileKit

import (
	"os"

	"github.com/duke-git/lancet/v2/fileutil"
)

// WriteToFile 将数据（[]byte, 字节流）写到文件中.
/*
@param filePath 目标文件的路径
				(1) 不存在的话，会创建一个新的文件（也会自动创建父目录）;
				(2) 存在且是个文件的话，会 "覆盖" 掉旧的（并不会加到该文件的最后面）.
*/
func WriteToFile(filePath string, content []byte, permArgs ...os.FileMode) error {
	var perm os.FileMode
	if len(permArgs) == 0 {
		perm = 0644 // 默认值
	} else {
		perm = permArgs[0]
	}

	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return err
	}
	if err := MkParentDirs(filePath); err != nil {
		return err
	}

	//return fileutil.WriteBytesToFile(filePath, content)
	return os.WriteFile(filePath, content, perm)
}

// WriteStringToFile 将数据（字符串）写到文件中.
/*
PS: 权限为0644.

@param filePath 目标文件的路径，	(1) 不存在的话，会创建一个新的文件;
							 	(2) 存在且是个文件的话，由 传参append 决定.
@param append 	true: 	追加在最后面
				false: 	覆盖
*/
func WriteStringToFile(filePath string, content string, append bool) error {
	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return err
	}
	if err := MkParentDirs(filePath); err != nil {
		return err
	}

	return fileutil.WriteStringToFile(filePath, content, append)
}
