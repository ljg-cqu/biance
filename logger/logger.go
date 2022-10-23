package logger

import "fmt"

// Field is the simple container for a single log field
type Field struct {
	Key   string
	Value interface{}
}

func AppendErrorFiled(err error, fields ...Field) []Field {
	if err == nil {
		return fields
	}
	var errorFields []Field
	errorFields = append(errorFields, fields...)
	errMsg := fmt.Sprintf("%v", err)
	errorFields = append(errorFields, Field{"error", errMsg})
	stackTrace := fmt.Sprintf("%+v", err) // work with github.com/pkg/errors
	errorFields = append(errorFields, Field{"stackTrace", stackTrace})
	return errorFields
}

// Logger declares base logging methods
type Logger interface {
	Debug(msg string, fields ...Field)
	DebugOnError(err error, msg string, fields ...Field)
	DebugNilError(err error, msg string, fields ...Field)
	DebugOnTrue(ok bool, msg string, fields ...Field)

	Info(msg string, fields ...Field)
	InfoOnError(err error, msg string, fields ...Field)
	InfoNilError(err error, msg string, fields ...Field)
	InfoOnTrue(ok bool, msg string, fields ...Field)

	Warn(msg string, fields ...Field)
	WarnOnError(err error, msg string, fields ...Field)
	WarnNilError(err error, msg string, fields ...Field)
	WarnOnTrue(ok bool, msg string, fields ...Field)

	Error(msg string, fields ...Field)
	ErrorOnError(err error, msg string, fields ...Field)
	ErrorNilError(err error, msg string, fields ...Field)
	ErrorOnTrue(ok bool, msg string, fields ...Field)

	Fatal(msg string, fields ...Field)
	FatalOnError(err error, msg string, fields ...Field)
	FatalNilError(err error, msg string, fields ...Field)
	FatalOnTrue(ok bool, msg string, fields ...Field)

	With(fields ...Field) Logger
}

// StdLogger declares standard logging methods
type StdLogger interface {
	Debugf(string, ...interface{})
	DebugOnErrorf(error, string, ...interface{})
	DebugNilErrorf(error, string, ...interface{})
	DebugOnTruef(bool, string, ...interface{})

	Infof(string, ...interface{})
	InfoOnErrorf(error, string, ...interface{})
	InfoNilErrorf(error, string, ...interface{})
	InfoOnTruef(bool, string, ...interface{})

	Warnf(string, ...interface{})
	WarnOnErrorf(error, string, ...interface{})
	WarnNilErrorf(error, string, ...interface{})
	WarnOnTruef(bool, string, ...interface{})

	Errorf(string, ...interface{})
	ErrorOnErrorf(error, string, ...interface{})
	ErrorNilErrorf(error, string, ...interface{})
	ErrorOnTruef(bool, string, ...interface{})

	Fatalf(string, ...interface{})
	FatalOnErrorf(error, string, ...interface{})
	FatalNilErrorf(error, string, ...interface{})
	FatalOnTruef(bool, string, ...interface{})
}

// BadgerDBLogger declares customized logging methods for BadgerDB
// as defined in https://github.com/dgraph-io/badger/blob/master/logger.go.
// Copy it here just for clarity.
type BadgerDBLogger interface {
	Errorf(string, ...interface{})
	Warningf(string, ...interface{})
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
}
