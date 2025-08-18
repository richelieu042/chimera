package gtypeKit

import "github.com/gogf/gf/v2/container/gtype"

var (
	NewUint func(value ...uint) *gtype.Uint = gtype.NewUint

	NewUint32 func(value ...uint32) *gtype.Uint32 = gtype.NewUint32

	NewUint64 func(value ...uint64) *gtype.Uint64 = gtype.NewUint64
)
