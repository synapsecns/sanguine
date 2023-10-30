package api_test

import (
	gosql "database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"math"
	"math/big"
	"sort"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
)

//nolint:cyclop
func (g APISuite) TestAddressRanking() {
	var chainID uint32
	chainIDs := []uint32{g.chainIDs[0], g.chainIDs[1], g.chainIDs[2]}
	destinationChainIDA := g.chainIDs[3]
	destinationChainIDB := g.chainIDs[4]
	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()

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

		currentTime := uint64(time.Now().Unix())

		// change up chainID (1/3 chance of using a new chain)
		chainID = chainIDs[gofakeit.Number(0, 2)]
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
			InsertTime:         1,
			ChainID:            chainID,
			Recipient:          gosql.NullString{String: address.String(), Valid: true},
			DestinationChainID: big.NewInt(int64(destinationChainID)),
			BlockNumber:        blockNumber,
			TxHash:             txHash.String(),
			EventIndex:         gofakeit.Uint64(),
			Token:              tokenAddr,
			Sender:             tokenAddr,
			TimeStamp:          &currentTime,
		})
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
			ChainID:         chainID,
			TokenAddress:    tokenAddr,
			ContractAddress: contractAddress,
			TokenIndex:      1,
		})

		// add the tokenAddr inserted to the test map (for validation later)
		addressesTried[tokenAddr]++

		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainIDs[0], blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), chainIDs[1], blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), chainIDs[2], blockNumber, uint64(time.Now().Unix())*blockNumber)
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

	blockNumberInit := uint64(10)
	err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumberInit, uint64(time.Now().Unix())*blockNumberInit)
	Nil(g.T(), err)

	result, err := g.client.GetAddressRanking(g.GetTestContext(), nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	// check if the length of the response is same to the number of unique addresses inserted into test db
	Equal(g.T(), len(addressesTried), len(result.Response))

	// Validate contents of response by comparing to addressesTried
	for k, v := range addressesTried {
		for _, res := range result.Response {
			if *res.Address == k {
				Equal(g.T(), v, *res.Count)
			}
		}
	}
}

//nolint:cyclop
func (g APISuite) TestGetCountByChainID() {
	chainID := g.chainIDs[0]
	chainID2 := g.chainIDs[1]
	chainID3 := g.chainIDs[2]
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	tokenAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		var destinationChainID int64
		var inputChain uint32
		destinationChainID = int64(g.chainIDs[1])
		inputChain = chainID
		if blockNumber > 1 {
			if blockNumber%2 == 0 {
				inputChain = chainID2
				destinationChainID = 0
			} else {
				inputChain = chainID3
				destinationChainID = int64(g.chainIDs[0])
			}
		}

		currentTime := uint64(time.Now().Unix())
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
			ChainID:            inputChain,
			EventType:          gofakeit.Uint8(),
			DestinationChainID: big.NewInt(destinationChainID),
			Recipient:          gosql.NullString{String: address.String(), Valid: true},
			BlockNumber:        blockNumber,
			TxHash:             txHash.String(),
			EventIndex:         gofakeit.Uint64(),
			TimeStamp:          &currentTime,
			ContractAddress:    contractAddress,
			Token:              tokenAddress,
		})
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
			ChainID:         chainID,
			TokenAddress:    tokenAddress,
			ContractAddress: contractAddress,
			TokenIndex:      1,
		})

		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), chainID2, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), chainID3, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
	}

	addressRef := address.String()
	directionRef := model.DirectionOut
	resultOut, err := g.client.GetCountByChainID(g.GetTestContext(), nil, &addressRef, &directionRef, nil)
	Nil(g.T(), err)
	// There should be 3 chains, 2 for the destination chain IDs and 1 for the source chain ID.
	Equal(g.T(), 1, len(resultOut.Response))
	// The source chain ID should have 10 events out, and the destination chain IDs should have 0 events out.
	var reached = 0
	for _, res := range resultOut.Response {
		switch *res.ChainID {
		case int(chainID):
			Equal(g.T(), 1, *res.Count)
			reached++
		case int(chainID2):
			Equal(g.T(), 5, *res.Count)
			reached++
		case int(chainID3):
			Equal(g.T(), 4, *res.Count)
			reached++
		}
	}
	Equal(g.T(), 1, reached)

	directionRef = model.DirectionIn
	resultIn, err := g.client.GetCountByChainID(g.GetTestContext(), nil, &addressRef, &directionRef, nil)
	Nil(g.T(), err)
	// Again, there should be 3 chains, 2 for the destination chain IDs and 1 for the source chain ID.
	Equal(g.T(), 2, len(resultIn.Response))
	// The source chain ID should have 0 events in, and the destination chain IDs should have 5 events in.
	reached = 0
	for _, res := range resultIn.Response {
		switch *res.ChainID {
		case int(chainID):
			Equal(g.T(), 1, *res.Count)
			reached++
		case int(chainID2):
			Equal(g.T(), 5, *res.Count)
			reached++
		case int(chainID3):
			Equal(g.T(), 4, *res.Count)
			reached++
		}
	}
	Equal(g.T(), 2, reached)
}

// nolint (needed for testing all possibilities)
func (g APISuite) TestGetCountByTokenAddress() {
	chainID := g.chainIDs[0]
	destinationChainID := g.chainIDs[1]
	tokenAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	tokenAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		var tokenAddress common.Address
		if blockNumber%2 == 0 {
			tokenAddress = tokenAddressA
			destinationChainID = g.chainIDs[1]
		} else {
			tokenAddress = tokenAddressB
			destinationChainID = 0
		}
		currentTime := uint64(time.Now().Unix())
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
			ChainID:            chainID,
			EventType:          gofakeit.Uint8(),
			Recipient:          gosql.NullString{String: address.String(), Valid: true},
			DestinationChainID: big.NewInt(int64(destinationChainID)),
			Token:              tokenAddress.String(),
			BlockNumber:        blockNumber,
			TxHash:             txHash.String(),
			EventIndex:         gofakeit.Uint64(),
			TimeStamp:          &currentTime,
			ContractAddress:    contractAddress,
		})
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
			ChainID:         chainID,
			TokenAddress:    tokenAddress.String(),
			ContractAddress: contractAddress,
			TokenIndex:      1,
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

	Equal(g.T(), 1, len(resultOut.Response))
	reached := 0
	for _, res := range resultOut.Response {
		if *res.ChainID == int(chainID) {
			if *res.TokenAddress == tokenAddressA.String() {
				Equal(g.T(), 5, *res.Count)
				reached++
			}
			if *res.TokenAddress == tokenAddressB.String() {
				Equal(g.T(), 5, *res.Count)
				reached++
			}
		}
		if *res.ChainID == int(destinationChainID) {
			if *res.TokenAddress == tokenAddressA.String() {
				Equal(g.T(), 5, *res.Count)
				reached++
			}
			if *res.TokenAddress == tokenAddressB.String() {
				Equal(g.T(), 5, *res.Count)
				reached++
			}
		}
	}
	Equal(g.T(), 1, reached)

	directionRef = model.DirectionIn
	resultIn, err := g.client.GetCountByTokenAddress(g.GetTestContext(), nil, nil, &directionRef, nil)
	Nil(g.T(), err)

	Equal(g.T(), 1, len(resultIn.Response))
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
				Equal(g.T(), 5, *res.Count)
				reached++
			}
			if *res.TokenAddress == tokenAddressB.String() {
				Equal(g.T(), 5, *res.Count)
				reached++
			}
		}
	}
	Equal(g.T(), 1, reached)
}

// TODO add other platforms to make this test more exhaustive
// nolint:cyclop
func (g APISuite) TestDailyStatisticsByChain() {
	chainID := g.chainIDs[0]
	destinationChainIDA := g.chainIDs[1]
	destinationChainIDB := g.chainIDs[2]
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	nowTime := time.Now().Unix()
	senders := []string{common.BigToHash(big.NewInt(gofakeit.Int64())).String(), common.BigToHash(big.NewInt(gofakeit.Int64())).String(), common.BigToHash(big.NewInt(gofakeit.Int64())).String()}
	cumulativePrice := []float64{}
	contract := common.BigToHash(big.NewInt(gofakeit.Int64()))
	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		var destinationChainID uint32
		if blockNumber%2 == 0 {
			destinationChainID = destinationChainIDA
		} else {
			destinationChainID = destinationChainIDB
		}
		price := float64(gofakeit.Number(1, 300))
		cumulativePrice = append(cumulativePrice, price)
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))

		timestamp := uint64(nowTime) - (10*blockNumber)*86400
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
			ChainID:            chainID,
			ContractAddress:    contract.String(),
			EventType:          gofakeit.Uint8(),
			Recipient:          gosql.NullString{String: address.String(), Valid: true},
			DestinationChainID: big.NewInt(int64(destinationChainID)),
			BlockNumber:        blockNumber,
			TxHash:             txHash.String(),
			EventIndex:         gofakeit.Uint64(),
			Token:              tokenAddr,
			Amount:             big.NewInt(int64(gofakeit.Number(1, 300))),
			AmountUSD:          &price,
			Sender:             senders[blockNumber%3],
			TimeStamp:          &timestamp,
		})
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
			ChainID:      chainID,
			TokenAddress: tokenAddr,
			TokenIndex:   1,
		})
		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDA, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDB, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
	}
	total := 0.0
	for _, v := range cumulativePrice {
		total += v
	}
	platform := model.PlatformBridge
	days := model.DurationAllTime
	typeArg := model.DailyStatisticTypeVolume
	result, err := g.client.GetDailyStatisticsByChain(g.GetTestContext(), nil, &typeArg, &days, &platform, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), cumulativePrice[len(cumulativePrice)-1], *result.Response[0].Total)
	Equal(g.T(), len(cumulativePrice), len(result.Response))

	typeArg = model.DailyStatisticTypeAddresses
	result, err = g.client.GetDailyStatisticsByChain(g.GetTestContext(), nil, &typeArg, &days, &platform, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), float64(1), *result.Response[0].Total)
	Equal(g.T(), len(cumulativePrice), len(result.Response))

	typeArg = model.DailyStatisticTypeTransactions
	result, err = g.client.GetDailyStatisticsByChain(g.GetTestContext(), nil, &typeArg, &days, &platform, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), float64(1), *result.Response[0].Total)
	Equal(g.T(), len(cumulativePrice), len(result.Response))
}

// TODO add swap txs.
func (g APISuite) TestGetBridgeTransactions() {
	chainID := g.chainIDs[0]
	destinationChainID := g.chainIDs[1]
	contractAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	tokenAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	senderAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	senderString := senderAddress.String()
	txHashA := common.BigToHash(big.NewInt(gofakeit.Int64()))
	txHashB := common.BigToHash(big.NewInt(gofakeit.Int64()))
	kappaString := crypto.Keccak256Hash(txHashA.Bytes()).String()
	txHashString := txHashA.String()
	amount := big.NewInt(int64(gofakeit.Uint64()))
	amountUSD := float64(gofakeit.Number(1, 300))
	tokenDecimals := uint8(gofakeit.Number(0, 3))
	tokenSymbol := gofakeit.Word()
	timestamp := uint64(time.Now().Unix())
	page := 1

	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
		InsertTime:         1,
		ContractAddress:    common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		ChainID:            chainID,
		EventType:          gofakeit.Uint8(),
		Sender:             senderString,
		Recipient:          gosql.NullString{String: address.String(), Valid: true},
		DestinationChainID: big.NewInt(int64(destinationChainID)),
		Token:              tokenAddress,
		BlockNumber:        1,
		TxHash:             txHashA.String(),
		DestinationKappa:   kappaString,
		EventIndex:         gofakeit.Uint64(),
		Amount:             amount,
		AmountUSD:          &amountUSD,
		TokenDecimal:       &tokenDecimals,
		TokenSymbol:        gosql.NullString{String: tokenSymbol, Valid: true},
		TimeStamp:          &timestamp,
	})
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
		ChainID:         chainID,
		TokenAddress:    tokenAddress,
		TokenIndex:      1,
		ContractAddress: contractAddr,
	})
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.BridgeEvent{
		InsertTime:      1,
		ChainID:         destinationChainID,
		EventType:       gofakeit.Uint8(),
		Recipient:       gosql.NullString{String: address.String(), Valid: true},
		Token:           tokenAddress,
		BlockNumber:     1,
		TxHash:          txHashB.String(),
		Kappa:           gosql.NullString{String: kappaString, Valid: true},
		SwapSuccess:     big.NewInt(1),
		EventIndex:      gofakeit.Uint64(),
		Amount:          amount,
		AmountUSD:       &amountUSD,
		TokenDecimal:    &tokenDecimals,
		Sender:          gofakeit.Word(),
		TokenSymbol:     gosql.NullString{String: tokenSymbol, Valid: true},
		TimeStamp:       &timestamp,
		ContractAddress: contractAddr,
	})
	g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
		ChainID:         destinationChainID,
		TokenAddress:    tokenAddress,
		ContractAddress: contractAddr,
		TokenIndex:      1,
	})
	err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, 1, timestamp)
	Nil(g.T(), err)
	err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainID, 1, timestamp)
	Nil(g.T(), err)
	pending := false
	//nolint:dupword
	originRes, err := g.client.GetBridgeTransactions(g.GetTestContext(), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &txHashString, nil, &pending, &page, nil, nil, nil)

	Nil(g.T(), err)
	Equal(g.T(), 1, len(originRes.Response))
	originResOne := *originRes.Response[0]
	Equal(g.T(), kappaString, *originResOne.Kappa)
	// do pending
	Equal(g.T(), *originResOne.SwapSuccess, true)

	fromInfo := *originResOne.FromInfo
	Equal(g.T(), int(chainID), *fromInfo.ChainID)
	Equal(g.T(), address.String(), *fromInfo.Address)
	Equal(g.T(), txHashA.String(), *fromInfo.TxnHash)
	Equal(g.T(), amount.String(), *fromInfo.Value)
	Equal(g.T(), amountUSD, *fromInfo.USDValue)
	formattedValue := uint64((float64(amount.Int64()) / math.Pow10(int(tokenDecimals))) * 1000000)
	Equal(g.T(), formattedValue, uint64(*fromInfo.FormattedValue*1000000))
	Equal(g.T(), tokenSymbol, *fromInfo.TokenSymbol)
	Equal(g.T(), tokenAddress, *fromInfo.TokenAddress)
	Equal(g.T(), 1, *fromInfo.BlockNumber)
	Equal(g.T(), int(timestamp), *fromInfo.Time)

	toInfo := *originResOne.ToInfo
	Equal(g.T(), int(destinationChainID), *toInfo.ChainID)
	Equal(g.T(), address.String(), *toInfo.Address)
	Equal(g.T(), txHashB.String(), *toInfo.TxnHash)
	Equal(g.T(), amount.String(), *toInfo.Value)
	Equal(g.T(), amountUSD, *toInfo.USDValue)
	Equal(g.T(), formattedValue, uint64(*toInfo.FormattedValue*1000000))
	Equal(g.T(), tokenSymbol, *toInfo.TokenSymbol)
	Equal(g.T(), tokenAddress, *toInfo.TokenAddress)
	Equal(g.T(), 1, *toInfo.BlockNumber)
	Equal(g.T(), int(timestamp), *toInfo.Time)

	pending = false
	//nolint:dupword
	destinationRes, err := g.client.GetBridgeTransactions(g.GetTestContext(), nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &kappaString, &pending, &page, nil, nil, nil)
	Nil(g.T(), err)
	Equal(g.T(), 1, len(destinationRes.Response))
	destinationResOne := *destinationRes.Response[0]
	Equal(g.T(), originResOne, destinationResOne)

	pending = true
	addressRes, err := g.client.GetBridgeTransactions(g.GetTestContext(), nil, nil, nil, &senderString, nil, nil, nil, nil, nil, nil, nil, nil, &pending, &page, nil, nil, nil)
	Nil(g.T(), err)
	Equal(g.T(), 1, len(addressRes.Response))

	addressResOne := *addressRes.Response[0]
	Equal(g.T(), originResOne, addressResOne)
}

func (g APISuite) TestLeaderboard() {
	chainID := g.chainIDs[0]
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	var addressNS gosql.NullString
	addressNS.String = address.String()
	addressNS.Valid = true

	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	senders := []string{common.BigToHash(big.NewInt(gofakeit.Int64())).String(), common.BigToHash(big.NewInt(gofakeit.Int64())).String(), common.BigToHash(big.NewInt(gofakeit.Int64())).String()}
	nowTime := time.Now().Unix()
	contract := common.BigToHash(big.NewInt(gofakeit.Int64()))
	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		price := float64(gofakeit.Number(1, 300))
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))

		timestamp := uint64(nowTime) - (10*blockNumber)*86400
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&MvBridgeEvent{
			FChainID:         chainID,
			FContractAddress: contract.String(),
			FEventType:       gofakeit.Uint8(),
			FBlockNumber:     blockNumber,
			FTxHash:          txHash.String(),
			FEventIndex:      gofakeit.Uint64(),
			FAmountUSD:       &price,
			FFeeAmountUSD:    &price,
			FSender:          senders[blockNumber%3],
			FTimeStamp:       &timestamp,
			TChainID:         chainID,
			TContractAddress: contract.String(),
			TEventType:       gofakeit.Uint8(),
			TBlockNumber:     blockNumber,
			TTxHash:          txHash.String(),
			TEventIndex:      gofakeit.Uint64(),
			TAmountUSD:       &price,
			TFeeAmountUSD:    &price,
			TSender:          senders[blockNumber%3],
			TTimeStamp:       &timestamp,
		})
		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
			ChainID:      chainID,
			TokenAddress: tokenAddr,
			TokenIndex:   1,
		})
		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
	}

	useMv := true
	page := 1
	duration := model.DurationAllTime
	result, err := g.client.GetLeaderboard(g.GetTestContext(), &duration, nil, &useMv, &page)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	for i := 0; i < len(result.Response); i++ {
		NotNil(g.T(), result.Response[i].Address)
		NotNil(g.T(), result.Response[i].VolumeUsd)
		NotNil(g.T(), result.Response[i].Fees)
		NotNil(g.T(), result.Response[i].Txs)
		NotNil(g.T(), result.Response[i].Rank)
		NotNil(g.T(), result.Response[i].AvgVolumeUsd)
	}
}

// TODO rewrite this test so that it is exhaustive with all platform and statistic types.
// nolint:cyclop
func (g APISuite) TestAmountStatistic() {
	chainID := g.chainIDs[0]
	destinationChainIDA := g.chainIDs[1]
	destinationChainIDB := g.chainIDs[2]
	address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()

	tokenAddr := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	sender := common.BigToAddress(big.NewInt(gofakeit.Int64())).String()
	cumulativePrice := []float64{}
	// Generate bridge events for different chain IDs.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		var destinationChainID uint32
		if blockNumber%2 == 0 {
			destinationChainID = destinationChainIDA
		} else {
			destinationChainID = destinationChainIDB
		}

		currentTime := uint64(time.Now().Unix())
		price := float64(gofakeit.Number(1, 300))
		cumulativePrice = append(cumulativePrice, price)
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))

		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&MvBridgeEvent{
			InsertTime:       1,
			FChainID:         chainID,
			FContractAddress: contractAddress,
			FEventType:       gofakeit.Uint8(),
			FBlockNumber:     blockNumber,
			FTxHash:          txHash.String(),
			FEventIndex:      gofakeit.Uint64(),
			FAmountUSD:       &price,
			FFeeAmountUSD:    &price,
			FRecipient:       gosql.NullString{String: address.String(), Valid: true},
			FSender:          sender,
			FTimeStamp:       &currentTime,
			TChainID:         destinationChainID,
			TContractAddress: contractAddress,
			TEventType:       gofakeit.Uint8(),
			TBlockNumber:     blockNumber,
			TTxHash:          txHash.String(),
			TEventIndex:      gofakeit.Uint64(),
			TAmountUSD:       &price,
			TFeeAmountUSD:    &price,
			TSender:          sender,
			TTimeStamp:       &currentTime,
		})

		g.db.UNSAFE_DB().WithContext(g.GetTestContext()).Create(&sql.TokenIndex{
			ChainID:         chainID,
			TokenAddress:    tokenAddr,
			ContractAddress: contractAddress,
			TokenIndex:      1,
		})
		// Set all times after current time, so we can get the events.
		err := g.eventDB.StoreBlockTime(g.GetTestContext(), chainID, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDA, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
		err = g.eventDB.StoreBlockTime(g.GetTestContext(), destinationChainIDB, blockNumber, uint64(time.Now().Unix())*blockNumber)
		Nil(g.T(), err)
	}

	total := 0.0
	for _, v := range cumulativePrice {
		total += v
	}
	count := float64(len(cumulativePrice))
	mean := total / count
	var median float64
	sort.Float64s(cumulativePrice)
	switch {
	case count == 0:
		median = 0.0
	case len(cumulativePrice)%2 == 0:
		median = (cumulativePrice[len(cumulativePrice)/2-1] + cumulativePrice[len(cumulativePrice)/2]) / 2
	default:
		median = cumulativePrice[len(cumulativePrice)/2]
	}

	statType := model.StatisticTypeTotalVolumeUsd
	duration := model.DurationAllTime
	platform := model.PlatformBridge
	// nolint:dupword
	result, err := g.client.GetAmountStatistic(g.GetTestContext(), statType, &platform, &duration, nil, nil, nil, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)

	Equal(g.T(), fmt.Sprintf("%f", total), *result.Response.Value)

	statType = model.StatisticTypeCountTransactions
	// nolint:dupword
	result, err = g.client.GetAmountStatistic(g.GetTestContext(), statType, &platform, &duration, nil, nil, nil, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), fmt.Sprintf("%f", count), *result.Response.Value)

	statType = model.StatisticTypeMeanVolumeUsd
	// nolint:dupword
	result, err = g.client.GetAmountStatistic(g.GetTestContext(), statType, &platform, &duration, nil, nil, nil, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), fmt.Sprintf("%f", mean), *result.Response.Value)

	statType = model.StatisticTypeMedianVolumeUsd
	result, err = g.client.GetAmountStatistic(g.GetTestContext(), statType, &platform, &duration, nil, nil, nil, nil)
	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), fmt.Sprintf("%f", median), *result.Response.Value)

	statType = model.StatisticTypeCountAddresses
	result, err = g.client.GetAmountStatistic(g.GetTestContext(), statType, &platform, &duration, nil, nil, nil, nil)

	Nil(g.T(), err)
	NotNil(g.T(), result)
	Equal(g.T(), "1.000000", *result.Response.Value)
}

func (g APISuite) TestGetBlockHeight() {
	chainID1 := 1
	chainID2 := 56

	type1 := model.ContractTypeCctp
	type2 := model.ContractTypeBridge

	contract1 := g.config.Chains[uint32(chainID1)].Contracts.CCTP
	contract2 := g.config.Chains[uint32(chainID2)].Contracts.Bridge

	block1 := uint64(3)
	block2 := uint64(4)

	contracts := []*model.ContractQuery{
		{
			ChainID: chainID1,
			Type:    type1,
		},
		{
			ChainID: chainID2,
			Type:    type2,
		},
	}

	// Store blocks in the database.
	err := g.db.StoreLastBlock(g.GetTestContext(), uint32(chainID1), block1, contract1)
	Nil(g.T(), err)

	err = g.db.StoreLastBlock(g.GetTestContext(), uint32(chainID2), block2, contract2)
	Nil(g.T(), err)

	results, err := g.client.GetBlockHeight(g.GetTestContext(), contracts)
	Nil(g.T(), err)
	Equal(g.T(), 2, len(results.Response))
	Equal(g.T(), int(block1), *results.Response[0].BlockNumber)
	Equal(g.T(), int(block2), *results.Response[1].BlockNumber)

}
