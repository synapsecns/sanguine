package rest

import (
	"context"
	"encoding/json"
	"fmt"

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
	conn         *websocket.Conn
	requestChan  chan *model.RelayerWsQuoteRequest
	responseChan chan *model.RelayerWsQuoteResponse
	doneChan     chan struct{}
}

func newWsClient(conn *websocket.Conn) *wsClient {
	return &wsClient{
		conn:         conn,
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
	}
	return nil
}

func (c *wsClient) ReceiveQuoteResponse(ctx context.Context) (*model.RelayerWsQuoteResponse, error) {
	response := <-c.responseChan
	return response, nil
}

const (
	pingOp         = "ping"
	pongOp         = "pong"
	requestQuoteOp = "request_quote"
	sendQuoteOp    = "send_quote"
)

func (c *wsClient) Run(ctx context.Context) (err error) {
	for {
		select {
		case <-ctx.Done():
			c.conn.Close()
			close(c.doneChan)
			return nil
		case data := <-c.requestChan:
			msg := model.ActiveRFQMessage{
				Op:      "send_quote",
				Content: data,
			}
			c.conn.WriteJSON(msg)
		default:
			// Read message from WebSocket
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				logger.Error("Error reading websocket message: %s", err)
				continue
			}

			var rfqMsg model.ActiveRFQMessage
			err = json.Unmarshal(msg, &rfqMsg)
			if err != nil {
				logger.Error("Error unmarshalling websocket message: %s", err)
				continue
			}

			switch rfqMsg.Op {
			case sendQuoteOp:
				// forward the response to the server
				resp, ok := rfqMsg.Content.(model.RelayerWsQuoteResponse)
				if !ok {
					logger.Error("Unexpected websocket message content for send_quote", "content", rfqMsg.Content)
					continue
				}
				c.responseChan <- &resp
			case pongOp:
				// TODO: keep connection alive
			default:
				logger.Errorf("Received unexpected operation from relayer: %s", rfqMsg.Op)
			}
		}
	}
}
