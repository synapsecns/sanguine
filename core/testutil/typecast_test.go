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

func (s SimulatedSuite) TestTypecastNotaryManager() {
	NotPanics(s.T(), func() {
		_, updaterManager := s.deployManager.GetNotaryManager(s.GetTestContext(), s.testBackend)
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

func (s SimulatedSuite) TestTypecastHeaderHarness() {
	NotPanics(s.T(), func() {
		_, headerHarness := s.deployManager.GetHeaderHarness(s.GetTestContext(), s.testBackend)
		NotNil(s.T(), headerHarness)
	})
}
