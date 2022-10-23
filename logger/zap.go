package logger

import (
	uberZap "go.uber.org/zap"
)

var _ Logger = (*zap)(nil)

type zap struct {
	l *uberZap.Logger
}

func NewZap(l *uberZap.Logger) Logger {
	return &zap{l}
}

func (l *zap) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, l.zapFields(fields...)...)
}

func (l *zap) DebugOnError(err error, msg string, fields ...Field) {
	if err != nil {
		l.l.Debug(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) DebugNilError(err error, msg string, fields ...Field) {
	if err == nil {
		l.l.Debug(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) DebugOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		l.l.Debug(msg, l.zapFields(fields...)...)
	}
}

func (l *zap) Info(msg string, fields ...Field) {
	l.l.Info(msg, l.zapFields(fields...)...)
}

func (l *zap) InfoOnError(err error, msg string, fields ...Field) {
	if err != nil {
		l.l.Info(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) InfoNilError(err error, msg string, fields ...Field) {
	if err == nil {
		l.l.Info(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) InfoOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		l.l.Info(msg, l.zapFields(fields...)...)
	}
}

func (l *zap) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, l.zapFields(fields...)...)
}

func (l *zap) WarnOnError(err error, msg string, fields ...Field) {
	if err != nil {
		l.l.Warn(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) WarnNilError(err error, msg string, fields ...Field) {
	if err == nil {
		l.l.Warn(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) WarnOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		l.l.Warn(msg, l.zapFields(fields...)...)
	}
}

func (l *zap) Error(msg string, fields ...Field) {
	l.l.Error(msg, l.zapFields(fields...)...)
}

func (l *zap) ErrorOnError(err error, msg string, fields ...Field) {
	if err != nil {

		l.l.Error(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) ErrorNilError(err error, msg string, fields ...Field) {
	if err == nil {

		l.l.Error(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) ErrorOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		l.l.Error(msg, l.zapFields(fields...)...)
	}
}

func (l *zap) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, l.zapFields(fields...)...)
}

func (l *zap) FatalOnError(err error, msg string, fields ...Field) {
	if err != nil {
		l.l.Fatal(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) FatalNilError(err error, msg string, fields ...Field) {
	if err == nil {
		l.l.Fatal(msg, l.zapFields(AppendErrorFiled(err, fields...)...)...)
	}
}

func (l *zap) FatalOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		l.l.Fatal(msg, l.zapFields(fields...)...)
	}
}

func (l *zap) With(fields ...Field) Logger {
	return NewZap(l.l.With(l.zapFields(fields...)...))
}

func (l *zap) zapFields(fields ...Field) []uberZap.Field {
	result := make([]uberZap.Field, len(fields))
	for i, f := range fields {
		result[i] = uberZap.Any(f.Key, f.Value)
	}
	return result
}
