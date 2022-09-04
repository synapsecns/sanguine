package proxy_test

import (
	"bytes"
	"errors"
	"github.com/Soft/iter"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/nsf/jsondiff"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy/mocks"
	"github.com/tidwall/pretty"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (p *ProxySuite) TestServeRequestNoChain() {
	prxy := proxy.NewProxy(config.Config{})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	prxy.Forward(c, 1)
	Equal(p.T(), w.Code, http.StatusBadRequest)
}

func (p *ProxySuite) TestCannotReadBody() {
	prxy := proxy.NewProxy(config.Config{})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	mockBody := new(mocks.BodyReader)
	mockBody.On("Read").Return(0, errors.New("could not read body"))

	prxy.Forward(c, 1)
	Equal(p.T(), w.Code, http.StatusBadRequest)
}

func (p *ProxySuite) TestJsonHash() {
	options := jsondiff.DefaultJSONOptions()
	for i := 0; i < 50; i++ {
		randomJSON := p.generateFakeJSON()
		url := gofakeit.URL()

		json1, err := proxy.NewRawResponse(fuzzFormatJSON(randomJSON), url)
		Nil(p.T(), err)

		json2, err := proxy.NewRawResponse(fuzzFormatJSON(randomJSON), url)
		Nil(p.T(), err)

		// make sure we confirm uniqueness in json diff
		// this will save us some time writing nonsense tests
		diff, _ := jsondiff.Compare(json1.Body(), json2.Body(), &options)
		Equal(p.T(), diff, jsondiff.FullMatch)
		Equal(p.T(), json1.Hash(), json2.Hash())
	}
}

// fuzzFormatJSON randomly formats json.
func fuzzFormatJSON(rawBody []byte) []byte {
	formatSetting := gofakeit.Number(0, 2)
	switch formatSetting {
	case 0:
		rawBody = pretty.Pretty(rawBody)
	case 1:
		rawBody = pretty.Ugly(rawBody)
	case 2:
		rawBody = pretty.PrettyOptions(rawBody, &pretty.Options{
			Width: gofakeit.Number(1, 200),
			// random indentation between 1 and 50
			Indent:   strings.Join(iter.ToSlice(iter.Take(iter.Repeat(" "), uint(gofakeit.Number(0, 50)))), ""),
			SortKeys: gofakeit.Bool(),
		})
	}
	return rawBody
}

func (p *ProxySuite) generateFakeJSON() []byte {
	rawBody, err := gofakeit.JSON(&gofakeit.JSONOptions{
		Type: "array",
		Fields: []gofakeit.Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
		},
		RowCount: gofakeit.Number(5, 20),
		Indent:   true,
	})
	Nil(p.T(), err)

	return rawBody
}

func (p *ProxySuite) TestMalformedRequestBody() {
	prxy := proxy.NewProxy(config.Config{})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(p.generateFakeJSON()))

	prxy.Forward(c, 1)
	Equal(p.T(), w.Code, http.StatusBadRequest)
}

// checkRequest is a helper method for checking requests.
// TODO: checkReq can be replaced w/ a confirmable call, we should do this once we're complete.
func (p *ProxySuite) checkRequest(makeReq func(client *ethclient.Client), checkReq func(body []byte)) {
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

	client, err := ethclient.DialContext(p.GetTestContext(), server.URL)
	Nil(p.T(), err)

	makeReq(client)

	<-doneChan
}

// Test parsing.
// nolint: maintidx
func (p *ProxySuite) TestParseRPCPayload() {
	/*
	  CHECK BLOCKS
	*/

	// check latest block, should not be confirmable since
	// rpcs might be on different  latest heights
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.BlockByNumber(p.GetTestContext(), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// check non-latest block, should be confirmable
	p.checkRequest(func(client *ethclient.Client) {
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
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.HeaderByNumber(p.GetTestContext(), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// check non-latest block, should be confirmable
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.HeaderByNumber(p.GetTestContext(), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	// eth_blockNumber should not be confirmable, can differ based on rpc
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.BlockNumber(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	/*
	  CHECK Transaction Methods
	*/
	p.checkRequest(func(client *ethclient.Client) {
		_, _, _ = client.TransactionByHash(p.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.TransactionCount(p.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.PendingTransactionCount(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.TransactionInBlock(p.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())), 1)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.TransactionReceipt(p.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	/*
	  Sync Methods
	*/

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.SyncProgress(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
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
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.BalanceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// pending
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.PendingBalanceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// non-latest block
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.BalanceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	// latest block, should not be confirmable
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.StorageAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.PendingStorageAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// non-latest block
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.StorageAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.CodeAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.PendingCodeAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// non-latest block
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.CodeAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	// storage:
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.NonceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.PendingNonceAt(p.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	// non-latest block
	p.checkRequest(func(client *ethclient.Client) {
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
	p.checkRequest(func(client *ethclient.Client) {
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
	p.checkRequest(func(client *ethclient.Client) {
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
	p.checkRequest(func(client *ethclient.Client) {
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
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.CallContract(p.GetTestContext(), ethereum.CallMsg{}, nil)
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.CallContract(p.GetTestContext(), ethereum.CallMsg{}, big.NewInt(1))
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		True(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.PendingCallContract(p.GetTestContext(), ethereum.CallMsg{})
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	/*
	  Gas Pricing
	*/
	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.SuggestGasPrice(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.SuggestGasTipCap(p.GetTestContext())
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})

	p.checkRequest(func(client *ethclient.Client) {
		_, _ = client.EstimateGas(p.GetTestContext(), ethereum.CallMsg{})
	}, func(body []byte) {
		confirmable, err := proxy.IsConfirmable(body)
		Nil(p.T(), err)
		False(p.T(), confirmable)
	})
}
