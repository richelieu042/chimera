package idKit

import (
	"github.com/google/uuid"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// NewUUID UUID v4
/*
推荐使用: NewXid or NewULID.

PS:
(1) 重复概率非常低，不建议用作分布式唯一id.
(2) 格式（5组）: {8}-{4}-{4}-{4}-{12}
(3) 长度: 36

e.g.
	() => "936eff5f-97c6-4f8b-b26d-9bab1f65ff55"
*/
func NewUUID() string {
	return uuid.New().String()
}

// NewSimpleUUID UUIDv4，去掉了其中所有"-"
/*
@return 长度32

e.g.
	() => "415ef754dc174b888b186873e093ced1"
*/
func NewSimpleUUID() string {
	return strKit.ReplaceAll(NewUUID(), "-", "")
}
