package redisKit

import (
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// IsConsumerGroupNameAlreadyExistError 适用场景: 创建Consumer group（XGroupCreate || XGroupCreateMkStream）时，返回error.
/*
TODO: 涉及 github.com/redis/go-redis/v9 源码， 后续看有没有好的解决方法.

@return true: 错误是由 "指定group已经存在（无需重复创建）" 导致的.
*/
func IsConsumerGroupNameAlreadyExistError(err error) bool {
	if err == nil {
		return false
	}

	err = errorKit.Cause(err)
	return strKit.ContainsIgnoreCase(err.Error(), "BUSYGROUP Consumer Group name already exists")
}

// IsNoStreamOrNoGroupError 适用场景: Consumer读取消息时，返回error.
/*
TODO: 涉及 github.com/redis/go-redis/v9 源码， 后续看有没有好的解决方法.

@return true: 错误是由 "指定stream不存在" 或 "指定stream存在，但指定group不存在" 导致的.
*/
func IsNoStreamOrNoGroupError(err error) bool {
	if err == nil {
		return false
	}

	err = errorKit.Cause(err)
	return strKit.ContainsIgnoreCase(err.Error(), "NOGROUP No such key") && strKit.ContainsIgnoreCase(err.Error(), "or consumer group")
}
