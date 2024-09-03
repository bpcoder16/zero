package logger

import (
	"context"
	"github.com/bpcoder16/zero/contrib/file/filerotatelogs"
	"github.com/bpcoder16/zero/contrib/log/zap"
	"github.com/bpcoder16/zero/core/log"
	"github.com/bpcoder16/zero/logit"
	"path"
	"time"
)

const (
	DownstreamKey = "downstream"
	MessageKey    = "type"
	LogIdKey      = "logId"
)

// 	logger.InitLogger(
//		"/Users/bpcoder/git/one",
//		"one",
//		log.FilterKey("password"),
//		log.FilterValue("bpcoder"),
//		log.FilterLevel(log.LevelDebug),
//		log.FilterFunc(func(level log.Level, keyValues ...interface{}) bool {
//			return true
//		}),
//	)

func InitLogger(rootPath, appName string, opts ...log.FilterOption) {
	debugWriter := filerotatelogs.NewWriter(
		path.Join(rootPath, "log", appName, appName+".debug.log"),
		time.Duration(86400*30)*time.Second,
		time.Duration(3600)*time.Second,
	)
	infoWriter := filerotatelogs.NewWriter(
		path.Join(rootPath, "log", appName, appName+".info.log"),
		time.Duration(86400*30)*time.Second,
		time.Duration(3600)*time.Second,
	)
	warnErrorFatalWriter := filerotatelogs.NewWriter(
		path.Join(rootPath, "log", appName, appName+".wf.log"),
		time.Duration(86400*30)*time.Second,
		time.Duration(3600)*time.Second,
	)

	logit.SetLogger(
		log.NewFilter(
			log.With(
				zap.NewLogger(debugWriter, infoWriter, warnErrorFatalWriter, MessageKey),
				LogIdKey,
				func() log.Valuer {
					return func(ctx context.Context) interface{} {
						logId := ctx.Value(LogIdKey)
						if logId == nil {
							return "NotSetLogId"
						}
						return logId
					}
				}(),
				MessageKey,
				func() log.Valuer {
					return func(ctx context.Context) interface{} {
						msg := ctx.Value(MessageKey)
						if msg == nil {
							return "Default"
						}
						return msg
					}
				}(),
				DownstreamKey,
				func() log.Valuer {
					return func(ctx context.Context) interface{} {
						msg := ctx.Value(DownstreamKey)
						if msg == nil {
							return "None"
						}
						return msg
					}
				}(),
				"caller",
				log.Caller(6),
			),
			opts...,
		),
	)
}
