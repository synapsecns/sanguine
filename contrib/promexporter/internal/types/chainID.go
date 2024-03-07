// Package types contains types specific to prom exporting.
//
// this comes from: https://github.com/synapsecns/sanguine/blob/5f5f6779a33ff40ad099a43874d401f9c9dfe02a/packages/synapse-interface/constants/chains/index.tsx and needs to be manually updated
// when https://github.com/synapsecns/sanguine/issues/1586 is implemented, this can be removed
// nolint: golint, revive
package types

// ChainID contains the chain id.
type ChainID int

// Int returns the int value of the chain id.
func (c ChainID) Int() int {
	return int(c)
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=ChainID

const (
	ETH       ChainID = 1
	ROPSTEN   ChainID = 3
	RINKEBY   ChainID = 4
	GOERLI    ChainID = 5
	OPTIMISM  ChainID = 10
	CRONOS    ChainID = 25
	KOVAN     ChainID = 42
	BSC       ChainID = 56
	POLYGON   ChainID = 137
	FANTOM    ChainID = 250
	BOBA      ChainID = 288
	METIS     ChainID = 1088
	MOONBEAM  ChainID = 1284
	MOONRIVER ChainID = 1285
	DOGECHAIN ChainID = 2000
	CANTO     ChainID = 7700
	KLAYTN    ChainID = 8217
	ARBITRUM  ChainID = 42161
	BASE      ChainID = 8453
	AVALANCHE ChainID = 43114
	DFK       ChainID = 53935
	AURORA    ChainID = 1313161554
	HARMONY   ChainID = 1666600000
)

// ToInts converts a list of chainIDs to a slice of ints.
// please note: this will not remove duplicate chainids.
func ToInts(ids ...ChainID) (out []int) {
	for _, id := range ids {
		out = append(out, int(id))
	}
	return out
}
