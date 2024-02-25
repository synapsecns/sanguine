package p2p

import (
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
)

func ethAddrFromPeer(p peer.ID) (common.Address, error) {
	pubkey, err := p.ExtractPublicKey()
	if err != nil {
		return common.Address{}, errors.Wrap(err, "extract pub key")
	}

	dbytes, _ := pubkey.Raw()
	k, err := secp256k1.ParsePubKey(dbytes)

	if err != nil {
		return common.Address{}, fmt.Errorf("parse pubkey: %w", err)
	}

	return crypto.PubkeyToAddress(*k.ToECDSA()), nil
}
