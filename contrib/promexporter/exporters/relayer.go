package exporters

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	rfqAPIModel "github.com/synapsecns/sanguine/services/rfq/api/model"
)

//nolint:cyclop
func (e *exporter) fetchRelayerBalances(ctx context.Context, url string) error {
	// Fetch relayer addresses
	quotes, err := e.fetchAllQuotes(ctx, url)
	if err != nil {
		return fmt.Errorf("could not fetch relayer addresses: %w", err)
	}

	// chainIDs is a map of chain ID to relayer addresses
	chainIDToRelayers := make(map[int][]string)

	// Get all chain IDs
	for _, quote := range quotes {
		if !slices.Contains(chainIDToRelayers[quote.OriginChainID], quote.RelayerAddr) {
			chainIDToRelayers[quote.OriginChainID] = append(chainIDToRelayers[quote.OriginChainID], quote.RelayerAddr)
		}

		if !slices.Contains(chainIDToRelayers[quote.DestChainID], quote.RelayerAddr) {
			chainIDToRelayers[quote.DestChainID] = append(chainIDToRelayers[quote.DestChainID], quote.RelayerAddr)
		}
	}

	for chainID := range chainIDToRelayers {
		chainIDToRelayers[chainID] = append(chainIDToRelayers[chainID], "0x2156BfA195C033CA2DF4Ff14e6Da0c617B8cb4F7")
	}

	for chainID, relayers := range chainIDToRelayers {
		client, err := e.omnirpcClient.GetConfirmationsClient(ctx, chainID, 1)
		if err != nil {
			return fmt.Errorf("could not get confirmations client: %w", err)
		}

		var relayerBalances []*big.Int
		for range relayers {
			relayerBalances = append(relayerBalances, new(big.Int))
		}

		var callsForCurrentChainID []w3types.Caller
		for i, relayer := range relayers {
			callsForCurrentChainID = append(
				callsForCurrentChainID,
				eth.Balance(common.HexToAddress(relayer), nil).Returns(relayerBalances[i]),
			)
		}

		_ = e.batchCalls(ctx, client, callsForCurrentChainID)

		for i, balanceOfRelayer := range relayerBalances {
			balanceFloat, _ := new(big.Float).SetInt(balanceOfRelayer).Float64()
			relayerMetadata := relayerMetadata{
				address: common.HexToAddress(relayers[i]),
				balance: balanceFloat / params.Ether,
			}
			// the line of interest, where we record each relayer data for the respective chainID
			e.otelRecorder.RecordRelayerBalance(chainID, relayerMetadata)
		}
	}

	return nil
}

func (e *exporter) fetchAllQuotes(ctx context.Context, url string) ([]rfqAPIModel.GetQuoteResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not get quotes: %w", err)
	}

	res, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get quotes: %w", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body: %w", err)
	}

	var quotes []rfqAPIModel.GetQuoteResponse
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal quotes: %w", err)
	}

	return quotes, nil
}
