package testcontracts_test

import . "github.com/stretchr/testify/assert"

// TestTypecast tests the typecast.
func (s SimulatedSuite) TestTypecast() {
	NotPanics(s.T(), func() {
		_, bridgeConfigHandle := s.deployManager.GetBridgeConfigV3(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), bridgeConfigHandle)
		_, bridgeHandle := s.deployManager.GetTestSynapseBridge(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), bridgeHandle)
		_, swapHandle := s.deployManager.GetTestSwapFlashLoan(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), swapHandle)
		_, messageHandle := s.deployManager.GetTestMessageBusUpgradeable(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), messageHandle)
		_, metaSwapHanlde := s.deployManager.GetTestMetaSwap(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), metaSwapHanlde)
	})
}
