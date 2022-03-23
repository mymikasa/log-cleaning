package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Panic(err error, kvs ...string) {
	logPanic(zap.L(), err, kvs...)
}

func Error(err error, kvs ...string) {
	logError(zap.L(), err, kvs...)
}

func Warn(err error, kvs ...string) {
	logWarn(zap.L(), err, kvs...)
}

func Info(msg string, kvs ...string) {
	logInfo(zap.L(), msg, kvs...)
}

func Debug(msg string, kvs ...string) {
	logDebug(zap.L(), msg, kvs...)
}

func zapFields(err error, kvs ...string) []zapcore.Field {
	var fields []zapcore.Field
	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	for i := 0; i < len(kvs); i += 2 {
		fields = append(fields, zap.String(kvs[i], kvs[i+1]))
	}

	return fields
}

func logPanic(l *zap.Logger, err error, kvs ...string) {
	fields := zapFields(err, kvs...)
	l.Panic("panic", fields...)
}

func logError(l *zap.Logger, err error, kvs ...string) {
	fields := zapFields(err, kvs...)
	l.Error("error", fields...)
}

func logWarn(l *zap.Logger, err error, kvs ...string) {
	fields := zapFields(err, kvs...)
	l.Warn("warn", fields...)
}

func logInfo(l *zap.Logger, msg string, kvs ...string) {
	fields := zapFields(nil, kvs...)
	l.Info(msg, fields...)
}

func logDebug(l *zap.Logger, msg string, kvs ...string) {
	fields := zapFields(nil, kvs...)
	l.Debug(msg, fields...)
}
