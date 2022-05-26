package evm_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"testing"
)

func (e *EVMSuite) TestGetters(t *testing.T) {
	name := "hi"
	testCfg := config.DomainConfig{
		DomainID: 1,
		RPCUrl:   e.testBackend.RPCAddress(),
	}

	testEvm, err := evm.NewEVM(e.GetTestContext(), name, testCfg)
	Nil(e.T(), err)
	Equal(t, testEvm.Config(), testCfg)
	Equal(t, testEvm.Name(), name)
}
