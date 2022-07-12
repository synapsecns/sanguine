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

func (s SimulatedSuite) TestTypecastMessageHarness() {
	NotPanics(s.T(), func() {
		_, messageHarness := s.deployManager.GetMessageHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), messageHarness)
	})
}

func (s SimulatedSuite) TestTypecastHomeHarness() {
	NotPanics(s.T(), func() {
		_, messageHarness := s.deployManager.GetHomeHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), messageHarness)
	})
}

func (s SimulatedSuite) TestTypecastUpdaterManager() {
	NotPanics(s.T(), func() {
		_, updaterManager := s.deployManager.GetUpdaterManager(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), updaterManager)
	})
}
