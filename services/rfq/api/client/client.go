// Package client provides a client for the RFQ quoting API.
// nolint:,wrapcheck
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	PutUserQuoteRequest(ctx context.Context, q *model.PutUserQuoteRequest) (*model.PutUserQuoteResponse, error)
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
	wsURL     *string
	reqSigner signer.Signer
}

// NewAuthenticatedClient creates a new client for the RFQ quoting API.
// TODO: @aurelius,  you don't actually need to be authed for GET Requests.
func NewAuthenticatedClient(metrics metrics.Handler, rfqURL string, wsURL *string, reqSigner signer.Signer) (AuthenticatedClient, error) {
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
		wsURL:                 wsURL,
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
	if c.wsURL == nil {
		return nil, fmt.Errorf("websocket URL is not set")
	}
	if len(req.ChainIDs) == 0 {
		return nil, fmt.Errorf("chain IDs are required")
	}

	reqURL := *c.wsURL + rest.QuoteRequestsRoute

	header := http.Header{}
	chainIDsJSON, err := json.Marshal(req.ChainIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal chain IDs: %w", err)
	}
	header.Set(rest.ChainsHeader, string(chainIDsJSON))
	authHeader, err := getAuthHeader(ctx, c.reqSigner)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth header: %w", err)
	}
	header.Set(rest.AuthorizationHeader, authHeader)

	conn, _, err := websocket.DefaultDialer.Dial(reqURL, header)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to websocket: %w", err)
	}

	respChan = make(chan *model.ActiveRFQMessage, 1000)

	// first, subscrbe to the given chains
	sub := model.SubscriptionParams{
		Chains: req.ChainIDs,
	}
	subJSON, err := json.Marshal(sub)
	if err != nil {
		return respChan, fmt.Errorf("error marshalling subscription params: %w", err)
	}
	conn.WriteJSON(model.ActiveRFQMessage{
		Op:      rest.SubscribeOp,
		Content: json.RawMessage(subJSON),
	})

	// make sure subscription is successful
	var resp model.ActiveRFQMessage
	conn.ReadJSON(&resp)
	if !resp.Success || resp.Op != rest.SubscribeOp {
		return nil, fmt.Errorf("subscription failed")
	}

	go func() {
		defer close(respChan)
		defer conn.Close()

		readChan := make(chan []byte, 1000)
		go func() {
			defer close(readChan)
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						logger.Warnf("websocket connection closed unexpectedly: %v", err)
					}
					return
				}
				readChan <- message
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-reqChan:
				if !ok {
					return
				}
				err := conn.WriteJSON(msg)
				if err != nil {
					logger.Warnf("error sending message to websocket: %v", err)
					return
				}
			case msg, ok := <-readChan:
				if !ok {
					return
				}
				var rfqMsg model.ActiveRFQMessage
				err = json.Unmarshal(msg, &rfqMsg)
				if err != nil {
					logger.Warn("error unmarshalling message: %v", err)
					continue
				}

				// automatically send the pong
				if rfqMsg.Op == rest.PingOp {
					reqChan <- &model.ActiveRFQMessage{
						Op: rest.PongOp,
					}
					continue
				}

				select {
				case respChan <- &rfqMsg:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return respChan, nil
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

func (c unauthenticatedClient) PutUserQuoteRequest(ctx context.Context, q *model.PutUserQuoteRequest) (*model.PutUserQuoteResponse, error) {
	var response model.PutUserQuoteResponse
	resp, err := c.rClient.R().
		SetContext(ctx).
		SetBody(q).
		SetResult(&response).
		Put(rest.PutQuoteRequestRoute)

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
