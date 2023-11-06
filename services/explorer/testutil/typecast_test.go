package testutil_test

import (
	"fmt"
	. "github.com/stretchr/testify/assert"
)

// TestTypecast tests the typecast.
func (s SimulatedSuite) TestTypecast() {
	NotPanics(s.T(), func() {
		_, bridgeConfigHandle := s.deployManager.GetBridgeConfigV3(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), bridgeConfigHandle)
		_, bridgeHandle := s.deployManager.GetSynapseBridge(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), bridgeHandle)
		_, swapHandle := s.deployManager.GetSwapFlashLoan(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), swapHandle)
		_, bridgeHandlev1 := s.deployManager.GetSynapseBridgeV1(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), bridgeHandlev1)
		_, messageBusHandle := s.deployManager.GetMessageBus(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), messageBusHandle)
		_, metaSwapHandle := s.deployManager.GetMetaSwap(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), metaSwapHandle)
		_, cctpHandle := s.deployManager.GetCCTP(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), cctpHandle)
		_, erc20Handle := s.deployManager.GetERC20(s.GetTestContext(), s.testBackend)
		fmt.Println("SSS", erc20Handle)
		NotNil(s.T(), erc20Handle)
	})
}
