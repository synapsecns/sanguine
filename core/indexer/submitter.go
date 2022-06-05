package indexer

import (
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

type UpdateSubmitter struct {
	// domain allows access to the home contract
	domain domains.DomainClient
	// db contains the db object
	db db.DB
	// intervalSeconds adds an interval
	intervalSeconds uint64
	// signer is the signer
	signer signer.Signer
}

// todo: next up you need to borrow the tx loop from synapse-node and test well
// myabe in ethergo? Should be agnostic and utilize nonce manager
