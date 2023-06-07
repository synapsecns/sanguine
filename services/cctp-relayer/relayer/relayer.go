package relayer

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/db"
	scribeDb "github.com/synapsecns/sanguine/services/scribe/db"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type usdcMessage struct {
	message       []byte // raw bytes of message produced by Circle's MessageTransmitter
	auxillaryData []byte // auxillary data emitted by SynapseCCTP
	signature     []byte // attestation produced by Circle's API: https://developers.circle.com/stablecoin/reference/getattestation
}

// chainRelayer is a struct that contains the necessary information for each chain level executor.
type chainRelayer struct {
	// chainID is the chain ID of the chain that this executor is responsible for.
	chainID uint32
	// closeConnection is a channel that is used to close the connection.
	closeConnection chan bool
	// stopListenChan is a channel that is used to stop listening to the log channel.
	stopListenChan chan bool
	// usdcMsgRecvChain contains incoming usdc messages yet to be signed.
	usdcMsgRecvChan chan<- *usdcMessage
	// usdcMsgSendChan contains outgoing usdc messages that are signed.
	usdcMsgSendChan <-chan *usdcMessage
}

type CCTPRelayer struct {
	cfg          config.Config
	db           CCTPRelayerDBReader
	scribeClient client.ScribeClient
	grpcClient   pbscribe.ScribeServiceClient
	grpcConn     *grpc.ClientConn
	// chainRelayers is a map from chain ID -> chain executor.
	chainRelayers map[uint32]*chainRelayer
	// handler is the metrics handler.
	handler metrics.Handler
}

const usdcMsgChanSize = 1000

func NewCCTPRelayer(ctx context.Context, cfg config.Config, scribeClient client.ScribeClient, handler metrics.Handler) (*CCTPRelayer, error) {
	chainRelayers := make(map[uint32]*chainRelayer)
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
	)
	if err != nil {
		return nil, fmt.Errorf("could not dial grpc: %w", err)
	}

	grpcClient := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := grpcClient.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return nil, fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return nil, fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	// Build chainRelayers.
	for _, chain := range cfg.Chains {
		chainRelayers[chain.ChainID] = &chainRelayer{
			chainID:         chain.ChainID,
			closeConnection: make(chan bool, 1),
			stopListenChan:  make(chan bool, 1),
			usdcMsgRecvChan: make(chan *usdcMessage, usdcMsgChanSize),
			usdcMsgSendChan: make(chan *usdcMessage, usdcMsgChanSize),
		}
	}

	return &CCTPRelayer{
		cfg:          cfg,
		scribeClient: scribeClient,
		grpcClient:   grpcClient,
		grpcConn:     conn,
		handler:      handler,
	}, nil
}

// Run starts the executor agent. It calls `Start` and `Listen`.
func (c CCTPRelayer) Run(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	// Listen for USDC burn events on origin chains.
	for _, chain := range c.cfg.Chains {
		chain := chain
		g.Go(func() error {
			return c.streamLogs(ctx, c.grpcClient, c.grpcConn, chain.ChainID, chain.OriginAddress, nil)
		})

		g.Go(func() error {
			return c.submitReceiveCircleToken(ctx)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error in executor agent: %w", err)
	}

	return nil
}

// Stop stops the CCTPRelayer.
func (c CCTPRelayer) Stop(chainID uint32) {
	c.chainRelayers[chainID].closeConnection <- true
	c.chainRelayers[chainID].stopListenChan <- true
}

type CCTPRelayerDBReader interface {
	// GetLastBlockNumber gets the last block number that had a message in the database.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	// RetrieveReceiptsWithFilter gets the receipts with the given filter.
	RetrieveReceiptsWithFilter(ctx context.Context, receiptFilter db.ReceiptFilter, page int) ([]types.Receipt, error)
}

// Listens for USDC send events on origin chain, and registers usdcMessages to be signed.
func (c CCTPRelayer) streamLogs(ctx context.Context, grpcClient pbscribe.ScribeServiceClient, conn *grpc.ClientConn, chainID uint32, address string, toBlockNumber *uint64) error {
	lastStoredBlock, err := c.db.GetLastBlockNumber(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not get last stored block: %w", err)
	}

	fromBlock := strconv.FormatUint(lastStoredBlock, 16)

	toBlock := "latest"
	if toBlockNumber != nil {
		toBlock = strconv.FormatUint(*toBlockNumber, 16)
	}

	stream, err := grpcClient.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
		Filter: &pbscribe.LogFilter{
			ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: address}},
			ChainId:         chainID,
		},
		FromBlock: fromBlock,
		ToBlock:   toBlock,
	})
	if err != nil {
		return fmt.Errorf("could not stream logs: %w", err)
	}

	for {
		select {
		case <-c.chainRelayers[chainID].closeConnection:
			err := stream.CloseSend()
			if err != nil {
				return fmt.Errorf("could not close stream: %w", err)
			}

			err = conn.Close()
			if err != nil {
				return fmt.Errorf("could not close connection: %w", err)
			}

			return nil
		default:
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return nil
			}
			if err != nil {
				return fmt.Errorf("could not receive: %w", err)
			}

			msg, err := scribeResponseToMsg(ctx, response, c.db)
			if err != nil {
				return err
			}

			c.chainRelayers[chainID].usdcMsgRecvChan <- msg
		}
	}
}

// Converts a scribe response to a usdcMessage.
func scribeResponseToMsg(ctx context.Context, response *pbscribe.StreamLogsResponse, db CCTPRelayerDBReader) (*usdcMessage, error) {
	receipts, err := db.RetrieveReceiptsWithFilter(ctx, scribeDb.ReceiptFilter{}, 0)
	if err != nil {
		return nil, err
	}

	if len(receipts) != 1 {
		err = fmt.Errorf("expected one receipt; got %d", len(receipts))
		return nil, err
	}

	// TODO(dwasse): parse the logs from the receipt to populate
	// msg.message and msg.auxillaryData
	msg := &usdcMessage{}

	return msg, nil
}

// Completes a USDC bridging sequence by calling ReceiveCircleToken() on the destination chain.
func (c CCTPRelayer) submitReceiveCircleToken(ctx context.Context) error { return nil }

type contractType int

type eventType int

const (
	synapseCCTPContract contractType = iota
)

const (
	// MessageSent event emitted by Circle's contracts.
	messageSent eventType = iota
	// CircleRequestSent event with auxillary data emitted by SynapseCCTP.
	circleRequestSent
)

type contractEventType struct {
	contractType contractType
	eventType    eventType
}
