package db

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"math/big"
)

type TransactionSent struct {
	TransactionID common.Hash
	// EncodedTX is the encoded transaction.
	EncodedTX []byte
	// SrcChainID is the source chain id.
	SrcChainID *big.Int
	// DstChainID is the destination chain id.
	DstChainID *big.Int
	// Status is the status of the transaction.
	Status ExecutableStatus
	// Options is the options of the transaction.
	Options interchainclient.OptionsV1
}
