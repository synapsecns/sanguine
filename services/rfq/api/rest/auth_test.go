package rest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestEIP191AuthRejectsFutureTimestamp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	farFuture := strconv.FormatInt(time.Now().Unix()+365*24*60*60, 10)
	data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(farFuture)) + farFuture
	digest := crypto.Keccak256([]byte(data))
	sig, err := crypto.Sign(digest, key)
	require.NoError(t, err)

	auth := farFuture + ":" + hexutil.Encode(sig)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/quotes", nil)
	c.Request.Header.Set(AuthorizationHeader, auth)

	deadline := time.Now().Unix() - 1000
	_, err = EIP191Auth(c, deadline)
	require.Error(t, err)
	require.Contains(t, err.Error(), "future")
	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestEIP191AuthAcceptsFreshTimestamp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	now := strconv.FormatInt(time.Now().Unix(), 10)
	data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(now)) + now
	digest := crypto.Keccak256([]byte(data))
	sig, err := crypto.Sign(digest, key)
	require.NoError(t, err)

	auth := fmt.Sprintf("%s:%s", now, hexutil.Encode(sig))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/quotes", nil)
	c.Request.Header.Set(AuthorizationHeader, auth)

	deadline := time.Now().Unix() - 1000
	addr, err := EIP191Auth(c, deadline)
	require.NoError(t, err)
	require.Equal(t, crypto.PubkeyToAddress(key.PublicKey), addr)
}
