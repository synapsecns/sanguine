package graph

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"time"
)

func (r *queryResolver) getChainIDs(ctx context.Context, chainID *int) ([]uint32, error) {
	var chainIDs []uint32
	// if the chain ID is not specified, get all chain IDs
	if chainID == nil {
		chainIDsInt, err := r.DB.GetAllChainIDs(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get all chain IDs: %w", err)
		}
		chainIDs = append(chainIDs, chainIDsInt...)
	} else {
		chainIDs = append(chainIDs, uint32(*chainID))
	}
	return chainIDs, nil
}

func (r *queryResolver) getDirectionIn(direction *model.Direction) bool {
	var directionIn bool
	if direction != nil {
		directionIn = *direction == model.DirectionIn
	} else {
		directionIn = true
	}
	return directionIn
}

func (r *queryResolver) getTargetTime(hours *int) uint64 {
	var targetTime uint64
	if hours == nil {
		targetTime = uint64(time.Now().Add(-time.Hour * 24).Unix())
	} else {
		targetTime = uint64(time.Now().Add(-time.Hour * time.Duration(*hours)).Unix())
	}
	return targetTime
}

// func (r *queryResolver) getTokenAddressesByChainID(ctx context.Context, chainID uint32) ([]string, error) {
//	tokenAddresses, err := r.DB.GetTokenAddressesByChainID(ctx, chainID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to get token addresses by chain ID: %w", err)
//	}
//	return tokenAddresses, nil
//}

func (r *queryResolver) originToDestinationBridge(ctx context.Context, address *string, kappa *string, includePending *bool, page *int, tokenAddress *string, fromInfos []*model.PartialInfo) ([]*model.BridgeTransaction, error) {
	var results []*model.BridgeTransaction
	for _, fromInfo := range fromInfos {
		// blockTime, err := r.Fetcher.FetchClient.GetBlockTime(ctx, *fromInfo.ChainID, *fromInfo.BlockNumber)
		// if err != nil {
		//	return nil, fmt.Errorf("failed to get block time: %w", err)
		//}
		// fromInfo.Time = blockTime.Response
		txHash := common.HexToHash(*fromInfo.TxnHash)
		destinationKappa := crypto.Keccak256Hash(txHash.Bytes()).String()
		if kappa != nil {
			destinationKappa = *kappa
		}
		toInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, nil, address, tokenAddress, &destinationKappa, nil, *page)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		switch len(toInfos) {
		case 1:
			toInfo := toInfos[0]
			// blockTime, err = r.Fetcher.FetchClient.GetBlockTime(ctx, *toInfo.ChainID, *toInfo.BlockNumber)
			// if err != nil {
			//	return nil, fmt.Errorf("failed to get block time: %w", err)
			//}
			// toInfo.Time = blockTime.Response
			swapSuccess, err := r.DB.GetSwapSuccess(ctx, destinationKappa, uint32(*toInfo.ChainID))
			if err != nil {
				return nil, fmt.Errorf("failed to get swap success: %w", err)
			}
			pending := false
			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfo,
				ToInfo:      toInfo,
				Kappa:       &destinationKappa,
				Pending:     &pending,
				SwapSuccess: swapSuccess,
			})
		case 0:
			if *includePending {
				results = append(results, &model.BridgeTransaction{
					FromInfo:    fromInfo,
					ToInfo:      nil,
					Kappa:       &destinationKappa,
					Pending:     includePending,
					SwapSuccess: nil,
				})
			}
		default:
			return nil, fmt.Errorf("multiple toInfos found for kappa %s", destinationKappa)
		}
	}
	return results, nil
}

func (r *queryResolver) destinationToOriginBridge(ctx context.Context, address *string, txnHash *string, kappa *string, page *int, tokenAddress *string, toInfos []*model.PartialInfo) ([]*model.BridgeTransaction, error) {
	var results []*model.BridgeTransaction
	pending := false
	for _, toInfo := range toInfos {
		// blockTime, err := r.Fetcher.FetchClient.GetBlockTime(ctx, *toInfo.ChainID, *toInfo.BlockNumber)
		// if err != nil {
		//	return nil, fmt.Errorf("failed to get block time: %w", err)
		//}
		// toInfo.Time = blockTime.Response
		swapSuccess, err := r.DB.GetSwapSuccess(ctx, *kappa, uint32(*toInfo.ChainID))
		if err != nil {
			return nil, fmt.Errorf("failed to get swap success: %w", err)
		}
		originTxHash, err := r.DB.GetTxHashFromKappa(ctx, *kappa)
		if txnHash != nil {
			originTxHash = txnHash
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get origin tx hash: %w", err)
		}
		fromInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, nil, address, tokenAddress, nil, originTxHash, *page)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		switch {
		case len(fromInfos) > 1:
			return nil, fmt.Errorf("multiple fromInfos found for kappa %s", *kappa)
		case len(fromInfos) == 1:
			fromInfo := fromInfos[0]
			// blockTime, err = r.Fetcher.FetchClient.GetBlockTime(ctx, *fromInfo.ChainID, *fromInfo.BlockNumber)
			// if err != nil {
			//	return nil, fmt.Errorf("failed to get block time: %w", err)
			//}
			// fromInfo.Time = blockTime.Response

			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfo,
				ToInfo:      toInfo,
				Kappa:       kappa,
				Pending:     &pending,
				SwapSuccess: swapSuccess,
			})
		case len(fromInfos) == 0:
			return nil, fmt.Errorf("no fromInfo found for kappa %s", *kappa)
		}
	}
	return results, nil
}

func (r *queryResolver) originOrDestinationBridge(ctx context.Context, chainID *uint32, address *string, txnHash *string, kappa *string, includePending *bool, page *int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
	var results []*model.BridgeTransaction
	var toInfos []*model.PartialInfo
	var fromInfos []*model.PartialInfo
	infos, err := r.DB.PartialInfosFromIdentifiers(ctx, chainID, address, tokenAddress, kappa, txnHash, *page)
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
	}
	for _, info := range infos {
		kappa, err := r.DB.GetKappaFromTxHash(ctx, *info.TxnHash, chainID)
		if err != nil {
			return nil, fmt.Errorf("failed to get kappa from tx hash: %w", err)
		}
		if kappa == nil {
			fromInfos = append(fromInfos, info)
		} else {
			toInfos = append(toInfos, info)
		}
	}
	originResults, err := r.originToDestinationBridge(ctx, nil, nil, includePending, page, tokenAddress, fromInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to get origin -> destination bridge transactions: %w", err)
	}
	destinationResults, err := r.destinationToOriginBridge(ctx, nil, nil, nil, page, tokenAddress, toInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to get destination -> origin bridge transactions: %w", err)
	}
	results = r.mergeBridgeTransactions(originResults, destinationResults)

	return results, nil
}

func (r *queryResolver) mergeBridgeTransactions(origin []*model.BridgeTransaction, destination []*model.BridgeTransaction) []*model.BridgeTransaction {
	var results []*model.BridgeTransaction
	uniqueBridgeTransactions := make(map[*model.BridgeTransaction]bool)
	for _, originTx := range origin {
		uniqueBridgeTransactions[originTx] = true
	}
	for _, destinationTx := range destination {
		uniqueBridgeTransactions[destinationTx] = true
	}
	for tx := range uniqueBridgeTransactions {
		results = append(results, tx)
	}
	return results
}
