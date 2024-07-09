package relapi_test

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/synapsecns/sanguine/core/retry"
	submitterdb "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

func (c *RelayerClientSuite) TestHealth() {
	ok, err := c.Client.Health(c.GetTestContext())
	c.NoError(err)
	c.True(ok)
}

func (c *RelayerClientSuite) TestGetQuoteRequestStatusByTxHash() {
	testReq := c.underlying.getTestQuoteRequest(reldb.Seen)
	err := c.underlying.database.StoreQuoteRequest(c.GetTestContext(), testReq)
	c.Require().NoError(err)

	resp, err := c.Client.GetQuoteRequestStatusByTxHash(c.GetTestContext(), testReq.OriginTxHash.String())
	c.Require().NoError(err)

	c.Equal(resp.Status, testReq.Status.String())
	c.Equal(resp.TxID, hexutil.Encode(testReq.TransactionID[:]))
	c.Equal(resp.DestTxHash, testReq.DestTxHash.String())
	c.Equal(resp.OriginTxHash, testReq.OriginTxHash.String())
}

func (c *RelayerClientSuite) TestGetQuoteRequestStatusByTxID() {
	testReq := c.underlying.getTestQuoteRequest(reldb.Seen)
	err := c.underlying.database.StoreQuoteRequest(c.GetTestContext(), testReq)
	c.Require().NoError(err)

	resp, err := c.Client.GetQuoteRequestStatusByTxID(c.GetTestContext(), hexutil.Encode(testReq.TransactionID[:]))
	c.Require().NoError(err)

	c.Equal(resp.Status, testReq.Status.String())
	c.Equal(resp.TxID, hexutil.Encode(testReq.TransactionID[:]))
	c.Equal(resp.DestTxHash, testReq.DestTxHash.String())
	c.Equal(resp.OriginTxHash, testReq.OriginTxHash.String())
}

func (c *RelayerClientSuite) TestRetryTransaction() {
	testReq := c.underlying.getTestQuoteRequest(reldb.Seen)
	err := c.underlying.database.StoreQuoteRequest(c.GetTestContext(), testReq)
	c.Require().NoError(err)

	resp, err := c.Client.RetryTransaction(c.GetTestContext(), testReq.OriginTxHash.String())
	c.Require().NoError(err)

	c.Equal(resp.TxID, hexutil.Encode(testReq.TransactionID[:]))
}

func (c *RelayerClientSuite) TestEthWithdraw() {
	backend := c.underlying.testBackends[uint64(c.underlying.originChainID)]

	startBalance, err := backend.BalanceAt(c.GetTestContext(), testWithdrawalAddress, nil)
	c.Require().NoError(err)

	withdrawalAmount := big.NewInt(50)

	_, err = c.Client.Withdraw(c.GetTestContext(), &relapi.WithdrawRequest{
		ChainID:      uint32(backend.GetChainID()),
		To:           testWithdrawalAddress,
		Amount:       withdrawalAmount.String(),
		TokenAddress: chain.EthAddress,
	})
	c.Require().NoError(err)

	// Wait for the transaction to be mined
	err = retry.WithBackoff(c.GetTestContext(), func(ctx context.Context) error {
		balance, err := backend.BalanceAt(ctx, testWithdrawalAddress, nil)
		if err != nil {
			return fmt.Errorf("could not fetch balance %w", err)
		}

		expectedBalance := new(big.Int).Add(startBalance, withdrawalAmount)

		if balance.Cmp(expectedBalance) != 0 {
			return fmt.Errorf("balance not updated")
		}

		return nil
	})
	c.Require().NoError(err)
}

func (c *RelayerClientSuite) TestERC20Withdraw() {
	backend := c.underlying.testBackends[uint64(c.underlying.originChainID)]

	_, erc20 := c.underlying.deployManager.GetMockERC20(c.GetTestContext(), backend)

	startBalance, err := erc20.BalanceOf(&bind.CallOpts{Context: c.GetTestContext()}, testWithdrawalAddress)
	c.Require().NoError(err)

	withdrawalAmount := big.NewInt(50)

	_, err = c.Client.Withdraw(c.GetTestContext(), &relapi.WithdrawRequest{
		ChainID:      uint32(backend.GetChainID()),
		To:           testWithdrawalAddress,
		Amount:       withdrawalAmount.String(),
		TokenAddress: erc20.Address(),
	})
	c.Require().NoError(err)

	// Wait for the transaction to be mined
	err = retry.WithBackoff(c.GetTestContext(), func(ctx context.Context) error {
		balance, err := erc20.BalanceOf(&bind.CallOpts{Context: ctx}, testWithdrawalAddress)
		if err != nil {
			return fmt.Errorf("could not get balance %w", err)
		}

		expectedBalance := new(big.Int).Add(startBalance, withdrawalAmount)

		if balance.Cmp(expectedBalance) != 0 {
			return fmt.Errorf("balance not updated")
		}

		return nil
	})
	c.Require().NoError(err)
}

func (c *RelayerClientSuite) TestGetTxHash() {
	backend := c.underlying.testBackends[uint64(c.underlying.originChainID)]

	_, erc20 := c.underlying.deployManager.GetMockERC20(c.GetTestContext(), backend)

	withdrawalAmount := big.NewInt(50)

	withdrawRes, err := c.Client.Withdraw(c.GetTestContext(), &relapi.WithdrawRequest{
		ChainID:      uint32(backend.GetChainID()),
		To:           testWithdrawalAddress,
		Amount:       withdrawalAmount.String(),
		TokenAddress: erc20.Address(),
	})
	c.Require().NoError(err)

	var txHashRes *relapi.TxHashByNonceResponse

	// Wait for the transaction to be mined
	err = retry.WithBackoff(c.GetTestContext(), func(ctx context.Context) error {
		txHashRes, err = c.Client.GetTxHashByNonce(
			c.GetTestContext(),
			&relapi.GetTxByNonceRequest{
				ChainID: uint32(backend.GetChainID()),
				Nonce:   withdrawRes.Nonce,
			},
		)
		if err != nil {
			return fmt.Errorf("could not get hash %w", err)
		}
		return nil
	})

	c.Require().NoError(err)
	c.Require().NotEmpty(txHashRes.Hash)
}

func (c *RelayerClientSuite) TestEthWithdrawCLI() {
	res, err := c.Client.Withdraw(c.GetTestContext(), &relapi.WithdrawRequest{
		ChainID:      c.underlying.originChainID,
		To:           common.HexToAddress(testWithdrawalAddress.String()),
		Amount:       "1000000000000000000",
		TokenAddress: chain.EthAddress,
	})
	c.Require().NoError(err)

	// Wait for the transaction to be mined
	err = retry.WithBackoff(c.GetTestContext(), func(ctx context.Context) error {
		status, err := c.underlying.database.SubmitterDB().
			GetNonceStatus(
				ctx,
				c.underlying.wallet.Address(),
				big.NewInt(int64(c.underlying.originChainID)),
				res.Nonce,
			)
		if err != nil {
			return fmt.Errorf("could not get status %w", err)
		}

		if status != submitterdb.Stored {
			return fmt.Errorf("transaction not mined")
		}

		return nil
	})
	c.Require().NoError(err)
	c.Require().NotNil(res)
}

func (c *RelayerClientSuite) TestERC20WithdrawCLI() {
	backend := c.underlying.testBackends[uint64(c.underlying.originChainID)]

	_, erc20 := c.underlying.deployManager.GetMockERC20(c.GetTestContext(), backend)

	startBalance, err := erc20.BalanceOf(&bind.CallOpts{Context: c.GetTestContext()}, testWithdrawalAddress)
	c.Require().NoError(err)

	withdrawalAmount := big.NewInt(1000000000000000000)

	res, err := c.Client.Withdraw(c.GetTestContext(), &relapi.WithdrawRequest{
		ChainID:      c.underlying.originChainID,
		To:           common.HexToAddress(testWithdrawalAddress.String()),
		Amount:       withdrawalAmount.String(),
		TokenAddress: erc20.Address(),
	})
	c.Require().NoError(err)

	// Wait for the transaction to be mined
	err = retry.WithBackoff(c.GetTestContext(), func(ctx context.Context) error {
		balance, err := erc20.BalanceOf(&bind.CallOpts{Context: ctx}, testWithdrawalAddress)
		if err != nil {
			return fmt.Errorf("could not fetch balance %w", err)
		}

		expectedBalance := new(big.Int).Add(startBalance, withdrawalAmount)

		if balance.Cmp(expectedBalance) != 0 {
			return fmt.Errorf("balance not updated")
		}

		return nil
	})

	c.Require().NoError(err)
	c.Require().NotNil(res)
}
func (c *RelayerClientSuite) TestGetQuoteByTX() {
	testReq := c.underlying.getTestQuoteRequest(reldb.Seen)
	err := c.underlying.database.StoreQuoteRequest(c.GetTestContext(), testReq)
	c.Require().NoError(err)

	resp, err := c.Client.GetQuoteRequestByTXID(c.GetTestContext(), hexutil.Encode(testReq.TransactionID[:]))
	c.Require().NoError(err)

	c.Equal(len(common.Hex2Bytes(resp.QuoteRequestRaw)), len(testReq.RawRequest))
}
