package graph

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
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

func (r *queryResolver) originToDestinationBridge(ctx context.Context, address *string, kappa *string, includePending *bool, page *int, tokenAddress *string, fromInfos []*model.PartialInfo) ([]*model.BridgeTransaction, error) {
	var results []*model.BridgeTransaction
	for _, fromInfo := range fromInfos {
		txHash := common.HexToHash(*fromInfo.TxnHash)
		destinationKappa := crypto.Keccak256Hash(txHash.Bytes()).String()
		if kappa != nil {
			destinationKappa = *kappa
		}
		toInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, generatePartialInfoQuery(nil, address, tokenAddress, &destinationKappa, nil, *page))
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		switch len(toInfos) {
		case 1:
			toInfo := toInfos[0]
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
		fromInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, generatePartialInfoQuery(nil, address, tokenAddress, nil, originTxHash, *page))
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		switch {
		case len(fromInfos) > 1:
			return nil, fmt.Errorf("multiple fromInfos found for kappa %s", *kappa)
		case len(fromInfos) == 1:
			fromInfo := fromInfos[0]
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
	infos, err := r.DB.PartialInfosFromIdentifiers(ctx, generatePartialInfoQuery(chainID, address, tokenAddress, kappa, txnHash, *page))
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
	}
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	for _, info := range infos {
		txHashSpecifier := generateSingleSpecifierStringSQL(info.TxnHash, sql.TxHashFieldName, &firstFilter, "")
		query := fmt.Sprintf(`SELECT * FROM bridge_events %s%s`, chainIDSpecifier, txHashSpecifier)
		kappa, err := r.DB.GetKappaFromTxHash(ctx, query)
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

// generateAddressSpecifierSQL generates a where function with an string.
func generateAddressSpecifierSQL(address *string, firstFilter *bool, tablePrefix string) string {
	if address != nil {
		if *firstFilter {
			*firstFilter = false
			return fmt.Sprintf(" WHERE (%s%s = '%s' OR  %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
		}
		return fmt.Sprintf(" AND (%s%s = '%s' OR %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	}
	return ""
}

// generateSingleSpecifierI32SQL generates a where function with an uint32.
func generateSingleSpecifierI32SQL(value *uint32, field string, firstFilter *bool, tablePrefix string) string {
	if value != nil {
		if *firstFilter {
			*firstFilter = false
			return fmt.Sprintf(" WHERE %s%s = %d", tablePrefix, field, *value)
		}
		return fmt.Sprintf(" AND %s%s = %d", tablePrefix, field, *value)
	}
	return ""
}

// generateBlockSpecifierSQL generates a where function with an uint64.
func generateBlockSpecifierSQL(value *uint64, field string, firstFilter *bool, tablePrefix string) string {
	if value != nil {
		if *firstFilter {
			*firstFilter = false
			return fmt.Sprintf(" WHERE %s%s >= %d", tablePrefix, field, *value)
		}
		return fmt.Sprintf(" AND %s%s >= %d", tablePrefix, field, *value)
	}
	return ""
}

// GenerateSingleSpecifierStringSQL generates a where function with a string.
func generateSingleSpecifierStringSQL(value *string, field string, firstFilter *bool, tablePrefix string) string {
	if value != nil {
		if *firstFilter {
			*firstFilter = false
			return fmt.Sprintf(" WHERE %s%s = '%s'", tablePrefix, field, *value)
		}
		return fmt.Sprintf(" AND %s%s = '%s'", tablePrefix, field, *value)
	}
	return ""
}

// generatePartialInfoQuery returns the query for making the PartialInfo query.
func generatePartialInfoQuery(chainID *uint32, address, tokenAddress, kappa, txHash *string, page int) string {
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "t1.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "t1.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "t1.")
	kappaSpecifier := generateSingleSpecifierStringSQL(kappa, sql.KappaFieldName, &firstFilter, "t1.")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "t1.")

	pageSpecifier := fmt.Sprintf(" ORDER BY %s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.PageSize, (page-1)*sql.PageSize)

	compositeIdentifiers := chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + txHashSpecifier + pageSpecifier
	selectParameters := fmt.Sprintf(
		`%s,%s,%s,%s,%s,%s,%s,%s,%s,%s, max(%s)`,
		sql.ContractAddressFieldName, sql.ChainIDFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.DestinationKappaFieldName,
		sql.SenderFieldName, sql.TxHashFieldName, sql.InsertTimeFieldName,
	)
	groupByParameters := fmt.Sprintf(
		`%s,%s,%s,%s,%s,%s,%s,%s,%s,%s`,
		sql.TxHashFieldName, sql.ContractAddressFieldName, sql.ChainIDFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.DestinationKappaFieldName, sql.SenderFieldName,
	)
	joinOnParameters := fmt.Sprintf(
		`t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s
		AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = insert_max_time`,
		sql.TxHashFieldName, sql.TxHashFieldName, sql.ContractAddressFieldName, sql.ContractAddressFieldName, sql.ChainIDFieldName,
		sql.ChainIDFieldName, sql.EventTypeFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.TokenFieldName, sql.AmountFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.EventIndexFieldName,
		sql.DestinationKappaFieldName, sql.DestinationKappaFieldName, sql.SenderFieldName, sql.SenderFieldName, sql.InsertTimeFieldName,
	)
	query := fmt.Sprintf(
		`
		SELECT t1.* FROM bridge_events t1
    	JOIN (
    	SELECT %s AS insert_max_time
    	FROM bridge_events GROUP BY %s) t2
    	    ON (%s) %s `,
		selectParameters, groupByParameters, joinOnParameters, compositeIdentifiers)
	return query
}

// generateBridgeEventCountQuery creates the query for bridge event count.
func generateBridgeEventCountQuery(chainID uint32, address *string, tokenAddress *string, directionIn bool, firstBlock *uint64) string {
	chainField := sql.ChainIDFieldName
	if directionIn {
		chainField = sql.DestinationChainIDFieldName
	}
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(&chainID, chainField, &firstFilter, "")
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.RecipientFieldName, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	blockSpecifier := generateBlockSpecifierSQL(firstBlock, sql.BlockNumberFieldName, &firstFilter, "")

	query := fmt.Sprintf(`SELECT COUNT(DISTINCT (%s, %s)) FROM bridge_events %s%s%s%s`,
		sql.TxHashFieldName, sql.EventIndexFieldName, chainIDSpecifier, addressSpecifier, tokenAddressSpecifier, blockSpecifier)
	return query
}
