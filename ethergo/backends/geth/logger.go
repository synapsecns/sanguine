package geth

import (
	gethLog "github.com/ethereum/go-ethereum/log"
	"github.com/ipfs/go-log"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/synapsecns/sanguine/core"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var logger = log.Logger("geth-logger")

var ethLogger *gethLog.GlogHandler

// do this on init to get the default log level, as any global log level will not have been set yet.
func init() {
	setupEthLogger()
}

func getEthLogLevel() gethLog.Lvl {
	for _, level := range core.LogLevels {
		if logger.Desugar().Core().Enabled(level) {
			switch level {
			case zapcore.DebugLevel, zapcore.InvalidLevel:
				return gethLog.LvlDebug
			case zapcore.InfoLevel:
				return gethLog.LvlInfo
			case zapcore.WarnLevel:
				return gethLog.LvlWarn
			case zapcore.ErrorLevel:
				return gethLog.LvlError
			case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
				return gethLog.LvlCrit
			}
		}
	}
	logger.Warn("could not get geth log level from ipfs logger")
	// return info otherwise
	return gethLog.LvlInfo
}

// setupEthLogger sets up the eth global logger.
func setupEthLogger() {
	// eth sets up global logging through it's internal rpc analytics. here we setup some helpers to set it up for us
	ethLogger = gethLog.NewGlogHandler(gethLog.StreamHandler(os.Stderr, gethLog.TerminalFormat(false)))
	ethLogger.Verbosity(getEthLogLevel())

	// create writer
	output := io.Writer(os.Stderr)

	// enable color logging if available
	usecolor := (isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb"
	if usecolor {
		output = colorable.NewColorableStderr()
	}
	ethLogger.SetHandler(gethLog.StreamHandler(output, gethLog.TerminalFormat(usecolor)))

	// set the global handler to the new logger
	gethLog.Root().SetHandler(ethLogger)
}
