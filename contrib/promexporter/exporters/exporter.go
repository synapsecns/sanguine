// Package exporters contains the exporters for the prometheus exporter.
package exporters

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/lmittmann/w3/module/eth"
	"github.com/synapsecns/sanguine/contrib/promexporter/config"
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/gql/dfk"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/instrumentation"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

var logger = log.Logger("proxy-logger")

// meterName is the name of the meter used by this package.
// TODO: figure out how to autoset

const meterName = "github.com/synapsecns/sanguine/contrib/promexporter/exporters"

// makeHTTPClient makes a tracing http client.
func makeHTTPClient(handler metrics.Handler) *http.Client {
	httpClient := new(http.Client)
	handler.ConfigureHTTPClient(httpClient)

	httpClient.Transport = instrumentation.NewCaptureTransport(httpClient.Transport, handler)

	return httpClient
}

type exporter struct {
	client        *http.Client
	metrics       metrics.Handler
	cfg           config.Config
	omnirpcClient omnirpcClient.RPCClient
}

// StartExporterServer starts the exporter server.
func StartExporterServer(ctx context.Context, handler metrics.Handler, cfg config.Config) error {
	// the main server serves metrics since this is only a prom exporter
	_ = os.Setenv(metrics.MetricsPortEnabledEnv, "false")

	router := ginhelper.New(logger)
	router.Use(handler.Gin())
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
	}

	// register dfk metrics
	for _, pending := range cfg.DFKPending {
		// heroes on both chains
		err = exp.stuckHeroCount(common.HexToAddress(pending.Owner), pending.ChainName)
		if err != nil {
			return fmt.Errorf("could setup metric: %w", err)
		}
	}

	// register gas check metrics
	for _, gasCheck := range cfg.SubmitterChecks {
		for _, chainID := range gasCheck.ChainIDs {
			err := exp.submitterStats(common.HexToAddress(gasCheck.Address), chainID, gasCheck.Name)
			if err != nil {
				return fmt.Errorf("could setup metric: %w", err)
			}
		}
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not start exporter server: %w", err)
	}

	return nil
}

func (e *exporter) stuckHeroCount(owner common.Address, chainName string) error {
	meter := e.metrics.Meter(meterName)
	attributes := attribute.NewSet(attribute.String("chain_name", chainName))

	stuckCount, err := meter.Int64ObservableGauge(stuckHeroMetric)
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	if _, err := meter.RegisterCallback(func(parentCtx context.Context, o metric.Observer) (err error) {
		ctx, span := e.metrics.Tracer().Start(parentCtx, "dfk_stats", trace.WithAttributes(
			attribute.String("chain_name", chainName),
		))

		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()

		dfkClient := dfk.NewClient(e.client, e.cfg.DFKUrl)

		stuckHeroes, err := dfkClient.StuckHeroes(ctx, core.PtrTo[int64](0), core.PtrTo(owner.String()))
		if err != nil {
			return fmt.Errorf("could not get stuck hero count: %w", err)
		}

		// TODO: this maxes out at 100 now. Need binary search or something.
		o.ObserveInt64(stuckCount, int64(len(stuckHeroes.Heroes)), metric.WithAttributeSet(attributes))

		return nil
	}, stuckCount); err != nil {
		return fmt.Errorf("registering callback on instruments: %w", err)
	}

	return nil
}

const stuckHeroMetric = "dfk_pending_heroes"
const gasBalance = "gas_balance"
const nonce = "nonce"

// note: this kind of check should be deprecated in favor of submitter metrics once everything has been moved over.
func (e *exporter) submitterStats(address common.Address, chainID int, name string) error {
	meter := e.metrics.Meter(fmt.Sprintf("%s_%d", meterName, chainID))

	balanceGauge, err := meter.Float64ObservableGauge(gasBalance)
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	nonceGauge, err := meter.Int64ObservableGauge(nonce)
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	attributes := attribute.NewSet(attribute.Int(metrics.ChainID, chainID), attribute.String(metrics.EOAAddress, address.String()), attribute.String("name", name))

	if _, err := meter.RegisterCallback(func(parentCtx context.Context, o metric.Observer) (err error) {
		ctx, span := e.metrics.Tracer().Start(parentCtx, "relayer_stats", trace.WithAttributes(
			attribute.Int(metrics.ChainID, chainID),
			attribute.String(metrics.EOAAddress, address.String()),
		))

		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		client, err := e.omnirpcClient.GetConfirmationsClient(ctx, chainID, 1)
		if err != nil {
			return fmt.Errorf("could not get confirmations client: %w", err)
		}

		var nonce uint64
		var balance big.Int

		err = client.BatchWithContext(ctx,
			eth.Nonce(address, nil).Returns(&nonce),
			eth.Balance(address, nil).Returns(&balance),
		)

		if err != nil {
			return fmt.Errorf("could not get balance: %w", err)
		}

		ethBalance := new(big.Float).Quo(new(big.Float).SetInt(&balance), new(big.Float).SetInt64(params.Ether))
		truncEthBalance, _ := ethBalance.Float64()

		o.ObserveFloat64(balanceGauge, truncEthBalance, metric.WithAttributeSet(attributes))
		o.ObserveInt64(nonceGauge, int64(nonce), metric.WithAttributeSet(attributes))

		return nil
	}, balanceGauge, nonceGauge); err != nil {
		return fmt.Errorf("registering callback on instruments: %w", err)
	}

	return nil
}
