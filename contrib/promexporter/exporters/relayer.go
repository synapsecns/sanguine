package exporters

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	rfqAPIModel "github.com/synapsecns/sanguine/services/rfq/api/model"
)

func (e *exporter) fetchRelayerBalances(ctx context.Context) error {

	// Fetch relayer addresses
	quotes := fetchRelayerAddresses()

	// chainIDs is a map of chain ID to relayer addresses
	chainIDToRelayers := make(map[int][]string)

	// Get all chain IDs
	for _, quote := range quotes {
		chainIDToRelayers[quote.OriginChainID] = append(chainIDToRelayers[quote.OriginChainID], quote.RelayerAddr)
		chainIDToRelayers[quote.DestChainID] = append(chainIDToRelayers[quote.DestChainID], quote.RelayerAddr)
	}

	for chainID, relayers := range chainIDToRelayers {
		client, err := e.omnirpcClient.GetConfirmationsClient(ctx, chainID, 1)
		if err != nil {
			return fmt.Errorf("could not get confirmations client: %w", err)
		}

		var callsForCurrentChainID []w3types.Caller
		relayerBalances := make([]*big.Int, 0, len(relayers))
		for i, relayer := range relayers {
			callsForCurrentChainID = append(
				callsForCurrentChainID,
				eth.Balance(common.HexToAddress(relayer), nil).Returns(relayerBalances[i]),
			)
		}

		_ = e.batchCalls(ctx, client, callsForCurrentChainID)

		for _, balanceOfRelayer := range relayerBalances {
			balanceFloat, _ := new(big.Float).SetInt(balanceOfRelayer).Float64()
			e.otelRecorder.RecordRelayerBalance(chainID, balanceFloat)
		}
	}

	return nil
}

func fetchRelayerAddresses() []rfqAPIModel.GetQuoteResponse {
	url := "https://rfq-api.omnirpc.io/quotes"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	var quotes []rfqAPIModel.GetQuoteResponse
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	return quotes
}
