package p2p

import "github.com/ethereum/go-ethereum/common"

type TXAnnouncement struct {
	Transaction common.Hash `json:"transaction"`
	ChainID     int         `json:"chain_id"`
}
