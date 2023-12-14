// Quoter submits quotes to the RFQ API for which assets the relayer is willing to relay.
package quoter

type Quoter interface {
	// SubmitQuote submits a quote to the RFQ API.
	SubmitQuote() error
	// GetQuote gets a quote from the RFQ API.
}
