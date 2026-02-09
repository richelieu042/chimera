package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu042/chimera/v3/src/consts/key"
	"github.com/richelieu042/chimera/v3/src/core/mapKit"
)

// SetTag
/*
@param pMsg	不能为nil
@param tag	可以为""
*/
func SetTag(pMsg *pulsar.ProducerMessage, tag string) {
	pMsg.Properties = mapKit.SetSafely(pMsg.Properties, key.Tag, tag)
}

// GetTag
/*
@param msg	不能为nil
@return 可能为""
*/
func GetTag(msg pulsar.Message) string {
	return mapKit.Get(msg.Properties(), key.Tag)
}
