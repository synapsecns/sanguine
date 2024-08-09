package util

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.opentelemetry.io/otel/attribute"
)

// EthAddress is the address of a chain's native gas token.
var EthAddress = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

// QuoteRequestToAttributes converts a quote request to attributes.
func QuoteRequestToAttributes(request reldb.QuoteRequest) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.Int64("block_number", int64(request.BlockNumber)),
		attribute.Int64("origin_token_decimals", int64(request.OriginTokenDecimals)),
		attribute.Int64("dest_token_decimals", int64(request.DestTokenDecimals)),
		attribute.String("transaction_id", fmt.Sprintf("%x", request.TransactionID)),
		attribute.String("sender", request.Sender.Hex()),
		attribute.String("origin_chain_id", fmt.Sprint(request.Transaction.OriginChainId)),
		attribute.String("origin_token", request.Transaction.OriginToken.Hex()),
		attribute.String("origin_amount", request.Transaction.OriginAmount.String()),
		attribute.String("status", request.Status.String()),
		attribute.String("origin_tx_hash", request.OriginTxHash.Hex()),
		attribute.Int64("relay_nonce", int64(request.RelayNonce)),
	}
}

// IsGasToken returns true if the given token is the gas token.
func IsGasToken(token common.Address) bool {
	return token == EthAddress
}
