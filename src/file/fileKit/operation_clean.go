package fileKit

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/richelieu042/chimera/v3/src/log/commonLogKit"
	"go.uber.org/zap"
)

// Predicate 判断文件是否应该被删除的回调函数类型
type Predicate func(info os.FileInfo) bool

// Clean 递归清理路径下满足所有 predicate 条件的文件，并删除空目录
/*
@param path 		文件或目录的路径（如果不存在，将返回nil）
@param predicates	（1）所有 predicate 返回 true 时才删除文件（AND 逻辑）
					（2）如果不传，将整个删除
*/
func Clean(log commonLogKit.Logger, path string, predicates ...Predicate) error {
	if log == nil {
		log = zap.NewNop().Sugar() // 丢弃输出
	}

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("stat %s: %w", path, err)
	}

	if info.IsDir() {
		return cleanDirectory(path, predicates)
	}
	return cleanFile(path, info, predicates)
}

func cleanDirectory(dirPath string, predicates []Predicate) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("read dir %s: %w", dirPath, err)
	}

	for _, entry := range entries {
		entryPath := filepath.Join(dirPath, entry.Name())
		if entry.IsDir() {
			if err := cleanDirectory(entryPath, predicates); err != nil {
				return err
			}
		} else {
			info, err := entry.Info()
			if err != nil {
				return fmt.Errorf("get info %s: %w", entryPath, err)
			}
			if err := cleanFile(entryPath, info, predicates); err != nil {
				return err
			}
		}
	}

	// 子项处理完毕后，若目录为空则删除
	remaining, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("re-read dir %s: %w", dirPath, err)
	}
	if len(remaining) == 0 {
		if err := os.Remove(dirPath); err != nil {
			return fmt.Errorf("remove empty dir %s: %w", dirPath, err)
		}
		fmt.Printf("removed empty dir : %s\n", dirPath)
	}

	return nil
}

func cleanFile(filePath string, info os.FileInfo, predicates []Predicate) error {
	for _, predicate := range predicates {
		if !predicate(info) {
			return nil // 任一 predicate 返回 false，跳过删除
		}
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("remove file %s: %w", filePath, err)
	}
	fmt.Printf("removed file: %s\n", filePath)
	return nil
}

// OlderThan 内置 predicate：文件修改时间超过指定时长
func OlderThan(d time.Duration) Predicate {
	return func(info os.FileInfo) bool {
		return time.Since(info.ModTime()) >= d
	}
}

// SmallerThan 内置 predicate：文件大小小于指定字节数
func SmallerThan(bytes int64) Predicate {
	return func(info os.FileInfo) bool {
		return info.Size() < bytes
	}
}

// LargerThan 内置 predicate：文件大小大于指定字节数
func LargerThan(bytes int64) Predicate {
	return func(info os.FileInfo) bool {
		return info.Size() > bytes
	}
}
