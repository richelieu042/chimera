package asynqKit

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu042/chimera/v3/src/component/database/nosql/redis/redisKit"
)

type redisConnOptImpl struct {
	asynq.RedisConnOpt

	uc redis.UniversalClient
}

func (impl *redisConnOptImpl) MakeRedisClient() interface{} {
	return impl.uc
}

// NewAsynqClient
/*
!!!: 必需先setup redis.
*/
func NewAsynqClient() (*asynq.Client, error) {
	client, err := redisKit.GetClient()
	if err != nil {
		return nil, err
	}
	impl := &redisConnOptImpl{
		uc: client.GetUniversalClient(),
	}
	return asynq.NewClient(impl), nil
}
