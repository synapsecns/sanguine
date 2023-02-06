package prom

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grafana-tools/sdk"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
	"net/http"
	"time"
)

// pushInterval is how often metrics should be pushed to the push gateway (if enabled).
const pushInterval = time.Second * 3

// NodeIDName is the name of the node id.
const NodeIDName = "node_id"

// MetricHandler stores metrics and handles exposing them.
type MetricHandler struct {
	// registry is the prometheus registry
	registry *prometheus.Registry
	// config is the metric handler config
	config *Config
	// nodeID is the identifier of the node
	nodeID string
}

// NewMetricHandler creates a new metrics handler.
func NewMetricHandler(nodeID string, cfg *Config) *MetricHandler {
	return &MetricHandler{nodeID: nodeID, registry: prometheus.NewRegistry(), config: cfg}
}

// RegisterMetrics registers metrics in the registry.
func (m *MetricHandler) RegisterMetrics(cs ...prometheus.Collector) {
	m.registry.MustRegister(cs...)
}

// enablePushGateway enables a push gateway for metrics and pushes every pushInterval until context is canceled.
func (m *MetricHandler) enablePushGateway(ctx context.Context) error {
	// synchronously run the first push
	pushMetrics := func() error {
		if err := push.New(m.config.PushGateway, m.nodeID).
			Gatherer(m.registry).
			Grouping(NodeIDName, m.nodeID).
			Push(); err != nil {
			return fmt.Errorf("could not push metrics to push gateway")
		}
		return nil
	}

	err := pushMetrics()
	if err != nil {
		return fmt.Errorf("could not commit initial push to gateway: %w", err)
	}

	go func() {
		for {
			var err error
			var finished bool

			select {
			case <-ctx.Done():
				// on quit, run a last push
				err = pushMetrics()
				finished = true
			case <-time.After(pushInterval):
				err = pushMetrics()
			}
			if err != nil {
				logger.Errorf("could not push metrics: %v", err)
			}

			if finished {
				return
			}
		}
	}()
	return nil
}

// Start starts the metric handler backround services with context.
func (m *MetricHandler) Start(ctx context.Context) (err error) {
	if m.config.Enabled && m.config.PushGateway != "" {
		err = m.enablePushGateway(ctx)
	}

	if err != nil {
		return fmt.Errorf("could not start push gateway")
	}
	return err
}

// Handler gets the http handler for the metrics handler. This can be used in combination with or instead of
// the push gateway.
func (m *MetricHandler) Handler() http.Handler {
	return promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{
		ErrorLog:            NewPromLogger(logger),
		ErrorHandling:       promhttp.HTTPErrorOnError,
		Registry:            m.registry,
		DisableCompression:  false,
		MaxRequestsInFlight: 50,
		Timeout:             10 * time.Second,
		EnableOpenMetrics:   false,
	})
}

// PutDashboard creates a dashboard. TODO replace by title/verify grafana config.
func (m *MetricHandler) PutDashboard(ctx context.Context, dashboard *sdk.Board) error {
	c, err := sdk.NewClient(m.config.GrafanaHost, m.config.GrafanaKey, sdk.DefaultHTTPClient)
	if err != nil {
		return fmt.Errorf("could not add doashboard: %w", err)
	}

	_, err = c.SetDashboard(ctx, *dashboard, sdk.SetDashboardParams{
		Overwrite: true,
	})
	if err != nil {
		return fmt.Errorf("could not create dashboard: %w", err)
	}
	return nil
}

// boardHandler is an http handler that returns an http dashboard.
type boardHandler struct {
	dashboard *sdk.Board
}

// ServeHTTP serves up the dashboard.
func (b boardHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	board, err := json.Marshal(b.dashboard)
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		_, err = writer.Write([]byte("could not convert board to json"))
	} else {
		_, err = writer.Write(board)
		writer.WriteHeader(http.StatusOK)
	}

	if err != nil {
		logger.Errorf("got error while writing response: %v", err)
	}
}

var _ http.Handler = &boardHandler{}

// CreateBoardJSONHandler creates a handler that returns the json
// for a grafana dashboard.
func CreateBoardJSONHandler(board *sdk.Board) http.Handler {
	return boardHandler{board}
}
