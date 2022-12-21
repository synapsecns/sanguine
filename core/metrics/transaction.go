package metrics

// Transaction is a transaction used fot racking metrics
type Transaction interface {
	// End ends the transaction
	End()
	// NewGoroutine creates a new goroutine transaction
	NewGoroutine() Transaction
}
