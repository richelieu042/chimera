package redisKit

import (
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
)

var (
	NotSetupError = errorKit.Newf("haven’t been set up correctly")
)

var client *Client

// MustSetUp
/*
PS:
在使用go-redis连接Redis Cluster集群时，尽管通常仅需传入主节点的地址即可，但如果同时传入了所有主节点和从节点的地址，理论上不会造成任何问题。
客户端在连接时会首先尝试连接提供的地址列表中的任何一个节点，无论它是主节点还是从节点。一旦连接成功，客户端会通过与该节点交互获取整个集群的拓扑信息，包括所有主节点和从节点的详细情况。
*/
func MustSetUp(config *Config) {
	if err := SetUp(config); err != nil {
		console.Fatalf("Fail to set up, error: %+v", err)
	}
}

func SetUp(config *Config) (err error) {
	client, err = NewClient(config)
	return
}

// GetClient
/*
前提: 成功调用 SetUp() || MustSetUp().
*/
func GetClient() (*Client, error) {
	if client == nil {
		return nil, NotSetupError
	}
	return client, nil
}
