package adapter

import (
	"fmt"
	"github.com/avalido/mpc-controller/logger"
	"strings"
)

var _ logger.StdLogger = (*StdLoggerAdapter)(nil)

type StdLoggerAdapter struct {
	logger.Logger
}

func (l *StdLoggerAdapter) Debugf(format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.Debug(msg)
}

func (l *StdLoggerAdapter) DebugOnErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.DebugOnError(err, msg)
}

func (l *StdLoggerAdapter) DebugNilErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.DebugNilError(err, msg)
}

func (l *StdLoggerAdapter) DebugOnTruef(ok bool, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.DebugOnTrue(ok, msg)
}

func (l *StdLoggerAdapter) Infof(format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.Info(msg)
}

func (l *StdLoggerAdapter) InfoOnErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.InfoOnError(err, msg)
}

func (l *StdLoggerAdapter) InfoNilErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.InfoNilError(err, msg)
}

func (l *StdLoggerAdapter) InfoOnTruef(ok bool, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.InfoOnTrue(ok, msg)
}

func (l *StdLoggerAdapter) Warnf(format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.Warn(msg)
}

func (l *StdLoggerAdapter) WarnOnErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.WarnOnError(err, msg)
}

func (l *StdLoggerAdapter) WarnNilErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.WarnNilError(err, msg)
}

func (l *StdLoggerAdapter) WarnOnTruef(ok bool, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.WarnOnTrue(ok, msg)
}

func (l *StdLoggerAdapter) Errorf(format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.Error(msg)
}

func (l *StdLoggerAdapter) ErrorOnErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.ErrorOnError(err, msg)
}

func (l *StdLoggerAdapter) ErrorNilErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.ErrorNilError(err, msg)
}

func (l *StdLoggerAdapter) ErrorOnTruef(ok bool, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.ErrorOnTrue(ok, msg)
}

func (l *StdLoggerAdapter) Fatalf(format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.Fatalf(msg)
}

func (l *StdLoggerAdapter) FatalOnErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.FatalOnError(err, msg)
}

func (l *StdLoggerAdapter) FatalNilErrorf(err error, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.FatalNilError(err, msg)
}

func (l *StdLoggerAdapter) FatalOnTruef(ok bool, format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.FatalOnTrue(ok, msg)
}
