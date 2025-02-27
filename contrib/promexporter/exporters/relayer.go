package exporters

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/decoders"
)

var relayerAddresses = []string{
	"0xDc927Bd56CF9DfC2e3779C7E3D6d28dA1C219969",
	"0xDD50676F81f607fD8bA7Ed3187DdF172DB174CD3",
	"0xbe75079fd259a82054cAAB2CE007cd0c20b177a8",
	"0x2156BfA195C033CA2DF4Ff14e6Da0c617B8cb4F7",
}
var usdcAddresses = map[int]string{
	1:      "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", // mainnet
	10:     "0x0b2c639c533813f4aa9d7837caf62653d097ff85", // optimism
	42161:  "0xaf88d065e77c8cC2239327C5EDb3A432268e5831", // arbitrum
	8453:   "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913", // base
	534352: "0x06eFdBFf2a14a7c8E15944D1F4A48F9F95F663A4", // scroll
	59144:  "0x176211869cA2b568f2A7D4EE941E073a821EE1ff", // linea
	480:    "0x79A02482A880bCE3F13e09Da970dC34db4CD24d1", // world
	80094:  "0x549943e04f40284185054145c6e4e9568c1d3241", // bera
	130:    "0x078d782b760474a361dda0af3839290b0ef57ad6", // unichain

}

//nolint:cyclop
func (e *exporter) fetchRelayerBalances(ctx context.Context, _ string) error {
	chainIDToRelayers := make(map[int][]string)
	for chainid := range usdcAddresses {
		chainIDToRelayers[chainid] = relayerAddresses
	}

	for chainID, relayers := range chainIDToRelayers {
		client, err := e.omnirpcClient.GetChainClient(ctx, chainID)
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
			// fmt.Printf("chainid=%d, address=%s, balance=%f, usdcBalance=%f\n", chainID, relayers[i], balanceFloat/params.Ether, usdcBalanceFloat/1e6)
			e.otelRecorder.RecordRelayerBalance(chainID, relayerMetadata)
		}
		// fmt.Println("-----------------------")
	}

	return nil
}

func chainIDsToRelayers() map[int][]string {
	quotesMap := make(map[int][]string)

	resp, err := http.Get("https://rfq-api.omnirpc.io/quotes")
	if err != nil {
		fmt.Printf("Error fetching quotes: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil
	}

	var quotes []struct {
		OriginChainID int    `json:"origin_chain_id"`
		RelayerAddr   string `json:"relayer_addr"`
		DestChainID   int    `json:"dest_chain_id"`
	}

	if err := json.Unmarshal(body, &quotes); err != nil {
		fmt.Printf("Error unmarshalling quotes: %v\n", err)
		return nil
	}

	for _, quote := range quotes {
		quotesMap[quote.OriginChainID] = append(quotesMap[quote.OriginChainID], quote.RelayerAddr)
		quotesMap[quote.DestChainID] = append(quotesMap[quote.DestChainID], quote.RelayerAddr)
	}

	return quotesMap

}
