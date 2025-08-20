package fileKit

import "os"

var (
	// OpenRoot
	/*
		https://cloud.tencent.com/developer/article/2522181
		https://cloud.tencent.com/developer/article/2502915

		os.Root 可以锁定工作目录。 使用户无法打开目录外的文件，例如 ../../../etc/passwd 。
		可以算一种 安全保护，最重要的是 强制约束用户， 限制用户行为， 检查计划外的使用逻辑 。 免得和煞笔瞎掰扯， 浪费时间。
	*/
	OpenRoot func(dir string) (*os.Root, error) = os.OpenRoot

	OpenInRoot func(dir, name string) (*os.File, error) = os.OpenInRoot
)
