//go:build go1.18 && amd64 && sonic && avx

package jsonKit

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/klauspost/cpuid/v2"
	"github.com/richelieu042/chimera/v3/src/core/cpuKit"
	"github.com/richelieu042/chimera/v3/src/core/osKit"
)

func init() {
	library = "bytedance/sonic"
	defaultApi = sonic.ConfigDefault
	stdApi = sonic.ConfigStd

	// 并非 amd64 CPU 就行了，还需要支持 avx指令集 等.（e.g.yozo某台amd64内网机就不行）
	if !cpuKit.AnyOfFeature(cpuid.AVX, cpuid.AVX2) {
		text := fmt.Sprintf("AVX isn't supported with os(%s) and arch(%s)", osKit.OS, osKit.ARCH)
		panic(text)
	}

	// 只要支持 AVX 或 AVX2 其中之一，sonic 就能跑（内部会自动挑更快的那套）
	// AVX2 是 AVX 的超集，支持 AVX2 的 CPU 一定支持 AVX
}
