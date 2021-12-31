package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CoreLogger is the core of the package.
// It is based on zap.Logger in the current release;
// but it could be exchanged for a different one in
// future releases.
type CoreLogger struct {
	zap.Logger
}

// Logger is where all the logging happens
// A Logger can be a initialised by `NewLogger`.
type Logger struct {
	level      LogLevel
	coreLogger *CoreLogger
	fields     map[string]interface{}
}

// NewLogger returns a Logger instance.
func NewLogger(level LogLevel, cl *CoreLogger, fields map[string]interface{}) *Logger {
	return &Logger{
		level:      level,
		coreLogger: cl,
		fields:     fields,
	}
}

// NewCoreLogger returns a CoreLogger instance.
func NewCoreLogger() *CoreLogger {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	cl := CoreLogger{*zap.New(core)}
	return &cl
}

// SetField returns a new Logger with the field added
func (l *Logger) SetField(key string, value interface{}) *Logger {
	l.fields[key] = value
	return l
}

// WithFields returns a new Logger with the fields added
func (l Logger) WithFields(fields map[string]interface{}) FieldLogger {
	for k, v := range l.fields {
		fields[k] = v
	}
	return FieldLogger{
		inner:  l,
		fields: fields,
	}
}

// Info logs a message at level logger.LogLevelInfo
func (l Logger) Info(msg string) {
	if l.level <= LogLevelInfo {
		l.coreLogger.Sugar().Infow(msg, mapFieldsToArr(l.fields)...)
	}
}

// Error logs a message at level logger.LogLevelInfo
func (l Logger) Error(msg string) {
	if l.level <= LogLevelError {
		l.coreLogger.Sugar().Errorw(msg, mapFieldsToArr(l.fields)...)
	}
}

// Warning logs a message at level logger.LogLevelWarn
func (l Logger) Warning(msg string) {
	if l.level <= LogLevelWarn {
		l.coreLogger.Sugar().Warnw(msg, mapFieldsToArr(l.fields)...)
	}
}

// Debug logs a message at level logger.LogLevelDebug
func (l Logger) Debug(msg string) {
	if l.level <= LogLevelDebug {
		l.coreLogger.Sugar().Debugw(msg, mapFieldsToArr(l.fields)...)
	}
}

// FieldLogger ...
type FieldLogger struct {
	inner  Logger
	fields map[string]interface{}
}

// Info logs a message at level logger.LogLevelInfo
func (l FieldLogger) Info(msg string) {
	if l.inner.level <= LogLevelInfo {
		l.inner.coreLogger.Sugar().Infow(msg, mapFieldsToArr(l.fields)...)
	}
}

// Error logs a message at level logger.LogLevelInfo
func (l FieldLogger) Error(msg string) {
	if l.inner.level <= LogLevelError {
		l.inner.coreLogger.Sugar().Errorw(msg, mapFieldsToArr(l.fields)...)
	}
}

// Warning logs a message at level logger.LogLevelWarn
func (l FieldLogger) Warning(msg string) {
	if l.inner.level <= LogLevelWarn {
		l.inner.coreLogger.Sugar().Warnw(msg, mapFieldsToArr(l.fields)...)
	}
}

// Debug logs a message at level logger.LogLevelDebug
func (l FieldLogger) Debug(msg string) {
	if l.inner.level <= LogLevelDebug {
		l.inner.coreLogger.Sugar().Debugw(msg, mapFieldsToArr(l.fields)...)
	}
}

func mapFieldsToArr(fields map[string]interface{}) []interface{} {
	fieldArr := make([]interface{}, (len(fields) * 2))
	i := 0
	for k, v := range fields {
		fieldArr[i] = k
		fieldArr[i+1] = v
		i += 2
	}
	return fieldArr
}

// LogLevel ...
type LogLevel int

const (
	// LogLevelDebug ...
	LogLevelDebug LogLevel = iota
	// LogLevelInfo ...
	LogLevelInfo
	// LogLevelWarn ...
	LogLevelWarn
	// LogLevelError ...
	LogLevelError
)

// String returns the log level as string
func (ll LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[ll]
}

// LogLevelFromStr converts string to LogLevel
func LogLevelFromStr(l string) LogLevel {
	levelStr := strings.ToUpper(l)
	switch levelStr {
	case "DEBUG":
		return LogLevelDebug
	case "INFO":
		return LogLevelInfo
	case "WARN":
		return LogLevelWarn
	case "ERROR":
		return LogLevelError
	}
	return LogLevelInfo
}
