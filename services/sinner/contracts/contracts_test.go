package contracts_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"math/big"

	"github.com/synapsecns/sanguine/services/sinner/contracts/destination"
	"github.com/synapsecns/sanguine/services/sinner/contracts/origin"

	sinnerTypes "github.com/synapsecns/sanguine/services/sinner/types"
)

func (t *ContractsSuite) TestOriginNewParser() {
	_, err := origin.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)
}

func (t *ContractsSuite) TestOriginUpdateTxMap() {
	parser, err := origin.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)

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
	parser, err := origin.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)

	err = parser.ParseAndStore(t.GetTestContext(), t.originTestLog)
	Nil(t.T(), err)
}

func (t *ContractsSuite) TestOriginParseSent() {
	parser, err := origin.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)

	parsedLog, err := parser.ParseSent(t.originTestLog)
	Nil(t.T(), err)
	Equal(t.T(), t.originTestLog.TxHash.String(), parsedLog.TxHash)
}

func (t *ContractsSuite) TestDestinationNewParser() {
	// Mock values for test
	addr := common.Address{}

	_, err := destination.NewParser(addr, t.db, t.originChainID)
	Nil(t.T(), err)
}

func (t *ContractsSuite) TestDestinationUpdateTxMap() {
	parser, err := destination.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)

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
	parser, err := destination.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)

	err = parser.ParseAndStore(t.GetTestContext(), t.desTestLog)
	Nil(t.T(), err)
}

func (t *ContractsSuite) TestDestinationParseExecuted() {
	parser, err := destination.NewParser(common.Address{}, t.db, t.originChainID)
	Nil(t.T(), err)

	parsedLog, err := parser.ParseExecuted(t.desTestLog)
	Nil(t.T(), err)
	Equal(t.T(), t.desTestLog.TxHash.String(), parsedLog.TxHash)
}
