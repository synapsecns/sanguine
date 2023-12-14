package fastbridge

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// TODO: consider not exporting to avoid accidental mutation
var (
	// BridgeRequestedTopic is the event topic for a bridge request
	BridgeRequestedTopic common.Hash
	// BridgeRelayedTopic is the topic emitted by a bridge relay
	BridgeRelayedTopic common.Hash
)

// static checks to make sure topics actually exist
func init() {
	var err error

	parsedABI, err := abi.JSON(strings.NewReader(FastBridgeMetaData.ABI))
	if err != nil {
		panic(err)
	}

	BridgeRequestedTopic = parsedABI.Events["BridgeRequested"].ID
	BridgeRelayedTopic = parsedABI.Events["BridgeRelayed"].ID

	_, err = parsedABI.EventByID(BridgeRequestedTopic)
	if err != nil {
		panic(err)
	}

	_, err = parsedABI.EventByID(BridgeRelayedTopic)
	if err != nil {
		panic(err)
	}

}
