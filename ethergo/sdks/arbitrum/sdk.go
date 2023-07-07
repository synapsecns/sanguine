package arbitrum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/synapsecns/sanguine/ethergo/sdks/arbitrum/contracts/arbgasinfo"
	"github.com/synapsecns/sanguine/ethergo/sdks/arbitrum/contracts/nodeinterface"
)

type ArbitrumSDK interface {
	// TODO:
}

type arbitrumSDKImpl struct {
	client        bind.ContractBackend
	nodeInterface nodeinterface.INodeInterface
	gasInfo       arbgasinfo.IArbGasInfo
}

// NewArbitrumSDK creates a new ArbitrumSDK.
func NewArbitrumSDK(client bind.ContractBackend, options ...ArbitrumOption) (ArbitrumSDK, error) {
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	nodeInterface, err := nodeinterface.NewNodeInterfaceRef(opts.nodeInterfaceAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create node interface: %w", err)
	}

	gasInfo, err := arbgasinfo.NewArbGasInfo(opts.gasInfoAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create gas info: %w", err)
	}
	nodeInterface.GasEstimateComponents()
	return &arbitrumSDKImpl{
		client:        client,
		nodeInterface: nodeInterface,
		gasInfo:       gasInfo,
	}, nil
}
