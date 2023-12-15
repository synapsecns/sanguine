package relconfig

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
)

// TODO: validation function.
type Config struct {
	// ChainID: address
	// TODO(aurelius): move under ChainConfig
	// TODO: this can actually be replaced by quotable tokens.
	Tokens map[int][]string `yaml:"tokens"`
	// ChainID: bridge
	Bridges        map[int]ChainConfig `yaml:"bridges"`
	OmnirpcURL     string              `yaml:"omnirpc_url"`
	RfqAPIURL      string              `yaml:"rfq_api_url"`
	DBConfig       string
	QuotableTokens map[string][]string `yaml:"quotable_tokens"`
	// TODO: remove, replace w/ pkey recover
	RelayerAddress  common.Address
	Signer          config.SignerConfig
	SubmitterConfig submitterConfig.Config
}

type ChainConfig struct {
	// Bridge is the bridge confrimation count.
	Bridge string `yaml:"address"`
	// Confirmations is the number of required confirmations
	Confirmations uint64 `yaml:"confirmations"`
}
