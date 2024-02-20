package p2p_test

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/committee/db/connect"
	"github.com/synapsecns/sanguine/committee/p2p"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"sync"
	"testing"
	"time"
)

type P2PTestSuite struct {
	*testsuite.TestSuite
	handler metrics.Handler
}

func NewP2PSuite(t *testing.T) *P2PTestSuite {
	return &P2PTestSuite{
		TestSuite: testsuite.NewTestSuite(t),
	}
}

func TestP2PSuite(t *testing.T) {
	suite.Run(t, NewP2PSuite(t))
}

func (s *P2PTestSuite) SetupSuite() {
	s.TestSuite.SetupSuite()
	s.handler = metrics.NewNullHandler()
}

func (s *P2PTestSuite) TestLibP2PManager() {
	m1 := s.makeManager()
	m2 := s.makeManager()
	m3 := s.makeManager()

	managers := []p2p.LibP2PManager{m1, m2, m3}
	peers := combineHostAddresses(managers...)
	addresses := managersToValidators(managers...)

	var wg sync.WaitGroup
	wg.Add(len(managers))

	for _, manager := range managers {
		manager := manager
		go func() {
			defer wg.Done()

			err := manager.Start(s.GetTestContext(), peers)
			s.Require().NoError(err)
			time.Sleep(time.Second)

			err = manager.AddValidators(s.GetTestContext(), addresses...)
			s.Require().NoError(err)
		}()
	}
	wg.Wait()

	time.Sleep(time.Second * 2)
	const (
		chainID = 1
		nonce   = 2
	)

	signature := []byte(gofakeit.Word())

	// m1.DoSomething()
	err := m1.PutSignature(s.GetTestContext(), chainID, nonce, signature)
	s.Require().NoError(err)

	time.Sleep(time.Second * 1)
	for {
		time.Sleep(time.Second)
		if realSig, err := m2.GetSignature(s.GetTestContext(), m1.Address(), chainID, nonce); err == nil {
			s.Require().Equal(signature, realSig)
			break
		}
	}
}

func (s *P2PTestSuite) makeManager() p2p.LibP2PManager {
	wall, err := wallet.FromRandom()
	s.Require().NoError(err)

	signer := localsigner.NewSigner(wall.PrivateKey())

	tmpDir := filet.TmpDir(s.T(), "")

	db, err := connect.Connect(s.GetTestContext(), dbcommon.Sqlite, tmpDir, s.handler)
	s.Require().NoError(err)

	manager, err := p2p.NewLibP2PManager(s.GetTestContext(), s.handler, signer, db, freeport.GetPort())
	s.Require().NoError(err)

	return manager
}

func combineHostAddresses(hostLikes ...p2p.LibP2PManager) []string {
	var addresses []string
	for _, hl := range hostLikes {
		for _, addr := range hl.Host().Addrs() {
			addresses = append(addresses, fmt.Sprintf("%s/p2p/%s", addr, hl.Host().ID()))
		}
	}
	return addresses
}

func managersToValidators(managers ...p2p.LibP2PManager) []common.Address {
	var validators []common.Address
	for _, manager := range managers {
		validators = append(validators, manager.Address())
	}
	return validators
}
