package internal

type GetRFQByTxIDResponse struct {
	Bridge        Bridge        `json:"Bridge"`
	BridgeRequest BridgeRequest `json:"BridgeRequest"`
	BridgeRelay   BridgeRelay   `json:"BridgeRelay"`
	BridgeProof   BridgeProof   `json:"BridgeProof"`
	BridgeClaim   BridgeClaim   `json:"BridgeClaim"`
}

type Bridge struct {
	TransactionID         string `json:"transactionId"`
	OriginChain           string `json:"originChain"`
	DestChain             string `json:"destChain"`
	OriginChainID         int    `json:"originChainId"`
	DestChainID           int    `json:"destChainId"`
	OriginToken           string `json:"originToken"`
	DestToken             string `json:"destToken"`
	OriginAmountFormatted string `json:"originAmountFormatted"`
	DestAmountFormatted   string `json:"destAmountFormatted"`
	Sender                string `json:"sender"`
	SendChainGas          int    `json:"sendChainGas"`
}

type BridgeRequest struct {
	BlockNumber     string `json:"blockNumber"`
	BlockTimestamp  int64  `json:"blockTimestamp"`
	TransactionHash string `json:"transactionHash"`
}

type BridgeRelay struct {
	BlockNumber     string `json:"blockNumber"`
	BlockTimestamp  int64  `json:"blockTimestamp"`
	TransactionHash string `json:"transactionHash"`
	Relayer         string `json:"relayer"`
	To              string `json:"to"`
}

type BridgeProof struct {
	BlockNumber     string `json:"blockNumber"`
	BlockTimestamp  int64  `json:"blockTimestamp"`
	TransactionHash string `json:"transactionHash"`
	Relayer         string `json:"relayer"`
}

type BridgeClaim struct {
	BlockNumber     string `json:"blockNumber"`
	BlockTimestamp  int64  `json:"blockTimestamp"`
	TransactionHash string `json:"transactionHash"`
	To              string `json:"to"`
	Relayer         string `json:"relayer"`
	AmountFormatted string `json:"amountFormatted"`
}
