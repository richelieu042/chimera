package fileKit

import "os"

// IsDirEmpty 判断目录是否为空.
/*
!!!: Windows 下会忽略系统自动生成的文件（e.g. desktop.ini、Thumbs.db），不将其计入目录内容.

@param dirPath 目录路径（必须存在且为目录）
*/
func IsDirEmpty(dirPath string) (bool, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return false, err
	}
	for _, entry := range entries {
		if !isWindowsSystemFile(entry.Name()) {
			return false, nil
		}
	}
	return true, nil
}

func isWindowsSystemFile(name string) bool {
	switch name {
	case "desktop.ini", "Thumbs.db", "thumbs.db", "$RECYCLE.BIN", "System Volume Information":
		return true
	}
	return false
}
