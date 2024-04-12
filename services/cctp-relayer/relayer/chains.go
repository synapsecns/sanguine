package relayer

import "fmt"

// TODO: this file can be moved to somewhere common

// Mainnet chain IDs.
const ethereumChainID = 1
const avalancheChainID = 43114
const optimismChainID = 10
const arbitrumChainID = 42161
const baseChainID = 8453
const polygonChainID = 137

// Testnet chain IDs.
const sepoliaChainID = 11155111
const avalancheFujiChainID = 43113
const opSepoliaChainID = 11155420
const arbitrumSepoliaChainID = 421614
const baseSepoliaChainID = 84532
const polygonAmoyChainID = 80002

// see https://developers.circle.com/stablecoins/docs/evm-smart-contracts#mainnet-contract-addresses
var mainnetDomainMap = map[uint32]uint32{
	0: ethereumChainID,
	1: avalancheChainID,
	2: optimismChainID,
	3: arbitrumChainID,
	6: baseChainID,
	7: polygonChainID,
}

// see https://developers.circle.com/stablecoins/docs/evm-smart-contracts#testnet-contract-addresses
var testnetDomainMap = map[uint32]uint32{
	0: sepoliaChainID,
	1: avalancheFujiChainID,
	2: opSepoliaChainID,
	3: arbitrumSepoliaChainID,
	6: baseSepoliaChainID,
	7: polygonAmoyChainID,
}

var testnetChainIDMap map[uint32]uint32 // Maps chainID to domain for testnet
var mainnetChainIDMap map[uint32]uint32 // Maps chainID to domain for mainnet

// IsTestnetChainID returns true if the chain ID is a testnet chain ID.
func IsTestnetChainID(chainID uint32) bool {
	return chainID == sepoliaChainID || chainID == opSepoliaChainID || chainID == arbitrumSepoliaChainID
}

// CircleDomainToChainID converts a Circle domain to a chain ID.
func CircleDomainToChainID(domain uint32, isTestnet bool) (chainID uint32, err error) {
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

// ChainIDToCircleDomain converts a chain ID to a Circle domain.
func ChainIDToCircleDomain(chainID uint32, isTestnet bool) (domain uint32, err error) {
	if testnetChainIDMap == nil || mainnetChainIDMap == nil {
		initReverseMaps()
	}

	if isTestnet {
		domain, ok := testnetChainIDMap[chainID]
		if !ok {
			return 0, fmt.Errorf("chain ID %d not found in testnet chain ID map", chainID)
		}
		return domain, nil
	}

	domain, ok := mainnetChainIDMap[chainID]
	if !ok {
		return 0, fmt.Errorf("chain ID %d not found in mainnet chain ID map", chainID)
	}
	return domain, nil
}

func initReverseMaps() {
	testnetChainIDMap = make(map[uint32]uint32)
	mainnetChainIDMap = make(map[uint32]uint32)

	for domain, chainID := range testnetDomainMap {
		testnetChainIDMap[chainID] = domain
	}

	for domain, chainID := range mainnetDomainMap {
		mainnetChainIDMap[chainID] = domain
	}
}
