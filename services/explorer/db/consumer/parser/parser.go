package parser

import (
	"context"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type Parser interface {
	// ParseAndStore parses the logs and stores them in the database.
	ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error
}
