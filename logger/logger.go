package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	level  LogLevel
	logger *zap.Logger
	fields map[string]interface{}
}

func (l *Logger) SetField(key string, value interface{}) *Logger {
	l.fields[key] = value
	return l
}

func (l Logger) WithFields(fields map[string]interface{}) FieldLogger {
	for k, v := range l.fields {
		fields[k] = v
	}
	return FieldLogger{
		inner:  l,
		fields: fields,
	}
}

func (l Logger) Info(msg string) {
	if l.level <= LogLevelInfo {
		l.logger.Sugar().Infow(msg, mapFieldsToArr(l.fields)...)
	}
}

func (l Logger) Error(msg string) {
	if l.level <= LogLevelError {
		l.logger.Sugar().Errorw(msg, mapFieldsToArr(l.fields)...)
	}
}

func (l Logger) Warning(msg string) {
	if l.level <= LogLevelWarn {
		l.logger.Sugar().Warnw(msg, mapFieldsToArr(l.fields)...)
	}
}

func (l Logger) Debug(msg string) {
	if l.level <= LogLevelDebug {
		l.logger.Sugar().Debugw(msg, mapFieldsToArr(l.fields)...)
	}
}

type FieldLogger struct {
	inner  Logger
	fields map[string]interface{}
}

func (l FieldLogger) Info(msg string) {
	if l.inner.level <= LogLevelInfo {
		l.inner.logger.Sugar().Infow(msg, mapFieldsToArr(l.fields)...)
	}
}

func (l FieldLogger) Error(msg string) {
	if l.inner.level <= LogLevelError {
		l.inner.logger.Sugar().Errorw(msg, mapFieldsToArr(l.fields)...)
	}
}

func (l FieldLogger) Warning(msg string) {
	if l.inner.level <= LogLevelWarn {
		l.inner.logger.Sugar().Warnw(msg, mapFieldsToArr(l.fields)...)
	}
}

func (l FieldLogger) Debug(msg string) {
	fmt.Println("here")
	if l.inner.level <= LogLevelDebug {
		l.inner.logger.Sugar().Debugw(msg, mapFieldsToArr(l.fields)...)
	}
}

func mapFieldsToArr(fields map[string]interface{}) []interface{} {
	fieldArr := make([]interface{}, (len(fields) * 2))
	i := 0
	for k, v := range fields {
		fieldArr[i] = k
		fieldArr[i+1] = v
		i = i + 2
	}
	return fieldArr
}

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

func (ll LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[ll]
}
