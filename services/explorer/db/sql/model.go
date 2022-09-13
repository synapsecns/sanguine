package sql

// SwapEvent stores data for emitted events from the Swap contract.
type SwapEvent struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;auto_increment:false"`
	// EventType is the type of the event
	EventType uint8 `gorm:"column:event_type;primaryKey;auto_increment:false"`
}

// BridgeEvent stores data for emitted events from the Bridge contract.
type BridgeEvent struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;auto_increment:false"`
	// EventType is the type of the event
	EventType uint8 `gorm:"column:event_type;primaryKey;auto_increment:false"`
}
