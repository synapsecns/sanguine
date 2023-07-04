package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

func getPushTx(hash string, origin uint32) (*http.Response, error) {
	// Endpoint URL
	txURL := "http://localhost:8080/push_tx"

	// URL parameters
	params := url.Values{}
	params.Set("hash", hash)
	params.Set("origin", strconv.Itoa(int(origin)))

	// Add URL parameters to the endpoint URL
	reqURL := fmt.Sprintf("%s?%s", txURL, params.Encode())

	// Send GET request
	return http.Get(reqURL)
}

func (s *RelayerAPISuite) TestPendingTx() {
	reqChan := make(chan *api.RelayRequest, 1000)
	server := api.NewRelayerAPIServer(8080, "localhost", s.testStore, reqChan)
	ctx, cancel := context.WithCancel(s.GetTestContext())
	go server.Start(ctx)
	defer cancel()

	// store pending tx
	msg := s.mockMessage(1, relayTypes.Pending)
	err := s.testStore.StoreMessage(ctx, msg)
	s.Nil(err)

	// make api request
	resp, err := getPushTx(msg.OriginTxHash, msg.OriginChainID)
	s.Nil(err)
	defer resp.Body.Close()

	// verify response
	body, err := ioutil.ReadAll(resp.Body)
	s.Nil(err)
	var relayerResp api.RelayerResponse
	err = json.Unmarshal(body, &relayerResp)
	s.Nil(err)
	s.True(relayerResp.Success)
	resultMap, ok := relayerResp.Result.(map[string]interface{})
	s.True(ok)
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
	server := api.NewRelayerAPIServer(8080, "localhost", s.testStore, reqChan)
	ctx, cancel := context.WithCancel(s.GetTestContext())
	go server.Start(ctx)
	defer cancel()

	// create mock message, but don't store it
	msg := s.mockMessage(1, relayTypes.Pending)

	// make api request
	resp, err := getPushTx(msg.OriginTxHash, msg.OriginChainID)
	s.Nil(err)
	defer resp.Body.Close()

	// verify response
	body, err := ioutil.ReadAll(resp.Body)
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
