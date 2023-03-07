package types

import "math/big"

// State is the state interface.
type State interface {
	// Root is the root of the Origin Merkle Tree.
	Root() [32]byte
	// Origin is the domain where Origin is located.
	Origin() uint32
	// Nonce is the amount of dispatched messages.
	Nonce() uint32
	// BlockNumber is the block of the last dispatched message.
	BlockNumber() *big.Int
	// Timestamp is the unix time of the last dispatched message.
	Timestamp() *big.Int
}

type state struct {
	root        [32]byte
	origin      uint32
	nonce       uint32
	blockNumber *big.Int
	timestamp   *big.Int
}

// NewState creates a new state.
func NewState(root [32]byte, origin uint32, nonce uint32, blockNumber *big.Int, timestamp *big.Int) State {
	return &state{
		root:        root,
		origin:      origin,
		nonce:       nonce,
		blockNumber: blockNumber,
		timestamp:   timestamp,
	}
}

func (s state) Root() [32]byte {
	return s.root
}

func (s state) Origin() uint32 {
	return s.origin
}

func (s state) Nonce() uint32 {
	return s.nonce
}

func (s state) BlockNumber() *big.Int {
	return s.blockNumber
}

func (s state) Timestamp() *big.Int {
	return s.timestamp
}

var _ State = state{}
