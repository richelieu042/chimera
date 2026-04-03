package fileKit

import "os"

// IsDirEmpty 判断目录是否为空.
/*
@param dirPath 目录路径（必须存在且为目录）
*/
func IsDirEmpty(dirPath string) (bool, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return false, err
	}
	return len(entries) == 0, nil
}
