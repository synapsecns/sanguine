package prometheus

import (
	"fmt"
	//nolint: staticcheck
	kitLog "github.com/go-kit/kit/log"
	"github.com/ipfs/go-log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var logger = log.Logger("synapse-metrics")

// NewPromLogger creates a new prometheus error logger.
func NewPromLogger(logClient *log.ZapEventLogger) PromErrorLogger {
	return PromErrorLogger{logClient}
}

// PromErrorLogger is the error logger used for prometheus http.
type PromErrorLogger struct {
	*log.ZapEventLogger
}

// Log prints the output of prom to the logger.
func (p PromErrorLogger) Log(keyvals ...interface{}) error {
	msg := fmt.Sprintln(keyvals...)
	p.Debugf(msg[:len(msg)-1])
	return nil
}

// Println prints the output of prom to the error logger.
func (p PromErrorLogger) Println(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	p.Errorf(msg[:len(msg)-1])
}

var _ promhttp.Logger = PromErrorLogger{}

// used for push gateway.
var _ kitLog.Logger = PromErrorLogger{}
