package dbcommon

import (
	"github.com/synapsecns/sanguine/core"
	goLog "log"
	"os"
	"time"

	"github.com/ipfs/go-log"
	"go.uber.org/zap/zapcore"
	gormLogger "gorm.io/gorm/logger"
)

// GetGormLogger gets a gorm logger at the correct level
// TODO investigate https://github.com/moul/zapgorm, we want to use the same write group.
// TODO: otel logging.
func GetGormLogger(zapLogger *log.ZapEventLogger) gormLogger.Interface {
	// use a more performant logger in prod since we don't care about logging misses, etc
	return gormLogger.New(
		goLog.New(os.Stdout, "", goLog.LstdFlags),
		gormLogger.Config{
			SlowThreshold:             200 * time.Millisecond,
			Colorful:                  true,
			IgnoreRecordNotFoundError: zapLogger.Desugar().Core().Enabled(zapcore.DebugLevel),
			LogLevel:                  getGormLogLevel(zapLogger),
		},
	)
}

// GetGormLogger converts the ipfs subsystem logger log level
// to a gorm level.
func getGormLogLevel(zapLogger *log.ZapEventLogger) gormLogger.LogLevel {
	for _, level := range core.LogLevels {
		if zapLogger.Desugar().Core().Enabled(level) {
			switch level {
			case zapcore.DebugLevel, zapcore.InfoLevel, zapcore.InvalidLevel:
				return gormLogger.Info
			case zapcore.WarnLevel:
				return gormLogger.Warn
			case zapcore.ErrorLevel:
				return gormLogger.Error
			case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
				return gormLogger.Error
			}
		}
	}

	zapLogger.Warn("could not get gorm log level from ipfs logger")
	// return info otherwise
	return gormLogger.Warn
}
