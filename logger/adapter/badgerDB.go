package adapter

import (
	"fmt"
	"github.com/ljg-cqu/biance/logger"
	"strings"
)

var _ logger.BadgerDBLogger = (*BadgerDBLoggerAdapter)(nil)

type BadgerDBLoggerAdapter struct {
	logger.Logger
}

func (l *BadgerDBLoggerAdapter) Debugf(format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.Debug(msg)
}

func (l *BadgerDBLoggerAdapter) Infof(format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.Info(msg)
}

func (l *BadgerDBLoggerAdapter) Warningf(format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.Warn(msg)
}

func (l *BadgerDBLoggerAdapter) Errorf(format string, a ...interface{}) {
	msg := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	l.Error(msg)
}
