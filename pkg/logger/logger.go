package logger

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	KibanaHost  string
	KibanaPort  string
	KibanaIndex string
}

func New(l *Logger, mode string) (logger *zap.Logger, err error) {
	switch mode {
	case "debug": // all err stacktrace
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}

	case "dev": // only fatal stacktrace
		logger, err = zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
		if err != nil {
			return nil, err
		}

	case "prod":
		conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", l.KibanaHost, l.KibanaPort))
		if err != nil {
			return nil, err
		}

		encodeConfig := zapcore.EncoderConfig{
			LevelKey:       "level",
			TimeKey:        "ts",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encodeConfig),
			zapcore.AddSync(conn),
			zap.InfoLevel)

		logger = zap.New(
			core,
			zap.AddCaller(),
			zap.AddStacktrace(zap.FatalLevel),
		).With(zap.String("index", l.KibanaIndex))

	default:
		return nil, fmt.Errorf("unknown logger mode")
	}

	return logger, nil
}

func ToContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, Logger{}, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(Logger{}).(*zap.Logger)
	if !ok {
		return zap.L()
	}
	return logger
}
