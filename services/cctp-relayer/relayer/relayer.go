package relayer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmessagetransmitter"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/db"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type usdcMessage struct {
	message       []byte      // raw bytes of message produced by Circle's MessageTransmitter
	auxillaryData []byte      // auxillary data emitted by SynapseCCTP
	txHash        common.Hash // hash of the USDC burn transaction
	signature     []byte      // attestation produced by Circle's API: https://developers.circle.com/stablecoin/reference/getattestation
}

// chainRelayer is a struct that contains the necessary information for each chain level relayer.
type chainRelayer struct {
	// chainID is the chain ID of the chain that this relayer is responsible for.
	chainID uint32
	// closeConnection is a channel that is used to close the connection.
	closeConnection chan bool
	// stopListenChan is a channel that is used to stop listening to the log channel.
	stopListenChan chan bool
	// usdcMsgRecvChan contains incoming usdc messages yet to be signed.
	usdcMsgRecvChan chan *usdcMessage
	// usdcMsgSendChan contains outgoing usdc messages that are signed.
	usdcMsgSendChan chan *usdcMessage
}

type CCTPRelayer struct {
	cfg           config.Config
	db            CCTPRelayerDBReader
	scribeClient  client.ScribeClient
	grpcClient    pbscribe.ScribeServiceClient
	grpcConn      *grpc.ClientConn
	client        *http.Client
	omnirpcClient omniClient.RPCClient
	// chainRelayers is a map from chain ID -> chain relayer.
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
		cfg:           cfg,
		chainRelayers: chainRelayers,
		scribeClient:  scribeClient,
		grpcClient:    grpcClient,
		grpcConn:      conn,
		client:        &http.Client{},
		handler:       handler,
	}, nil
}

// Run starts the CCTPRelayer.
func (c CCTPRelayer) Run(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	// Listen for USDC burn events on origin chains.
	for _, chain := range c.cfg.Chains {
		chain := chain
		g.Go(func() error {
			return c.streamLogs(ctx, c.grpcClient, c.grpcConn, chain.ChainID, chain.OriginAddress, nil)
		})

		g.Go(func() error {
			return c.submitReceiveCircleToken(ctx, chain.ChainID)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error in cctp relayer: %w", err)
	}

	return nil
}

// Stop stops the CCTPRelayer.
func (c CCTPRelayer) Stop(chainID uint32) {
	c.chainRelayers[chainID].closeConnection <- true
	c.chainRelayers[chainID].stopListenChan <- true
}

// TODO(dwasse): impl db interactions
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
			if err != nil {
				return fmt.Errorf("could not receive: %w", err)
			}

			err = c.handleLog(ctx, response.Log.ToLog(), chainID)
			if err != nil {
				return err
			}
		}
	}
}

// Converts a scribe response to a usdcMessage.
// This takes ina  log from the SynapseCCTP contract, determines the topic and then performs an action based on that topic
// this could be a send or receive
func (c CCTPRelayer) handleLog(ctx context.Context, log *types.Log, originChain uint32) error {
	if log == nil {
		return fmt.Errorf("log is nil")
	}

	// shouldn't be possible: maybe remove?
	if len(log.Topics) == 0 {
		return fmt.Errorf("not enough topics")
	}

	switch log.Topics[0] {
	// since this is the last stopic that comes out of the message, we use it to kick off the send loop
	case cctp.CircleRequestSentTopic:
		// TODO: figure out if we want to use scribe here, for now, we'll keep it simple w/ omnirpc gettxreceipt

	case cctp.CircleRequestFulfilledTopic:
		// TODO mark request as fulfilled
	default:
		// TODO; just continue
		logger.Warnf("unknown topic %s", log.Topics[0])
		return nil
	}
	return nil
}

func (c CCTPRelayer) handleSendRequest(parentCtx context.Context, txhash common.Hash, originChain uint32) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "handleSendRequest", trace.WithAttributes(
		attribute.String(metrics.TxHash, txhash.String()),
		attribute.Int(metrics.ChainID, int(originChain)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	ethClient, err := c.omnirpcClient.GetChainClient(ctx, int(originChain))
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}

	// TODO: consider pulling from scribe
	receipt, err := ethClient.TransactionReceipt(ctx, txhash)
	if err != nil {
		return fmt.Errorf("could not get transaction receipt: %w", err)
	}

	// from this receipt, we expect two different logs. One is message sent
	// message sent tells us: TODO fill me in
	// circleRequestSentEvent tells us: TODO fill me in?
	var messageSentEvent *mockmessagetransmitter.MessageTransmitterEventsMessageSent
	var circleRequestSentEvent *cctp.SynapseCCTPEventsCircleRequestSent

	for _, log := range receipt.Logs {
		// this should never happen
		if len(log.Topics) == 0 {
			continue
		}

		switch log.Topics[0] {
		case cctp.CircleRequestSentTopic:
			// TODO: do we need to make sure log.Address matches our log.Address?
			eventParser, err := cctp.NewSynapseCCTPEvents(log.Address, ethClient)
			if err != nil {
				return fmt.Errorf("could not create event parser: %w", err)
			}

			circleRequestSentEvent, err = eventParser.ParseCircleRequestSent(*log)
			if err != nil {
				return fmt.Errorf("could not parse circle request sent: %w", err)
			}
			// TODO: this shouldn't be coming from a mock contract, generate from the abstract contract itself
		case mockmessagetransmitter.MessageSentTopic:
			eventParser, err := mockmessagetransmitter.NewMessageTransmitterEvents(log.Address, ethClient)
			if err != nil {
				return fmt.Errorf("could not create event parser: %w", err)
			}

			messageSentEvent, err = eventParser.ParseMessageSent(*log)
			if err != nil {
				return fmt.Errorf("could not parse message sent: %w", err)
			}

		}
	}

	if messageSentEvent == nil {
		return fmt.Errorf("no message sent event found")
	}

	if circleRequestSentEvent == nil {
		return fmt.Errorf("no circle request sent event found")
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case c.chainRelayers[originChain].usdcMsgSendChan <- &usdcMessage{
		txHash:        txhash,
		auxillaryData: circleRequestSentEvent.Request,
		message:       messageSentEvent.Message,
		//signature: //comes from the api
	}:
	}
	return nil
}

// Completes a USDC bridging sequence by calling ReceiveCircleToken() on the destination chain.
func (c CCTPRelayer) submitReceiveCircleToken(ctx context.Context, chainID uint32) (err error) {
	for {
		select {
		// Receive a raw message from the receive channel.
		case msg := <-c.chainRelayers[chainID].usdcMsgRecvChan:
			go func() {
				// Fetch the circle attestation in a new goroutine so that we are not blocked from future requests.
				// TODO(dwasse): configure this backoff
				backoff.Retry(func() (err error) {
					msg.signature, err = getCircleAttestation(ctx, c.client, msg.txHash)
					return
				}, backoff.WithMaxRetries(backoff.NewConstantBackOff(time.Second), 5))
				if err != nil {
					logger.Errorf("could not get circle attestation: %w", err)
					return
				}

				// Send the completed message back through the send channel.
				c.chainRelayers[chainID].usdcMsgSendChan <- msg
			}()
		case <-c.chainRelayers[chainID].usdcMsgSendChan:
			// Submit the message to the destination chain.
			// TODO(dwasse): implement
		case <-ctx.Done():
			return nil
		}
	}
}

const circleAttestationURL = "https://iris-api-sandbox.circle.com/v1/attestations"

type circleAttestationResponse struct {
	Data struct {
		Attestation string `json:"attestation"`
		Status      string `json:"status"`
	} `json:"data"`
}

func getCircleAttestation(ctx context.Context, client *http.Client, txHash common.Hash) (signature []byte, err error) {
	url := fmt.Sprintf("%s/%s", circleAttestationURL, txHash.String())
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var attestationResp circleAttestationResponse
	err = json.Unmarshal(body, &attestationResp)
	if err != nil {
		err = fmt.Errorf("could not unmarshal body: %w", err)
		return
	}

	signature, err = hex.DecodeString(attestationResp.Data.Attestation)
	if err != nil {
		err = fmt.Errorf("could not decode signature: %w", err)
		return
	}
	return
}

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
