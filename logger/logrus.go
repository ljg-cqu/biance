package logger

import (
	sirupsenLogrus "github.com/sirupsen/logrus"
)

var _ Logger = (*logrus)(nil)

type logrus struct {
	l *sirupsenLogrus.Logger
}

func NewLogrus(l *sirupsenLogrus.Logger) Logger {
	return &logrus{l}
}

func (l *logrus) Debug(msg string, fields ...Field) {
	fieldsl := sirupsenLogrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl).Debug(msg)
}

func (l *logrus) DebugOnError(err error, msg string, fields ...Field) {
	if err != nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Debug(msg)
	}
}

func (l *logrus) DebugNilError(err error, msg string, fields ...Field) {
	if err == nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Debug(msg)
	}
}

func (l *logrus) DebugOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Debug(msg)
	}
}

func (l *logrus) Info(msg string, fields ...Field) {
	fieldsl := sirupsenLogrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl).Info(msg)
}

func (l *logrus) InfoOnError(err error, msg string, fields ...Field) {
	if err != nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Info(msg)
	}
}

func (l *logrus) InfoNilError(err error, msg string, fields ...Field) {
	if err == nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Info(msg)
	}
}

func (l *logrus) InfoOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Info(msg)
	}
}

func (l *logrus) Warn(msg string, fields ...Field) {
	fieldsl := sirupsenLogrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl).Warn(msg)
}

func (l *logrus) WarnOnError(err error, msg string, fields ...Field) {
	if err != nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Warn(msg)
	}
}

func (l *logrus) WarnNilError(err error, msg string, fields ...Field) {
	if err == nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Warn(msg)
	}
}

func (l *logrus) WarnOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Warn(msg)
	}
}

func (l *logrus) Error(msg string, fields ...Field) {
	fieldsl := sirupsenLogrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl).Error(msg)
}

func (l *logrus) ErrorOnError(err error, msg string, fields ...Field) {
	if err != nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Error(msg)
	}
}

func (l *logrus) ErrorNilError(err error, msg string, fields ...Field) {
	if err == nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Error(msg)
	}
}

func (l *logrus) ErrorOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Error(msg)
	}
}

func (l *logrus) Fatal(msg string, fields ...Field) {
	fieldsl := sirupsenLogrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl).Fatal(msg)
}

func (l *logrus) FatalOnError(err error, msg string, fields ...Field) {
	if err != nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Fatal(msg)
	}
}

func (l *logrus) FatalNilError(err error, msg string, fields ...Field) {
	if err == nil {
		fields = AppendErrorFiled(err, fields...)
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Fatal(msg)
	}
}

func (l *logrus) FatalOnTrue(ok bool, msg string, fields ...Field) {
	if ok {
		fieldsl := sirupsenLogrus.Fields{}
		for _, f := range fields {
			fieldsl[f.Key] = f.Value
		}
		l.l.WithFields(fieldsl).Fatal(msg)
	}
}

func (l *logrus) With(fields ...Field) Logger {
	fieldsl := sirupsenLogrus.Fields{}
	for _, f := range fields {
		fieldsl[f.Key] = f.Value
	}
	l.l.WithFields(fieldsl)
	return l
}
