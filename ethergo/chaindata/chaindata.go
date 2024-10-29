// Package chaindata provides the chain metadata.
// TODO: this more elegantly.
// nolint: mnd
package chaindata

import "strings"

// ChainMetadata represents the chain metadata.
type ChainMetadata struct {
	// ChainID is the chain id.
	ChainID int64
	// ChainName is the chain name.
	ChainName string
	// Explorer is the chain explorer.
	Explorer string
}

// ChainMetadataList is the list of chain metadata.
var ChainMetadataList = []ChainMetadata{
	{
		ChainID:   1,
		ChainName: "ethereum",
		Explorer:  "https://etherscan.io",
	},
	{
		ChainID:   56,
		ChainName: "bsc",
		Explorer:  "https://bscscan.com",
	},
	{
		ChainID:   137,
		ChainName: "polygon",
		Explorer:  "https://polygonscan.com",
	},
	{
		ChainID:   80001,
		ChainName: "mumbai",
		Explorer:  "https://mumbai.polygonscan.com",
	},
	{
		ChainID:   250,
		ChainName: "fantom",
		Explorer:  "https://ftmscan.com",
	},
	{
		ChainID:   43114,
		ChainName: "avalanche",
		Explorer:  "https://snowtrace.io/",
	},
	{
		ChainID:   42161,
		ChainName: "arbitrum",
		Explorer:  "https://arbiscan.io",
	},
	{
		ChainID:   1337,
		ChainName: "local",
		Explorer:  "http://localhost:1337",
	},
	{
		ChainID:   42220,
		ChainName: "celo",
		Explorer:  "https://explorer.celo.org",
	},
	{
		ChainID:   128,
		ChainName: "heco",
		Explorer:  "https://hecoinfo.com",
	},
	{
		ChainID:   66,
		ChainName: "okexchain",
		Explorer:  "https://www.oklink.com/okexchain",
	},
	{
		ChainID:   100,
		ChainName: "xdai",
		Explorer:  "https://blockscout.com/poa/xdai",
	},
	{
		ChainID:   10,
		ChainName: "optimism",
		Explorer:  "https://optimistic.etherscan.io",
	},
	{
		ChainID:   25,
		ChainName: "cronos",
		Explorer:  "https://cronoscan.com",
	},
	{
		ChainID:   1285,
		ChainName: "moonriver",
		Explorer:  "https://moonriver.moonscan.io",
	},
	{
		ChainID:   1284,
		ChainName: "moonbeam",
		Explorer:  "https://moonbeam.moonscan.io",
	},
	{
		ChainID:   66,
		ChainName: "okc",
		Explorer:  "https://www.oklink.com/en/okc",
	},
	{
		ChainID:   9001,
		ChainName: "evmos",
		Explorer:  "https://evm.evmos.org",
	},
	{
		ChainID:   7700,
		ChainName: "canto",
		Explorer:  "https://tuber.build/",
	},
	{
		ChainID:   53935,
		ChainName: "dfk chain",
		Explorer:  "https://subnets.avax.network/defi-kingdoms",
	},
	{
		ChainID:   8217,
		ChainName: "klaytn",
		Explorer:  "https://scope.klaytn.com",
	},
	{
		ChainID:   288,
		ChainName: "boba",
		Explorer:  "https://bobascan.com",
	},
	{
		ChainID:   1088,
		ChainName: "metis",
		Explorer:  "https://andromeda-explorer.metis.io",
	},
	{
		ChainID:   1313161554,
		ChainName: "aurora",
		Explorer:  "https://explorer.mainnet.aurora.dev",
	},
	{
		ChainID:   1666600000,
		ChainName: "harmony",
		Explorer:  "https://explorer.harmony.one",
	},
	{
		ChainID:   2000,
		ChainName: "dogechain",
		Explorer:  "https://explorer.dogechain.dog",
	},
	{
		ChainID:   8453,
		ChainName: "base",
		Explorer:  "https://basescan.org",
	},
	{
		ChainID:   81457,
		ChainName: "blast",
		Explorer:  "https://blastscan.io",
	},
	{
		ChainID:   534352,
		ChainName: "scroll",
		Explorer:  "https://scrollscan.com",
	},
	{
		ChainID:   59144,
		ChainName: "linea",
		Explorer:  "https://lineascan.build/",
	},
}

// ChainNameToChainID converts the chain name to the chain id.
// It returns 0 if the chain name is not found.
func ChainNameToChainID(chainName string) uint64 {
	for _, chainMetadata := range ChainMetadataList {
		if strings.EqualFold(chainMetadata.ChainName, chainName) {
			return uint64(chainMetadata.ChainID)
		}
	}
	return 0
}

// ChainIDToChainName converts the chain id to the chain name.
func ChainIDToChainName(chainID int64, isUpper bool) string {
	for _, chainMetadata := range ChainMetadataList {
		if chainMetadata.ChainID == chainID {
			// upper the first letter
			if isUpper {
				return strings.ToUpper(chainMetadata.ChainName)
			}
			return chainMetadata.ChainName
		}
	}
	return ""
}

// ChainIDToExplorer converts the chain id to the chain explorer.
func ChainIDToExplorer(chainID int64) string {
	for _, chainMetadata := range ChainMetadataList {
		if chainMetadata.ChainID == chainID {
			return chainMetadata.Explorer
		}
	}
	return ""
}

// ToTXLink converts the chain id and hash to the explorer link.
func ToTXLink(chainID int64, hash string) string {
	explorer := ChainIDToExplorer(chainID)
	if explorer == "" {
		return ""
	}
	return explorer + "/tx/" + hash
}
