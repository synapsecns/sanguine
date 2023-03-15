package anvil_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"math/big"
)

const defaultMnemonic = "sound practice disease erupt basket pumpkin truck file gorilla behave find exchange napkin boy congress address city net prosper crop chair marine chase seven"

func (a *AnvilSuite) TestHardforkMnemonic() {
	backendRPCAddress := core.GetEnv("ETH_URL", "https://rpc.ankr.com/eth")
	options := anvil.NewAnvilOptionBuilder()
	err := options.SetForkURL(backendRPCAddress)
	Nil(a.T(), err)

	backend := anvil.NewAnvilBackend(a.GetTestContext(), a.T(), options)

	wall, err := wallet.FromSeedPhrase(defaultMnemonic, accounts.DefaultRootDerivationPath)
	NoError(a.T(), err)

	fmt.Println(wall.Address().Hex())

	bal, err := backend.BalanceAt(a.GetTestContext(), wall.Address(), nil)
	NoError(a.T(), err)

	True(a.T(), bal.Cmp(big.NewInt(0)) > 0)
}
