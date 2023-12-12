package utils

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
)

// Encode encodes a bridge transaction into a request.
func Encode(bridgeTransaction *bindings.IFastBridgeBridgeTransaction) (request []byte, err error) {
	// encode from bridge transaction interface using abi
	parsedABI, err := abi.JSON(strings.NewReader(bindings.FastBridgeABI))
	if err != nil {
		fmt.Println("Failed to parse ABI:", err)
		return
	}

	method, exist := parsedABI.Methods["getBridgeTransaction"]
	if !exist {
		err = fmt.Errorf("method '%s' not found", "getBridgeTransaction")
		return
	}
	request, err = method.Outputs.Pack(bridgeTransaction)
	return
}

// Decode decodes a request into a bridge transaction.
func Decode(request []byte) (bridgeTransaction *bindings.IFastBridgeBridgeTransaction, err error) {
	// decode into bridge transaction interface using abi
	parsedABI, err := abi.JSON(strings.NewReader(bindings.FastBridgeABI))
	if err != nil {
		fmt.Println("Failed to parse ABI:", err)
		return
	}

	bridgeTransaction = new(bindings.IFastBridgeBridgeTransaction)
	if err = parsedABI.UnpackIntoInterface(&bridgeTransaction, "getBridgeTransaction", request); err != nil {
		fmt.Println("Error unpacking bridge transaction:", err)
		return
	}

	return
}

// TransactionID returns the transaction id for a given request.
func TransactionID(request []byte) common.Hash {
	return crypto.Keccak256Hash(request)
}

// IsBridgeRequested checks if the topic is the BridgeRequested event.
func IsBridgeRequested(topic common.Hash, abi abi.ABI) bool {
	return topic == abi.Events["BridgeRequested"].ID
}

// IsBridgeRelayed checks if the topic is the BridgeRelayed event.
func IsBridgeRelayed(topic common.Hash, abi abi.ABI) bool {
	return topic == abi.Events["BridgeRelayed"].ID
}

// ParseBridgeRequested parses the BridgeRequested event.
func ParseBridgeRequested(log types.Log, abi abi.ABI) (*bindings.FastBridgeBridgeRequested, error) {
	event := new(bindings.FastBridgeBridgeRequested)
	err := abi.UnpackIntoInterface(event, "BridgeRequested", log.Data)
	if err != nil {
		return nil, fmt.Errorf("could not unpack BridgeRequested event: %w", err)
	}

	return event, nil
}

// ParseBridgeRelayed parses the BridgeRelayed event.
func ParseBridgeRelayed(log types.Log, abi abi.ABI) (*bindings.FastBridgeBridgeRelayed, error) {
	event := new(bindings.FastBridgeBridgeRelayed)
	err := abi.UnpackIntoInterface(event, "BridgeRelayed", log.Data)
	if err != nil {
		return nil, fmt.Errorf("could not unpack BridgeRelayed event: %w", err)
	}
	return event, nil
}
