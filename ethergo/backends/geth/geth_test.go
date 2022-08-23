package geth_test

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"math/big"
)

func (g *GethSuite) TestGetFullBackend() {
	be := geth.NewEmbeddedBackendForChainID(g.GetTestContext(), g.T(), big.NewInt(1337))
	client, err := ethclient.DialContext(g.GetTestContext(), be.HTTPEndpoint())
	Nil(g.T(), err)
	height, err := client.HeaderByNumber(g.GetTestContext(), nil)
	Nil(g.T(), err)
	Equal(g.T(), height.Number.Uint64(), uint64(0))
}

func (g *GethSuite) TestFaucet() {
	ctx, cancel := context.WithCancel(g.GetTestContext())
	defer cancel()

	be := geth.NewEmbeddedBackend(ctx, g.T())
	be.EnableTenderly()

	targetBalance := big.NewInt(params.Ether)

	fundedAcct := be.GetFundedAccount(g.GetTestContext(), targetBalance)

	acctBalance, err := be.Client().BalanceAt(g.GetTestContext(), fundedAcct.Address, nil)
	Nil(g.T(), err)

	Equal(g.T(), acctBalance, targetBalance)
}
