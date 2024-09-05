package redis

import (
	"github.com/bpcoder16/zero/contrib/goredis"
	"github.com/bpcoder16/zero/core/log"
	"github.com/redis/go-redis/v9"
)

var defaultRedisManager *goredis.RedisManager

func SetManager(configPath string, logger *log.Helper) {
	defaultRedisManager = goredis.NewRedisManager(configPath, logger)
}

func DefaultClient() *redis.Client {
	return defaultRedisManager.Client()
}
