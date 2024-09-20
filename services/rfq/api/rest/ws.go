package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// WsClient is a client for the WebSocket API.
type WsClient interface {
	Run(ctx context.Context) error
	SendQuoteRequest(ctx context.Context, quoteRequest *model.RelayerWsQuoteRequest) error
	ReceiveQuoteResponse(ctx context.Context) (*model.RelayerWsQuoteResponse, error)
}

type wsClient struct {
	relayerAddr  string
	conn         *websocket.Conn
	pubsub       PubSubManager
	requestChan  chan *model.RelayerWsQuoteRequest
	responseChan chan *model.RelayerWsQuoteResponse
	doneChan     chan struct{}
}

func newWsClient(relayerAddr string, conn *websocket.Conn, pubsub PubSubManager) *wsClient {
	return &wsClient{
		relayerAddr:  relayerAddr,
		conn:         conn,
		pubsub:       pubsub,
		requestChan:  make(chan *model.RelayerWsQuoteRequest),
		responseChan: make(chan *model.RelayerWsQuoteResponse),
		doneChan:     make(chan struct{}),
	}
}

func (c *wsClient) SendQuoteRequest(ctx context.Context, quoteRequest *model.RelayerWsQuoteRequest) error {
	select {
	case c.requestChan <- quoteRequest:
		// successfully sent
	case <-c.doneChan:
		return fmt.Errorf("websocket client is closed")
	case <-ctx.Done():
		return nil
	}
	return nil
}

func (c *wsClient) ReceiveQuoteResponse(ctx context.Context) (resp *model.RelayerWsQuoteResponse, err error) {
	for {
		select {
		case resp = <-c.responseChan:
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
	// PongOp is the operation for a pong message
	PongOp = "pong"
	// PingOp is the operation for a ping message
	PingOp = "ping"
	// SubscribeOp is the operation for a subscribe message
	SubscribeOp = "subscribe"
	// UnsubscribeOp is the operation for an unsubscribe message
	UnsubscribeOp = "unsubscribe"
	// RequestQuoteOp is the operation for a request quote message
	RequestQuoteOp = "request_quote"
	// SendQuoteOp is the operation for a send quote message
	SendQuoteOp = "send_quote"
	// PingPeriod is the period for a ping message
	PingPeriod = 15 * time.Second
)

func (c *wsClient) Run(ctx context.Context) (err error) {
	messageChan := make(chan []byte)
	pingTicker := time.NewTicker(PingPeriod)
	defer pingTicker.Stop()

	lastPong := time.Now()

	// Goroutine to read messages from WebSocket and send to channel
	go func() {
		defer close(messageChan)
		for {
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				logger.Error("Error reading websocket message: %s", err)
				return
			}
			messageChan <- msg
		}
	}()

	for {
		select {
		case <-ctx.Done():
			err = c.conn.Close()
			if err != nil {
				return fmt.Errorf("error closing websocket connection: %w", err)
			}
			close(c.doneChan)
			return nil
		case data := <-c.requestChan:
			rawData, err := json.Marshal(data)
			if err != nil {
				logger.Error("Error marshaling quote request: %s", err)
				continue
			}
			msg := model.ActiveRFQMessage{
				Op:      RequestQuoteOp,
				Content: json.RawMessage(rawData),
			}
			err = c.conn.WriteJSON(msg)
			if err != nil {
				logger.Error("Error sending quote request: %s", err)
				continue
			}
		case msg := <-messageChan:
			var rfqMsg model.ActiveRFQMessage
			err = json.Unmarshal(msg, &rfqMsg)
			if err != nil {
				logger.Error("Error unmarshaling websocket message: %s", err)
				continue
			}

			switch rfqMsg.Op {
			case SubscribeOp:
				resp := c.handleSubscribe(rfqMsg.Content)
				err = c.conn.WriteJSON(resp)
				if err != nil {
					logger.Error("Error sending subscribe response: %s", err)
				}
			case UnsubscribeOp:
				resp := c.handleUnsubscribe(rfqMsg.Content)
				err = c.conn.WriteJSON(resp)
				if err != nil {
					logger.Error("Error sending unsubscribe response: %s", err)
				}
			case SendQuoteOp:
				// forward the response to the server
				var resp model.RelayerWsQuoteResponse
				err = json.Unmarshal(rfqMsg.Content, &resp)
				if err != nil {
					logger.Error("Unexpected websocket message content for send_quote", "content", rfqMsg.Content)
					continue
				}
				c.responseChan <- &resp
			case PongOp:
				lastPong = time.Now()
			default:
				logger.Errorf("Received unexpected operation from relayer: %s", rfqMsg.Op)
			}
		case <-pingTicker.C:
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
				logger.Error("Error sending ping message: %s", err)
				err = c.conn.Close()
				if err != nil {
					return fmt.Errorf("error closing websocket connection: %w", err)
				}
				close(c.doneChan)
				return fmt.Errorf("error closing websocket connection: %w", err)
			}
		}
	}
}

func (c *wsClient) handleSubscribe(content json.RawMessage) (resp model.ActiveRFQMessage) {
	var sub model.SubscriptionParams
	err := json.Unmarshal(content, &sub)
	if err != nil {
		return getErrorResponse(SubscribeOp, fmt.Errorf("could not unmarshal subscription params: %w", err))
	}
	err = c.pubsub.AddSubscription(c.relayerAddr, sub)
	if err != nil {
		return getErrorResponse(SubscribeOp, fmt.Errorf("error adding subscription: %w", err))
	}
	return getSuccessResponse(SubscribeOp)
}

func (c *wsClient) handleUnsubscribe(content json.RawMessage) (resp model.ActiveRFQMessage) {
	var sub model.SubscriptionParams
	err := json.Unmarshal(content, &sub)
	if err != nil {
		return getErrorResponse(UnsubscribeOp, fmt.Errorf("could not unmarshal subscription params: %w", err))
	}
	err = c.pubsub.RemoveSubscription(c.relayerAddr, sub)
	if err != nil {
		return getErrorResponse(UnsubscribeOp, fmt.Errorf("error removing subscription: %w", err))
	}
	return getSuccessResponse(UnsubscribeOp)
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
