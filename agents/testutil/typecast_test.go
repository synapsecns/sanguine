package testutil_test

import . "github.com/stretchr/testify/assert"

func (s SimulatedSuite) TestTypecastOrigin() {
	NotPanics(s.T(), func() {
		_, originHandle := s.deployManager.GetOrigin(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), originHandle)
	})
}

func (s SimulatedSuite) TestTypecastMessageHarness() {
	NotPanics(s.T(), func() {
		_, messageHarness := s.deployManager.GetMessageHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), messageHarness)
	})
}

func (s SimulatedSuite) TestTypecastOriginHarness() {
	NotPanics(s.T(), func() {
		_, messageHarness := s.deployManager.GetOriginHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), messageHarness)
	})
}

func (s SimulatedSuite) TestTypecastSummit() {
	NotPanics(s.T(), func() {
		_, summitRef := s.deployManager.GetSummit(s.GetTestContext(), s.testSynBackend)
		NotNil(s.T(), summitRef)
	})
}

func (s SimulatedSuite) TestTypecastTipsHarness() {
	NotPanics(s.T(), func() {
		_, tipsHarness := s.deployManager.GetTipsHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), tipsHarness)
	})
}

func (s SimulatedSuite) TestTypecastDestination() {
	NotPanics(s.T(), func() {
		_, destination := s.deployManager.GetDestination(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), destination)
	})
}

func (s SimulatedSuite) TestTypecastDestinationHarness() {
	NotPanics(s.T(), func() {
		_, destinationHarness := s.deployManager.GetDestinationHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), destinationHarness)
	})
}

func (s SimulatedSuite) TestTypecastHeaderHarness() {
	NotPanics(s.T(), func() {
		_, headerHarness := s.deployManager.GetHeaderHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), headerHarness)
	})
}
