package contracts_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"math/big"

	"github.com/synapsecns/sanguine/services/sinner/contracts/destination"
	"github.com/synapsecns/sanguine/services/sinner/contracts/origin"

	"github.com/synapsecns/sanguine/services/sinner/db/mocks"
	sinnerTypes "github.com/synapsecns/sanguine/services/sinner/types"
)

func (t *ContractsSuite) TestOriginNewParser() {
	// Mock values for test
	dbMock := &mocks.EventDB{}
	addr := common.Address{}

	_, err := origin.NewParser(addr, dbMock, t.originChainID)
	Nil(t.T(), err)
}

func (t *ContractsSuite) TestOriginUpdateTxMap() {
	parser, _ := origin.NewParser(common.Address{}, &mocks.EventDB{}, t.originChainID)
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

	txMap := map[string]sinnerTypes.TxSupplementalInfo{
		txHash: {
			TxHash:    txHash,
			Sender:    common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
			Timestamp: 0,
		},
	}

	parser.UpdateTxMap(txMap)
	Equal(t.T(), txMap, parser.TxMap)
}

func (t *ContractsSuite) TestOriginParseAndStore() {
	parser, _ := origin.NewParser(common.Address{}, &mocks.EventDB{}, t.originChainID)

	err := parser.ParseAndStore(t.GetTestContext(), t.originTestLog)

	Equal(t.T(), "error while parsing origin sent event. Err: could not parse sent log. err: topic/field count mismatch", err.Error()) // TODO add correct byte code to test log
}

func (t *ContractsSuite) TestOriginParseSent() {
	parser, _ := origin.NewParser(common.Address{}, &mocks.EventDB{}, t.originChainID)

	_, err := parser.ParseSent(t.originTestLog)
	Equal(t.T(), "could not parse sent log. err: topic/field count mismatch", err.Error()) // TODO add correct byte code to test log
}

func (t *ContractsSuite) TestDestinationNewParser() {
	// Mock values for test
	dbMock := &mocks.EventDB{}
	addr := common.Address{}

	_, err := destination.NewParser(addr, dbMock, t.originChainID)
	Nil(t.T(), err)
}

func (t *ContractsSuite) TestDestinationUpdateTxMap() {
	parser, _ := destination.NewParser(common.Address{}, &mocks.EventDB{}, t.originChainID)
	txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

	txMap := map[string]sinnerTypes.TxSupplementalInfo{
		txHash: {
			TxHash:    txHash,
			Sender:    common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
			Timestamp: 0,
		},
	}

	parser.UpdateTxMap(txMap)
	Equal(t.T(), txMap, parser.TxMap)
}

func (t *ContractsSuite) TestDestinationParseAndStore() {
	parser, _ := destination.NewParser(common.Address{}, &mocks.EventDB{}, t.originChainID)

	err := parser.ParseAndStore(t.GetTestContext(), t.desTestLog)
	Equal(t.T(), "error while parsing origin sent event. Err: could not parse sent log. err: topic/field count mismatch", err.Error()) // TODO add correct byte code to test log
}

func (t *ContractsSuite) TestDestinationParseExecuted() {
	parser, _ := destination.NewParser(common.Address{}, &mocks.EventDB{}, t.originChainID)

	_, err := parser.ParseExecuted(t.desTestLog)
	Equal(t.T(), "could not parse sent log. err: topic/field count mismatch", err.Error()) // TODO add correct byte code to test log
}
