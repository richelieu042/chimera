package fsnotifyKit

import "github.com/gogf/gf/v2/os/gfsnotify"

var (
	// Add 添加监听，监控指定文件/目录的改变，如文件的增加、删除、修改、重命名等操作.
	/*
		@param path 文件 || 目录
	*/
	Add func(path string, callbackFunc func(event *gfsnotify.Event), option ...gfsnotify.WatchOption) (callback *gfsnotify.Callback, err error) = gfsnotify.Add

	// Remove 移除监听.
	Remove func(path string) error = gfsnotify.Remove

	// RemoveCallback 移除监听.
	RemoveCallback func(callbackId int) error = gfsnotify.RemoveCallback
)
