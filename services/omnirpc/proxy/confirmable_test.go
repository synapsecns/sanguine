package proxy_test

import (
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3/module/eth"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
)

// checkRequest is a helper method for checking requests.
// TODO: checkReq can be replaced w/ a confirmable call, we should do this once we're complete.
func (p *ProxySuite) checkRequest(makeReq func(client client.EVMClient), checkReq func(body []byte)) {
	doneChan := make(chan bool)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		Nil(p.T(), err)

		checkReq(body)

		go func() {
			doneChan <- true
		}()
	}))

	defer server.Close()

	evmClient, err := client.NewClientFromChainID(p.GetTestContext(), server.URL, big.NewInt(1))
	Nil(p.T(), err)

	makeReq(evmClient)

	<-doneChan
}

// Test parsing.
//
//nolint:maintidx
func (p *ProxySuite) TestParseRPCPayload() {
	/*
	  CHECK BLOCKS
	*/

	// check latest block, should not be confirmable since
	// rpcs might be on different  latest heights
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.BlockByNumber(p.GetTestContext(), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// check non-latest block, should be confirmable
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.BlockByNumber(p.GetTestContext(), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	/*
	  CHECK HEADERS
	*/

	// do the same thing, but check headers this time
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.HeaderByNumber(p.GetTestContext(), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// check non-latest block, should be confirmable
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.HeaderByNumber(p.GetTestContext(), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	// eth_blockNumber should not be confirmable, can differ based on rpc
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.BlockNumber(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	/*
	  CHECK Transaction Methods
	*/
	p.checkRequest(func(client client.EVMClient) {
		_, _, _ = client.TransactionByHash(p.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.TransactionCount(p.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.PendingTransactionCount(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.TransactionInBlock(p.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())), 1)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.TransactionReceipt(p.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	/*
	  Sync Methods
	*/

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.SyncProgress(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.NetworkID(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	/*
	  Accessor Methods
	*/

	// latest block, should not be confirmable
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.BalanceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// pending
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.PendingBalanceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// non-latest block
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.BalanceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	// latest block, should not be confirmable
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.StorageAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.PendingStorageAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// non-latest block
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.StorageAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.CodeAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.PendingCodeAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// non-latest block
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.CodeAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	// storage:
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.NonceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.PendingNonceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// non-latest block
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.NonceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	/*
	  Filter Logs
	*/
	// this one's a bit more complicated. It should only return true if fromblock is set and toblock is not set if blockhash is not set
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.FilterLogs(p.GetTestContext(), ethereum.FilterQuery{
			FromBlock: nil,
			ToBlock:   nil,
			Addresses: []common.Address{common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))},
			Topics:    [][]common.Hash{{common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))}},
		})
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// set just to block
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.FilterLogs(p.GetTestContext(), ethereum.FilterQuery{
			FromBlock: nil,
			ToBlock:   big.NewInt(1),
			Addresses: []common.Address{common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))},
			Topics:    [][]common.Hash{{common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))}},
		})
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	// set only block hash
	p.checkRequest(func(client client.EVMClient) {
		bhash := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		_, _ = client.FilterLogs(p.GetTestContext(), ethereum.FilterQuery{
			BlockHash: &bhash,
			Addresses: []common.Address{common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))},
			Topics:    [][]common.Hash{{common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))}},
		})
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	/*
	  Call Contract
	*/
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.CallContract(p.GetTestContext(), ethereum.CallMsg{}, nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.CallContract(p.GetTestContext(), ethereum.CallMsg{}, big.NewInt(1))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.PendingCallContract(p.GetTestContext(), ethereum.CallMsg{})
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	/*
	  Gas Pricing
	*/
	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.SuggestGasPrice(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.SuggestGasTipCap(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client client.EVMClient) {
		_, _ = client.EstimateGas(p.GetTestContext(), ethereum.CallMsg{})
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// try a confirmable batch
	p.checkRequest(func(client client.EVMClient) {
		_ = client.BatchContext(p.GetTestContext(), eth.ChainID().Returns(nil))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	// try a non-confirmable batch
	// try a confirmable batch
	p.checkRequest(func(client client.EVMClient) {
		_ = client.BatchContext(p.GetTestContext(), eth.ChainID().Returns(nil), eth.BlockNumber().Returns(nil))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})
}
