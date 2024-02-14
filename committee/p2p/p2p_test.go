package p2p_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/committee/p2p"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"testing"
	"time"
)

type P2PTestSuite struct {
	*testsuite.TestSuite
}

func NewP2PSuite(t *testing.T) *P2PTestSuite {
	return &P2PTestSuite{
		TestSuite: testsuite.NewTestSuite(t),
	}
}

func TestP2PSuite(t *testing.T) {
	suite.Run(t, NewP2PSuite(t))
}

func (s *P2PTestSuite) TestLibP2PManager() {
	m1 := s.makeManager()
	m2 := s.makeManager()
	m3 := s.makeManager()

	managers := []p2p.LibP2PManager{m1, m2, m3}
	addresses := combineHostAddresses(managers...)

	for _, manager := range managers {
		err := manager.Start(s.GetTestContext(), addresses)
		s.Require().NoError(err)
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 2)
	m1.DoSomething()
	time.Sleep(time.Second * 1)
	//m3.DoSomething()
	//m4.DoSomethingElse()
	for {
		m1.DoSomething()
		m3.DoSomething()
		if m2.DoSomethingElse() {
			break
		}
	}
}

func (s *P2PTestSuite) makeManager() p2p.LibP2PManager {
	wall, err := wallet.FromRandom()
	s.Require().NoError(err)

	signer := localsigner.NewSigner(wall.PrivateKey())

	manager, err := p2p.NewLibP2PManager(s.GetTestContext(), signer)
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
