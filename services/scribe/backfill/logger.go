package backfill

import "github.com/ipfs/go-log"

type BackfillLoggerStruct struct {
	chainID string
}

type logLevel int

const (
	InfoLevel logLevel = iota
	WarnLevel
	ErrorLevel
)

var logger = log.Logger("synapse-backfiller")
var loggerBlocktime = log.Logger("synapse-blocktime-backfiller")

func GenerateLog(msg string, logData BackfillLoggerStruct, level logLevel) {
	switch level {
	case InfoLevel:
		logger.Infof("ChainID: %s", msg, logData)
	case WarnLevel:
		logger.Warnf("ChainID: %s", msg, logData)
	case ErrorLevel:
		logger.Errorf("ChainID: %s", msg, logData)
	default:
		logger.Infof("ChainID: %s", msg, logData)
	}
}
