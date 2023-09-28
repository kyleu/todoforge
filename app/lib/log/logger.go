// Package log - Content managed by Project Forge, see [projectforge.md] for details.
package log

import (
	"runtime"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/kyleu/todoforge/app/util"
)

const keyCustom = "custom"

var (
	RecentLogs []*zapcore.Entry
	recentMU   = &sync.Mutex{}
)

func InitLogging(debug bool) (util.Logger, error) {
	var logger *zap.Logger
	var err error
	lf := util.GetEnv("logging_format")
	switch {
	case strings.EqualFold(lf, "json"):
		logger, err = initJSONLogging(getLevel(zap.InfoLevel))
	case debug:
		logger, err = initDevLogging(getLevel(zap.DebugLevel))
	default:
		logger, err = initSimpleLogging(getLevel(zap.DebugLevel))
	}
	if err != nil {
		return nil, errors.Wrap(err, "error initializing logging")
	}
	return logger.Sugar(), nil
}

func CreateTestLogger() (util.Logger, error) {
	return InitLogging(false)
}

func initDevLogging(lvl zapcore.Level) (*zap.Logger, error) {
	_ = zap.RegisterEncoder(keyCustom, func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return newEncoder(cfg, runtime.GOOS != "js"), nil
	})
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig = zapcore.EncoderConfig{}
	config.Level = zap.NewAtomicLevelAt(lvl)
	config.Encoding = keyCustom
	return config.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCaller())
}

func initJSONLogging(lvl zapcore.Level) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.FunctionKey = "func"
	config.Level = zap.NewAtomicLevelAt(lvl)
	return config.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCaller())
}

func initSimpleLogging(lvl zapcore.Level) (*zap.Logger, error) {
	_ = zap.RegisterEncoder(keyCustom, func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return createSimpleEncoder(cfg), nil
	})
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig = zapcore.EncoderConfig{}
	config.Level = zap.NewAtomicLevelAt(lvl)
	config.Encoding = keyCustom
	return config.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCaller())
}

func getLevel(dflt zapcore.Level) zapcore.Level {
	l := util.GetEnv("logging_level")
	switch {
	case strings.EqualFold(l, "debug"):
		return zap.DebugLevel
	case strings.EqualFold(l, "info"):
		return zap.InfoLevel
	case strings.EqualFold(l, "warn"):
		return zap.WarnLevel
	case strings.EqualFold(l, "error"):
		return zap.ErrorLevel
	}
	return dflt
}
