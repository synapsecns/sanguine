package ganache_test

import (
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/ganache"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"math/big"
	"os"
)

// Tests ganache.
func (g *GanacheSuite) TestGanacheE2E() {
	if os.Getenv("CI") == "" && os.Getenv("GANACHE_KEYS") == "" && os.Getenv("GANACHE_RPC_URL") == "" {
		g.T().Skip()
	}
	ganacheBe := ganache.NewGanacheBackend(g.GetTestContext(), g.T(), params.RinkebyChainConfig, os.Getenv("GANACHE_RPC_URL"), "rinkeby", os.Getenv("GANACHE_KEYS"))
	// WaitForConfirmation makes sure this worked
	newAddress := mocks.MockAddress()

	balance := big.NewInt(params.Ether)
	ganacheBe.FundAccount(g.GetTestContext(), newAddress, *balance)
	zeroAddressBalance, err := ganacheBe.BalanceAt(g.GetTestContext(), newAddress, nil)
	Nil(g.T(), err)

	Equal(g.T(), zeroAddressBalance, balance)
}
