package metrics

import (
	"context"
	"github.com/synapsecns/sanguine/core/config"
)

var globalHandler Handler

// by defualt, use the null handler.
func init() {
	globalHandler = NewNullHandler()
}

// Setup sets up the global metrics handler. In general, we discourage globals
// but because of the ubiquitiy of global variables and the tangential nature
// of metrics, we allow this.
func Setup(ctx context.Context, buildInfo config.BuildInfo) error {
	handler, err := NewFromEnv(ctx, buildInfo)
	if err != nil {
		return err
	}

	globalHandler = handler
	return nil
}

// Get gets the global handler
func Get() Handler {
	return globalHandler
}
