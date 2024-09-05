package nonblockinglock

import (
	"context"
	"github.com/bpcoder16/zero/logit"
	"github.com/bpcoder16/zero/redis"
	"strconv"
	"time"
)

func RedisLock(ctx context.Context, lockName string, deadLockExpireTime time.Duration) (result bool, redisUnlockFunc func(ctx context.Context)) {
	redisUnlockFunc = func(unlockCtx context.Context) {
		redis.DefaultClient().Del(unlockCtx, lockName)
	}

	timeNow := time.Now()
	expireTimeStr := strconv.Itoa(int(timeNow.Add(deadLockExpireTime).Unix()))
	var err error
	result, err = redis.DefaultClient().SetNX(ctx, lockName, expireTimeStr, deadLockExpireTime).Result()

	if err != nil {
		logit.Context(ctx).WarnW("non-blocking-lock.RedisLock", err.Error())
		return
	}

	// 防止死锁
	if !result {
		var errRedis error
		if expireTimeStr, errRedis = redis.DefaultClient().Get(ctx, lockName).Result(); errRedis == nil {
			if expireTimeRedis, errStr := strconv.Atoi(expireTimeStr); errStr == nil {
				if timeNow.Unix() > int64(expireTimeRedis+5) {
					redis.DefaultClient().Del(ctx, lockName)
				}
			} else {
				redis.DefaultClient().Del(ctx, lockName)
			}
		}
	}

	return
}
