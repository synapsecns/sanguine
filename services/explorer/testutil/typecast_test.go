package testutil_test

import . "github.com/stretchr/testify/assert"

func (s SimulatedSuite) TestTypecastOrigin() {
	NotPanics(s.T(), func() {
		_, originHandle := s.deployManager.GetBridgeConfigV3(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), originHandle)
	})
}
