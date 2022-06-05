package indexer

// type DomainIndexer interface {
//	SyncMessages(ctx context.Context) error
//}

// func NewDomainIndexer(db db.DB, domain domains.DomainClient) DomainIndexer {
//	return NewDomainIndexer(db, domain)
//}

// MaxUint32 exports maxUint32 for testing.
func MaxUint32(x, y uint32) uint32 {
	return maxUint32(x, y)
}
