package livefill_test

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/livefill"
)

// TestContractLivefill tests the livefill of the contract livefiller.
func (l LivefillSuite) TestContractLivefill() {
	// Get simulated blockchain, deploy the test contract, and set up test variables.
	simulatedChain := simulated.NewSimulatedBackendWithChainID(l.GetSuiteContext(), l.T(), big.NewInt(142))
	simulatedChain.FundAccount(l.GetTestContext(), l.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := l.manager.GetTestContract(l.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(l.GetTestContext(), nil)

	// Set config.
	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}

	go func() {
		livefiller, err := livefill.NewContractLivefiller(142, contractConfig.Address, l.testDB, simulatedChain)
		Nil(l.T(), err)
		// Activate the livefiller.
		err = livefiller.Livefill(l.GetTestContext())
		Nil(l.T(), err)
	}()

	time.Sleep(1 * time.Second)

	// Emit 3 events from 2 transactions.
	tx, err := testRef.EmitEventA(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(l.T(), err)
	simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)
	tx, err = testRef.EmitEventAandB(transactOpts.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	Nil(l.T(), err)
	simulatedChain.WaitForConfirmation(l.GetTestContext(), tx)

	time.Sleep(1 * time.Second)

	// Ensure that the livefiller has stored the events.
	receipts, err := l.testDB.UnsafeRetrieveAllReceipts(l.GetTestContext(), false, 0)
	Nil(l.T(), err)
	Equal(l.T(), 2, len(receipts))
	logs, err := l.testDB.UnsafeRetrieveAllLogs(l.GetTestContext(), false, 0, common.Address{})
	Nil(l.T(), err)
	Equal(l.T(), 3, len(logs))
}
