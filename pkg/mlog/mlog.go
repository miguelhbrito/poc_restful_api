package mlog

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/stone_assignment/pkg/api"
)

const (
	debug = "debug"
	info  = "info"
	error = "error"
	warn  = "warn"
	panic = "panic"
)

type (
	logInfo struct {
		cpf string
	}
)

func buildLog(ctx context.Context, level string) *zerolog.Event {
	logger := zerolog.New(os.Stdout)

	if ctx != nil {
		li := loadLogInfo(ctx)

		if li.cpf != "" {
			logger = logger.With().
				Str(string(api.CpfCtxKey), li.cpf).Logger()
		}
	}

	switch level {
	case info:
		return logger.Info()
	case debug:
		return logger.Debug()
	case error:
		return logger.Error()
	case warn:
		return logger.Warn()
	case panic:
		return logger.Panic()
	default:
		return logger.Info()
	}
}

func loadLogInfo(ctx context.Context) logInfo {
	var li logInfo

	if v := ctx.Value(api.CpfCtxKey); v != nil {
		li.cpf = v.(api.Cpf).String()
	}
	return li
}

func Info(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, info)
}

func Error(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, error)
}

func Debug(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, debug)
}

func Warn(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, warn)
}

func Panic(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, panic)
}
