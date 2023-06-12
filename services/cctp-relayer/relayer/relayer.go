package relayer

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	db2 "github.com/synapsecns/sanguine/services/cctp-relayer/db"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sqlite"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/cctp-relayer/api"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmessagetransmitter"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// UsdcMessage contains data necessary to be posted on the destination chain.
type UsdcMessage struct {
	ChainID          uint32      // chain ID of the destination chain
	Message          []byte      // raw bytes of message produced by Circle's MessageTransmitter
	MessageHash      common.Hash // keccak256 hash of message bytes
	Signature        []byte      // attestation produced by Circle's API: https://developers.circle.com/stablecoin/reference/getattestation
	RequestVersion   uint32      // version of the request
	FormattedRequest []byte      // formatted request produced by SynapseCCTP
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
	usdcMsgRecvChan chan *UsdcMessage
	// usdcMsgSendChan contains outgoing usdc messages that are signed.
	usdcMsgSendChan chan *UsdcMessage
}

// CCTPRelayer listens for USDC burn events on origin chains,
// fetches attestations from Circle's API, and posts the necessary data
// on the destination chain to complete the USDC bridging process.
type CCTPRelayer struct {
	cfg           config.Config
	db            db2.CCTPRelayerDBReader
	scribeClient  client.ScribeClient
	grpcClient    pbscribe.ScribeServiceClient
	grpcConn      *grpc.ClientConn
	httpBackoff   backoff.BackOff
	omnirpcClient omniClient.RPCClient
	// chainRelayers is a map from chain ID -> chain relayer.
	chainRelayers map[uint32]*chainRelayer
	// handler is the metrics handler.
	handler metrics.Handler
	// attestationAPI is the client for Circle's REST API.
	attestationAPI api.AttestationAPI
	// txSubmitter is the tx submission service
	txSubmitter submitter.TransactionSubmitter
	// boundSynapseCCTPs is a map from chain ID -> SynapseCCTP.
	boundSynapseCCTPs map[uint32]*cctp.SynapseCCTP
}

const usdcMsgChanSize = 1000

// NewCCTPRelayer creates a new CCTPRelayer.
func NewCCTPRelayer(ctx context.Context, cfg config.Config, scribeClient client.ScribeClient, omniRPCClient omniClient.RPCClient, handler metrics.Handler, attestationAPI api.AttestationAPI) (*CCTPRelayer, error) {
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

	// Build chainRelayers and bound contracts.
	chainRelayers := make(map[uint32]*chainRelayer)
	boundSynapseCCTPs := make(map[uint32]*cctp.SynapseCCTP)
	for _, chain := range cfg.Chains {
		chainRelayers[chain.ChainID] = &chainRelayer{
			chainID:         chain.ChainID,
			closeConnection: make(chan bool, 1),
			stopListenChan:  make(chan bool, 1),
			usdcMsgRecvChan: make(chan *UsdcMessage, usdcMsgChanSize),
			usdcMsgSendChan: make(chan *UsdcMessage, usdcMsgChanSize),
		}
		client, err := omniRPCClient.GetClient(ctx, big.NewInt(int64(chain.ChainID)))
		if err != nil {
			return nil, fmt.Errorf("could not get client: %w", err)
		}
		boundSynapseCCTPs[chain.ChainID], err = cctp.NewSynapseCCTP(chain.GetDestinationAddress(), client)
		if err != nil {
			return nil, fmt.Errorf("could not build bound contract: %w", err)
		}
		fmt.Printf("Set synapseCCTP on chain %v at address %v\n", chain.ChainID, chain.GetDestinationAddress())
	}

	httpBackoff := backoff.NewExponentialBackOff()
	httpBackoff.InitialInterval = time.Duration(cfg.HTTPBackoffInitialIntervalMs) * time.Millisecond
	httpBackoff.MaxElapsedTime = time.Duration(cfg.HTTPBackoffMaxElapsedTimeMs) * time.Millisecond

	signer, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not make cctp signer: %w", err)
	}

	db, err := sqlite.NewSqliteStore(ctx, cfg.DBPrefix, handler, false)
	if err != nil {
		return nil, fmt.Errorf("could not make cctp db: %w", err)
	}

	txSubmitter := submitter.NewTransactionSubmitter(handler, signer, omniRPCClient, db.SubmitterDB(), &cfg.SubmitterConfig)

	return &CCTPRelayer{
		cfg:               cfg,
		omnirpcClient:     omniRPCClient,
		chainRelayers:     chainRelayers,
		scribeClient:      scribeClient,
		grpcClient:        grpcClient,
		grpcConn:          conn,
		httpBackoff:       httpBackoff,
		handler:           handler,
		attestationAPI:    attestationAPI,
		txSubmitter:       txSubmitter,
		boundSynapseCCTPs: boundSynapseCCTPs,
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
			return c.processBridgeEvents(ctx, chain.ChainID)
		})

		g.Go(func() error {
			return c.txSubmitter.Start(ctx)
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

// Listens for USDC send events on origin chain, and registers UsdcMessages to be signed.
//
//nolint:cyclop
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

// This takes in a log from the SynapseCCTP contract, determines the topic and then performs an action based on that topic.
// Note that the log could correspond to a send or receive event.
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
		err := c.handleCircleRequestSent(ctx, log.TxHash, originChain)
		if err != nil {
			return err
		}
	case cctp.CircleRequestFulfilledTopic:
		// TODO mark request as fulfilled
	default:
		// TODO; just continue
		logger.Warnf("unknown topic %s", log.Topics[0])
		return nil
	}
	return nil
}

//nolint:cyclop
func (c CCTPRelayer) handleCircleRequestSent(parentCtx context.Context, txhash common.Hash, originChain uint32) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "handleCircleRequestSent", trace.WithAttributes(
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

	// From this receipt, we expect two different logs:
	// - `messageSentEvent` gives us the raw bytes of the CCTP message
	// - `circleRequestSentEvent` gives us auxiliary data for SynapseCCTP
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
		err = ctx.Err()
		if err != nil {
			err = fmt.Errorf("error handling circle request sent: %w", err)
		}
		return err
	default:
		msg := UsdcMessage{
			ChainID:     uint32(circleRequestSentEvent.ChainId.Int64()),
			Message:     messageSentEvent.Message,
			MessageHash: crypto.Keccak256Hash(messageSentEvent.Message),
			//Signature: //comes from the api
			RequestVersion:   circleRequestSentEvent.RequestVersion,
			FormattedRequest: circleRequestSentEvent.FormattedRequest,
		}
		c.chainRelayers[originChain].usdcMsgRecvChan <- &msg
	}
	return nil
}

// Completes a USDC bridging sequence by calling ReceiveCircleToken() on the destination chain.
func (c CCTPRelayer) processBridgeEvents(ctx context.Context, chainID uint32) (err error) {
	for {
		select {
		// Receive a raw message from the receive channel.
		case msg := <-c.chainRelayers[chainID].usdcMsgRecvChan:
			// Fetch the circle attestation in a new goroutine so that we are not blocked from future requests.
			go c.fetchAttestation(ctx, chainID, msg)
		case msg := <-c.chainRelayers[chainID].usdcMsgSendChan:
			// Submit the message to the destination chain.
			go c.submitReceiveCircleToken(ctx, msg)
		case <-ctx.Done():
			return nil
		}
	}
}

func (c CCTPRelayer) fetchAttestation(parentCtx context.Context, chainID uint32, msg *UsdcMessage) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "fetchAttestation", trace.WithAttributes(
		attribute.String("messageHash", msg.MessageHash.String()),
		attribute.Int(metrics.ChainID, int(chainID)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	err = backoff.Retry(func() (err error) {
		msg.Signature, err = c.attestationAPI.GetAttestation(ctx, msg.MessageHash)
		return
	}, c.httpBackoff)
	if err != nil {
		return
	}

	// Send the completed message back through the send channel.
	c.chainRelayers[chainID].usdcMsgSendChan <- msg
	return
}

func (c CCTPRelayer) submitReceiveCircleToken(parentCtx context.Context, msg *UsdcMessage) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "submitReceiveCircleToken", trace.WithAttributes(
		attribute.String("messageHash", msg.MessageHash.String()),
		attribute.Int(metrics.ChainID, int(msg.ChainID)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	_, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(msg.ChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		fmt.Printf("ChainID: %v\n", msg.ChainID)
		contract := c.boundSynapseCCTPs[msg.ChainID]
		fmt.Printf("transactor: %v\n", transactor)
		fmt.Printf("contract: %v\n", contract)
		fmt.Printf("msg: %v\n", msg)
		fmt.Printf("synapsecctps: %v\n", c.boundSynapseCCTPs)
		return contract.ReceiveCircleToken(transactor, msg.Message, msg.Signature, msg.RequestVersion, msg.FormattedRequest)
	})
	if err != nil {
		err = fmt.Errorf("could not submit transaction: %w", err)
	}
	return
}
