package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/ethergo/client"
	listenerDB "github.com/synapsecns/sanguine/ethergo/listener/db"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/listener"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	ethereumNodeURL = "https://eth.llamarpc.com"                   // Replace with your Ethereum node URL
	contractAddress = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48" // Replace with the contract address you want to listen to
	dbPath          = ":memory:"
	initialBlock    = 0 // Replace with the block number you want to start listening from
)

func main() {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(listenerDB.GetAllModels()...)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Create a simple metrics handler (you may want to implement a proper one)
	metricsHandler := metrics.NewNullHandler()

	store := listenerDB.NewChainListenerStore(db, metricsHandler)

	_, err = store.LatestBlockForChain(context.Background(), 1)
	if err != nil {
		if !errors.Is(err, listenerDB.ErrNoLatestBlockForChainID) {
			log.Fatalf("Failed to get latest block: %v", err)
		}

		// pick a reasonable start block, will default to 0.
		err = store.PutLatestBlock(context.Background(), 1, 20612563)
		if err != nil {
			log.Fatalf("Failed to put latest block): %v", err)
		}
	}

	ethClient, err := client.DialBackend(context.Background(), ethereumNodeURL, metricsHandler)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}

	contractAddr := common.HexToAddress(contractAddress)

	chainListener, err := listener.NewChainListener(
		ethClient,
		store,
		contractAddr,
		uint64(initialBlock),
		metricsHandler,
	)
	if err != nil {
		log.Fatalf("Failed to create chain listener: %v", err)
	}

	ctx := context.Background()

	fmt.Println("Starting listener...")
	err = chainListener.Listen(ctx, handleLog)
	if err != nil {
		log.Fatalf("Listener error: %v", err)
	}
}

func handleLog(_ context.Context, log types.Log) error {
	fmt.Printf("New log: BlockNumber=%d, TxHash=%s\n", log.BlockNumber, log.TxHash.Hex())
	// Here you can process the log as needed
	return nil
}
