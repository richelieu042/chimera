package rueidisKit

import (
	"github.com/redis/rueidis"
	"github.com/richelieu042/chimera/v3/src/component/database/nosql/redis/redisKit"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
)

// NewClient TODO: 没深入使用过.
/*
Instantiating a new Redis Client
	https://github.com/redis/rueidis?tab=readme-ov-file#instantiating-a-new-redis-client
*/
func NewClient(config *redisKit.Config) (client rueidis.Client, err error) {
	if err = config.Validate(); err != nil {
		return
	}

	option := rueidis.ClientOption{
		Username: config.UserName,
		Password: config.Password,
	}

	switch config.Mode {
	case redisKit.ModeSingle:
		option.InitAddress = []string{config.Single.Addr}
		option.SelectDB = config.Single.DB
	case redisKit.ModeMasterSlave:
		return nil, errorKit.Newf("mode(%s) is supported", config.Mode)
	case redisKit.ModeSentinel:
		option.InitAddress = config.Sentinel.Addrs
		option.SelectDB = config.Sentinel.DB
		option.Sentinel = rueidis.SentinelOption{
			MasterSet: config.Sentinel.MasterName,
			Username:  config.UserName,
			Password:  config.Password,
		}
	case redisKit.ModeCluster:
		option.InitAddress = config.Cluster.Addrs
		if config.Cluster.UseReplicasForReadOperations {
			// use replicas for read operations
			option.SendToReplicas = func(cmd rueidis.Completed) bool {
				return cmd.IsReadOnly()
			}
		} else {
			option.ShuffleInit = true
		}
	default:
		return nil, errorKit.Newf("mode(%s) is invalid", config.Mode)
	}
	return rueidis.NewClient(option)
}
