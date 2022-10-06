package api_test

import (
	gosql "database/sql"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"math/big"
	"time"
)

// nolint:cyclop
func (g APISuite) TestGetCountByChainID() {
	chainID := gofakeit.Uint32()
	destinationChainIDA := gofakeit.Uint32()
	destinationChainIDB := gofakeit.Uint32()
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		var destinationChainID uint32
		if blockNumber%2 == 0 {
			destinationChainID = destinationChainIDA
		} else {
			destinationChainID = destinationChainIDB
		}
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		g.db.DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
			ChainID:            chainID,
			Recipient:          gosql.NullString{address.String(), true},
			DestinationChainID: big.NewInt(int64(destinationChainID)),
			BlockNumber:        blockNumber,
			TxHash:             txHash.String(),
			EventIndex:         gofakeit.Uint64(),
		})
		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDA, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDB, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
	}
	addressRef := address.String()
	directionRef := model.DirectionOut
	resultOut, err := g.client.GetCountByChainID(g.GetTestContext(), nil, &addressRef, &directionRef, nil)
	Nil(g.T(), err)
	// There should be 3 chains, 2 for the destination chain IDs and 1 for the source chain ID.
	Equal(g.T(), len(resultOut.Response), 3)
	// The source chain ID should have 10 events out, and the destination chain IDs should have 0 events out.
	var reached = 0
	for _, res := range resultOut.Response {
		switch *res.ChainID {
		case int(chainID):
			Equal(g.T(), *res.Count, 10)
			reached++
		case int(destinationChainIDA):
			Equal(g.T(), *res.Count, 0)
			reached++
		case int(destinationChainIDB):
			Equal(g.T(), *res.Count, 0)
			reached++
		}
	}
	Equal(g.T(), reached, 3)

	directionRef = model.DirectionIn
	resultIn, err := g.client.GetCountByChainID(g.GetTestContext(), nil, &addressRef, &directionRef, nil)
	Nil(g.T(), err)
	// Again, there should be 3 chains, 2 for the destination chain IDs and 1 for the source chain ID.
	Equal(g.T(), len(resultIn.Response), 3)
	// The source chain ID should have 0 events in, and the destination chain IDs should have 5 events in.
	reached = 0
	for _, res := range resultIn.Response {
		switch *res.ChainID {
		case int(chainID):
			Equal(g.T(), *res.Count, 0)
			reached++
		case int(destinationChainIDA):
			Equal(g.T(), *res.Count, 5)
			reached++
		case int(destinationChainIDB):
			Equal(g.T(), *res.Count, 5)
			reached++
		}
	}
	Equal(g.T(), reached, 3)
}

func (g APISuite) TestGetCountByTokenAddress() {
	chainID := gofakeit.Uint32()
	destinationChainID := gofakeit.Uint32()
	tokenAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	tokenAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		var tokenAddress common.Address
		if blockNumber%2 == 0 {
			tokenAddress = tokenAddressA
		} else {
			tokenAddress = tokenAddressB
		}
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		g.db.DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
			ChainID:            chainID,
			Recipient:          gosql.NullString{address.String(), true},
			DestinationChainID: big.NewInt(int64(destinationChainID)),
			Token:              tokenAddress.String(),
			BlockNumber:        blockNumber,
			TxHash:             txHash.String(),
			EventIndex:         gofakeit.Uint64(),
		})
		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
	}
	addressRef := address.String()
	directionRef := model.DirectionOut
	resultOut, err := g.client.GetCountByTokenAddress(g.GetTestContext(), nil, &addressRef, &directionRef, nil)
	Nil(g.T(), err)
	// There should be 4 results, two for each token on two chain. Each on the source chain ID should have 5 events,
	// while each on the destination chain ID should have 0 events.
	Equal(g.T(), len(resultOut.Response), 4)
	reached := 0
	for _, res := range resultOut.Response {
		if *res.ChainID == int(chainID) {
			if *res.TokenAddress == tokenAddressA.String() {
				Equal(g.T(), *res.Count, 5)
				reached++
			}
			if *res.TokenAddress == tokenAddressB.String() {
				Equal(g.T(), *res.Count, 5)
				reached++
			}
		}
		if *res.ChainID == int(destinationChainID) {
			if *res.TokenAddress == tokenAddressA.String() {
				Equal(g.T(), *res.Count, 0)
				reached++
			}
			if *res.TokenAddress == tokenAddressB.String() {
				Equal(g.T(), *res.Count, 0)
				reached++
			}
		}
	}
	Equal(g.T(), reached, 4)

	directionRef = model.DirectionIn
	resultIn, err := g.client.GetCountByTokenAddress(g.GetTestContext(), nil, nil, &directionRef, nil)
	Nil(g.T(), err)
	// Again, there should be 4 results, two for each token on two chain. Each on the source chain ID should have 0 events,
	// while each on the destination chain ID should have 5 events.
	Equal(g.T(), len(resultIn.Response), 4)
	reached = 0
	for _, res := range resultIn.Response {
		if *res.ChainID == int(destinationChainID) {
			if *res.TokenAddress == tokenAddressA.String() {
				Equal(g.T(), *res.Count, 5)
				reached++
			}
			if *res.TokenAddress == tokenAddressB.String() {
				Equal(g.T(), *res.Count, 5)
				reached++
			}
		}
		if *res.ChainID == int(chainID) {
			if *res.TokenAddress == tokenAddressA.String() {
				Equal(g.T(), *res.Count, 0)
				reached++
			}
			if *res.TokenAddress == tokenAddressB.String() {
				Equal(g.T(), *res.Count, 0)
				reached++
			}
		}
	}
	Equal(g.T(), reached, 4)
}

func (g APISuite) TestGetBridgeTransactions() {
	chainID := gofakeit.Uint32()
	destinationChainID := gofakeit.Uint32()
	tokenAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	senderAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	senderString := senderAddress.String()
	txHashA := common.BigToHash(big.NewInt(gofakeit.Int64()))
	txHashB := common.BigToHash(big.NewInt(gofakeit.Int64()))
	kappaString := crypto.Keccak256Hash(txHashA.Bytes()).String()
	txHashString := txHashA.String()
	page := 1

	g.db.DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
		ChainID:            chainID,
		Sender:             senderString,
		Recipient:          gosql.NullString{address.String(), true},
		DestinationChainID: big.NewInt(int64(destinationChainID)),
		Token:              tokenAddress.String(),
		BlockNumber:        1,
		TxHash:             txHashA.String(),
		DestinationKappa:   kappaString,
		EventIndex:         gofakeit.Uint64(),
	})
	g.db.DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
		ChainID:            destinationChainID,
		Recipient:          gosql.NullString{address.String(), true},
		DestinationChainID: big.NewInt(int64(chainID)),
		Token:              tokenAddress.String(),
		BlockNumber:        1,
		TxHash:             txHashB.String(),
		Kappa:              gosql.NullString{kappaString, true},
		SwapSuccess:        big.NewInt(1),
		EventIndex:         gofakeit.Uint64(),
	})
	timestamp := uint64(time.Now().Unix())
	err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, 1, timestamp)
	Nil(g.T(), err)
	err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainID, 1, timestamp)
	Nil(g.T(), err)

	originRes, err := g.client.GetBridgeTransactions(g.GetTestContext(), nil, nil, &txHashString, nil, nil, &page, nil)
	Nil(g.T(), err)
	Equal(g.T(), len(originRes.Response), 1)
	originResOne := *originRes.Response[0]
	Equal(g.T(), *originResOne.Kappa, kappaString)
	// do pending
	Equal(g.T(), *originResOne.SwapSuccess, true)

	fromInfo := *originResOne.FromInfo
	Equal(g.T(), *fromInfo.ChainID, int(chainID))
	Equal(g.T(), *fromInfo.Address, address.String())
	Equal(g.T(), *fromInfo.TxnHash, txHashA.String())
	// do value
	// do formatted value
	// do usd value
	Equal(g.T(), *fromInfo.TokenAddress, tokenAddress.String())
	// do token symbol
	Equal(g.T(), *fromInfo.BlockNumber, 1)
	Equal(g.T(), *fromInfo.Time, int(timestamp))

	toInfo := *originResOne.ToInfo
	Equal(g.T(), *toInfo.ChainID, int(destinationChainID))
	Equal(g.T(), *toInfo.Address, address.String())
	Equal(g.T(), *toInfo.TxnHash, txHashB.String())
	// do value
	// do formatted value
	// do usd value
	Equal(g.T(), *toInfo.TokenAddress, tokenAddress.String())
	// do token symbol
	Equal(g.T(), *toInfo.BlockNumber, 1)
	Equal(g.T(), *toInfo.Time, int(timestamp))

	destinationRes, err := g.client.GetBridgeTransactions(g.GetTestContext(), nil, nil, nil, &kappaString, nil, &page, nil)
	Nil(g.T(), err)
	Equal(g.T(), len(destinationRes.Response), 1)
	destinationResOne := *destinationRes.Response[0]
	Equal(g.T(), originResOne, destinationResOne)

	addressRes, err := g.client.GetBridgeTransactions(g.GetTestContext(), nil, &senderString, nil, nil, nil, &page, nil)
	Nil(g.T(), err)
	Equal(g.T(), len(addressRes.Response), 1)
	addressResOne := *addressRes.Response[0]
	Equal(g.T(), originResOne, addressResOne)
}

// nolint:cyclop
func (g APISuite) TestAddressRanking() {
	var chainID uint32
	chainIDs := []uint32{gofakeit.Uint32(), gofakeit.Uint32(), gofakeit.Uint32()}
	destinationChainIDA := gofakeit.Uint32()
	destinationChainIDB := gofakeit.Uint32()
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))

	// used for validation later
	var addressesTried = make(map[string]int)

	// this counter lets us have a random variation in address occurrence
	resetTokenAddrCounter := gofakeit.Number(1, 3)
	// random token addr
	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	// for holding the current token addr in line the gofakeit.Bool() decides to pass true
	lastTokenAddr := tokenAddr
	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		var destinationChainID uint32
		if blockNumber%2 == 0 {
			destinationChainID = destinationChainIDA
		} else {
			destinationChainID = destinationChainIDB
		}

		// if the token counter is zero reset it
		if resetTokenAddrCounter == 0 {
			tokenAddr = common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
			lastTokenAddr = tokenAddr
			resetTokenAddrCounter = gofakeit.Number(1, 3)
		} else {
			// before using the current token addr, let throw in some randomness
			if gofakeit.Bool() {
				tokenAddr = common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
			} else {
				resetTokenAddrCounter--
			}
		}
		// change up chainID (1/3 chance of using a new chain)
		chainID = chainIDs[gofakeit.Number(0, 2)]
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		g.db.DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
			ChainID:            chainID,
			Recipient:          gosql.NullString{address.String(), true},
			DestinationChainID: big.NewInt(int64(destinationChainID)),
			BlockNumber:        blockNumber,
			TxHash:             txHash.String(),
			EventIndex:         gofakeit.Uint64(),
			Token:              tokenAddr,
		})

		// add the tokenAddr inserted to the test map (for validation later)
		addressesTried[tokenAddr]++

		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDA, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDB, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)

		// if a random address was inserted, revert to address corresponding to resetTokenAddrCounter
		if lastTokenAddr != tokenAddr {
			tokenAddr = lastTokenAddr
		}
	}
	result, err := g.client.GetAddressRanking(g.GetTestContext(), nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)

	// check if the length of the response is same to the number of unique addresses inserted into test db
	Equal(g.T(), len(result.Response), len(addressesTried))

	// Validate contents of response by comparing to addressesTried
	for k, v := range addressesTried {
		for _, res := range result.Response {
			if *res.Address == k {
				Equal(g.T(), v, *res.Count)
			}
		}
	}
}
