package logger

import (
	uberZap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var DevMode bool
var UseConsoleEncoder bool

// Default return a Logger right depending on go.uber.org/zap Logger.
func Default() Logger {
	return DefaultWithCallerSkip(1)
}

// DefaultWithCallerSkip is same as Default except caller skip can be specified.
func DefaultWithCallerSkip(skip int) Logger {
	var logger *uberZap.Logger
	var logConfig uberZap.Config

	if DevMode {
		logConfig = uberZap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeTime = iso8601LocalTimeEncoder
		if UseConsoleEncoder {
			logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		} else {
			logConfig.Encoding = "json"
		}
		logger, _ = logConfig.Build(uberZap.AddCallerSkip(skip))
	} else {
		logConfig = uberZap.NewProductionConfig()
		logConfig.EncoderConfig.EncodeTime = iso8601LocalTimeEncoder
		logger, _ = logConfig.Build(uberZap.AddCallerSkip(skip))
	}
	return NewZap(logger)
}

// A UTC variation of ZapCore.ISO8601TimeEncoder with nanosecond precision
func iso8601UTCTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000000Z"))
}

// A Local variation of ZapCore.ISO8601TimeEncoder with nanosecond precision
func iso8601LocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Local().Format("2006-01-02T15:04:05.000000000Z"))
}

func Debug(msg string, fields ...Field) {
	logger := Default()
	logger.Debug(msg, fields...)
}

func DebugOnError(err error, msg string, fields ...Field) {
	logger := Default()
	logger.DebugOnError(err, msg, fields...)
}

func DebugOnTrue(ok bool, msg string, fields ...Field) {
	logger := Default()
	logger.DebugOnTrue(ok, msg, fields...)
}

func Info(msg string, fields ...Field) {
	logger := Default()
	logger.Info(msg, fields...)
}

func InfoOnError(err error, msg string, fields ...Field) {
	logger := Default()
	logger.InfoOnError(err, msg, fields...)
}

func InfoOnTrue(ok bool, msg string, fields ...Field) {
	logger := Default()
	logger.InfoOnTrue(ok, msg, fields...)
}

func Warn(msg string, fields ...Field) {
	logger := Default()
	logger.Warn(msg, fields...)
}

func WarnOnError(err error, msg string, fields ...Field) {
	logger := Default()
	logger.WarnOnError(err, msg, fields...)
}

func WarnOnTrue(ok bool, msg string, fields ...Field) {
	logger := Default()
	logger.WarnOnTrue(ok, msg, fields...)
}

func Error(msg string, fields ...Field) {
	logger := Default()
	logger.Error(msg, fields...)
}

func ErrorOnError(err error, msg string, fields ...Field) {
	logger := Default()
	logger.ErrorOnError(err, msg, fields...)
}

func ErrorOnTrue(ok bool, msg string, fields ...Field) {
	logger := Default()
	logger.ErrorOnTrue(ok, msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	logger := Default()
	logger.Fatal(msg, fields...)
}

func FatalOnError(err error, msg string, fields ...Field) {
	logger := Default()
	logger.FatalOnError(err, msg, fields...)
}

func FatalOnTrue(ok bool, msg string, fields ...Field) {
	logger := Default()
	logger.FatalOnTrue(ok, msg, fields...)
}

func With(fields ...Field) Logger {
	logger := Default()
	logger.With(fields...)
	return logger
}
