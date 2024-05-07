package p2p_test

import (
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

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
)

type P2PTestSuite struct {
	*testsuite.TestSuite
	handler metrics.Handler
}

func NewP2PSuite(t *testing.T) *P2PTestSuite {
	t.Helper()

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
	var managers []p2p.LibP2PManager

	// make 10 Hosts
	for i := 0; i < 10; i++ {
		managers = append(managers, s.makeManager())
	}
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
	)

	nonce := common.BigToHash(big.NewInt(2))
	signature := []byte(gofakeit.Word())

	// m1 signs and broadcasts the signature
	m1 := managers[0]

	err := m1.PutSignature(s.GetTestContext(), chainID, nonce, signature)
	s.Require().NoError(err)
	time.Sleep(time.Second * 1)

	for {
		time.Sleep(time.Second)
		for _, manager := range managers {
			sig, err := manager.GetSignature(s.GetTestContext(), m1.Address(), chainID, nonce)
			s.Require().NoError(err)
			s.Require().Equal(signature, sig)
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

	manager, err := p2p.NewLibP2PManager(s.GetTestContext(), s.handler, signer, db, freeport.GetPort(), true)
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
