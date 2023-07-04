package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	db2 "github.com/synapsecns/sanguine/services/cctp-relayer/db"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type RelayerAPIServer struct {
	port             uint16
	host             string
	db               db2.CCTPRelayerDB
	relayRequestChan chan *RelayRequest
}

func NewRelayerAPIServer(port uint16, host string, db db2.CCTPRelayerDB, relayRequestChan chan *RelayRequest) *RelayerAPIServer {
	return &RelayerAPIServer{
		port:             port,
		host:             host,
		db:               db,
		relayRequestChan: relayRequestChan,
	}
}

func (r RelayerAPIServer) Start(ctx context.Context) error {
	engine := gin.Default()
	engine.GET("/push_tx", func(ctx *gin.Context) {
		r.GetPushTx(ctx)
	})
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", r.port),
		Handler: engine,
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil {
			return fmt.Errorf("stopped serving: %w", err)
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("error while serving: %w", err)
	}

	return nil
}

type RelayRequest struct {
	Origin uint32
	TxHash common.Hash
}

func (r RelayerAPIServer) GetPushTx(ctx *gin.Context) {
	var err error

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
			Origin: uint32(origin),
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

type MessageResult struct {
	OriginHash      string `json:"origin_hash"`
	DestinationHash string `json:"destination_hash"`
	Origin          uint32 `json:"origin"`
	Destination     uint32 `json:"destination"`
	RequestID       string `json:"request_id"`
	State           string `json:"state"`
}

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
