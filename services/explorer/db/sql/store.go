package sql

// Store is the clickhouse store.

//// NewClickHouseStore creates a new clickhouse database.
// func NewClickHouseStore(ctx context.Context, connString string) (db.ConsumerDB, error) {
//	logger.Debug("creating clickhouse store")
//
//	chdb, err := gorm.Open(clickhouse.Open(connString), &gorm.Config{})
//	if err != nil {
//		return nil, fmt.Errorf("could not create clickhouse connection: %w", err)
//	}
//
//}

// GetAllModels gets all models to migrate.
func GetAllModels() (allModels []interface{}) {
	allModels = append(allModels,
		&SwapEvent{}, &BridgeEvent{},
	)
	return allModels
}
