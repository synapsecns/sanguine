package metrics

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	ngrin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/synapsecns/sanguine/core/config"
	nrcontrib "github.com/synapsecns/sanguine/core/metrics/newrelic"
	"gorm.io/gorm"
)

type newRelicHandler struct {
	*baseHandler
	app       *newrelic.Application
	startMux  sync.Mutex
	buildInfo config.BuildInfo
}

// NewRelicMetricsHandler creates a new newrelic metrics handler.
func NewRelicMetricsHandler(buildInfo config.BuildInfo) Handler {
	logger.Warn("new relic metrics handler is not fully implemented, please add an otel bride")

	return &newRelicHandler{
		buildInfo:   buildInfo,
		baseHandler: newBaseHandler(buildInfo),
	}
}

func (n *newRelicHandler) AddGormCallbacks(db *gorm.DB) {
	nrcontrib.AddGormCallbacks(db, func(name string, opts ...newrelic.TraceOption) *newrelic.Transaction {
		return n.app.StartTransaction(name, opts...)
	})
}

func (n *newRelicHandler) Gin() gin.HandlerFunc {
	return ngrin.Middleware(n.app)
}

func (n *newRelicHandler) Start(_ context.Context) (err error) {
	n.startMux.Lock()
	defer n.startMux.Unlock()

	if n.app == nil {
		n.app, err = newrelic.NewApplication(
			newrelic.ConfigAppName(n.buildInfo.Name()),
			newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
			newrelic.ConfigAppLogForwardingEnabled(true),
			newrelic.ConfigAppLogEnabled(true),
			newrelic.ConfigCodeLevelMetricsEnabled(true),
			nrzap.ConfigLogger(logger.Desugar()),
			func(c *newrelic.Config) {
				c.Labels = map[string]string{
					"version": n.buildInfo.Version(),
					"commit":  n.buildInfo.Commit(),
				}
			},
			// optional overrides
			newrelic.ConfigFromEnvironment(),
		)
		if err != nil {
			return fmt.Errorf("could not create new relic application: %w", err)
		}
	}

	return nil
}

func (n *newRelicHandler) ConfigureHTTPClient(client *http.Client, opts ...otelhttp.Option) {
	// note: opts are ignored here
	// use the newrelic transport
	nrTransport := newrelic.NewRoundTripper(client.Transport)
	client.Transport = nrRoundTripper{app: n.app, inner: nrTransport}
}

type nrRoundTripper struct {
	inner http.RoundTripper
	app   *newrelic.Application
}

func (n nrRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	txn := newrelic.FromContext(req.Context())
	if txn == nil {
		txn = n.app.StartTransaction(req.URL.String())
		req = newrelic.RequestWithTransactionContext(req, txn)
	}
	defer txn.End()
	resp, err := n.inner.RoundTrip(req)
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}
	return resp, nil
}
