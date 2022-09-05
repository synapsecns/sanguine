package chain_test

import (
	"bytes"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"math/big"
	"time"
)

// TestSimulatedBlockByTime runs the test block by time method
// on a simulated backend to catch any exceptions we might encounter with a non mocked chain.
func (s ChainSuite) TestSimulatedBlockByTime() {
	simulatedBackend := simulated.NewSimulatedBackend(s.GetTestContext(), s.T())

	const mockedBlocks = 300

	// create mockedBlocks an hour apart
	for i := 1; i < mockedBlocks+1; i++ {
		simulatedBackend.EmptyBlock(time.Now().Add(-1 * time.Hour * time.Duration(mockedBlocks-i)))
	}

	// make the test chain. The simulated backend contains one, but we want to create a new one
	// when testing the chain module
	testChain, err := chain.NewFromClient(s.GetTestContext(), &client.Config{
		ChainID: int(simulatedBackend.GetChainID()),
	}, simulatedBackend.Client())
	Nil(s.T(), err)

	latestBlock, err := simulatedBackend.BlockByNumber(s.GetTestContext(), nil)
	Nil(s.T(), err)

	// test start block enabled and disabled
	for _, useStartBlock := range []bool{false, true} {
		for i := 1; i <= int(latestBlock.Number().Uint64()); i++ {
			targetBlock, err := simulatedBackend.Chain.HeaderByNumber(s.GetTestContext(), big.NewInt(int64(i)))
			Nil(s.T(), err)

			// get the target as a time
			target := s.fuzzTime(targetBlock)

			var startBlock *big.Int
			if useStartBlock && i != 1 && i != int(latestBlock.NumberU64()) {
				startBlockNumber := gofakeit.Number(1, int(targetBlock.Number.Uint64())-1)
				startBlock = big.NewInt(0).Sub(targetBlock.Number, big.NewInt(int64(startBlockNumber)))
			}

			resultingBlock, err := testChain.HeaderByTime(s.GetTestContext(), startBlock, target)
			Nil(s.T(), err)

			if targetBlock.Number.Uint64() != resultingBlock.Number.Uint64() {
				_, err := testChain.HeaderByTime(s.GetTestContext(), startBlock, target)
				Nil(s.T(), err)
			}
			Equal(s.T(), targetBlock.Number.Uint64(), resultingBlock.Number.Uint64())
			True(s.T(), bytes.Equal(resultingBlock.Hash().Bytes(), targetBlock.Hash().Bytes()))
		}
	}
}

// fuzzTime fuzzes the time up to half an hour in either direction.
// 1/3 chance it's the same.
func (s ChainSuite) fuzzTime(block *types.Header) time.Time {
	parsedTime := time.Unix(int64(block.Time), 0)
	fuzzDirection := gofakeit.Number(-1, 1)
	if fuzzDirection == 0 {
		return parsedTime
	}

	fuzzInterval := time.Duration(gofakeit.Number(1, 29)) * time.Minute
	parsedTime = parsedTime.Add(fuzzInterval * time.Duration(fuzzDirection))

	return parsedTime
}
