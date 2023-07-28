package exporters

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/contrib/promexporter/config"
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/gql/dfk"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/instrumentation"
	"go.opentelemetry.io/otel/metric"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"time"
)

var logger = log.Logger("proxy-logger")

// makeHTTPClient makes a tracing http client.
func makeHTTPClient(handler metrics.Handler) *http.Client {
	httpClient := new(http.Client)
	handler.ConfigureHTTPClient(httpClient)

	httpClient.Transport = instrumentation.NewCaptureTransport(httpClient.Transport, handler)

	return httpClient
}

type exporter struct {
	client  *http.Client
	metrics metrics.Handler
	meter   metric.Meter
	cfg     config.Config
}

func StartExporterServer(ctx context.Context, handler metrics.Handler, cfg config.Config) error {
	router := ginhelper.New(logger)
	router.Use(handler.Gin())
	router.GET("/metrics", gin.WrapH(handler.Handler()))

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return fmt.Errorf("could not listen on port %d", cfg.Port)
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		//nolint: gosec
		// TODO: consider setting timeouts here:  https://ieftimov.com/posts/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/
		err := http.Serve(listener, router)
		if err != nil {
			return fmt.Errorf("could not serve http: %w", err)
		}

		return nil
	})

	metermaid := handler.Meter("github.com/synapsecns/sanguine/contrib/promexporter/exporters")

	exp := exporter{
		client:  makeHTTPClient(handler),
		metrics: handler,
		meter:   metermaid,
		cfg:     cfg,
	}

	// heroes on both chains
	err = exp.stuckHeroCount(common.HexToAddress("0x739B1666c2956f601f095298132773074c3E184b"), "dfk")
	if err != nil {
		return fmt.Errorf("could setup metric: %w", err)
	}
	err = exp.stuckHeroCount(common.HexToAddress("0xEE258eF5F4338B37E9BA9dE6a56382AdB32056E2"), "klatyn")
	if err != nil {
		return fmt.Errorf("could setup metric: %w", err)
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not start exporter server: %w", err)
	}

	return nil
}

const stuckHeroMetric = "stuck_heroes_"

func (e *exporter) stuckHeroCount(owner common.Address, chainName string) error {
	stuckCount, err := e.meter.Int64ObservableGauge(fmt.Sprintf("%s%s", stuckHeroMetric, chainName))
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	if _, err := e.meter.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()

		dfkClient := dfk.NewClient(e.client, e.cfg.DFKUrl)

		stuckHeroes, err := dfkClient.StuckHeroes(ctx, core.PtrTo[int64](0), core.PtrTo(owner.String()))
		if err != nil {
			return fmt.Errorf("could not get stuck hero count: %w", err)
		}

		// TODO: this maxes out at 100 now. Need binary search or something.
		o.ObserveInt64(stuckCount, int64(len(stuckHeroes.Heroes)))

		return nil
	}, stuckCount); err != nil {
		return fmt.Errorf("registering callback on instruments: %s", err)
	}

	return nil
}
