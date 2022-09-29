package debug

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/cast"
	"github.com/tenderly/tenderly-cli/commands/export"
	"github.com/tenderly/tenderly-cli/config"
	"github.com/tenderly/tenderly-cli/ethereum"
	"net/url"
	"sync"
)

// glob hash mux ensures we don't duplicate write (not supported by imported lib).
var globHasMux sync.Mutex

// MakeClient makes a evm processor client
// project slug is optional.
func MakeClient(rpcURL, chainID, projectSlug string, chainConfig *params.ChainConfig) (*ethereum.Client, error) {
	// set the project network
	globHasMux.Lock()
	exports := cast.ToStringMap(config.MaybeGetString(config.Exports))
	rpcAddress, err := url.Parse(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("could not parse rpc address: %w", err)
	}

	exports[chainID] = config.ExportNetwork{
		Name:          chainID,
		ProjectSlug:   projectSlug,
		RpcAddress:    rpcAddress.Host,
		Protocol:      rpcAddress.Scheme,
		ForkedNetwork: "", // TODO ?
		ChainConfig:   chainConfig,
	}
	config.SetProjectConfig(config.Exports, exports)
	globHasMux.Unlock()
	exportNetwork := export.GetNetwork(chainID)
	client, err := ethereum.Dial(exportNetwork.RpcAddress, exportNetwork.Protocol)
	if err != nil {
		return nil, fmt.Errorf("could not connect to rpc server: %w", err)
	}
	return client, nil
}
