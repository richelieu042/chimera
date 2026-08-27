package cpuKit

import (
	"runtime"

	"github.com/klauspost/cpuid/v2"
)

// GetCpuNumber returns the number of logical CPUs usable by the current process.
var GetCpuNumber func() int = runtime.NumCPU

// HasFeature CPU是否支持特定指令集？
/*
@param id e.g. cpuid.AVX
*/
func HasFeature(id cpuid.FeatureID) bool {
	return cpuid.CPU.Has(id)
}

// InVirtualMachine 是否在虚拟机中？
var InVirtualMachine func() bool = cpuid.CPU.VM

func GetVendorID() cpuid.Vendor {
	return cpuid.CPU.VendorID
}

// GetVendorString CPU供应商
/*
@return e.g."Apple"
*/
func GetVendorString() string {
	return cpuid.CPU.VendorString
}

// GetBrandName CPU品牌名称
/*
@return e.g."Apple M1 Pro"
*/
func GetBrandName() string {
	return cpuid.CPU.BrandName
}

func GetPhysicalCores() int {
	return cpuid.CPU.PhysicalCores
}

func GetThreadsPerCore() int {
	return cpuid.CPU.ThreadsPerCore
}

func GetLogicalCores() int {
	return cpuid.CPU.LogicalCores
}

// GetFeatureSet 获取CPU支持的指令集s.
/*
Linux命令: cat /proc/cpuinfo
*/
var GetFeatureSet func() []string = cpuid.CPU.FeatureSet

func GetFamily() int {
	return cpuid.CPU.Family
}

func GetModel() int {
	return cpuid.CPU.Model
}

func GetFrequency() int64 {
	return cpuid.CPU.Hz
}
