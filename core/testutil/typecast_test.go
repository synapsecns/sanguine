package testutil_test

import . "github.com/stretchr/testify/assert"

func (s SimulatedSuite) TestTypecastHome() {
	NotPanics(s.T(), func() {
		_, homeHandle := s.deployManager.GetHome(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), homeHandle)
	})
}

func (s SimulatedSuite) TestTypecastXappConfig() {
	NotPanics(s.T(), func() {
		_, xappHandle := s.deployManager.GetXAppConfig(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), xappHandle)
	})
}
