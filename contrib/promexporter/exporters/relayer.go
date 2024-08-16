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
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/decoders"
	rfqAPIModel "github.com/synapsecns/sanguine/services/rfq/api/model"
)

// TODO: This is ugly. We can probably get this from the config.
var usdcAddresses = map[int]string{
	1:      "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	10:     "0x0b2c639c533813f4aa9d7837caf62653d097ff85",
	42161:  "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
	8453:   "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
	534352: "0x06eFdBFf2a14a7c8E15944D1F4A48F9F95F663A4",
	59144:  "0x176211869cA2b568f2A7D4EE941E073a821EE1ff",
}

// TODO: this function does too many things.
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

	for chainID, relayers := range chainIDToRelayers {
		client, err := e.omnirpcClient.GetConfirmationsClient(ctx, chainID, 1)
		if err != nil {
			return fmt.Errorf("could not get confirmations client: %w", err)
		}

		var relayerBalances []*big.Int
		var usdcBalances []*big.Int
		for range relayers {
			relayerBalances = append(relayerBalances, new(big.Int))
			usdcBalances = append(usdcBalances, new(big.Int))
		}

		var callsForCurrentChainID []w3types.Caller
		for i, relayer := range relayers {
			callsForCurrentChainID = append(
				callsForCurrentChainID,
				eth.Balance(common.HexToAddress(relayer), nil).Returns(relayerBalances[i]),
			)
			callsForCurrentChainID = append(
				callsForCurrentChainID,
				eth.CallFunc(
					decoders.FuncBalanceOf(),
					common.HexToAddress(usdcAddresses[chainID]),
					common.HexToAddress(relayer)).Returns(usdcBalances[i]),
			)
		}

		_ = e.batchCalls(ctx, client, callsForCurrentChainID)

		for i := range relayerBalances {
			balanceFloat, _ := new(big.Float).SetInt(relayerBalances[i]).Float64()
			usdcBalanceFloat, _ := new(big.Float).SetInt(usdcBalances[i]).Float64()
			relayerMetadata := relayerMetadata{
				address:     common.HexToAddress(relayers[i]),
				balance:     balanceFloat / params.Ether,
				usdcBalance: usdcBalanceFloat / 1e6,
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
