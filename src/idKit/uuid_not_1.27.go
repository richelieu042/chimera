//go:build !go1.27

package idKit

import (
	"github.com/google/uuid"
)

// NewUuid UUID v4
/*
推荐使用: NewXid or NewUlid.

PS:
(1) 重复概率非常低，不建议用作分布式唯一id.
(2) 格式（5组）: {8}-{4}-{4}-{4}-{12}
(3) 长度: 36

e.g.
	() => "936eff5f-97c6-4f8b-b26d-9bab1f65ff55"
*/
func NewUuid() string {
	return uuid.New().String()
}

func NewUuidV7() string {
	return uuid.Must(uuid.NewV7()).String()
}
