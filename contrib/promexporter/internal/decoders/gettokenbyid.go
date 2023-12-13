package decoders

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/ethergo/parser/abiutil"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
)

// TokenConfigGetToken returns the token config get token function.
// Must return into bridgeconfig.BridgeConfigV3Token.
func TokenConfigGetToken() w3types.Func {
	return tokenConfigGetToken
}

var tokenConfigGetToken w3types.Func

type customDecodedFunc struct {
	*w3.Func
}

func (c *customDecodedFunc) DecodeReturns(output []byte, returns ...any) (err error) {
	err = returnChecks(c.Func, output, returns...)
	if err != nil {
		return fmt.Errorf("could not decode returns: %w", err)
	}

	raw, err := bridgeConfigV3ABI.Unpack("getTokenByID", output)
	if err != nil {
		return fmt.Errorf("could not unpack returns: %w", err)
	}

	abi.ConvertType(raw[0], returns[0])
	if err != nil {
		return fmt.Errorf("could not decode returns: %w", err)
	}
	return nil
}

var bridgeConfigV3ABI *abi.ABI

func init() {
	// create the token config selector
	underlying := w3.MustNewFunc("getTokenByID(string,uint256)", "uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,bool")

	tokenConfigGetToken = &customDecodedFunc{underlying}

	// make sure it matches expected (we didn't mess up anything). this is a sanity check at boot time
	tokenIDSelector, err := abiutil.GetSelectorByName("getTokenByID", bridgeconfig.BridgeConfigV3MetaData)
	if err != nil {
		panic("could not get token id selector")
	}
	if underlying.Selector != tokenIDSelector {
		panic("incorrect token selector")
	}

	bridgeConfigV3ABI, err = bridgeconfig.BridgeConfigV3MetaData.GetAbi()
	if err != nil {
		panic("could not get bridge config v3 abi")
	}
}
