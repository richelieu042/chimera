package redisKit

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/richelieu042/chimera/v3/src/core/error/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

type (
	Client struct {
		mode Mode

		// universalClient go-redis的客户端
		universalClient redis.UniversalClient

		// keyPrefix 所有key的前缀（可以为""）
		/*
			在go-redis库中，你可以通过在每个键前面添加一个字符串来设置键的前缀。但是，go-redis库本身并没有提供直接设置键前缀的功能。
			如果你需要在所有键前面添加一个公共的前缀，你可能需要自己实现这个功能。一种可能的方法是: 创建一个包装器函数，该函数接受一个键作为参数，然后返回一个带有前缀的键。
		*/
		keyPrefix string
	}
)

func (client *Client) GetPrefix() string {
	return client.keyPrefix
}

func (client *Client) GetKeyWithPrefix(key string) string {
	return client.keyPrefix + key
}

func (client *Client) Close() error {
	if client != nil && client.universalClient != nil {
		return client.universalClient.Close()
	}
	return nil
}

func (client *Client) GetMode() Mode {
	return client.mode
}

// GetUniversalClient 返回go-redis客户端
func (client *Client) GetUniversalClient() redis.UniversalClient {
	return client.universalClient
}

// NewClient 新建一个go-redis客户端（内置连接池，调用方无需额外考虑并发问题）.
/*
!!!: 每一个命令都会重新取得一个连接，执行后立即回收，而且回收到资源池的顺序类似于堆. https://www.cnblogs.com/yangqi7/p/13289232.html

连接哨兵集群的demo: https://blog.csdn.net/supery071/article/details/109491404

@return	(1) 两个返回值，必定有一个为nil，另一个非nil；
		(2) Cluster模式下，第1个返回值的类型: *redis.ClusterClient.
*/
func NewClient(config *Config) (client *Client, err error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	var opts *redis.UniversalOptions
	switch config.Mode {
	case ModeSingle:
		opts, err = newSingleOptions(config)
	case ModeMasterSlave:
		opts, err = newMasterSlaveOptions(config)
	case ModeSentinel:
		opts, err = newSentinelOptions(config)
	case ModeCluster:
		opts, err = newClusterOptions(config)
	default:
		err = errorKit.Newf("mode(%s) is invalid", config.Mode)
	}
	if err != nil {
		return
	}

	//opts.OnConnect = func(ctx context.Context, conn *redis.Conn) error {
	//	fmt.Printf("conn: %v", conn)
	//	return nil
	//}

	universalClient := redis.NewUniversalClient(opts)
	client = &Client{
		mode:            config.Mode,
		universalClient: universalClient,
		keyPrefix:       config.KeyPrefix,
	}
	defer func() {
		if err != nil {
			_ = client.Close()
			client = nil
		}
	}()

	// 简单测试是否Redis服务可用
	pingTimeout := time.Second * 3
	ctx, cancel := context.WithTimeout(context.TODO(), pingTimeout)
	defer cancel()
	str, err := client.Ping(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			err = errorKit.Newf("initial ping timeout(%s)", pingTimeout)
		} else {
			err = errorKit.Wrapf(err, "initial ping fails")
		}
		return
	}
	if !strKit.EqualsIgnoreCase(str, "PONG") {
		err = errorKit.Newf("result(%s) of initial ping is invalid", str)
		return
	}
	return
}

func newBaseOptions(config *Config) *redis.UniversalOptions {
	return &redis.UniversalOptions{
		Username: config.UserName,
		Password: config.Password,
	}
}

// newSingleOptions 单点模式
func newSingleOptions(config *Config) (*redis.UniversalOptions, error) {
	c := config.Single

	opts := newBaseOptions(config)
	opts.Addrs = []string{c.Addr}
	opts.DB = c.DB
	return opts, nil
}

// newMasterSlaveOptions 主从模式
func newMasterSlaveOptions(config *Config) (*redis.UniversalOptions, error) {
	return nil, errorKit.Newf("mode(%s) is unsupported now", config.Mode)
}

// newSentinelOptions 哨兵模式
func newSentinelOptions(config *Config) (*redis.UniversalOptions, error) {
	c := config.Sentinel

	opts := newBaseOptions(config)
	opts.MasterName = strKit.EmptyToDefault(c.MasterName, DefaultMasterName, true)
	opts.Addrs = c.Addrs
	opts.DB = c.DB
	return opts, nil
}

// newClusterOptions cluster模式
func newClusterOptions(config *Config) (*redis.UniversalOptions, error) {
	opts := newBaseOptions(config)
	opts.Addrs = config.Cluster.Addrs

	/*
		Enables read-only commands on slave nodes.

		在 ClusterOptions 中设置 ReadOnly 为 false，这样客户端即使在执行只读操作时也不会尝试连接从节点。
		这有助于保持整体逻辑清晰，即所有读、写操作都直接发送到主节点，而不会由于意外的读取配置而间接导致写入从节点。
	*/
	//opts.ReadOnly = false
	opts.ReadOnly = config.Cluster.UseReplicasForReadOperations

	/*
		RouteByLatency 和 RouteRandomly 互斥.
	*/
	opts.RouteByLatency = true
	opts.RouteRandomly = false

	return opts, nil
}
