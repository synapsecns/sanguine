package client_test

import (
	"github.com/richardwilkes/toolbox/collection"
	"github.com/synapsecns/sanguine/ethergo/client"
	"math/big"
)

func (s *TestClientSuite) TestClient() {
	chainIDs := collection.NewSet[int]()

	for _, backend := range s.testBackends {
		chainID := int(backend.GetChainID())
		// TODO: make sure chainid is correct

		// test default client, this should work
		evmClient, err := s.client.GetChainClient(s.GetTestContext(), chainID)
		s.Require().NoError(err)
		s.testBlockFetch(evmClient, false)
		s.validateChainID(evmClient, chainID)

		// test 1 conf client, this should work
		evmClient, err = s.client.GetConfirmationsClient(s.GetTestContext(), chainID, 1)
		s.Require().NoError(err)
		s.testBlockFetch(evmClient, false)
		s.validateChainID(evmClient, chainID)

		// test 2 conf client, this should break since we only have 1
		evmClient, err = s.client.GetConfirmationsClient(s.GetTestContext(), chainID, 2)
		s.Require().NoError(err)
		s.testBlockFetch(evmClient, true)

		chainIDs.Add(chainID)
	}

	res, err := s.client.GetChainIDs(s.GetTestContext())
	s.Require().NoError(err)

	s.Require().Equal(chainIDs.Len(), len(res))
	for _, chainID := range res {
		s.Require().True(chainIDs.Contains(chainID))
	}
}

func (s *TestClientSuite) validateChainID(evmClient client.EVM, expectedChainID int) {
	chainID, err := evmClient.ChainID(s.GetTestContext())
	s.Require().NoError(err)

	s.Require().Equal(expectedChainID, int(chainID.Int64()))
}

func (s *TestClientSuite) testBlockFetch(evmClient client.EVM, shouldErr bool) {
	_, err := evmClient.BlockByNumber(s.GetTestContext(), big.NewInt(1))
	if shouldErr {
		s.Require().Error(err)
	} else {
		s.Require().NoError(err)
	}
}
