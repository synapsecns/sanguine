package relayer_test

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/mockerc20"
)

// func TestExecuteDuneQuery(t *testing.T) {
// 	resp, err := stiprelayer.ExecuteDuneQuery()
// 	if err != nil {
// 		t.Fatalf("Failed to execute Dune query: %v", err)
// 	}

// 	if resp.StatusCode != 200 {
// 		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("Failed to read response body: %v", err)
// 	}

// 	if len(body) == 0 {
// 		t.Error("Expected non-empty response body, got empty")
// 	}

// 	fmt.Println(string(body))
// }

// func TestGetExecutionResults(t *testing.T) {
// 	resp, err := stiprelayer.ExecuteDuneQuery()
// 	if err != nil {
// 		t.Fatalf("Failed to execute Dune query: %v", err)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("Failed to read response body: %v", err)
// 	}

// 	var result map[string]string
// 	err = json.Unmarshal(body, &result)
// 	if err != nil {
// 		t.Fatalf("Failed to unmarshal response body: %v", err)
// 	}

// 	execution_id, ok := result["execution_id"]
// 	if !ok {
// 		t.Fatal("No execution_id found in the response")
// 	}

// 	time.Sleep(20000 * time.Millisecond)

// 	resp, err = stiprelayer.GetExecutionResults(execution_id)
// 	if err != nil {
// 		t.Fatalf("Failed to get execution results: %v", err)
// 	}

// 	if resp.StatusCode != 200 {
// 		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
// 	}

// 	getResultsBody, err := ioutil.ReadAll(resp.Body)

// 	fmt.Println(string(getResultsBody))

// }

func (c *STIPRelayerSuite) TestStartRelayer() {
	go func() {
		_ = c.stipRelayer.Run(c.GetTestContext())
	}()

	time.Sleep(30000 * time.Millisecond)
	results, err := c.database.GetSTIPTransactionsNotRebated(c.GetTestContext())
	c.Require().NoError(err)

	fmt.Println("LENGTH: " + strconv.Itoa(len(results)))

	time.Sleep(30000 * time.Millisecond)

	arbERC20Instance, err := mockerc20.NewMockERC20(c.arbERC20Address, c.arbitrumSimulatedBackend)
	c.Require().NoError(err)
	balance, err := arbERC20Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x119bde4540d7703c2f12d37aba39a24cc49d74e8"))
	c.Require().NoError(err)
	c.Require().Equal(balance.String(), "1000000000000000000")
	fmt.Println("BALANCE: " + balance.String())
}

// func (c *STIPRelayerSuite) TestQueryAndStore() {
// 	c.stipRelayer.ProcessExecutionResults(c.GetTestContext())
// 	resultsFirst, err := c.database.GetSTIPTransactionsNotRebated(c.GetTestContext())
// 	c.Require().NoError(err)

// 	fmt.Println("LENGTH: " + strconv.Itoa(len(resultsFirst)))

// 	c.stipRelayer.ProcessExecutionResults(c.GetTestContext())

// 	resultsSecond, err := c.database.GetSTIPTransactionsNotRebated(c.GetTestContext())
// 	c.Require().NoError(err)

// 	fmt.Println("LENGTH: " + strconv.Itoa(len(resultsSecond)))
// 	c.Require().Equal(len(resultsFirst), len(resultsSecond))
// }
