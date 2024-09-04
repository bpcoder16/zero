package bootstrap

import (
	"context"
	"github.com/bpcoder16/zero/core/log"
	"github.com/bpcoder16/zero/logit"
	"github.com/bpcoder16/zero/modules/appconfig"
	"github.com/bpcoder16/zero/modules/appconfig/env"
	"github.com/bpcoder16/zero/modules/zaplogger"
	"github.com/bpcoder16/zero/mysql"
)

func MustInit(ctx context.Context, config *appconfig.AppConfig) {
	initLoggers(ctx, config)
	if config.MySQLSupport {
		initMySQL()
	}
}

func initLoggers(_ context.Context, config *appconfig.AppConfig) {
	logit.SetLogger(zaplogger.GetZapLogger(
		env.RootPath(),
		env.AppName(),
		env.AppName(),
		true,
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

func initMySQL() {
	mysql.SetManager(env.RootPath()+"/conf/mysql.json", log.NewHelper(
		zaplogger.GetZapLogger(
			env.RootPath(),
			env.AppName(),
			env.AppName(),
			false,
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
