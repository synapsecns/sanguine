package bridge

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/agents/types"
)

// BridgeRef is a bound synapse bridge config v2 contract that returns the address of that contract
// nolint: golint
type BridgeRef struct {
	*SynapseBridge
	address common.Address
}

// Address is the contract address.
func (s BridgeRef) Address() common.Address {
	return s.address
}

// NewBridgeRef gets a bound synapse bridge config contract that returns the address of the contract
// nolint: golint
func NewBridgeRef(address common.Address, backend bind.ContractBackend) (*BridgeRef, error) {
	bridge, err := NewSynapseBridge(address, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeRef{
		SynapseBridge: bridge,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &BridgeRef{}

// GetKappa returns the kappa for a cross-chain event log.
func GetKappa(event types.CrossChainUserEventLog) (kappa [32]byte) {
	// get the kappa as a slice
	return KappaFromIdentifier(event.GetIdentifier())
}

// KappaFromTX gets the kappa value from the tx.
// Deprecated: used only for historical parity.
func KappaFromTX(txhash common.Hash) (kappa [32]byte) {
	// get the kappa as a slice
	rawKappa := crypto.Keccak256([]byte(txhash.String()))
	copy(kappa[:], rawKappa)
	return kappa
}

// KappaToSlice converts a kappa value toa  byte slice.
func KappaToSlice(kappa [32]byte) []byte {
	rawKappa := make([]byte, len(kappa))
	copy(rawKappa, kappa[:])
	return rawKappa
}

// KappaFromIdentifier derive sa kappa from a string identifier.
func KappaFromIdentifier(identifier string) (kappa [32]byte) {
	rawKappa := crypto.Keccak256([]byte(identifier))
	copy(kappa[:], rawKappa)
	return kappa
}
