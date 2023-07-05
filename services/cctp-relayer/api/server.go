package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	db2 "github.com/synapsecns/sanguine/services/cctp-relayer/db"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// RelayerAPIServer exposes API endpoints that interact with CCTPRelayer.
type RelayerAPIServer struct {
	port             uint16
	host             string
	db               db2.CCTPRelayerDB
	relayRequestChan chan *RelayRequest
}

// NewRelayerAPIServer creates a new RelayerAPIServer.
func NewRelayerAPIServer(port uint16, host string, db db2.CCTPRelayerDB, relayRequestChan chan *RelayRequest) *RelayerAPIServer {
	return &RelayerAPIServer{
		port:             port,
		host:             host,
		db:               db,
		relayRequestChan: relayRequestChan,
	}
}

// Start starts the RelayerAPIServer.
func (r RelayerAPIServer) Start(ctx context.Context) error {
	engine := gin.Default()
	engine.GET("/tx", func(ctx *gin.Context) {
		r.GetTx(ctx)
	})
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", r.port),
		ReadHeaderTimeout: 5 * time.Second,
		Handler:           engine,
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil {
			return fmt.Errorf("stopped serving: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		// shutdown server if parent context is canceled
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			return fmt.Errorf("error during server shutdown: %w", err)
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("error while serving: %w", err)
	}

	return nil
}

// RelayRequest is a request to relay a transaction.
type RelayRequest struct {
	Origin uint32
	TxHash common.Hash
}

// GetTx handles the /tx endpoint.
// If the transaction is found in the db, return information about the transaction.
// Otherwise, queue the corresponding Message for relay.
func (r RelayerAPIServer) GetTx(ctx *gin.Context) {
	// parse params
	origin, err := getOriginParam(ctx)
	if err != nil {
		encodeError(ctx, http.StatusBadRequest, err)
		return
	}
	hash, err := getHashParam(ctx)
	if err != nil {
		encodeError(ctx, http.StatusBadRequest, err)
		return
	}

	// fetch corresponding hash from db
	msg, err := r.db.GetMessageByOriginHash(ctx, common.HexToHash(hash))
	if err == nil {
		// return if found
		resp := RelayerResponse{
			Success: true,
			Result: MessageResult{
				OriginHash:      hash,
				DestinationHash: msg.DestTxHash,
				Origin:          msg.OriginChainID,
				Destination:     msg.DestChainID,
				RequestID:       msg.RequestID,
				State:           msg.State.String(),
			},
		}
		ctx.JSON(http.StatusOK, resp)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// enqueue new pending message if not found
		r.relayRequestChan <- &RelayRequest{
			Origin: origin,
			TxHash: common.HexToHash(hash),
		}
		resp := RelayerResponse{
			Success: true,
			Result:  fmt.Sprintf("Successfully queued relay request from chain %d: %s", origin, hash),
		}
		ctx.JSON(http.StatusOK, resp)
		return
	}

	encodeError(ctx, http.StatusInternalServerError, err)
}

// MessageResult is the result of a successful /tx request.
type MessageResult struct {
	OriginHash      string `json:"origin_hash"`
	DestinationHash string `json:"destination_hash"`
	Origin          uint32 `json:"origin"`
	Destination     uint32 `json:"destination"`
	RequestID       string `json:"request_id"`
	State           string `json:"state"`
}

// RelayerResponse is a wrapper struct for a relayer API response.
type RelayerResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

type errorResult struct {
	Reason string `json:"reason"`
}

func encodeError(ctx *gin.Context, status int, err error) {
	resp := RelayerResponse{
		Success: false,
		Result: errorResult{
			Reason: err.Error(),
		},
	}
	ctx.JSON(status, resp)
}
