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
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer/api"
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

func (s *RelayerAPISuite) TestPendingTx(t *testing.T) {
	reqChan := make(chan *api.RelayRequest, 1000)
	server := api.NewRelayerAPIServer(8080, "localhost", s.testStore, reqChan)
	ctx, cancel := context.WithCancel(s.GetTestContext())
	go server.Start(ctx)
	defer cancel()

	// store pending tx
	msg := s.mockMessage(1, 2, 1)
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
	expectedResp := api.RelayerResponse{
		Success: true,
		Result: api.MessageResult{
			OriginHash:      msg.OriginTxHash,
			DestinationHash: msg.DestTxHash,
			Origin:          msg.OriginChainID,
			Destination:     msg.DestChainID,
			RequestID:       msg.RequestID,
			State:           msg.State.String(),
		},
	}
	s.Equal(expectedResp, relayerResp)
}
