package relayer

import "fmt"

// TODO: this file can be moved to somewhere common

// Mainnet chain IDs.
//
//nolint:deadcode
const ethereumChainID = 1
const optimismChainID = 10
const arbitrumChainID = 42161

// Testnet chain IDs.
//
//nolint:deadcode
const sepoliaChainID = 11155111
const opSepoliaChainID = 11155420
const arbitrumSepoliaChainID = 421614

func isTestnetChainID(chainID uint32) bool {
	return chainID == sepoliaChainID || chainID == opSepoliaChainID || chainID == arbitrumSepoliaChainID
}

// see https://developers.circle.com/stablecoins/docs/evm-smart-contracts#mainnet-contract-addresses
var mainnetDomainMap = map[uint32]uint32{
	0: sepoliaChainID,
	2: opSepoliaChainID,
	3: arbitrumSepoliaChainID,
}

// see https://developers.circle.com/stablecoins/docs/evm-smart-contracts#testnet-contract-addresses
var testnetDomainMap = map[uint32]uint32{
	0: sepoliaChainID,
	2: opSepoliaChainID,
	3: arbitrumSepoliaChainID,
}

// circleDomainToChainID converts a Circle domain to a chain ID.
func circleDomainToChainID(domain uint32, isTestnet bool) (chainID uint32, err error) {
	if isTestnet {
		chainID, ok := testnetDomainMap[domain]
		if !ok {
			return 0, fmt.Errorf("domain %d not found in testnet domain map", domain)
		}
		return chainID, nil
	}

	chainID, ok := mainnetDomainMap[domain]
	if !ok {
		return 0, fmt.Errorf("domain %d not found in mainnet domain map", domain)
	}
	return chainID, nil
}
