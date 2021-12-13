package aggregate

import (
	"time"

	"github.com/airdb/sls-bbhj/pkg/cache"
)

type RedisAggr interface {
	Get(key string) (string, error)
	Set(key string) error
}

type redisAggr struct {
	redis *cache.Redis
}

func newRedis() *redisAggr {
	return &redisAggr{redis: cache.NewRedis()}
}

func (a *redisAggr) Get(key string) (string, error) {
	return a.redis.Get(key)
}

func (a *redisAggr) Set(key, value string, expires time.Duration) error {
	return a.redis.Set(key, value, expires)
}
