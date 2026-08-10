package atomicKit

import (
	"time"
	"unsafe"
	"go.uber.org/atomic"
)

var (
	NewBool func(val bool) *atomic.Bool = atomic.NewBool

	NewFloat32 func(val float32) *atomic.Float32 = atomic.NewFloat32
	NewFloat64 func(val float64) *atomic.Float64 = atomic.NewFloat64

	NewInt32  func(val int32) *atomic.Int32   = atomic.NewInt32
	NewInt64  func(val int64) *atomic.Int64   = atomic.NewInt64
	NewUint32 func(val uint32) *atomic.Uint32 = atomic.NewUint32
	NewUint64 func(val uint64) *atomic.Uint64 = atomic.NewUint64

	NewString func(val string) *atomic.String = atomic.NewString

	NewTime     func(val time.Time) *atomic.Time         = atomic.NewTime
	NewDuration func(val time.Duration) *atomic.Duration = atomic.NewDuration

	NewUintptr       func(val uintptr) *atomic.Uintptr              = atomic.NewUintptr
	NewUnsafePointer func(val unsafe.Pointer) *atomic.UnsafePointer = atomic.NewUnsafePointer
	NewError         func(val error) *atomic.Error                  = atomic.NewError
)

func NewPointer[T any](v *T) *atomic.Pointer[T] {
	return atomic.NewPointer(v)
}
