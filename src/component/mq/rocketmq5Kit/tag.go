package rocketmq5Kit

import "github.com/richelieu042/chimera/v3/src/core/sliceKit"

func MixTags(tags ...string) string {
	return sliceKit.Join(tags, "||")
}

func GetTagString(tag *string) string {
	if tag == nil {
		return "<nil>"
	}
	return *tag
}
