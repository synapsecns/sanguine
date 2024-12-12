// Package client provides a client for the RFQ quoting API.
// nolint:,wrapcheck
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ipfs/go-log"

	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
)

var logger = log.Logger("rfq-client")

const (
	pingPeriod = 20 * time.Second
	chanBuffer = 1000
)

// AuthenticatedClient is an interface for the RFQ API.
// It provides methods for creating, retrieving and updating quotes.
type AuthenticatedClient interface {
	PutQuote(ctx context.Context, q *model.PutRelayerQuoteRequest) error
	PutBulkQuotes(ctx context.Context, q *model.PutBulkQuotesRequest) error
	PutRelayAck(ctx context.Context, req *model.PutAckRequest) (*model.PutRelayAckResponse, error)
	SubscribeActiveQuotes(ctx context.Context, req *model.SubscribeActiveRFQRequest, reqChan chan *model.ActiveRFQMessage) (respChan chan *model.ActiveRFQMessage, err error)
	UnauthenticatedClient
}

// UnauthenticatedClient is an interface for the RFQ API.
type UnauthenticatedClient interface {
	GetAllQuotes(ctx context.Context) ([]*model.GetQuoteResponse, error)
	GetSpecificQuote(ctx context.Context, q *model.GetQuoteSpecificRequest) ([]*model.GetQuoteResponse, error)
	GetQuoteByRelayerAddress(ctx context.Context, relayerAddr string) ([]*model.GetQuoteResponse, error)
	GetRFQContracts(ctx context.Context) (*model.GetContractsResponse, error)
	PutRFQRequest(ctx context.Context, q *model.PutRFQRequest) (*model.PutRFQResponse, error)
	resty() *resty.Client
}

type unauthenticatedClient struct {
	rClient *resty.Client
}

func (c unauthenticatedClient) resty() *resty.Client {
	return c.rClient
}

type clientImpl struct {
	UnauthenticatedClient
	rClient   *resty.Client
	reqSigner signer.Signer
}

// NewAuthenticatedClient creates a new client for the RFQ quoting API.
// TODO: @aurelius,  you don't actually need to be authed for GET Requests.
func NewAuthenticatedClient(metrics metrics.Handler, rfqURL string, reqSigner signer.Signer) (AuthenticatedClient, error) {
	unauthedClient, err := NewUnauthenticatedClient(metrics, rfqURL)
	if err != nil {
		return nil, fmt.Errorf("could not create unauthenticated client: %w", err)
	}

	// since this is a pointer, all requests going forward will be authenticated. This is assigned
	// to a new variable for clarity.
	authedClient := unauthedClient.resty().
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			authHeader, err := getAuthHeader(request.Context(), reqSigner)
			if err != nil {
				return fmt.Errorf("failed to get auth header: %w", err)
			}
			request.SetHeader(rest.AuthorizationHeader, authHeader)
			return nil
		})

	return &clientImpl{
		UnauthenticatedClient: unauthedClient,
		rClient:               authedClient,
		reqSigner:             reqSigner,
	}, nil
}

func getAuthHeader(ctx context.Context, reqSigner signer.Signer) (string, error) {
	// if request.Method == "PUT" && request.URL == rfqURL+rest.QUOTE_ROUTE {
	// i.e. signature (hex encoded) = keccak(bytes.concat("\x19Ethereum Signed Message:\n", len(strconv.Itoa(time.Now().Unix()), strconv.Itoa(time.Now().Unix())))
	// so that full auth header string: auth = strconv.Itoa(time.Now().Unix()) + ":" + signature
	// Get the current Unix timestamp as a string.
	now := strconv.Itoa(int(time.Now().Unix()))

	// Prepare the data to be signed.
	data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(now)) + now

	sig, err := reqSigner.SignMessage(ctx, []byte(data), true)

	if err != nil {
		return "", fmt.Errorf("failed to sign request: %w", err)
	}

	return fmt.Sprintf("%s:%s", now, hexutil.Encode(signer.Encode(sig))), nil
}

// NewUnauthenticatedClient creates a new client for the RFQ quoting API.
func NewUnauthenticatedClient(metricHandler metrics.Handler, rfqURL string) (UnauthenticatedClient, error) {
	client := resty.New().
		SetBaseURL(rfqURL).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			request.Header.Add(ginhelper.RequestIDHeader, uuid.New().String())
			return nil
		})
	client.SetTransport(
		otelhttp.NewTransport(client.GetClient().Transport,
			otelhttp.WithTracerProvider(
				metricHandler.GetTracerProvider()),
			otelhttp.WithSpanNameFormatter(
				func(_ string, r *http.Request) string {
					return fmt.Sprintf("rfq-api %s", r.Method)
				},
			),
		),
	)
	return &unauthenticatedClient{client}, nil
}

// PutQuote puts a new quote in the RFQ quoting API.
func (c *clientImpl) PutQuote(ctx context.Context, q *model.PutRelayerQuoteRequest) error {
	res, err := c.rClient.R().
		SetContext(ctx).
		SetBody(q).
		Put(rest.QuoteRoute)

	// TODO: Figure out if there's anything to do with the response, right now it's result: Status Code 200 OK
	_ = res

	return err
}

// PutBulkQuotes puts multiple new quotes in the RFQ quoting API.
func (c *clientImpl) PutBulkQuotes(ctx context.Context, q *model.PutBulkQuotesRequest) error {
	res, err := c.rClient.R().
		SetContext(ctx).
		SetBody(q).
		Put(rest.BulkQuotesRoute)

	// TODO: Figure out if there's anything to do with the response, right now it's result: Status Code 200 OK
	_ = res

	return err
}

func (c *clientImpl) PutRelayAck(ctx context.Context, req *model.PutAckRequest) (*model.PutRelayAckResponse, error) {
	var ack *model.PutRelayAckResponse
	resp, err := c.rClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&ack).
		Put(rest.AckRoute)

	if err != nil {
		return nil, fmt.Errorf("error from server: %s %w", getStatus(resp), err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", getStatus(resp))
	}

	return ack, nil
}

func (c *clientImpl) SubscribeActiveQuotes(ctx context.Context, req *model.SubscribeActiveRFQRequest, reqChan chan *model.ActiveRFQMessage) (respChan chan *model.ActiveRFQMessage, err error) {
	conn, err := c.connectWebsocket(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to websocket: %w", err)
	}
	// first, subscrbe to the given chains
	sub := model.SubscriptionParams{
		Chains: req.ChainIDs,
	}
	subJSON, err := json.Marshal(sub)
	if err != nil {
		return respChan, fmt.Errorf("error marshaling subscription params: %w", err)
	}
	err = conn.WriteJSON(model.ActiveRFQMessage{
		Op:      rest.SubscribeOp,
		Content: json.RawMessage(subJSON),
	})
	if err != nil {
		return nil, fmt.Errorf("error sending subscribe message: %w", err)
	}
	// make sure subscription is successful
	var resp model.ActiveRFQMessage
	err = conn.ReadJSON(&resp)
	if err != nil {
		return nil, fmt.Errorf("error reading subscribe response: %w", err)
	}
	if !resp.Success || resp.Op != rest.SubscribeOp {
		return nil, fmt.Errorf("subscription failed")
	}

	respChan = make(chan *model.ActiveRFQMessage, chanBuffer)
	go func() {
		wsErr := c.processWebsocket(ctx, conn, reqChan, respChan)
		if wsErr != nil {
			logger.Error("Error running websocket listener: %s", wsErr)
		}
	}()

	return respChan, nil
}

func (c *clientImpl) connectWebsocket(ctx context.Context, req *model.SubscribeActiveRFQRequest) (conn *websocket.Conn, err error) {
	if len(req.ChainIDs) == 0 {
		return nil, fmt.Errorf("chain IDs are required")
	}

	header, err := c.getWsHeaders(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth header: %w", err)
	}

	reqURL := strings.Replace(c.rClient.BaseURL, "http", "ws", 1) + rest.RFQStreamRoute
	conn, httpResp, err := websocket.DefaultDialer.Dial(reqURL, header)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to websocket: %w", err)
	}
	err = httpResp.Body.Close()
	if err != nil {
		logger.Warnf("error closing websocket connection: %v", err)
	}

	return conn, nil
}

func (c *clientImpl) getWsHeaders(ctx context.Context, req *model.SubscribeActiveRFQRequest) (header http.Header, err error) {
	header = http.Header{}
	chainIDsJSON, err := json.Marshal(req.ChainIDs)
	if err != nil {
		return header, fmt.Errorf("failed to marshal chain IDs: %w", err)
	}
	header.Set(rest.ChainsHeader, string(chainIDsJSON))
	authHeader, err := getAuthHeader(ctx, c.reqSigner)
	if err != nil {
		return header, fmt.Errorf("failed to get auth header: %w", err)
	}
	header.Set(rest.AuthorizationHeader, authHeader)
	return header, nil
}

func (c *clientImpl) processWebsocket(ctx context.Context, conn *websocket.Conn, reqChan, respChan chan *model.ActiveRFQMessage) (err error) {
	defer func() {
		close(respChan)
		err := conn.Close()
		if err != nil {
			logger.Warnf("error closing websocket connection: %v", err)
		}
		panic("processWebsocket exited")
	}()

	readChan := make(chan []byte, chanBuffer)
	go c.listenWsMessages(ctx, conn, readChan)
	go c.sendPings(ctx, reqChan)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg, ok := <-reqChan:
			if !ok {
				return fmt.Errorf("error reading from reqChan: %w", ctx.Err())
			}
			err := conn.WriteJSON(msg)
			if err != nil {
				return fmt.Errorf("error sending message to websocket: %w", err)
			}
		case msg, ok := <-readChan:
			if !ok {
				return nil
			}
			err = c.handleWsMessage(ctx, msg, respChan)
			if err != nil {
				return fmt.Errorf("error handling websocket message: %w", err)
			}
		}
	}
}

func (c *clientImpl) sendPings(ctx context.Context, reqChan chan *model.ActiveRFQMessage) {
	pingTicker := time.NewTicker(pingPeriod)
	defer pingTicker.Stop()

	for {
		select {
		case <-pingTicker.C:
			pingMsg := model.ActiveRFQMessage{
				Op: rest.PingOp,
			}
			reqChan <- &pingMsg
		case <-ctx.Done():
			return
		}
	}
}
func (c *clientImpl) listenWsMessages(ctx context.Context, conn *websocket.Conn, readChan chan []byte) {
	defer close(readChan)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Warnf("websocket connection closed unexpectedly: %v", err)
			}
			return
		}
		select {
		case readChan <- message:
		case <-ctx.Done():
			return
		}
	}
}

func (c *clientImpl) handleWsMessage(ctx context.Context, msg []byte, respChan chan *model.ActiveRFQMessage) (err error) {
	var rfqMsg model.ActiveRFQMessage
	err = json.Unmarshal(msg, &rfqMsg)
	if err != nil {
		return fmt.Errorf("error unmarshaling message: %w", err)
	}

	select {
	case respChan <- &rfqMsg:
	case <-ctx.Done():
		return nil
	}
	return nil
}

// GetAllQuotes retrieves all quotes from the RFQ quoting API.
func (c *unauthenticatedClient) GetAllQuotes(ctx context.Context) ([]*model.GetQuoteResponse, error) {
	var quotes []*model.GetQuoteResponse
	resp, err := c.rClient.R().
		SetContext(ctx).
		SetResult(&quotes).
		Get(rest.QuoteRoute)

	if err != nil {
		return nil, fmt.Errorf("error from server: %s: %w", getStatus(resp), err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", getStatus(resp))
	}

	return quotes, nil
}

// GetSpecificQuote retrieves a specific quote from the RFQ quoting API.
func (c *unauthenticatedClient) GetSpecificQuote(ctx context.Context, q *model.GetQuoteSpecificRequest) ([]*model.GetQuoteResponse, error) {
	var quotes []*model.GetQuoteResponse
	resp, err := c.rClient.R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"originChainId":   strconv.Itoa(q.OriginChainID),
			"originTokenAddr": q.OriginTokenAddr,
			"destChainId":     strconv.Itoa(q.DestChainID),
			"destTokenAddr":   q.DestTokenAddr,
		}).
		SetResult(&quotes).
		Get(rest.QuoteRoute)

	if err != nil {
		return nil, fmt.Errorf("error from server: %s: %w", getStatus(resp), err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", getStatus(resp))
	}

	return quotes, nil
}

func (c *unauthenticatedClient) GetQuoteByRelayerAddress(ctx context.Context, relayerAddr string) ([]*model.GetQuoteResponse, error) {
	var quotes []*model.GetQuoteResponse
	resp, err := c.rClient.R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"relayerAddr": relayerAddr,
		}).
		SetResult(&quotes).
		Get(rest.QuoteRoute)

	if err != nil {
		return nil, fmt.Errorf("error from server: %s %w", getStatus(resp), err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", getStatus(resp))
	}

	return quotes, nil
}

func (c unauthenticatedClient) GetRFQContracts(ctx context.Context) (*model.GetContractsResponse, error) {
	var contracts *model.GetContractsResponse
	resp, err := c.rClient.R().
		SetContext(ctx).
		SetResult(&contracts).
		Get(rest.ContractsRoute)

	if err != nil {
		return nil, fmt.Errorf("error from server: %s %w", getStatus(resp), err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", getStatus(resp))
	}

	return contracts, nil
}

func (c unauthenticatedClient) PutRFQRequest(ctx context.Context, q *model.PutRFQRequest) (*model.PutRFQResponse, error) {
	var response model.PutRFQResponse
	resp, err := c.rClient.R().
		SetContext(ctx).
		SetBody(q).
		SetResult(&response).
		Put(rest.RFQRoute)

	if err != nil {
		return nil, fmt.Errorf("error from server: %s: %w", getStatus(resp), err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error from server: %s", getStatus(resp))
	}

	return &response, nil
}

func getStatus(resp *resty.Response) string {
	if resp == nil {
		return "http status unavailable"
	}
	return resp.Status()
}
