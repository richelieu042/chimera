//go:build go1.17 && amd64 && sonic && avx

package jsonKit

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/klauspost/cpuid/v2"
	"github.com/richelieu-yang/chimera/v3/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
)

func init() {
	library = "bytedance/sonic"
	defaultApi = sonic.ConfigDefault
	stdApi = sonic.ConfigStd

	// 并非 amd64 CPU 就行了，还需要支持 avx指令集 等.（e.g.yozo某台amd64内网机就不行）
	if !cpuKit.HasFeature(cpuid.AVX) {
		text := fmt.Sprintf("AVX isn't supported with os(%s) and arch(%s)", osKit.OS, osKit.ARCH)
		panic(text)
		return
	}
}
