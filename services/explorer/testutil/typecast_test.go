package testutil_test

import . "github.com/stretchr/testify/assert"

func (s SimulatedSuite) TestTypecastOrigin() {
	NotPanics(s.T(), func() {
		_, bridgeConfigHandle := s.deployManager.GetBridgeConfigV3(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), bridgeConfigHandle)
		_, bridgeHandle := s.deployManager.GetSynapseBridge(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), bridgeHandle)
		_, swapHandle := s.deployManager.GetSwapFlashLoan(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), swapHandle)
	})
}
