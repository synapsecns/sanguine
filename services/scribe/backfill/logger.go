package backfill

import (
	"fmt"
	"github.com/ipfs/go-log"
)

// LogData holds all the data passed to LogEvent to be logged.
type LogData map[string]interface{}

type logLevel int

const (
	// InfoLevel prints a log at the info level.
	InfoLevel logLevel = iota
	// WarnLevel prints a log at the warn level.
	WarnLevel
	// ErrorLevel prints a log at the error level.
	ErrorLevel
)

var logger = log.Logger("scribe-backfiller")
var keyToTitle = map[string]string{
	"cid": "ChainID",
	"bn":  "Block Number",
	"tx":  "TX Hash",
	"la":  "Log Address",
	"ca":  "Contract Address",
	"sh":  "Start Height",
	"eh":  "End Height",
	"lc":  "Logs Chan",
	"bt":  "BlockTime Log",
	"bd":  "Backoff Duration",
	"lb":  "Last Block Stored",
	"a":   "Backoff Attempt",
	"t":   "Time Elapsed",
	"ts":  "Time Elapsed (Seconds)",
	"cn":  "Client Number",
	"e":   "Error"}

// LogEvent formats and logs an event.
func LogEvent(level logLevel, msg string, logData LogData) {
	switch level {
	case InfoLevel:
		logger.Infof("Message: %s%s", msg, generateLog(logData))
	case WarnLevel:
		logger.Warnf("Message: %s%s", msg, generateLog(logData))
	case ErrorLevel:
		logger.Errorf("Message: %s%s", msg, generateLog(logData))
	default:
		logger.Infof("Message: %s%s", msg, generateLog(logData))
	}
}

func generateLog(logData LogData) string {
	var logString string

	for k, v := range logData {
		title, ok := keyToTitle[k]
		if !ok {
			title = k
		}
		logString += "\n" + title + ": " + fmt.Sprintf("%v", v)
	}

	return logString
}
