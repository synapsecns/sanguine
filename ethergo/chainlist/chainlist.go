package chainlist

import (
	// embeds the chainsJSON.
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"reflect"
	"strings"
	"sync"
)

//go:embed chains.json
var chainsJSON string

// ChainIDs contains info about each chain.
var ChainIDs ChainList

// init creates the chainIDs object.
// nolint: cyclop
func init() {
	var chains []*Chain
	decoder := json.NewDecoder(strings.NewReader(chainsJSON))
	err := decoder.Decode(&chains)

	// panic on start if the chain json could not be parsed
	if err != nil {
		panic(err)
	}

	// override ropsten explorer url
	for _, chain := range chains {
		switch chain.ChainID {
		case 1:
			chain.APIs = append(chain.APIs, API{
				Name:     "etherscan",
				URL:      "https://api.etherscan.io/api?",
				Standard: EtherscanAPIStandard,
			})
		case 3:
			chain.Explorers = append(chain.Explorers, Explorer{
				Name:     "etherscan",
				URL:      "https://ropsten.etherscan.io",
				Standard: EIP3091Standard,
			})
			chain.APIs = append(chain.APIs, API{
				Name:     "etherscan (ropsten)",
				URL:      "https://api-ropsten.etherscan.io/api?",
				Standard: EtherscanAPIStandard,
			})
		case 4:
			chain.Explorers = append(chain.Explorers, Explorer{
				Name:     "etherscan",
				URL:      "https://rinkeby.etherscan.io",
				Standard: EIP3091Standard,
			})
			chain.APIs = append(chain.APIs, API{
				Name:     "rinkeby",
				URL:      "https://api-rinkeby.etherscan.io/api?",
				Standard: EtherscanAPIStandard,
			})
		case 10:
			chain.APIs = append(chain.APIs, API{
				Name:     "Optimistic Ethereum",
				URL:      "https://api-optimistic.etherscan.io",
				Standard: EtherscanAPIStandard,
			})
		case 25:
			chain.APIs = append(chain.APIs, API{
				Name:     "Cronos Scan",
				URL:      "https://api.cronoscan.com/",
				Standard: EtherscanAPIStandard,
			})
		case 56:
			chain.APIs = append(chain.APIs, API{
				Name:     "bscscan",
				URL:      "https://api.bscscan.com/api?",
				Standard: EtherscanAPIStandard,
			})
		case 97:
			chain.APIs = append(chain.APIs, API{
				Name:     "bscscan (testnet)",
				URL:      "https://api-testnet.bscscan.com/api?",
				Standard: EtherscanAPIStandard,
			})
		case 128:
			chain.APIs = append(chain.APIs, API{
				Name:     "hecoscan",
				URL:      "https://api.hecoinfo.com/api?",
				Standard: EtherscanAPIStandard,
			})
		case 137:
			chain.Explorers = append(chain.Explorers, Explorer{
				Name:     "polygonscan",
				URL:      "https://polygonscan.com",
				Standard: EIP3091Standard,
			})
			chain.APIs = append(chain.APIs, API{
				Name:     "polygonscan",
				URL:      "https://api.polygonscan.com/api?",
				Standard: EtherscanAPIStandard,
			})
		case 250:
			chain.Explorers = append(chain.Explorers, Explorer{
				Name:     "fantom",
				URL:      "https://ftmscan.com",
				Standard: EIP3091Standard,
			})
			chain.APIs = append(chain.APIs, API{
				Name:     "ftmscan",
				URL:      "https://api.ftmscan.com/api?",
				Standard: EtherscanAPIStandard,
			})
		case 288:
			chain.Explorers = append(chain.Explorers, Explorer{
				Name:     "Boba",
				URL:      "https://blockexplorer.boba.network/",
				Standard: EIP3091Standard,
			})
			chain.APIs = append(chain.APIs, API{
				Name:     "Boba Blockexplorer",
				URL:      "https://blockexplorer.boba.network/api/",
				Standard: BlockscoutAPIStandard,
			})
		case 1284:
			chain.Explorers = append(chain.Explorers, Explorer{
				Name:     "Moonbeam Explorer",
				URL:      "https://moonbeam.moonscan.io/",
				Standard: EtherscanAPIStandard,
			})
		case 1285:
			chain.Explorers = append(chain.Explorers, Explorer{
				Name:     "Moonriver Explorer",
				URL:      "https://blockscout.moonriver.moonbeam.network/",
				Standard: BlockscoutAPIStandard,
			})
		case 42161:
			chain.APIs = append(chain.APIs, API{
				Name:     "ArbiScan",
				URL:      "https://api.arbiscan.io/api?",
				Standard: EtherscanAPIStandard,
			})
		case 43114:
			chain.Explorers = append(chain.Explorers, Explorer{
				Name:     "Avalanche",
				URL:      "https://snowtrace.io/",
				Standard: EIP3091Standard,
			})
			chain.APIs = append(chain.APIs, API{
				Name:     "Snow Trace",
				URL:      "https://snowtrace.io/api?",
				Standard: EtherscanAPIStandard,
			})
		}
	}

	ChainIDs = ChainList{chainList: chains}
}

// Chain is a chain from the chains.json file.
type Chain struct {
	// Name of the chain
	Name string `json:"name"`
	// ChainID of the chain (e.g. 1)
	ChainID int `json:"chainId"`
	// ShortName of the chain (e.g. Ethereum->eth)
	ShortName string `json:"shortName"`
	// Chain is the chain name (e.g. ethereum)
	Chain string `json:"chain"`
	// Network is the network
	Network string `json:"network"`
	// NetworkID- https://medium.com/@pedrouid/chainid-vs-networkid-how-do-they-differ-on-ethereum-eec2ed41635b
	NetworkID int `json:"networkId"`
	// NativeCurrency of the chain
	NativeCurrency struct {
		// Name of the native currency (e.g. Ether)
		Name string `json:"name"`
		// Symbol of the chain (e.g. ETH)
		Symbol string `json:"symbol"`
		// Decimals ont the chain (e.g. 18)
		Decimals int `json:"decimals"`
	} `json:"nativeCurrency"`
	// RPC is the rpc url
	RPC []string `json:"rpc"`
	// Faucets available for the chain
	Faucets []interface{} `json:"faucets"`
	// Explorers available for the chain
	Explorers []Explorer `json:"explorers"`
	// APIs are the apis aviailable for each chain. These are not (yet) available from chainids
	APIs []API `json:"API"`
	// InfoURL: website for the protocol
	InfoURL string `json:"infoURL"`
}

// Explorer is the explorer item.
type Explorer struct {
	// Name of the explorer
	Name string `json:"name"`
	// URL of the explorer
	URL string `json:"url"`
	// Standard of the explorer (right now the only eip for this is: https://eips.ethereum.org/EIPS/eip-3091)
	Standard string `json:"standard"`
}

// API is an analytics for retreiving data.
type API struct {
	// Name is the name of the analytics: e.g. etherscan
	Name string `json:"name"`
	// URL of the explorer
	URL string `json:"url"`
	// Standard of the explorer (right now the only eip for this is: https://eips.ethereum.org/EIPS/eip-3091)
	Standard string `json:"standard"`
}

// ChainList is a list of evm chains by id.
type ChainList struct {
	chainList []*Chain
	mux       sync.RWMutex
}

// ChainByID gets a chain by id from the chain list
// if no chain is found, nil is returned.
func (c *ChainList) ChainByID(chainID uint) *Chain {
	c.mux.RLock()
	defer c.mux.RUnlock()

	for _, chain := range c.chainList {
		if chain.ChainID == int(chainID) {
			return chain
		}
	}

	logger.Warnf("could not get chainid config from %d", chainID)

	return nil
}

// PutChain allows you to add or overwrite an existing chain. This is useful
// for chains not covered by https://chainid.network/
func (c *ChainList) PutChain(newChain *Chain) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for i, chain := range c.chainList {
		if chain.ChainID == newChain.ChainID {
			// delete the old item
			c.chainList = append(c.chainList[:i], c.chainList[i+1:]...)
		}
	}
	c.chainList = append(c.chainList, newChain)
}

// EIP3091Standard is the standard supported by explorer url
// see: https://eips.ethereum.org/EIPS/eip-3091 for details
const EIP3091Standard = "EIP3091"

// EtherscanAPIStandard is an analytics that complies to the etherscan analytics https://etherscan.io/apidocs
const EtherscanAPIStandard = "etherscan"

// BlockscoutAPIStandard is the blockscout analytics standard.
const BlockscoutAPIStandard = "blockscout"

// explorerBaseURL gets the explorer base url.
func (c Chain) explorerBaseURL() string {
	var explorerBaseURL string
	// use the last explorer
	for _, explorer := range c.Explorers {
		if explorer.Standard == EIP3091Standard || explorer.Standard == EtherscanAPIStandard || explorer.Standard == BlockscoutAPIStandard {
			explorerBaseURL = explorer.URL
			break
		}
	}
	return explorerBaseURL
}

// ExplorerURL converts an ethereum type to an explorer url
// returns empty string if explorer url is not available.
// nolint: cyclop
func (c Chain) ExplorerURL(data interface{}) string {
	explorerBaseURL := c.explorerBaseURL()

	if explorerBaseURL == "" {
		logger.Warnf("no explorer url found for chain %s (chain id: %d)", c.Chain, c.ChainID)
		return ""
	}

	//nolint: gocritic
	switch data := data.(type) {
	case types.Block:
		return fmt.Sprintf("%s/block/%s", explorerBaseURL, data.Hash().String())
	case int, uint, uint64, int64:
		return fmt.Sprintf("%s/block/%d", explorerBaseURL, data)
	case *types.Block:
		return fmt.Sprintf("%s/block/%s", explorerBaseURL, data.Hash().String())
	case types.Transaction:
		return fmt.Sprintf("%s/tx/%s", c.explorerBaseURL(), data.Hash().String())
	case *types.Transaction:
		return fmt.Sprintf("%s/tx/%s", c.explorerBaseURL(), data.Hash().String())
	case common.Address:
		return fmt.Sprintf("%s/address/%s", explorerBaseURL, data.String())
	case *common.Address:
		return fmt.Sprintf("%s/address/%s", explorerBaseURL, data.String())
	default:
		logger.Warnf("could not convert %s (chainID: %d)", reflect.TypeOf(c.Chain), c.ChainID)
	}

	return ""
}

// TxHashURL gets the transaction hash explorer url for the chain.
func (c Chain) TxHashURL(tx common.Hash) string {
	return c.TxStringURL(tx.String())
}

// TxStringURL gets the tx url from a string.
func (c Chain) TxStringURL(tx string) string {
	return fmt.Sprintf("%s/tx/%s", c.explorerBaseURL(), tx)
}
