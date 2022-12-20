package prom

import (
	"fmt"
	//nolint: staticcheck
	kitLog "github.com/go-kit/kit/log"
	"github.com/ipfs/go-log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var logger = log.Logger("synapse-metrics")

// NewPromLogger creates a new prometheus error logger.
func NewPromLogger(logClient *log.ZapEventLogger) ErrorLogger {
	return ErrorLogger{logClient}
}

// ErrorLogger is the error logger used for prometheus http.
type ErrorLogger struct {
	*log.ZapEventLogger
}

// Log prints the output of prom to the logger.
func (p ErrorLogger) Log(keyvals ...interface{}) error {
	msg := fmt.Sprintln(keyvals...)
	p.Debugf(msg[:len(msg)-1])
	return nil
}

// Println prints the output of prom to the error logger.
func (p ErrorLogger) Println(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	p.Errorf(msg[:len(msg)-1])
}

var _ promhttp.Logger = ErrorLogger{}

// used for push gateway.
var _ kitLog.Logger = ErrorLogger{}
