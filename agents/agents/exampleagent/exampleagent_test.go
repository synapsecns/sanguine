package exampleagent_test

import (
	. "github.com/stretchr/testify/assert"
)

func (u ExampleAgentSuite) TestExampleAgentSimulatedTestSuite() {
	NotNil(u.T(), u.SimulatedBackendsTestSuite)
}
