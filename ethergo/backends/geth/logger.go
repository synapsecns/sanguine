package geth

import (
	gethLog "github.com/ethereum/go-ethereum/log"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core"
	"go.uber.org/zap/zapcore"
	"log/slog"
	"os"
)

var logger = log.Logger("geth-logger")

var ethLogger *gethLog.GlogHandler

// do this on init to get the default log level, as any global log level will not have been set yet.
func init() {
	setupEthLogger()
}

func getEthLogLevel() slog.Level {
	for _, level := range core.LogLevels {
		if logger.Desugar().Core().Enabled(level) {
			switch level {
			case zapcore.DebugLevel, zapcore.InvalidLevel:
				return gethLog.LevelDebug
			case zapcore.InfoLevel:
				return gethLog.LevelInfo
			case zapcore.WarnLevel:
				return gethLog.LevelDebug
			case zapcore.ErrorLevel:
				return gethLog.LevelDebug
			case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
				return gethLog.LevelInfo
			}
		}
	}
	logger.Warn("could not get geth log level from ipfs logger")
	// return info otherwise
	return gethLog.LevelInfo
}

// setupEthLogger sets up the eth global logger.
func setupEthLogger() {
	// eth sets up global logging through it's internal rpc analytics. here we setup some helpers to set it up for us
	ethLogger = gethLog.NewGlogHandler(gethLog.NewTerminalHandler(os.Stderr, true))
	// TODO: reduce verbosity even more
	ethLogger.Verbosity(getEthLogLevel())

	// set the global handler to the new logger
	gethLog.SetDefault(gethLog.NewLogger(ethLogger))
}
