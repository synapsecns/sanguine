package exporters

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hedzr/log"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// old content replaced: no more "hashmap" usage

type submitterMetadata struct {
	address common.Address
	name    string
	nonce   int64
	balance float64
}

type relayerMetadata struct {
	address     common.Address
	balance     float64
	usdcBalance float64
}

type tokenData struct {
	metadata        TokenConfig
	contractBalance *big.Int
	totalSuppply    *big.Int
	feeBalance      *big.Int
}

type otelRecorder struct {
	metrics metrics.Handler
	meter   metric.Meter

	// for vprice: map[int]float64
	vPrice sync.Map

	vpriceGauge metric.Float64ObservableGauge

	// for bridge: map[int]*sync.Map (the inner sync.Map is map[string]tokenData)
	td sync.Map

	// chainID -> float64
	gasBalance sync.Map

	gasBalanceGauge    metric.Float64ObservableGauge
	bridgeBalanceGauge metric.Float64ObservableGauge
	feeBalanceGauge    metric.Float64ObservableGauge
	totalSupplyGauge   metric.Float64ObservableGauge

	// dfk stats: map[string]int64
	stuckHeroes sync.Map

	stuckHeroesGauge metric.Int64ObservableGauge

	// submitter stats: map[int]submitterMetadata
	submitters sync.Map

	balanceGauge metric.Float64ObservableGauge
	nonceGauge   metric.Int64ObservableGauge

	// relayer stats: map[int]*sync.Map (the inner sync.Map is map[string]relayerMetadata)
	relayerBalance          sync.Map
	relayerBalanceGauge     metric.Float64ObservableGauge
	relayerUSDCBalanceGuage metric.Float64ObservableGauge
}

func newOtelRecorder(meterHandler metrics.Handler) iOtelRecorder {
	otr := &otelRecorder{
		metrics: meterHandler,
		meter:   meterHandler.Meter(meterName),
	}

	var err error
	if otr.vpriceGauge, err = otr.meter.Float64ObservableGauge("vpriceMetric"); err != nil {
		log.Warnf("failed to create vprice gauge: %v", err)
	}

	if otr.bridgeBalanceGauge, err = otr.meter.Float64ObservableGauge("bridgeBalanceMetric"); err != nil {
		log.Warnf("failed to create bridgeBalance gauge: %v", err)
	}

	if otr.feeBalanceGauge, err = otr.meter.Float64ObservableGauge("feeBalance_total"); err != nil {
		log.Warnf("failed to create feeBalance gauge: %v", err)
	}

	if otr.totalSupplyGauge, err = otr.meter.Float64ObservableGauge("totalSupply"); err != nil {
		log.Warnf("failed to create totalSupply gauge: %v", err)
	}

	if otr.gasBalanceGauge, err = otr.meter.Float64ObservableGauge("gasBalance"); err != nil {
		log.Warnf("failed to create gasBalance gauge: %v", err)
	}

	if otr.balanceGauge, err = otr.meter.Float64ObservableGauge("gas_balance"); err != nil {
		log.Warnf("failed to create balance gauge: %v", err)
	}

	if otr.nonceGauge, err = otr.meter.Int64ObservableGauge("nonce"); err != nil {
		log.Warnf("failed to create nonce gauge: %v", err)
	}

	if otr.stuckHeroesGauge, err = otr.meter.Int64ObservableGauge("dfk_pending_heroes"); err != nil {
		log.Warnf("failed to create stuckHeroes gauge: %v", err)
	}

	if otr.relayerBalanceGauge, err = otr.meter.Float64ObservableGauge("relayer_balance"); err != nil {
		log.Warnf("failed to create relayerBalance gauge: %v", err)
	}

	if otr.relayerUSDCBalanceGuage, err = otr.meter.Float64ObservableGauge("relayer_usdc_balance"); err != nil {
		log.Warnf("failed to create relayerUSDCBalance gauge: %v", err)
	}

	// register your callbacks
	if _, err = otr.meter.RegisterCallback(otr.recordVpriceGauge, otr.vpriceGauge); err != nil {
		log.Warnf("failed to register callback for vprice metrics: %v", err)
	}

	if _, err = otr.meter.RegisterCallback(otr.recordStuckHeroCount, otr.stuckHeroesGauge); err != nil {
		log.Warnf("failed to register callback for dfk stuck heroes metrics: %v", err)
	}

	if _, err = otr.meter.RegisterCallback(
		otr.recordTokenBalance,
		otr.bridgeBalanceGauge,
		otr.feeBalanceGauge,
		otr.totalSupplyGauge,
	); err != nil {
		log.Warnf("failed to register callback for bridge metrics : %v", err)
	}

	if _, err = otr.meter.RegisterCallback(otr.recordSubmitterStats, otr.balanceGauge, otr.nonceGauge); err != nil {
		log.Warnf("failed to register callback for submitter metrics: %v", err)
	}

	if _, err = otr.meter.RegisterCallback(otr.recordBridgeGasBalance, otr.gasBalanceGauge); err != nil {
		log.Warnf("failed to register callback for bridge gas balance metrics: %v", err)
	}

	if _, err = otr.meter.RegisterCallback(otr.recordRelayerBalance, otr.relayerBalanceGauge, otr.relayerUSDCBalanceGuage); err != nil {
		log.Warnf("failed to register callback for relayer balance metrics: %v", err)
	}

	return otr
}

// example usage for vPrice: store chainID -> float64
func (o *otelRecorder) RecordVPrice(chainid int, vPrice float64) {
	o.vPrice.Store(chainid, vPrice)
}

func (o *otelRecorder) recordVpriceGauge(_ context.Context, observer metric.Observer) error {
	o.vPrice.Range(func(k, v any) bool {
		chainid := k.(int)
		price := v.(float64)

		observer.ObserveFloat64(
			o.vpriceGauge,
			price,
			metric.WithAttributes(attribute.Int(metrics.ChainID, chainid)),
		)
		return true
	})
	return nil
}

// gasBalance: store chainID -> float64
func (o *otelRecorder) RecordBridgeGasBalance(chainid int, gasBalance float64) {
	o.gasBalance.Store(chainid, gasBalance)
}

func (o *otelRecorder) recordBridgeGasBalance(_ context.Context, observer metric.Observer) error {
	o.gasBalance.Range(func(k, v any) bool {
		chainid := k.(int)
		gb := v.(float64)
		observer.ObserveFloat64(
			o.gasBalanceGauge,
			gb,
			metric.WithAttributes(attribute.Int(metrics.ChainID, chainid)),
		)
		return true
	})
	return nil
}

// tokenData: map[int]*sync.Map, then inside that is map[string]tokenData
func (o *otelRecorder) RecordTokenBalance(chainID int, tData tokenData) {
	val, _ := o.td.LoadOrStore(chainID, &sync.Map{})
	chainMap := val.(*sync.Map)
	chainMap.Store(tData.metadata.TokenID, tData)
}

func (o *otelRecorder) recordTokenBalance(_ context.Context, observer metric.Observer) error {
	o.td.Range(func(k, v any) bool {
		// chainID := k.(int)
		chainMap := v.(*sync.Map)

		chainMap.Range(func(k2, v2 any) bool {
			// tokenID := k2.(string)
			td := v2.(tokenData)

			tokenAttributes := attribute.NewSet(
				attribute.String("tokenID", td.metadata.TokenID),
				attribute.Int(metrics.ChainID, td.metadata.ChainID),
			)

			bridgeBalance := core.BigToDecimals(td.contractBalance, td.metadata.TokenDecimals)
			observer.ObserveFloat64(
				o.bridgeBalanceGauge,
				bridgeBalance,
				metric.WithAttributeSet(tokenAttributes),
			)

			feeBalance := core.BigToDecimals(td.feeBalance, td.metadata.TokenDecimals)
			observer.ObserveFloat64(
				o.feeBalanceGauge,
				feeBalance,
				metric.WithAttributeSet(tokenAttributes),
			)

			totalSupply := core.BigToDecimals(td.totalSuppply, td.metadata.TokenDecimals)
			observer.ObserveFloat64(
				o.totalSupplyGauge,
				totalSupply,
				metric.WithAttributeSet(tokenAttributes),
			)

			return true
		})

		return true
	})
	return nil
}

// stuckHeroes: map[string]int64
func (o *otelRecorder) RecordStuckHeroCount(stuck int64, chainname string) {
	o.stuckHeroes.Store(chainname, stuck)
}

func (o *otelRecorder) recordStuckHeroCount(_ context.Context, observer metric.Observer) error {
	o.stuckHeroes.Range(func(k, v any) bool {
		chainName := k.(string)
		count := v.(int64)

		observer.ObserveInt64(
			o.stuckHeroesGauge,
			count,
			metric.WithAttributes(attribute.String("chain_name", chainName)),
		)
		return true
	})
	return nil
}

// submitters: map[int]submitterMetadata
func (o *otelRecorder) RecordSubmitterStats(chainid int, metadata submitterMetadata) {
	o.submitters.Store(chainid, metadata)
}

func (o *otelRecorder) recordSubmitterStats(_ context.Context, observer metric.Observer) error {
	o.submitters.Range(func(k, v any) bool {
		chainID := k.(int)
		submitter := v.(submitterMetadata)

		observer.ObserveInt64(
			o.nonceGauge,
			submitter.nonce,
			metric.WithAttributes(
				attribute.Int(metrics.ChainID, chainID),
				attribute.String(metrics.EOAAddress, submitter.address.String()),
				attribute.String("name", submitter.name),
			),
		)

		observer.ObserveFloat64(
			o.balanceGauge,
			submitter.balance,
			metric.WithAttributes(
				attribute.Int(metrics.ChainID, chainID),
				attribute.String(metrics.EOAAddress, submitter.address.String()),
				attribute.String("name", submitter.name),
			),
		)

		return true
	})
	return nil
}

// relayerBalance: map[int]*sync.Map, then inside that map[string]relayerMetadata
func (o *otelRecorder) RecordRelayerBalance(chainID int, relayer relayerMetadata) {
	val, _ := o.relayerBalance.LoadOrStore(chainID, &sync.Map{})
	chainMap := val.(*sync.Map)

	fmt.Printf("Setting chainid=%d, address=%s, balance=%f, usdcBalance=%f\n",
		chainID, relayer.address.String(), relayer.balance, relayer.usdcBalance)

	chainMap.Store(relayer.address.String(), relayer)

	// debug: show what's in there
	fmt.Println("-------------------------------")
	fmt.Println("This is what is currently stored in relayerBalance:")
	o.relayerBalance.Range(func(k, v any) bool {
		id := k.(int)
		m := v.(*sync.Map)
		fmt.Printf("chainid=%d\n", id)
		m.Range(func(k2, v2 any) bool {
			rAddr := k2.(string)
			rMeta := v2.(relayerMetadata)
			fmt.Printf("\t\taddress=%s, balance=%f, usdcBalance=%f\n", rAddr, rMeta.balance, rMeta.usdcBalance)
			return true
		})
		fmt.Println("=========")
		return true
	})
	fmt.Println("-------------------------------")
}

func (o *otelRecorder) recordRelayerBalance(_ context.Context, observer metric.Observer) error {
	fmt.Println(">>> recordRelayerBalance callback triggered <<<")

	o.relayerBalance.Range(func(k, v any) bool {
		chainID := k.(int)
		fmt.Println("in callback: chainID =", chainID)
		chainMap := v.(*sync.Map)

		chainMap.Range(func(k2, v2 any) bool {
			addr := k2.(string)
			meta := v2.(relayerMetadata)

			observer.ObserveFloat64(
				o.relayerBalanceGauge,
				meta.balance,
				metric.WithAttributes(
					attribute.Int(metrics.ChainID, chainID),
					attribute.String("relayer_address", addr),
				),
			)
			observer.ObserveFloat64(
				o.relayerUSDCBalanceGuage,
				meta.usdcBalance,
				metric.WithAttributes(
					attribute.Int(metrics.ChainID, chainID),
					attribute.String("relayer_address", addr),
				),
			)

			return true
		})
		return true
	})
	return nil
}
