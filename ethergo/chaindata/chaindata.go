// Package chaindata provides the chain metadata.
// TODO: this more elegantly.
// nolint: mnd
package chaindata

import "strings"

// ChainMetadata represents the chain metadata.
type ChainMetadata struct {
	// ChainId is the chain id.
	ChainId int64
	// ChainName is the chain name.
	ChainName string
	// Explorer is the chain explorer.
	Explorer string
}

var ChainMetadataList = []ChainMetadata{
	{
		ChainId:   1,
		ChainName: "ethereum",
		Explorer:  "https://etherscan.io",
	},
	{
		ChainId:   56,
		ChainName: "bsc",
		Explorer:  "https://bscscan.com",
	},
	{
		ChainId:   137,
		ChainName: "polygon",
		Explorer:  "https://polygonscan.com",
	},
	{
		ChainId:   80001,
		ChainName: "mumbai",
		Explorer:  "https://mumbai.polygonscan.com",
	},
	{
		ChainId:   250,
		ChainName: "fantom",
		Explorer:  "https://ftmscan.com",
	},
	{
		ChainId:   43114,
		ChainName: "avalanche",
		Explorer:  "https://snowtrace.io/",
	},
	{
		ChainId:   42161,
		ChainName: "arbitrum",
		Explorer:  "https://arbiscan.io",
	},
	{
		ChainId:   1337,
		ChainName: "local",
		Explorer:  "http://localhost:1337",
	},
	{
		ChainId:   42220,
		ChainName: "celo",
		Explorer:  "https://explorer.celo.org",
	},
	{
		ChainId:   128,
		ChainName: "heco",
		Explorer:  "https://hecoinfo.com",
	},
	{
		ChainId:   66,
		ChainName: "okexchain",
		Explorer:  "https://www.oklink.com/okexchain",
	},
	{
		ChainId:   100,
		ChainName: "xdai",
		Explorer:  "https://blockscout.com/poa/xdai",
	},
	{
		ChainId:   10,
		ChainName: "optimism",
		Explorer:  "https://optimistic.etherscan.io",
	},
	{
		ChainId:   25,
		ChainName: "cronos",
		Explorer:  "https://cronoscan.com",
	},
	{
		ChainId:   1285,
		ChainName: "moonriver",
		Explorer:  "https://moonriver.moonscan.io",
	},
	{
		ChainId:   1284,
		ChainName: "moonbeam",
		Explorer:  "https://moonbeam.moonscan.io",
	},
	{
		ChainId:   66,
		ChainName: "okc",
		Explorer:  "https://www.oklink.com/en/okc",
	},
	{
		ChainId:   9001,
		ChainName: "evmos",
		Explorer:  "https://evm.evmos.org",
	},
	{
		ChainId:   7700,
		ChainName: "canto",
		Explorer:  "https://tuber.build/",
	},
	{
		ChainId:   53935,
		ChainName: "dfk chain",
		Explorer:  "https://subnets.avax.network/defi-kingdoms",
	},
	{
		ChainId:   8217,
		ChainName: "klaytn",
		Explorer:  "https://scope.klaytn.com",
	},
	{
		ChainId:   288,
		ChainName: "boba",
		Explorer:  "https://bobascan.com",
	},
	{
		ChainId:   1088,
		ChainName: "metis",
		Explorer:  "https://andromeda-explorer.metis.io",
	},
	{
		ChainId:   1313161554,
		ChainName: "aurora",
		Explorer:  "https://explorer.mainnet.aurora.dev",
	},
	{
		ChainId:   1666600000,
		ChainName: "harmony",
		Explorer:  "https://explorer.harmony.one",
	},
	{
		ChainId:   2000,
		ChainName: "dogechain",
		Explorer:  "https://explorer.dogechain.dog",
	},
	{
		ChainId:   8453,
		ChainName: "base",
		Explorer:  "https://basescan.org",
	},
	{
		ChainId:   81457,
		ChainName: "blast",
		Explorer:  "https://blastscan.io",
	},
	{
		ChainId:   534352,
		ChainName: "scroll",
		Explorer:  "https://scrollscan.com",
	},
}

// ChainIDToChainName converts the chain id to the chain name.
func ChainIDToChainName(chainId int64, isUpper bool) string {
	for _, chainMetadata := range ChainMetadataList {
		if chainMetadata.ChainId == chainId {
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
func ChainIDToExplorer(chainId int64) string {
	for _, chainMetadata := range ChainMetadataList {
		if chainMetadata.ChainId == chainId {
			return chainMetadata.Explorer
		}
	}
	return ""
}

// ToExplorerLink converts the chain id and hash to the explorer link.
func ToExplorerLink(chainID int64, hash string) string {
	explorer := ChainIDToExplorer(chainID)
	if explorer == "" {
		return hash
	}
	return explorer + "/tx/" + hash
}
