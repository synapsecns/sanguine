package parser_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parsers"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
)

func (p *ParserSuite) TestCCTPHistorical() {
	// From: https://arbiscan.io/tx/0x8a94909499da5745f3738856fa7b1e71369ffc36c959bc7a5d6562433a0b911e
	// To: https://optimistic.etherscan.io/tx/0xf002c7773b783992f9efa067196214b280687ff61ebb35f69fb3c95283400ada
	sentTopic1 := common.HexToHash("0xc4980459837e213aedb84d9046eab1db050fec66cb9e046c4fe3b5578b01b20c")
	sentTopic2 := common.HexToHash("0x000000000000000000000000aeef513b28313c7f8e80891a7434bc2ae5104609")

	sentLog := types.Log{
		Address:     common.HexToAddress("0xfb2bfc368a7edfd51aa2cbec513ad50edea74e84"),
		BlockNumber: 147208580,
		Topics:      []common.Hash{sentTopic1, sentTopic2},
		Data:        common.FromHex("0x000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000003f46000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e58310000000000000000000000000000000000000000000000000000000129dd4cb0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000e047bc1a931e09d907eb0b8377ec9f148f0dd55698231898f644ef8188f479ac9200000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000003f46000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e58310000000000000000000000000000000000000000000000000000000129dd4cb0000000000000000000000000aeef513b28313c7f8e80891a7434bc2ae51046090000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000065501a3100000000000000000000000000000000000000000000010eb688d5b73536970c"),
		TxHash:      common.HexToHash("0x5763ef1207b1e043a94cd0cc74e98b743c7be11e826143c64eed3606c84d1222"),
		TxIndex:     1,
		BlockHash:   common.HexToHash("0x9471fc2799e60e69867e95c70edce213a1d0ed157a725948a0dcd86e4f8bf527"),
		Index:       1,
		Removed:     false,
	}

	fulfilledTopic1 := common.HexToHash("0x7864397c00beabf21ab17a04795e450354505d879a634dd2632f4fdc4b5ba04e")
	fulfilledTopic2 := common.HexToHash("0x000000000000000000000000aeef513b28313c7f8e80891a7434bc2ae5104609")

	fulfilledLog := types.Log{
		Address:     common.HexToAddress("0x5e69c336661dde70404e3345ba61f9c01ddb4c36"),
		BlockNumber: 147208580,
		Topics:      []common.Hash{fulfilledTopic1, fulfilledTopic2},
		Data:        common.FromHex("0x00000000000000000000000000000000000000000000000000000000000000030000000000000000000000000b2c639c533813f4aa9d7837caf62653d097ff8500000000000000000000000000000000000000000000000000000000001e80560000000000000000000000008c6f28f2f1a3c87f0f938b96d27520d9751ec8d900000000000000000000000000000000000000000000010efc269bbfa13434f347bc1a931e09d907eb0b8377ec9f148f0dd55698231898f644ef8188f479ac92"),
		TxHash:      common.HexToHash("0xf002c7773b783992f9efa067196214b280687ff61ebb35f69fb3c95283400ada"),
		TxIndex:     1,
		BlockHash:   common.HexToHash("0x5336bc07a348f95b4b266abf3569b6156ba4c395aa854de7798dde944cefca42"),
		Index:       1,
		Removed:     false,
	}

	sentChainID := uint32(42161)
	fulfilledChainID := uint32(10)

	err := p.eventDB.StoreBlockTime(p.GetTestContext(), sentChainID, sentLog.BlockNumber, 1)
	Nil(p.T(), err)
	err = p.eventDB.StoreBlockTime(p.GetTestContext(), fulfilledChainID, fulfilledLog.BlockNumber, 1)
	Nil(p.T(), err)
	// Parse logs
	arbCCTPParser, err := parser.NewCCTPParser(p.db, p.cctpContractArb, p.consumerFetcher, p.arbClient, p.tokenDataService, p.tokenPriceService, false)
	Nil(p.T(), err)

	opCCTPParser, err := parser.NewCCTPParser(p.db, p.cctpContractOp, p.consumerFetcher, p.opClient, p.tokenDataService, p.tokenPriceService, false)
	Nil(p.T(), err)

	parsedArb, err := arbCCTPParser.Parse(p.GetTestContext(), sentLog, sentChainID)
	Nil(p.T(), err)

	parsedOp, err := opCCTPParser.Parse(p.GetTestContext(), fulfilledLog, fulfilledChainID)
	Nil(p.T(), err)

	arbEvent, ok := parsedArb.(*sql.CCTPEvent)
	True(p.T(), ok)
	Equal(p.T(), sentChainID, arbEvent.ChainID)
	Equal(p.T(), sentLog.TxHash.String(), arbEvent.TxHash)
	Equal(p.T(), sentLog.BlockNumber, arbEvent.BlockNumber)
	// Test ERC20 retrieval
	Equal(p.T(), 6, int(*arbEvent.TokenDecimal))
	Equal(p.T(), "USDC", arbEvent.TokenSymbol)
	Equal(p.T(), "0xaf88d065e77c8cC2239327C5EDb3A432268e5831", arbEvent.Token)

	fmt.Println(arbEvent.TokenDecimal)

	opEvent, ok := parsedOp.(*sql.CCTPEvent)
	True(p.T(), ok)
	Equal(p.T(), fulfilledChainID, opEvent.ChainID)
	Equal(p.T(), fulfilledLog.TxHash.String(), opEvent.TxHash)
	Equal(p.T(), fulfilledLog.BlockNumber, opEvent.BlockNumber)
	// Test ERC20 retrieval
	Equal(p.T(), 18, int(*opEvent.TokenDecimal))
	Equal(p.T(), "sUSD", opEvent.TokenSymbol)
	Equal(p.T(), "0x8c6f28f2F1A3C87F0f938b96d27520d9751ec8d9", opEvent.Token)
}
