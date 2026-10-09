package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	// Remove
	/*
		Deprecated: please use RemoveFile or RemoveAll for explicit usage instead.
	*/
	Remove = gfile.Remove

	// RemoveFile 删除单个文件或空目录，不会递归删除非空目录。
	// 目标不存在、路径为空字符串或目录非空时返回错误。
	// 错误由 gfile 包装，判断目标不存在应使用 errors.Is(err, os.ErrNotExist)。
	// 明确只删除单个文件或空目录时使用此函数，避免误删整个目录树。
	RemoveFile = gfile.RemoveFile

	// RemoveAll 递归删除目标文件或目录及其全部内容，目录本身也会被删除。
	// 目标不存在或路径为空字符串时不执行删除，返回 nil。
	// 遇到错误时仍尽可能删除其他内容，最终返回遇到的第一个错误；因此可能部分删除。
	// 错误由 gfile 包装，可使用 errors.Is 判断底层错误。
	// 仅在需要删除整个目录树时使用；若需保留目录本身，请使用 EmptyDir。
	RemoveAll = gfile.RemoveAll
)
