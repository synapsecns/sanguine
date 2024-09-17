package rest

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// WsClient is a client for the WebSocket API.
//
//go:generate go run github.com/vektra/mockery/v2 --name WsClient --output ./mocks --case=underscore
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
		requestChan:  make(chan *model.RelayerWsQuoteRequest, 1),
		responseChan: make(chan *model.RelayerWsQuoteResponse, 1),
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

func (c *wsClient) ReceiveQuoteResponse(ctx context.Context) (resp *model.RelayerWsQuoteResponse, err error) {
	for {
		select {
		case resp = <-c.responseChan:
			// successfuly received
			return resp, nil
		case <-c.doneChan:
			return nil, fmt.Errorf("websocket client is closed")
		case <-ctx.Done():
			return nil, fmt.Errorf("expiration reached without response")
		}
	}
}

const (
	pingOp         = "ping"
	pongOp         = "pong"
	requestQuoteOp = "request_quote"
	sendQuoteOp    = "send_quote"
)

func (c *wsClient) Run(ctx context.Context) (err error) {
	messageChan := make(chan []byte)

	// Goroutine to read messages from WebSocket and send to channel
	go func() {
		defer close(messageChan)
		for {
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				logger.Error("Error reading websocket message: %s", err)
				continue
			}
			messageChan <- msg
		}
	}()

	for {
		select {
		case <-ctx.Done():
			c.conn.Close()
			close(c.doneChan)
			return nil
		case data := <-c.requestChan:
			rawData, err := json.Marshal(data)
			if err != nil {
				logger.Error("Error marshalling quote request: %s", err)
				continue
			}
			msg := model.ActiveRFQMessage{
				Op:      requestQuoteOp,
				Content: json.RawMessage(rawData),
			}
			c.conn.WriteJSON(msg)
		case msg := <-messageChan:
			var rfqMsg model.ActiveRFQMessage
			err = json.Unmarshal(msg, &rfqMsg)
			if err != nil {
				logger.Error("Error unmarshalling websocket message: %s", err)
				continue
			}

			switch rfqMsg.Op {
			case sendQuoteOp:
				// forward the response to the server
				var resp model.RelayerWsQuoteResponse
				err = json.Unmarshal(rfqMsg.Content, &resp)
				if err != nil {
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
