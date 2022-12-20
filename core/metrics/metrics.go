package metrics

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/config"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
)

// Handler collects metrics.
type Handler interface {
	Start(ctx context.Context) error
	// Gin gets a gin middleware for tracing.
	Gin() gin.HandlerFunc
	// ConfigureHTTPClient configures tracing on an http client
	ConfigureHTTPClient(client *http.Client)
	// AddGormCallbacks adds gorm callbacks for tracing.
	AddGormCallbacks(db *gorm.DB)
}

// HandlerType is the handler type to use
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=HandlerType -linecomment
type HandlerType uint8

// AllHandlerTypes is a list of all contract types. Since we use stringer and this is a testing library, instead
// of manually copying all these out we pull the names out of stringer. In order to make sure stringer is updated, we panic on
// any method called where the index is higher than the stringer array length.
var AllHandlerTypes []HandlerType

func init() {
	for i := 0; i < len(_HandlerType_index); i++ {
		contractType := HandlerType(i)
		AllHandlerTypes = append(AllHandlerTypes, contractType)
	}
}

const (
	// DataDog is the datadog driver.
	DataDog HandlerType = 0 // Datadog
	// NewRelic is the new relic driver.
	NewRelic HandlerType = iota // NewRelic
	// Null is a null data type handler.
	Null HandlerType = iota // Null
)

// Lower gets the lowercase version of the handler type. Useful for comparison
// in switch.
func (i HandlerType) Lower() string {
	return strings.ToLower(i.String())
}

// HandlerEnv is the driver to use for metrics.
const HandlerEnv = "METRICS_HANDLER"

// NewFromEnv sets up a metrics handler from environment variable.
// this will not set the global and generally, SetupFromEnv should be used instead.
func NewFromEnv(ctx context.Context, buildInfo config.BuildInfo) (handler Handler, err error) {
	metricsHandler := strings.ToLower(os.Getenv(HandlerEnv))
	//nolint: gocritic
	switch metricsHandler {
	case DataDog.Lower():
		handler = NewDatadogMetricsHandler(buildInfo)
	case NewRelic.Lower():
		handler = NewRelicMetricsHandler(buildInfo)
	case Null.Lower():
		handler = NewNullHandler()
	default:
		handler = NewNullHandler()
	}

	if handler != nil {
		err = handler.Start(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not start handler: %w", err)
		}
	}

	return handler, nil
}
