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
	lastPong      time.Time
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
	}
}

func (c *wsClient) SendQuoteRequest(ctx context.Context, quoteRequest *model.WsRFQRequest) error {
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
	// PingPeriod is the period for a ping message.
	PingPeriod = 15 * time.Second
)

// Run runs the WebSocket client.
func (c *wsClient) Run(ctx context.Context) (err error) {
	fmt.Println("running ws client")
	c.lastPong = time.Now()
	messageChan := make(chan []byte)
	pingTicker := time.NewTicker(PingPeriod)
	defer pingTicker.Stop()

	// poll messages from websocket in background
	go pollWsMessages(c.conn, messageChan)

	for {
		select {
		case <-ctx.Done():
			err = c.conn.Close()
			if err != nil {
				return fmt.Errorf("error closing websocket connection: %w", err)
			}
			close(c.doneChan)
			return nil
		case req := <-c.requestChan:
			fmt.Printf("sending quote request: %v\n", req)
			err = c.sendRelayerRequest(ctx, req)
			if err != nil {
				fmt.Printf("error sending quote request: %v\n", err)
				logger.Error("Error sending quote request: %s", err)
			}
		case msg := <-messageChan:
			fmt.Printf("received message: %s\n", msg)
			err = c.handleRelayerMessage(ctx, msg)
			if err != nil {
				fmt.Printf("error handling relayer message: %v\n", err)
				logger.Error("Error handling relayer message: %s", err)
			}
		case <-pingTicker.C:
			err = c.trySendPing(c.lastPong)
			if err != nil {
				logger.Error("Error sending ping message: %s", err)
			}
		}
	}
}

func pollWsMessages(conn *websocket.Conn, messageChan chan []byte) {
	defer close(messageChan)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logger.Error("Error reading websocket message: %s", err)
			return
		}
		messageChan <- msg
	}
}

func (c *wsClient) sendRelayerRequest(ctx context.Context, req *model.WsRFQRequest) (err error) {
	ctx, span := c.handler.Tracer().Start(ctx, "sendRelayerRequest", trace.WithAttributes(
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
	msg := model.ActiveRFQMessage{
		Op:      RequestQuoteOp,
		Content: json.RawMessage(rawData),
	}
	err = c.conn.WriteJSON(msg)
	if err != nil {
		return fmt.Errorf("error sending quote request: %w", err)
	}

	return nil
}

func (c *wsClient) handleRelayerMessage(ctx context.Context, msg []byte) (err error) {
	var rfqMsg model.ActiveRFQMessage
	err = json.Unmarshal(msg, &rfqMsg)
	if err != nil {
		return fmt.Errorf("error unmarshaling websocket message: %w", err)
	}

	switch rfqMsg.Op {
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
			return fmt.Errorf("error handling send quote: %w", err)
		}
	case PongOp:
		c.lastPong = time.Now()
	default:
		return fmt.Errorf("received unexpected operation from relayer: %s", rfqMsg.Op)
	}

	return nil
}

func (c *wsClient) handleSubscribe(ctx context.Context, content json.RawMessage) (resp model.ActiveRFQMessage) {
	fmt.Printf("handleSubscribe: %s\n", content)
	ctx, span := c.handler.Tracer().Start(ctx, "handleSubscribe", trace.WithAttributes(
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
		fmt.Printf("error adding subscription: %v\n", err)
		return getErrorResponse(SubscribeOp, fmt.Errorf("error adding subscription: %w", err))
	}
	fmt.Printf("successfully added subscription for chain ids: %v\n", sub.Chains)
	return getSuccessResponse(SubscribeOp)
}

func (c *wsClient) handleUnsubscribe(ctx context.Context, content json.RawMessage) (resp model.ActiveRFQMessage) {
	ctx, span := c.handler.Tracer().Start(ctx, "handleUnsubscribe", trace.WithAttributes(
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
	ctx, span := c.handler.Tracer().Start(ctx, "handleSendQuote", trace.WithAttributes(
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

func (c *wsClient) trySendPing(lastPong time.Time) (err error) {
	if time.Since(lastPong) > PingPeriod {
		err = c.conn.Close()
		if err != nil {
			return fmt.Errorf("error closing websocket connection: %w", err)
		}
		close(c.doneChan)
		return fmt.Errorf("pong not received in time")
	}
	pingMsg := model.ActiveRFQMessage{
		Op: PingOp,
	}
	err = c.conn.WriteJSON(pingMsg)
	if err != nil {
		return fmt.Errorf("error sending ping message: %w", err)
	}

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
