package commandline

import (
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
	"strings"
)

// generic flags

// LogLevel sets the log level for the node.
var LogLevel = cli.StringFlag{
	Name:        "log-level",
	Usage:       fmt.Sprintf("set the log level for the application to one of %s", LogLevelOptions()),
	Value:       zapcore.WarnLevel.String(),
	DefaultText: "",
}

// LogLevelOptions generates log level options and returns them as a string.
func LogLevelOptions() (res string) {
	for _, level := range core.LogLevels {
		res += fmt.Sprintf("\"%s\" ", level.String())
	}
	return res
}

// IsValidLogLevel determines if a log level is valid.
func IsValidLogLevel(level string) bool {
	parsedLevel, err := log.LevelFromString(strings.ToUpper(level))
	if err != nil {
		return false
	}

	for _, level := range core.LogLevels {
		if parsedLevel == log.LogLevel(level) {
			return true
		}
	}
	return false
}

// SetLogLevel parses the cli context and sets the global log level accordingly.
// Note: for this to work the log level has to be set in the command.
func SetLogLevel(c *cli.Context) {
	err := log.SetLogLevel("*", c.String(LogLevel.Name))
	if err != nil || !IsValidLogLevel(c.String(LogLevel.Name)) {
		fmt.Printf("could not set log level to %s, using default %s \n", LogLevel.Name, LogLevel.Value)
		_ = log.SetLogLevel("*", LogLevel.Value)
	} else {
		logger.Infof("log level set to %s", c.String(LogLevel.Name))
	}
}
