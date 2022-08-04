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

func (s SimulatedSuite) TestTypecastAttesationCollector() {
	NotPanics(s.T(), func() {
		_, attestationCollector := s.deployManager.GetAttestationCollector(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), attestationCollector)
	})
}

func (s SimulatedSuite) TestTypecastAttestationHarness() {
	NotPanics(s.T(), func() {
		_, attestationHarness := s.deployManager.GetAttestationHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), attestationHarness)
	})
}

func (s SimulatedSuite) TestTypecastTipsHarness() {
	NotPanics(s.T(), func() {
		_, attestationHarness := s.deployManager.GetTipsHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), attestationHarness)
	})
}

func (s SimulatedSuite) TestTypecastReplicaManager() {
	NotPanics(s.T(), func() {
		_, replicaManager := s.deployManager.GetReplicaManager(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), replicaManager)
	})
}

func (s SimulatedSuite) TestTypecastReplicaManagerHarness() {
	NotPanics(s.T(), func() {
		_, replicaManagerHarness := s.deployManager.GetReplicaManagerHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), replicaManagerHarness)
	})
}
