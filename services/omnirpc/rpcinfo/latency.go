package rpcinfo

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/lmittmann/w3/module/eth"
	"github.com/synapsecns/sanguine/core/metrics"
	ethClient "github.com/synapsecns/sanguine/ethergo/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
	"net/url"
	"sync"
	"time"
)

var logger = log.Logger("rpcinfo-logger")

// Result is the result of a latency check on a url.
type Result struct {
	// URL is the url of the latency being tested
	URL string
	// Latency is the latency time in seconds
	Latency time.Duration
	// BlockAge is the age of the block
	BlockAge time.Duration
	// HasError is wether or not the result has an error
	HasError bool
	// Error is the error recevied when trying to establish latency
	Error error
}

// GetRPCLatency gets latency from a list of rpcs.
func GetRPCLatency(parentCtx context.Context, timeout time.Duration, rpcList []string, handler metrics.Handler) (latSlice []Result) {
	var mux sync.Mutex

	timeCtx, cancel := context.WithTimeout(parentCtx, timeout)

	traceCtx, span := handler.Tracer().Start(timeCtx, "rpcinfo.GetRPCLatency", trace.WithAttributes(attribute.StringSlice("rpcList", rpcList)))
	defer func() {
		metrics.EndSpan(span)
		cancel()
	}()

	g, ctx := errgroup.WithContext(traceCtx)
	for _, rpcURL := range rpcList {
		// capture func literal
		rpcURL := rpcURL
		g.Go(func() error {
			latency := getLatency(ctx, rpcURL, handler)

			mux.Lock()
			latSlice = append(latSlice, latency)
			mux.Unlock()

			return nil
		})
	}

	// we don't error at all above
	_ = g.Wait()
	return latSlice
}

const meter = "github.com/synapsecns/sanguine/services/omnirpc/rpcinfo"
const blockNumber = "block_number"

func getLatency(ctx context.Context, rpcURL string, handler metrics.Handler) (l Result) {
	l = Result{URL: rpcURL, HasError: true}

	parsedURL, err := url.Parse(rpcURL)
	if err != nil {
		l.Error = fmt.Errorf("url invalid: %w", err)
		return l
	}

	// maybe we should allow this?
	if slices.Contains([]string{"ws", "wss"}, parsedURL.Scheme) {
		l.Error = errors.New("websockets not supported")
		return l
	}

	startTime := time.Now()

	client, err := ethClient.DialBackend(ctx, rpcURL, handler)
	if err != nil {
		l.Error = fmt.Errorf("could not create client: %w", err)
		return l
	}

	var chainID uint64
	var latestHeader types.Header

	err = client.BatchWithContext(ctx,
		eth.ChainID().Returns(&chainID),
		eth.HeaderByNumber(nil).Returns(&latestHeader),
	)

	if err != nil {
		l.Error = err
		l.HasError = true
		return l
	}

	endTime := time.Now()

	l.Latency = endTime.Sub(startTime)

	l.BlockAge = endTime.Sub(time.Unix(int64(latestHeader.Time), 0))

	l.HasError = false

	err = recordMetrics(ctx, handler, rpcURL, chainID, &latestHeader, l)
	if err != nil {
		logger.Warnf("could not record metrics: %w", err)
	}

	return l
}

// recordMetrics records metrics for a given url.
func recordMetrics(ctx context.Context, handler metrics.Handler, url string, chainID uint64, block *types.Header, r Result) error {
	attributeSet := attribute.NewSet(attribute.Int64(metrics.ChainID, int64(chainID)), attribute.String("rpc_url", url))

	blockNumberMetric, err := handler.Meter(meter).Int64Histogram(blockNumber)
	if err != nil {
		return fmt.Errorf("could not create histogram: %w", err)
	}

	blockNumberMetric.Record(ctx, block.Number.Int64(), metric.WithAttributeSet(attributeSet))

	latencyMetric, err := handler.Meter(meter).Float64Histogram("latency", metric.WithUnit("seconds"))
	if err != nil {
		return fmt.Errorf("could not create histogram: %w", err)
	}

	latencyMetric.Record(ctx, r.Latency.Seconds(), metric.WithAttributeSet(attributeSet))

	blockAgeMetric, err := handler.Meter(meter).Float64Histogram("block_age", metric.WithUnit("seconds"))
	if err != nil {
		return fmt.Errorf("could not create histogram: %w", err)
	}

	blockAgeMetric.Record(ctx, r.BlockAge.Seconds(), metric.WithAttributeSet(attributeSet))
	if err != nil {
		return fmt.Errorf("could not create histogram: %w", err)
	}
	return nil
}
