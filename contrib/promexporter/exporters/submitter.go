package exporters

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/lmittmann/w3/module/eth"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// note: this kind of check should be deprecated in favor of submitter metrics once everything has been moved over.
func (e *exporter) submitterStats(address common.Address, chainID int, name string) (err error) {
	ctx, span := e.metrics.Tracer().Start(
		context.Background(),
		"submitter_stats",
		trace.WithAttributes(
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

	if err = client.BatchWithContext(ctx,
		eth.Nonce(address, nil).Returns(&nonce),
		eth.Balance(address, nil).Returns(&balance),
	); err != nil {
		return fmt.Errorf("could not get balance: %w", err)
	}

	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(&balance), new(big.Float).SetInt64(params.Ether))
	truncEthBalance, _ := ethBalance.Float64()

	submitterMetadata := submitterMetadata{
		address: address,
		name:    name,
		nonce:   int64(nonce),
		balance: truncEthBalance,
	}

	e.otelRecorder.RecordSubmitterStats(chainID, submitterMetadata)

	return nil
}
