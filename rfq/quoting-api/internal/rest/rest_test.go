package rest_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/bindings"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/config"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/db"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/db/models"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/rest"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/testutil"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func getBridges(t *testing.T, testWallet wallet.Wallet) (bridges map[uint]*bindings.FastBridge) {
	t.Helper()
	bridges = make(map[uint]*bindings.FastBridge)

	testCtx := context.Background()
	clients := make(map[uint]bind.ContractBackend)

	// deploy bridges on anvil dest chain
	chainID := uint32(42161)

	anvilOpts := anvil.NewAnvilOptionBuilder()
	anvilOpts.SetChainID(uint64(chainID))
	anvilOpts.SetBlockTime(1 * time.Second)
	anvilBackend := anvil.NewAnvilBackend(testCtx, t, anvilOpts)

	evmClient := anvilBackend.Backend.Client()
	clients[uint(chainID)] = evmClient

	// deploys bridge contract and adds wallet as relayer
	testContractHandler, _ := testutil.NewTestContractHandlerImpl(testCtx, anvilBackend, testWallet, chainID)
	bridges[uint(chainID)], _ = bindings.NewFastBridge(testContractHandler.FastBridgeAddress(), anvilBackend)
	return
}

func getServer(t *testing.T, testWallet wallet.Wallet) *rest.APIServer {
	t.Helper()
	cfg := &config.Config{}
	d, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	err = d.AutoMigrate(&models.Quote{})
	assert.NoError(t, err, "expected no error on automigrate")

	database := &db.Database{DB: d}
	engine := gin.Default()
	bridges := getBridges(t, testWallet)
	return rest.NewTestRestAPIServer(cfg, database, engine, bridges)
}

// auth = strconv.Itoa(time.Now().Unix()) + ":" + signature
// signature (hex encoded) = keccak(bytes.concat("\x19Ethereum Signed Message:\n", len(strconv.Itoa(time.Now().Unix()), strconv.Itoa(time.Now().Unix()))).
func getAuthHeader(t *testing.T, testWallet wallet.Wallet) (header string) {
	t.Helper()
	now := strconv.Itoa(int(time.Now().Unix()))
	data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(now)) + now
	digest := crypto.Keccak256([]byte(data))

	sig, err := crypto.Sign(digest, testWallet.PrivateKey())
	if err != nil {
		t.Error(err)
		return
	}
	signature := hexutil.Encode(sig)
	header = now + ":" + signature
	return
}

func TestSetup(t *testing.T) {
	testWallet, err := wallet.FromRandom()
	assert.NoError(t, err)
	r := getServer(t, testWallet)
	r.Setup()

	w := httptest.NewRecorder()
	req := newRequest(t, "GET", "/ping", nil)
	r.Engine().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"result\":\"pong\"}", w.Body.String())
}

func newRequest(tb testing.TB, method, url string, body io.Reader) *http.Request {
	tb.Helper()
	// nolint: noctx
	req, err := http.NewRequest(method, url, body)
	assert.Nil(tb, err)
	return req
}

func TestCreateQuote(t *testing.T) {
	ctx := context.Background()
	testWallet, err := wallet.FromRandom()
	assert.NoError(t, err)
	r := getServer(t, testWallet)
	r.Setup()

	w := httptest.NewRecorder()
	req := newRequest(t, "POST", "/quote", bytes.NewBufferString(fmt.Sprintf("{\"relayer\":\"%s\", \"origin_chain_id\":1, \"dest_chain_id\":42161, \"origin_token\":\"0x1\", \"dest_token\":\"0x2\", \"origin_amount\":100, \"dest_amount\":200}", testWallet.Address().Hex())))
	req.Header.Add("Content-Type", gin.MIMEJSON)
	req.Header.Add("Authorization", getAuthHeader(t, testWallet))
	r.Engine().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"result\":1}", w.Body.String()) // id = 1

	originAmount, _ := decimal.NewFromString("100")
	destAmount, _ := decimal.NewFromString("200")

	// Created a dummy quote
	quote := models.Quote{
		Relayer:       testWallet.Address().Hex(),
		OriginChainID: 1,
		DestChainID:   42161,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	q, err := r.DB().GetQuote(ctx, uint(1))
	assert.NoError(t, err)

	quote.ID = q.ID
	quote.Price = q.Price
	quote.CreatedAt = q.CreatedAt
	quote.UpdatedAt = q.UpdatedAt
	quote.DeletedAt = q.DeletedAt

	assert.Equal(t, quote, q)
}

func TestReadQuotes(t *testing.T) {
	ctx := context.Background()
	testWallet, err := wallet.FromRandom()
	assert.NoError(t, err)
	r := getServer(t, testWallet)
	r.Setup()

	originAmount, _ := decimal.NewFromString("100")
	destAmount, _ := decimal.NewFromString("300")

	// add another quote
	q := models.Quote{
		Relayer:       testWallet.Address().Hex(),
		OriginChainID: 1,
		DestChainID:   42161,
		OriginToken:   "0x1",
		DestToken:     "0x3",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	_, err = r.DB().InsertQuote(ctx, &q)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req := newRequest(t, "GET", "/quote", nil)
	r.Engine().ServeHTTP(w, req)

	q1, err := r.DB().GetQuote(ctx, uint(1))
	assert.NoError(t, err)
	now1 := q1.CreatedAt.Format(time.RFC3339Nano)

	q2, err := r.DB().GetQuote(ctx, uint(2))
	assert.NoError(t, err)
	now2 := q2.CreatedAt.Format(time.RFC3339Nano)

	body := fmt.Sprintf("{\"result\":[{\"id\":2,\"relayer\":\"%s\",\"origin_chain_id\":1,\"dest_chain_id\":42161,\"origin_token\":\"0x1\",\"dest_token\":\"0x3\",\"origin_amount\":\"100\",\"dest_amount\":\"300\",\"price\":\"3\",\"created_at\":\"%s\",\"updated_at\":\"%s\",\"deleted_at\":null},{\"id\":1,\"relayer\":\"%s\",\"origin_chain_id\":1,\"dest_chain_id\":42161,\"origin_token\":\"0x1\",\"dest_token\":\"0x2\",\"origin_amount\":\"100\",\"dest_amount\":\"200\",\"price\":\"2\",\"created_at\":\"%s\",\"updated_at\":\"%s\",\"deleted_at\":null}]}", q2.Relayer, now2, now2, q1.Relayer, now1, now1)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, body, w.Body.String())
}

func TestReadQuote(t *testing.T) {
	ctx := context.Background()
	testWallet, err := wallet.FromRandom()
	assert.NoError(t, err)
	r := getServer(t, testWallet)
	r.Setup()

	w := httptest.NewRecorder()
	req := newRequest(t, "GET", "/quote/1", nil)
	r.Engine().ServeHTTP(w, req)

	q, err := r.DB().GetQuote(ctx, uint(1))
	assert.NoError(t, err)

	now := q.CreatedAt.Format(time.RFC3339Nano)
	body := fmt.Sprintf("{\"result\":{\"id\":1,\"relayer\":\"%s\",\"origin_chain_id\":1,\"dest_chain_id\":42161,\"origin_token\":\"0x1\",\"dest_token\":\"0x2\",\"origin_amount\":\"100\",\"dest_amount\":\"200\",\"price\":\"2\",\"created_at\":\"%s\",\"updated_at\":\"%s\",\"deleted_at\":null}}", q.Relayer, now, now)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, body, w.Body.String())
}

func TestUpdateQuote(t *testing.T) {
	ctx := context.Background()
	testWallet, err := wallet.FromRandom()
	assert.NoError(t, err)
	r := getServer(t, testWallet)
	r.Setup()

	originAmount, _ := decimal.NewFromString("100")
	destAmount, _ := decimal.NewFromString("400")

	// add another quote
	q := models.Quote{
		Relayer:       testWallet.Address().Hex(),
		OriginChainID: 1,
		DestChainID:   42161,
		OriginToken:   "0x1",
		DestToken:     "0x3",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	id, err := r.DB().InsertQuote(ctx, &q)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req := newRequest(t, "PUT", fmt.Sprintf("/quote/%d", id), bytes.NewBufferString(fmt.Sprintf("{\"relayer\":\"%s\", \"origin_chain_id\":1, \"dest_chain_id\":42161, \"origin_token\":\"0x1\", \"dest_token\":\"0x2\", \"origin_amount\":100, \"dest_amount\":150}", testWallet.Address().Hex())))
	req.Header.Add("Content-Type", gin.MIMEJSON)
	req.Header.Add("Authorization", getAuthHeader(t, testWallet))
	r.Engine().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, fmt.Sprintf("{\"result\":%d}", id), w.Body.String())

	originAmount, _ = decimal.NewFromString("100")
	destAmount, _ = decimal.NewFromString("150")
	price, _ := decimal.NewFromString("1.5")

	// Created a dummy quote
	quote := models.Quote{
		ID:            id,
		Relayer:       testWallet.Address().Hex(),
		OriginChainID: 1,
		DestChainID:   42161,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
		Price:         price,
	}
	q, err = r.DB().GetQuote(ctx, id)
	quote.CreatedAt = q.CreatedAt
	quote.UpdatedAt = q.UpdatedAt
	quote.DeletedAt = q.DeletedAt
	assert.NoError(t, err)
	assert.Equal(t, quote, q)
}

func TestDeleteQuote(t *testing.T) {
	ctx := context.Background()
	testWallet, err := wallet.FromRandom()
	assert.NoError(t, err)
	r := getServer(t, testWallet)
	r.Setup()

	originAmount, _ := decimal.NewFromString("100")
	destAmount, _ := decimal.NewFromString("400")

	// add another quote
	q := models.Quote{
		Relayer:       testWallet.Address().Hex(),
		OriginChainID: 1,
		DestChainID:   42161,
		OriginToken:   "0x1",
		DestToken:     "0x3",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	id, err := r.DB().InsertQuote(ctx, &q)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req := newRequest(t, "DELETE", fmt.Sprintf("/quote/%d", id), nil)
	req.Header.Add("Authorization", getAuthHeader(t, testWallet))
	r.Engine().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, fmt.Sprintf("{\"result\":%d}", id), w.Body.String())
}

func TestPingQuote(t *testing.T) {
	ctx := context.Background()
	testWallet, err := wallet.FromRandom()
	assert.NoError(t, err)
	r := getServer(t, testWallet)
	r.Setup()

	originAmount, _ := decimal.NewFromString("100")
	destAmount, _ := decimal.NewFromString("500")

	// add another quote
	q := models.Quote{
		Relayer:       testWallet.Address().Hex(),
		OriginChainID: 1,
		DestChainID:   42161,
		OriginToken:   "0x1",
		DestToken:     "0x3",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}
	id, err := r.DB().InsertQuote(ctx, &q)
	assert.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	// get timestamp before
	quote, err := r.DB().GetQuote(ctx, id)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req := newRequest(t, "POST", fmt.Sprintf("/quote/%d/ping", id), nil)
	req.Header.Add("Authorization", getAuthHeader(t, testWallet))
	r.Engine().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, fmt.Sprintf("{\"result\":%d}", id), w.Body.String())

	// check timestamp updated
	q, err = r.DB().GetQuote(ctx, id)
	assert.NoError(t, err)

	assert.Greater(t, q.UpdatedAt, quote.UpdatedAt)

	// reset cached updated at and check all other fields same
	quote.UpdatedAt = q.UpdatedAt
	assert.Equal(t, quote, q)
}
