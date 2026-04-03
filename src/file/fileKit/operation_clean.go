package fileKit

import (
	"os"
	"path/filepath"
	"time"

	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
	"go.uber.org/zap"
)

// Clean 递归清理路径下满足所有 predicate 条件的文件，并删除空目录.
/*
@param path 		文件或目录的路径（如果不存在，将返回nil）
@param predicates	（1）所有 predicate 返回 true 时才删除文件或空目录（AND 逻辑）
					（2）如果不传，将整个删除
*/
func Clean(log *zap.Logger, path string, predicates ...Predicate) error {
	if log == nil {
		log = zap.NewNop() // 丢弃输出
	}

	if len(predicates) == 0 {
		return RemoveAll(path)
	}

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return errKit.Wrapf(err, "fail to stat path(%s)", path)
	}

	if info.IsDir() {
		return cleanDirectory(log, path, predicates)
	}
	return cleanFile(log, path, info, predicates)
}

func cleanDirectory(logger *zap.Logger, dirPath string, predicates []Predicate) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return errKit.Wrapf(err, "fail to read dir(%s)", dirPath)
	}

	for _, entry := range entries {
		entryPath := filepath.Join(dirPath, entry.Name())
		if entry.IsDir() {
			if err := cleanDirectory(logger, entryPath, predicates); err != nil {
				return err
			}
		} else {
			info, err := entry.Info()
			if err != nil {
				return errKit.Wrapf(err, "fail to get info of entry(%s)", entryPath)
			}
			if err := cleanFile(logger, entryPath, info, predicates); err != nil {
				return err
			}
		}
	}

	// 子项处理完毕后，若目录为空则经过 predicates 判断后再删除
	remaining, err := os.ReadDir(dirPath)
	if err != nil {
		return errKit.Wrapf(err, "fail to re-read dir(%s)", dirPath)
	}
	if len(remaining) == 0 {
		info, err := os.Stat(dirPath)
		if err != nil {
			return errKit.Wrapf(err, "fail to stat dir(%s)", dirPath)
		}

		if !canDelete(info, predicates) {
			return nil // predicates 不允许删除
		}

		if err := os.Remove(dirPath); err != nil {
			return errKit.Wrapf(err, "fail to remove dir(%s)", dirPath)
		}
		logger.Debug("removed empty dir", zap.String("path", dirPath))
	}

	return nil
}

func cleanFile(logger *zap.Logger, filePath string, info os.FileInfo, predicates []Predicate) error {
	if !canDelete(info, predicates) {
		return nil // predicates 不允许删除
	}

	if err := os.Remove(filePath); err != nil {
		return errKit.Wrapf(err, "fail to remove file(%s)", filePath)
	}
	logger.Debug("removed file", zap.String("path", filePath))

	return nil
}

func canDelete(info os.FileInfo, predicates []Predicate) bool {
	for _, predicate := range predicates {
		if !predicate(info) {
			return false
		}
	}
	return true
}

// Predicate 判断文件是否应该被删除的回调函数类型
type Predicate func(info os.FileInfo) bool

// OlderThanModTime 内置 predicate：文件修改时间超过指定时长
func OlderThanModTime(d time.Duration) Predicate {
	return func(info os.FileInfo) bool {
		return time.Since(info.ModTime()) >= d
	}
}
