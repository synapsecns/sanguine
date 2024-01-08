package base

import (
	"github.com/puzpuzpuz/xsync/v2"
	"github.com/synapsecns/sanguine/ethergo/contracts"
)

// blacklist is a global used for skipping contract verification
// TODO: we should actually have a contractInfo not available state, but considering
// the tenderly integration will probably be removed and otterscan self verification is not yet available
// this seems likely likely to be removed entirely.
var blacklist *xsync.MapOf[string, bool]

func init() {
	blacklist = xsync.NewMapOf[bool]()
}

// AddToVerificationBlacklist adds a contract to a verification blacklist for the remained of the process.
// this should be used sparingly or not at all. This only exists for USDT since it has no combined json owing to
// it coming from solidity 0.4.0
//
// This method is not documented to discourage use.
func AddToVerificationBlacklist(c contracts.ContractType) {
	blacklist.Store(c.Name(), true)
}

// IsVerificationBlacklisted checks if a contract is blacklisted for verification.
func IsVerificationBlacklisted(c contracts.ContractType) bool {
	v, ok := blacklist.Load(c.Name())
	return v && ok
}
