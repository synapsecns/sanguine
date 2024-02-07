package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/phayes/freeport"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/cctp-relayer/api"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
)

func (s *RelayerAPISuite) mockMessage(originChainID uint32, state relayTypes.MessageState) relayTypes.Message {
	return relayTypes.Message{
		OriginTxHash:     mocks.NewMockHash(s.T()).String(),
		DestTxHash:       mocks.NewMockHash(s.T()).String(),
		OriginChainID:    originChainID,
		DestChainID:      originChainID + 1,
		Message:          []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		MessageHash:      mocks.NewMockHash(s.T()).String(),
		Attestation:      []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		RequestVersion:   0,
		FormattedRequest: []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		RequestID:        strings.TrimPrefix(mocks.NewMockHash(s.T()).String(), "0x"),
		BlockNumber:      1,
		State:            state,
	}
}

//nolint:gosec,wrapcheck
func getTx(ctx context.Context, hash string, origin uint32, port int) (*http.Response, error) {
	txURL := fmt.Sprintf("http://localhost:%d/tx", port)
	params := url.Values{}
	params.Set("hash", hash)
	params.Set("origin", strconv.Itoa(int(origin)))
	reqURL := fmt.Sprintf("%s?%s", txURL, params.Encode())
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func (s *RelayerAPISuite) TestPendingTx() {
	reqChan := make(chan *api.RelayRequest, 1000)

	port, err := freeport.GetFreePort()
	s.Nil(err)

	server := api.NewRelayerAPIServer(uint16(port), "localhost", s.testStore, reqChan)
	ctx, cancel := context.WithCancel(s.GetTestContext())
	//nolint:errcheck
	go server.Start(ctx)
	defer cancel()

	// store pending tx
	msg := s.mockMessage(1, relayTypes.Pending)
	err = s.testStore.StoreMessage(ctx, msg)
	s.Nil(err)

	// make api request
	resp, err := getTx(s.GetTestContext(), msg.OriginTxHash, msg.OriginChainID, port)
	s.Nil(err)
	defer func() {
		err := resp.Body.Close()
		s.Nil(err)
	}()
	s.Equal(resp.StatusCode, http.StatusOK)

	// verify response
	body, err := io.ReadAll(resp.Body)
	s.Nil(err)
	var relayerResp api.RelayerResponse
	err = json.Unmarshal(body, &relayerResp)
	s.Nil(err)
	s.True(relayerResp.Success)
	resultMap, ok := relayerResp.Result.(map[string]interface{})
	s.True(ok)
	//nolint:forcetypeassert
	result := api.MessageResult{
		OriginHash:      resultMap["origin_hash"].(string),
		DestinationHash: resultMap["destination_hash"].(string),
		Origin:          uint32(resultMap["origin"].(float64)),
		Destination:     uint32(resultMap["destination"].(float64)),
		RequestID:       resultMap["request_id"].(string),
		State:           resultMap["state"].(string),
	}
	expectedResult := api.MessageResult{
		OriginHash:      msg.OriginTxHash,
		DestinationHash: msg.DestTxHash,
		Origin:          msg.OriginChainID,
		Destination:     msg.DestChainID,
		RequestID:       msg.RequestID,
		State:           msg.State.String(),
	}
	s.Equal(expectedResult, result)
}

func (s *RelayerAPISuite) TestMissingTx() {
	reqChan := make(chan *api.RelayRequest, 1000)

	port, err := freeport.GetFreePort()
	s.Nil(err)

	server := api.NewRelayerAPIServer(uint16(port), "localhost", s.testStore, reqChan)
	ctx, cancel := context.WithCancel(s.GetTestContext())
	//nolint:errcheck
	go server.Start(ctx)
	defer cancel()

	// create mock message, but don't store it
	msg := s.mockMessage(1, relayTypes.Pending)

	// make api request
	resp, err := getTx(s.GetTestContext(), msg.OriginTxHash, msg.OriginChainID, port)
	s.Nil(err)
	defer func() {
		err := resp.Body.Close()
		s.Nil(err)
	}()
	s.Equal(resp.StatusCode, http.StatusOK)

	// verify response
	body, err := io.ReadAll(resp.Body)
	s.Nil(err)
	var relayerResp api.RelayerResponse
	err = json.Unmarshal(body, &relayerResp)
	s.Nil(err)
	s.True(relayerResp.Success)
	result, ok := relayerResp.Result.(string)
	s.True(ok)
	expectedResult := fmt.Sprintf("Successfully queued relay request from chain %d: %s", msg.OriginChainID, msg.OriginTxHash)
	s.Equal(expectedResult, result)

	// verify request was queued
	relayReq := <-reqChan
	expectedRelayReq := &api.RelayRequest{
		Origin: msg.OriginChainID,
		TxHash: common.HexToHash(msg.OriginTxHash),
	}
	s.Equal(expectedRelayReq, relayReq)
}

func (s *RelayerAPISuite) TestBadRequest() {
	reqChan := make(chan *api.RelayRequest, 1000)

	port, err := freeport.GetFreePort()
	s.Nil(err)

	server := api.NewRelayerAPIServer(uint16(port), "localhost", s.testStore, reqChan)
	ctx, cancel := context.WithCancel(s.GetTestContext())
	//nolint:errcheck
	go server.Start(ctx)
	defer cancel()

	// store pending tx
	msg := s.mockMessage(1, relayTypes.Pending)
	err = s.testStore.StoreMessage(ctx, msg)
	s.Nil(err)

	// make api request with no params
	txURL := fmt.Sprintf("http://localhost:%d/tx", port)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, txURL, nil)
	s.Nil(err)
	resp, err := client.Do(req)
	s.Nil(err)
	defer func() {
		err := resp.Body.Close()
		s.Nil(err)
	}()
	s.Equal(resp.StatusCode, http.StatusBadRequest)

	// verify response
	body, err := io.ReadAll(resp.Body)
	s.Nil(err)
	var relayerResp api.RelayerResponse
	err = json.Unmarshal(body, &relayerResp)
	s.Nil(err)
	s.False(relayerResp.Success)
	resultMap, ok := relayerResp.Result.(map[string]interface{})
	s.True(ok)
	//nolint:forcetypeassert
	reason := resultMap["reason"].(string)
	expectedReason := "required parameter 'origin' is missing"
	s.Equal(expectedReason, reason)
}
