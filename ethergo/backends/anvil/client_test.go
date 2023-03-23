package anvil_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"math/big"
	"time"
)

func (a *AnvilSuite) TestClientImpersonateAccount() {
	ogCount, err := a.counter.GetVitalikCount(&bind.CallOpts{Context: a.GetTestContext()})
	Nil(a.T(), err)

	err = a.client.ImpersonateAccount(a.GetTestContext(), vitalik)
	Nil(a.T(), err)

	increment, err := a.counter.VitalikIncrement(&bind.TransactOpts{
		From:   vitalik,
		Value:  big.NewInt(0),
		NoSend: true,
		Signer: anvil.ImpersonatedSigner,
	})
	Nil(a.T(), err)

	err = a.client.SendUnsignedTransaction(a.GetTestContext(), vitalik, increment)
	Nil(a.T(), err)

	vitalikCount, err := a.counter.GetVitalikCount(&bind.CallOpts{Context: a.GetTestContext()})
	Nil(a.T(), err)

	Equal(a.T(), ogCount.Uint64()+10, vitalikCount.Uint64())

	defer func() {
		err = a.client.StopImpersonatingAccount(a.GetTestContext(), vitalik)
		Nil(a.T(), err)
	}()
}

func (a *AnvilSuite) TestGetAutomine() {
	isAutomining, err := a.client.GetAutomine(a.GetTestContext())
	Nil(a.T(), err)
	True(a.T(), isAutomining)
}

func (a *AnvilSuite) TestMine() {
	for i := 0; i < 3; i++ {
		height, err := a.backend.BlockNumber(a.GetTestContext())
		Nil(a.T(), err)

		mineCount := uint64(gofakeit.Number(1, 10))

		err = a.client.Mine(a.GetTestContext(), uint(mineCount))
		Nil(a.T(), err)

		newHeight, err := a.backend.BlockNumber(a.GetTestContext())
		Nil(a.T(), err)

		Equal(a.T(), height+mineCount, newHeight)
	}
}

func (a *AnvilSuite) TestReset() {
	a.T().Skip("investigate why this is failing: method not implemented")
	err := a.client.Mine(a.GetTestContext(), 1)
	Nil(a.T(), err)

	height, err := a.backend.BlockNumber(a.GetTestContext())
	Nil(a.T(), err)

	err = a.client.Reset(a.GetTestContext())
	Nil(a.T(), err)

	newHeight, err := a.backend.BlockNumber(a.GetTestContext())
	Nil(a.T(), err)

	Less(a.T(), newHeight, height)
}

// just make sure there's no error here.
func (a *AnvilSuite) TestSetRPCURL() {
	err := a.client.SetRPCURL(a.GetTestContext(), a.forkAddress)
	Nil(a.T(), err)
}

func (a *AnvilSuite) TestSetBalance() {
	address := mocks.MockAddress()
	balance := gofakeit.Uint64()
	err := a.client.SetBalance(a.GetTestContext(), address, balance)
	Nil(a.T(), err)

	realBal, err := a.backend.BalanceAt(a.GetTestContext(), address, nil)
	Nil(a.T(), err)

	Equal(a.T(), balance, realBal.Uint64())
}

func (a *AnvilSuite) TestSetCode() {
	address := mocks.MockAddress()
	code := []byte("0x606060")
	err := a.client.SetCode(a.GetTestContext(), address, code)
	Nil(a.T(), err)

	realCode, err := a.backend.Client().CodeAt(a.GetTestContext(), address, nil)
	Nil(a.T(), err)
	Equal(a.T(), code, realCode)
}

func (a *AnvilSuite) TestSetNonce() {
	address := mocks.MockAddress()
	nonce := uint64(gofakeit.Number(1, 100))
	err := a.client.SetNonce(a.GetTestContext(), address, nonce)
	Nil(a.T(), err)

	realNonce, err := a.backend.NonceAt(a.GetTestContext(), address, nil)
	Nil(a.T(), err)
	Equal(a.T(), nonce, realNonce)
}

func (a *AnvilSuite) TestSetStorageAt() {
	address := mocks.MockAddress()
	key := common.HexToHash("0x0123456789abcdef")
	value := common.HexToHash("0xabcdef1234567890")
	err := a.client.SetStorageAt(a.GetTestContext(), address, key, value)
	Nil(a.T(), err)

	realValue, err := a.backend.StorageAt(a.GetTestContext(), address, key, nil)
	Nil(a.T(), err)

	Equal(a.T(), value, common.BytesToHash(realValue))
}

func (a *AnvilSuite) TestSetCoinbase() {
	coinbase := mocks.MockAddress()
	err := a.client.SetCoinbase(a.GetTestContext(), coinbase)
	Nil(a.T(), err)
}

func (a *AnvilSuite) TestSetLoggingEnabled() {
	err := a.client.SetLoggingEnabled(a.GetTestContext(), true)
	Nil(a.T(), err)
}

func (a *AnvilSuite) TestDropTransactions() {
	// TODO: test w/ a mempool
	err := a.client.DropTransaction(a.GetTestContext(), common.Hash{})
	Nil(a.T(), err)
}

func (a *AnvilSuite) TestSetMinGasPrice() {
	a.T().Skip("This is disabled when eip-1559 is active, we might want to spin up a separate docker container for this later")

	minGasPrice := big.NewInt(int64(gofakeit.Number(1, 100)))
	err := a.client.SetMinGasPrice(a.GetTestContext(), minGasPrice)
	Nil(a.T(), err)
}

func (a *AnvilSuite) TestSetNextBlockBaseFeePerGas() {
	baseFee := big.NewInt(int64(gofakeit.Number(1, 100)))
	err := a.client.SetNextBlockBaseFeePerGas(a.GetTestContext(), baseFee)
	Nil(a.T(), err)

	err = a.client.Mine(a.GetTestContext(), 1)
	Nil(a.T(), err)

	header, err := a.backend.Client().HeaderByNumber(a.GetTestContext(), nil)
	Nil(a.T(), err)
	Equal(a.T(), header.BaseFee, baseFee)
}

func (a *AnvilSuite) TestDumpState() {
	stateHex, err := a.client.DumpState(a.GetTestContext())
	Nil(a.T(), err)
	NotEqual(a.T(), "", stateHex)
}

func (a *AnvilSuite) TestLoadState() {
	stateHex, err := a.client.DumpState(a.GetTestContext())
	Nil(a.T(), err)

	err = a.client.LoadState(a.GetTestContext(), stateHex)
	Nil(a.T(), err)
}

func (a *AnvilSuite) TestNodeInfo() {
	nodeInfo, err := a.client.NodeInfo(a.GetTestContext())
	Nil(a.T(), err)
	NotNil(a.T(), nodeInfo)
}

func (a *AnvilSuite) TestSetAutomine() {
	err := a.client.SetAutomine(a.GetTestContext(), false)
	Nil(a.T(), err)

	isAutomining, err := a.client.GetAutomine(a.GetTestContext())
	Nil(a.T(), err)
	False(a.T(), isAutomining)

	err = a.client.SetAutomine(a.GetTestContext(), true)
	Nil(a.T(), err)

	isAutomining, err = a.client.GetAutomine(a.GetTestContext())
	Nil(a.T(), err)
	True(a.T(), isAutomining)
}

func (a *AnvilSuite) TestSetIntervalMining() {
	a.T().Skip("This is disabled because it interferes with other tests, we might want to spin up a separate docker container for this later")
	err := a.client.SetIntervalMining(a.GetTestContext(), 5)
	Nil(a.T(), err)
}

func (a *AnvilSuite) TestSnapshotAndRevert() {
	curBlock, err := a.backend.BlockNumber(a.GetTestContext())
	Nil(a.T(), err)

	snapshotID, err := a.client.Snapshot(a.GetTestContext())
	Nil(a.T(), err)

	err = a.client.Mine(a.GetTestContext(), 1)
	Nil(a.T(), err)

	err = a.client.Revert(a.GetTestContext(), snapshotID)
	Nil(a.T(), err)

	newBlock, err := a.backend.BlockNumber(a.GetTestContext())
	Nil(a.T(), err)

	Equal(a.T(), curBlock, newBlock)
}

func (a *AnvilSuite) TestIncreaseTime() {
	// TODO: test w/ contract call
	err := a.client.IncreaseTime(a.GetTestContext(), 60)
	Nil(a.T(), err)
}

func (a *AnvilSuite) TestSetNextBlockTimestamp() {
	timestamp := time.Now().Add(time.Hour * 10000).Unix()
	err := a.client.SetNextBlockTimestamp(a.GetTestContext(), timestamp)
	Nil(a.T(), err)

	err = a.client.Mine(a.GetTestContext(), 1)
	Nil(a.T(), err)

	header, err := a.backend.Client().HeaderByNumber(a.GetTestContext(), nil)
	Nil(a.T(), err)

	Equal(a.T(), header.Time, uint64(timestamp))
}

func (a *AnvilSuite) TestSetRemoveBlockTimestampInterval() {
	err := a.client.SetBlockTimestampInterval(a.GetTestContext(), 10)
	Nil(a.T(), err)

	// todo: can test this by mining some blocks and checking the timestamp

	err = a.client.RemoveBlockTimestampInterval(a.GetTestContext())
	Nil(a.T(), err)
}

func (a *AnvilSuite) TestSetBlockGasLimit() {
	gasLimit := uint64(10000000)
	err := a.client.SetBlockGasLimit(a.GetTestContext(), gasLimit)
	Nil(a.T(), err)

	err = a.client.Mine(a.GetTestContext(), 1)
	Nil(a.T(), err)

	header, err := a.backend.Client().HeaderByNumber(a.GetTestContext(), nil)
	Nil(a.T(), err)

	Equal(a.T(), header.GasLimit, gasLimit)
}

func (a *AnvilSuite) TestEvmMine() {
	height, err := a.backend.BlockNumber(a.GetTestContext())
	Nil(a.T(), err)

	err = a.client.EvmMine(a.GetTestContext())
	Nil(a.T(), err)

	newHeight, err := a.backend.BlockNumber(a.GetTestContext())
	Nil(a.T(), err)

	Equal(a.T(), height+1, newHeight)
}

func (a *AnvilSuite) TestEnableTraces() {
	a.T().Skip("not yet implemented")
	err := a.client.EnableTraces(a.GetTestContext())
	Nil(a.T(), err)
}
