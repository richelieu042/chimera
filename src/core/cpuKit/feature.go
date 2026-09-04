package cpuKit

import "github.com/klauspost/cpuid/v2"

// HasFeature CPU是否支持特定指令集？
/*
@param id e.g. cpuid.AVX
*/
func HasFeature(id cpuid.FeatureID) bool {
	return cpuid.CPU.Has(id)
}

// AnyOfFeatures （OR逻辑）CPU是否支持任意一个指令集？
func AnyOfFeatures(ids ...cpuid.FeatureID) bool {
	return cpuid.CPU.AnyOf(ids...)
}

// SupportFeatures （AND逻辑）检测当前 CPU 是否支持一个或多个指定的特性(如 AVX2、SSE4.2、AES-NI 等),只有全部指定特性都支持时才返回 true
func SupportFeatures(ids ...cpuid.FeatureID) bool {
	return cpuid.CPU.Supports(ids...)
}
