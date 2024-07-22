// Package exporters contains the exporters for the prometheus exporter.
package exporters

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/contrib/promexporter/config"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/instrumentation/httpcapture"
	"github.com/synapsecns/sanguine/core/retry"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"golang.org/x/sync/errgroup"
)

var logger = log.Logger("proxy-logger")

// meterName is the name of the meter used by this package.
// TODO: figure out how to autoset

const meterName = "github.com/synapsecns/sanguine/contrib/promexporter/exporters"

// makeHTTPClient makes a tracing http client.
func makeHTTPClient(handler metrics.Handler) *http.Client {
	httpClient := new(http.Client)
	handler.ConfigureHTTPClient(httpClient)

	httpClient.Transport = httpcapture.NewCaptureTransport(httpClient.Transport, handler)

	return httpClient
}

type exporter struct {
	client        *http.Client
	metrics       metrics.Handler
	cfg           config.Config
	omnirpcClient omnirpcClient.RPCClient

	otelRecorder iOtelRecorder
}

// StartExporterServer starts the exporter server.
// nolint: cyclop
func StartExporterServer(ctx context.Context, handler metrics.Handler, cfg config.Config) error {
	// the main server serves metrics since this is only a prom exporter
	_ = os.Setenv(metrics.MetricsPortEnabledEnv, "false")

	router := ginhelper.New(logger)
	router.Use(handler.Gin()...)
	router.GET(metrics.MetricsPathDefault, gin.WrapH(handler.Handler()))

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return fmt.Errorf("could not listen on port %d", cfg.Port)
	}

	// TODO: this can probably be removed
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		//nolint: gosec
		// TODO: consider setting timeouts here:  https://ieftimov.com/posts/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/
		err := http.Serve(listener, router)
		if err != nil {
			return fmt.Errorf("could not serve http: %w", err)
		}

		return nil
	})

	exp := exporter{
		client:        makeHTTPClient(handler),
		metrics:       handler,
		cfg:           cfg,
		omnirpcClient: omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes()),
		otelRecorder:  newOtelRecorder(handler),
	}

	g.Go(func() error {
		return exp.recordMetrics(ctx)
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not start exporter server: %w", err)
	}

	return nil
}

const defaultMetricsInterval = 10

func (e *exporter) recordMetrics(ctx context.Context) (err error) {
	g, _ := errgroup.WithContext(ctx)

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("could not record metrics: %w", ctx.Err())
		case <-time.After(defaultMetricsInterval * time.Second):
			// bridge token balances
			if err := e.getTokenBalancesStats(ctx); err != nil {
				return fmt.Errorf("could not record metrics: %w", err)
			}

			// dfk stuck heroes
			for _, pending := range e.cfg.DFKPending {
				if err = e.stuckHeroCountStats(ctx, common.HexToAddress(pending.Owner), pending.ChainName); err != nil {
					return fmt.Errorf("could setup metric: %w", err)
				}
			}

			// submitter stats
			for _, gasCheck := range e.cfg.SubmitterChecks {
				for _, chainID := range gasCheck.ChainIDs {
					if err := e.submitterStats(common.HexToAddress(gasCheck.Address), chainID, gasCheck.Name); err != nil {
						return fmt.Errorf("could setup metric: %w", err)
					}
				}
			}

			for chainID := range e.cfg.BridgeChecks {
				for _, token := range e.cfg.VpriceCheckTokens {
					chainID := chainID
					token := token // capture func literals

					g.Go(func() error {
						//nolint: wrapcheck
						return retry.WithBackoff(ctx, func(ctx context.Context) error {
							err = e.vpriceStats(ctx, chainID, token)
							if errors.Is(err, errPoolNotExist) {
								return nil
							}

							if err != nil {
								return fmt.Errorf("error starting vprice:%w", err)
							}

							return nil
						}, retry.WithMaxAttempts(-1), retry.WithMaxAttemptTime(time.Second*10), retry.WithMaxTotalTime(-1))
					})
				}
			}

		}
	}
}
