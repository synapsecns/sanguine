package core

import "go.uber.org/zap/zapcore"

// LogLevels stores available log levels in serverity order.
var LogLevels = []zapcore.Level{zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel}
