package zaplogger

import (
	"context"
	"github.com/bpcoder16/zero/contrib/file/filerotatelogs"
	"github.com/bpcoder16/zero/contrib/log/zap"
	"github.com/bpcoder16/zero/core/log"
	"path"
	"time"
)

// 	zaplogger.GetZapLogger(
//		"/Users/bpcoder/git/one",
//		"one",
//		log.FilterKey("password"),
//		log.FilterValue("bpcoder"),
//		log.FilterLevel(log.LevelDebug),
//		log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
//			return true
//		}),
//	)

func GetZapLogger(rootPath, appName, logName string, isSetCaller bool, opts ...log.FilterOption) log.Logger {
	debugWriter := filerotatelogs.NewWriter(
		path.Join(rootPath, "log", appName, logName+".debug.log"),
		time.Duration(86400*30)*time.Second,
		time.Duration(3600)*time.Second,
	)
	infoWriter := filerotatelogs.NewWriter(
		path.Join(rootPath, "log", appName, logName+".info.log"),
		time.Duration(86400*30)*time.Second,
		time.Duration(3600)*time.Second,
	)
	warnErrorFatalWriter := filerotatelogs.NewWriter(
		path.Join(rootPath, "log", appName, logName+".wf.log"),
		time.Duration(86400*30)*time.Second,
		time.Duration(3600)*time.Second,
	)

	kv := make([]interface{}, 0, 8)
	kv = append(kv, log.DefaultLogIdKey,
		func() log.Valuer {
			return func(ctx context.Context) interface{} {
				logId := ctx.Value(log.DefaultLogIdKey)
				if logId == nil {
					return "NotSetLogId"
				}
				return logId
			}
		}(),
		log.DefaultMessageKey,
		func() log.Valuer {
			return func(ctx context.Context) interface{} {
				msg := ctx.Value(log.DefaultMessageKey)
				if msg == nil {
					return "Default"
				}
				return msg
			}
		}(),
		log.DefaultDownstreamKey,
		func() log.Valuer {
			return func(ctx context.Context) interface{} {
				msg := ctx.Value(log.DefaultDownstreamKey)
				if msg == nil {
					return "None"
				}
				return msg
			}
		}(),
	)
	if isSetCaller {
		kv = append(kv, "caller", log.FileWithLineNumCaller())
	}

	return log.NewFilter(
		log.With(
			zap.NewLogger(debugWriter, infoWriter, warnErrorFatalWriter),
			kv...,
		),
		opts...,
	)
}
