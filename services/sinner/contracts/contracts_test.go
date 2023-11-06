package contracts_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/sinner/contracts/destination"
	"github.com/synapsecns/sanguine/services/sinner/contracts/origin"

	sinnerTypes "github.com/synapsecns/sanguine/services/sinner/types"
)

func (t *ContractsSuite) TestOriginNewParser() {
	_, err := origin.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)
}

func (t *ContractsSuite) TestOriginParseAndStore() {
	parser, err := origin.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)

	err = parser.ParseAndStore(t.GetTestContext(), t.originTestLog, sinnerTypes.TxSupplementalInfo{})
	Nil(t.T(), err)
}

func (t *ContractsSuite) TestDestinationNewParser() {
	// Mock values for test
	addr := common.Address{}

	_, err := destination.NewParser(addr, t.db, t.originChainID)
	Nil(t.T(), err)
}

func (t *ContractsSuite) TestDestinationParseAndStore() {
	parser, err := destination.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)

	err = parser.ParseAndStore(t.GetTestContext(), t.desTestLog, sinnerTypes.TxSupplementalInfo{})
	Nil(t.T(), err)
}
