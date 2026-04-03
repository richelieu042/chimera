package fileKit

import "os"

// IsDirEmpty 判断目录是否为空.
/*
!!!: macOS 下会忽略系统自动生成的文件（e.g. .DS_Store），不将其计入目录内容.

@param dirPath 目录路径（必须存在且为目录）
*/
func IsDirEmpty(dirPath string) (bool, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return false, err
	}
	for _, entry := range entries {
		if !isMacOSSystemFile(entry.Name()) {
			return false, nil
		}
	}
	return true, nil
}

func isMacOSSystemFile(name string) bool {
	switch name {
	case ".DS_Store", "._.DS_Store", ".Spotlight-V100", ".Trashes", ".fseventsd":
		return true
	}
	return false
}
