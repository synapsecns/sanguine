package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// WsClient is a client for the WebSocket API.
type WsClient interface {
	Run(ctx context.Context) error
	SendQuoteRequest(ctx context.Context, quoteRequest *model.WsRFQRequest) error
	ReceiveQuoteResponse(ctx context.Context, requestID string) (*model.WsRFQResponse, error)
}

type wsClient struct {
	handler       metrics.Handler
	relayerAddr   string
	conn          *websocket.Conn
	pubsub        PubSubManager
	requestChan   chan *model.WsRFQRequest
	responseChans *xsync.MapOf[string, chan *model.WsRFQResponse]
	doneChan      chan struct{}
	pingTicker    *time.Ticker
	lastPing      time.Time
}

func newWsClient(relayerAddr string, conn *websocket.Conn, pubsub PubSubManager, handler metrics.Handler) *wsClient {
	return &wsClient{
		handler:       handler,
		relayerAddr:   relayerAddr,
		conn:          conn,
		pubsub:        pubsub,
		requestChan:   make(chan *model.WsRFQRequest),
		responseChans: xsync.NewMapOf[chan *model.WsRFQResponse](),
		doneChan:      make(chan struct{}),
		pingTicker:    time.NewTicker(pingPeriod),
	}
}

func (c *wsClient) SendQuoteRequest(ctx context.Context, quoteRequest *model.WsRFQRequest) error {
	fmt.Printf("wsClient.SendQuoteRequest with data: %+v\n", quoteRequest.Data)
	select {
	case c.requestChan <- quoteRequest:
		// successfully sent, register a response channel
		c.responseChans.Store(quoteRequest.RequestID, make(chan *model.WsRFQResponse))
	case <-c.doneChan:
		return fmt.Errorf("websocket client is closed")
	case <-ctx.Done():
		return nil
	}
	return nil
}

func (c *wsClient) ReceiveQuoteResponse(ctx context.Context, requestID string) (resp *model.WsRFQResponse, err error) {
	responseChan, ok := c.responseChans.Load(requestID)
	if !ok {
		return nil, fmt.Errorf("no response channel for request %s", requestID)
	}
	defer c.responseChans.Delete(requestID)

	for {
		select {
		case resp = <-responseChan:
			// successfully received
			return resp, nil
		case <-c.doneChan:
			return nil, fmt.Errorf("websocket client is closed")
		case <-ctx.Done():
			return nil, fmt.Errorf("expiration reached without response")
		}
	}
}

const (
	// PongOp is the operation for a pong message.
	PongOp = "pong"
	// PingOp is the operation for a ping message.
	PingOp = "ping"
	// SubscribeOp is the operation for a subscribe message.
	SubscribeOp = "subscribe"
	// UnsubscribeOp is the operation for an unsubscribe message.
	UnsubscribeOp = "unsubscribe"
	// RequestQuoteOp is the operation for a request quote message.
	RequestQuoteOp = "request_quote"
	// SendQuoteOp is the operation for a send quote message.
	SendQuoteOp = "send_quote"
	// pingPeriod is the period for a ping message.
	pingPeriod = 1 * time.Minute
)

// Run runs the WebSocket client.
func (c *wsClient) Run(ctx context.Context) (err error) {
	messageChan := make(chan []byte)

	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := pollWsMessages(gctx, c.conn, messageChan)
		if err != nil {
			return fmt.Errorf("error polling websocket messages: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		err := c.processWs(gctx, messageChan)
		if err != nil {
			return fmt.Errorf("error processing websocket messages: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("error running websocket client: %w", err)
	}

	return nil
}

func pollWsMessages(ctx context.Context, conn *websocket.Conn, messageChan chan []byte) (err error) {
	defer close(messageChan)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("error reading websocket message: %w", err)
		}
		select {
		case <-ctx.Done():
			return nil
		case messageChan <- msg:
		}
	}
}

func (c *wsClient) processWs(ctx context.Context, messageChan chan []byte) (err error) {
	defer c.pingTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			err = c.conn.Close()
			if err != nil {
				return fmt.Errorf("error closing websocket connection: %w", err)
			}
			close(c.doneChan)
			return fmt.Errorf("websocket client is closed")
		case req := <-c.requestChan:
			err = c.sendRelayerRequest(ctx, req)
			if err != nil {
				logger.Error("Error sending quote request: %s", err)
			}
		case msg := <-messageChan:
			err = c.handleRelayerMessage(ctx, msg)
			if err != nil {
				logger.Error("Error handling relayer message: %s", err)
				return fmt.Errorf("error handling relayer message: %w", err)
			}
		case <-c.pingTicker.C:
			// ping timed out, close the connection
			_, span := c.handler.Tracer().Start(ctx, "pingTimeout")
			defer metrics.EndSpanWithErr(span, err)
		}
	}
}

func (c *wsClient) sendRelayerRequest(ctx context.Context, req *model.WsRFQRequest) (err error) {
	fmt.Printf("sendRelayerRequest with data: %+v\n", req.Data)
	_, span := c.handler.Tracer().Start(ctx, "sendRelayerRequest", trace.WithAttributes(
		attribute.String("relayer_address", c.relayerAddr),
		attribute.String("request_id", req.RequestID),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	rawData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("error marshaling quote request: %w", err)
	}
	fmt.Printf("rawData: %+v\n", string(rawData))
	msg := model.ActiveRFQMessage{
		Op:      RequestQuoteOp,
		Content: json.RawMessage(rawData),
	}
	fmt.Printf("writing raw msg: %+v\n", msg)
	err = c.conn.WriteJSON(msg)
	if err != nil {
		return fmt.Errorf("error sending quote request: %w", err)
	}

	return nil
}

// handleRelayerMessage handles messages from the relayer.
// An error returned will result in the websocket connection being closed.
//
//nolint:cyclop
func (c *wsClient) handleRelayerMessage(ctx context.Context, msg []byte) (err error) {
	_, span := c.handler.Tracer().Start(ctx, "handleRelayerMessage", trace.WithAttributes(
		attribute.String("relayer_address", c.relayerAddr),
		attribute.String("message", string(msg)),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var rfqMsg model.ActiveRFQMessage
	err = json.Unmarshal(msg, &rfqMsg)
	if err != nil {
		return fmt.Errorf("error unmarshaling websocket message: %w", err)
	}

	switch rfqMsg.Op {
	case PingOp:
		c.lastPing = time.Now()
		resp := c.handlePing(ctx)
		fmt.Printf("[%v] writing pong resp: %+v\n", time.Now(), resp)
		err = c.conn.WriteJSON(resp)
		if err != nil {
			return fmt.Errorf("error sending ping response: %w", err)
		}
		fmt.Printf("[%v] wrote pong resp: %+v\n", time.Now(), resp)
	case SubscribeOp:
		resp := c.handleSubscribe(ctx, rfqMsg.Content)
		err = c.conn.WriteJSON(resp)
		if err != nil {
			return fmt.Errorf("error sending subscribe response: %w", err)
		}
	case UnsubscribeOp:
		resp := c.handleUnsubscribe(ctx, rfqMsg.Content)
		err = c.conn.WriteJSON(resp)
		if err != nil {
			return fmt.Errorf("error sending unsubscribe response: %w", err)
		}
	case SendQuoteOp:
		err = c.handleSendQuote(ctx, rfqMsg.Content)
		if err != nil {
			logger.Errorf("error handling send quote: %v", err)
		}
	default:
		logger.Errorf("received unexpected operation from relayer: %s", rfqMsg.Op)
		return nil
	}

	return nil
}

func (c *wsClient) handlePing(ctx context.Context) (resp model.ActiveRFQMessage) {
	fmt.Printf("[%v] handlePing", time.Now())
	_, span := c.handler.Tracer().Start(ctx, "handlePing", trace.WithAttributes(
		attribute.String("relayer_address", c.relayerAddr),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	return getSuccessResponse(PongOp)
}

func (c *wsClient) handleSubscribe(ctx context.Context, content json.RawMessage) (resp model.ActiveRFQMessage) {
	_, span := c.handler.Tracer().Start(ctx, "handleSubscribe", trace.WithAttributes(
		attribute.String("relayer_address", c.relayerAddr),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	var sub model.SubscriptionParams
	err := json.Unmarshal(content, &sub)
	if err != nil {
		return getErrorResponse(SubscribeOp, fmt.Errorf("could not unmarshal subscription params: %w", err))
	}
	span.SetAttributes(attribute.IntSlice("chain_ids", sub.Chains))
	err = c.pubsub.AddSubscription(c.relayerAddr, sub)
	if err != nil {
		return getErrorResponse(SubscribeOp, fmt.Errorf("error adding subscription: %w", err))
	}
	return getSuccessResponse(SubscribeOp)
}

func (c *wsClient) handleUnsubscribe(ctx context.Context, content json.RawMessage) (resp model.ActiveRFQMessage) {
	_, span := c.handler.Tracer().Start(ctx, "handleUnsubscribe", trace.WithAttributes(
		attribute.String("relayer_address", c.relayerAddr),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	var sub model.SubscriptionParams
	err := json.Unmarshal(content, &sub)
	if err != nil {
		return getErrorResponse(UnsubscribeOp, fmt.Errorf("could not unmarshal subscription params: %w", err))
	}
	span.SetAttributes(attribute.IntSlice("chain_ids", sub.Chains))
	err = c.pubsub.RemoveSubscription(c.relayerAddr, sub)
	if err != nil {
		return getErrorResponse(UnsubscribeOp, fmt.Errorf("error removing subscription: %w", err))
	}
	return getSuccessResponse(UnsubscribeOp)
}

func (c *wsClient) handleSendQuote(ctx context.Context, content json.RawMessage) (err error) {
	_, span := c.handler.Tracer().Start(ctx, "handleSendQuote", trace.WithAttributes(
		attribute.String("relayer_address", c.relayerAddr),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	// forward the response to the server
	var resp model.WsRFQResponse
	err = json.Unmarshal(content, &resp)
	if err != nil {
		return fmt.Errorf("error unmarshaling websocket message: %w", err)
	}
	span.SetAttributes(
		attribute.String("request_id", resp.RequestID),
		attribute.String("dest_amount", resp.DestAmount),
	)
	responseChan, ok := c.responseChans.Load(resp.RequestID)
	if !ok {
		return fmt.Errorf("no response channel for request %s", resp.RequestID)
	}
	responseChan <- &resp

	return nil
}

func getSuccessResponse(op string) model.ActiveRFQMessage {
	return model.ActiveRFQMessage{
		Op:      op,
		Success: true,
	}
}

func getErrorResponse(op string, err error) model.ActiveRFQMessage {
	return model.ActiveRFQMessage{
		Op:      op,
		Success: false,
		Content: json.RawMessage(fmt.Sprintf("{\"reason\": \"%s\"}", err.Error())),
	}
}
