package relapi_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

type quoteRequestDB struct {
	reldb.Service
	quoteRequest *reldb.QuoteRequest
}

func (d quoteRequestDB) GetQuoteRequestByOriginTxHash(_ context.Context, _ common.Hash) (*reldb.QuoteRequest, error) {
	return d.quoteRequest, nil
}

func TestGetTxRetryUnknownChainID(t *testing.T) {
	const chainID uint32 = 999999

	db := quoteRequestDB{
		quoteRequest: &reldb.QuoteRequest{
			Transaction: fastbridge.IFastBridgeBridgeTransaction{
				DestChainId: chainID,
			},
		},
	}
	handler := relapi.NewHandler(
		metrics.NewNullHandler(),
		db,
		map[uint32]*chain.Chain{},
		relconfig.Config{},
		nil,
	)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/retry?hash=0x01", nil)

	handler.GetTxRetry(ctx)

	require.Equal(t, http.StatusInternalServerError, recorder.Code)
	require.JSONEq(
		t,
		fmt.Sprintf(`{"error":"No contract found for chain: %d"}`, chainID),
		recorder.Body.String(),
	)
}
