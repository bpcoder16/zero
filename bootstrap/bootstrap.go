package bootstrap

import (
	"context"
	"github.com/bpcoder16/zero/contrib/aliyun/oss"
	"github.com/bpcoder16/zero/core/log"
	"github.com/bpcoder16/zero/logit"
	"github.com/bpcoder16/zero/modules/appconfig"
	"github.com/bpcoder16/zero/modules/appconfig/env"
	"github.com/bpcoder16/zero/modules/zaplogger"
	"github.com/bpcoder16/zero/mysql"
	"github.com/bpcoder16/zero/redis"
	"io"
)

func MustInit(ctx context.Context, config *appconfig.AppConfig) {
	debugWriter, infoWriter, warnErrorFatalWriter := zaplogger.GetWriters(env.RootPath(), env.AppName(), env.AppName())
	initLoggers(ctx, config, debugWriter, infoWriter, warnErrorFatalWriter)
	if config.MySQLSupport {
		initMySQL(debugWriter, infoWriter, warnErrorFatalWriter)
	}
	if config.RedisSupport {
		initRedis(debugWriter, infoWriter, warnErrorFatalWriter)
	}
	if config.AliyunOSSSupport {
		initAliyunOSS()
	}
}

func initLoggers(_ context.Context, config *appconfig.AppConfig, debugWriter, infoWriter, warnErrorFatalWriter io.Writer) {
	logit.SetLogger(zaplogger.GetZapLogger(
		debugWriter, infoWriter, warnErrorFatalWriter,
		log.FileWithLineNumCaller(),
		log.FilterKey(config.FilterKeys...),
		log.FilterValue(config.FilterValues...),
		log.FilterLevel(func() log.Level {
			if env.RunMode() == env.RunModeRelease {
				return log.LevelInfo
			}
			return log.LevelDebug
		}()),
		//log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
		//	return false
		//}),
	))
}

func initMySQL(debugWriter, infoWriter, warnErrorFatalWriter io.Writer) {
	mysql.SetManager(env.RootPath()+"/conf/mysql.json", log.NewHelper(
		zaplogger.GetZapLogger(
			debugWriter, infoWriter, warnErrorFatalWriter,
			nil,
			log.FilterLevel(func() log.Level {
				if env.RunMode() == env.RunModeRelease {
					return log.LevelInfo
				}
				return log.LevelDebug
			}()),
			//log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
			//	return false
			//}),
		),
	))
}

func initRedis(debugWriter, infoWriter, warnErrorFatalWriter io.Writer) {
	redis.SetManager(env.RootPath()+"/conf/redis.json", log.NewHelper(
		zaplogger.GetZapLogger(
			debugWriter, infoWriter, warnErrorFatalWriter,
			log.FileWithLineNumCallerRedis(),
			log.FilterLevel(func() log.Level {
				if env.RunMode() == env.RunModeRelease {
					return log.LevelInfo
				}
				return log.LevelDebug
			}()),
			//log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
			//	return false
			//}),
		),
	))
}

func initAliyunOSS() {
	oss.InitAliyunOSS(env.RootPath() + "/conf/aliyun.json")
}
