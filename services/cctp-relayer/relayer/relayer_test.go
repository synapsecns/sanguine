package relayer_test

import (
	"context"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/cctp-relayer/attestation"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

func (s *CCTPRelayerSuite) mockMessage(originChainID, destinationChainID, blockNumber uint32) relayTypes.Message {
	return relayTypes.Message{
		OriginTxHash:     mocks.NewMockHash(s.T()).String(),
		DestTxHash:       mocks.NewMockHash(s.T()).String(),
		OriginChainID:    originChainID,
		DestChainID:      destinationChainID,
		Message:          []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		MessageHash:      mocks.NewMockHash(s.T()).String(),
		Attestation:      []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		RequestVersion:   0,
		FormattedRequest: []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		RequestID:        strings.TrimPrefix(mocks.NewMockHash(s.T()).String(), "0x"),
		BlockNumber:      uint64(blockNumber),
		State:            relayTypes.Submitted,
	}
}

func (s *CCTPRelayerSuite) TestFetchAttestation() {
	// create a new relayer
	mockAPI := attestation.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(s.testOmnirpc, s.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(s.GetTestContext(), s.GetTestConfig(), s.testStore, omniRPCClient, s.metricsHandler, mockAPI)
	s.Nil(err)

	// override mocked api call
	expectedSignature := "abc"
	mockAPI.SetGetAttestation(func(ctx context.Context, txHash string) (attestation []byte, err error) {
		return []byte(expectedSignature), nil
	})

	// fetch attestation
	testHash := "0x5dba62229dba62f233dca8f3fd14488fdc45d2a86537da2dea7a5683b5e7f622"
	msg := relayTypes.Message{
		Message:          []byte{},
		MessageHash:      testHash,
		FormattedRequest: []byte{},
	}
	completeMsg, err := relay.FetchAttestation(s.GetTestContext(), &msg)
	s.Nil(err)

	s.Equal(completeMsg.MessageHash, msg.MessageHash)
	s.Equal(completeMsg.Attestation, []byte(expectedSignature))
	s.Equal(completeMsg.State, relayTypes.Attested)

	// verify that the attested request is stored in the db
	var storedMsg relayTypes.Message
	err = s.testStore.DB().Where("origin_tx_hash = ?", completeMsg.OriginTxHash).First(&storedMsg).Error
	s.Nil(err)
	s.Equal(*completeMsg, storedMsg)
}
