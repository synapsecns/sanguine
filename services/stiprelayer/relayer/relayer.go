package relayer

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/contrib/screener-api/client"
	"math/big"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/stiprelayer/db"
	"github.com/synapsecns/sanguine/services/stiprelayer/stipapi"
	"github.com/synapsecns/sanguine/services/stiprelayer/stipconfig"
	"go.opentelemetry.io/otel/attribute"
	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"
)

// Check Dune Query
// Store in database

// Call database
// Submit transactions for corresponding rebate

// STIPRelayer is the main struct for the STIP relayer service.
type STIPRelayer struct {
	cfg           stipconfig.Config
	db            db.STIPDB
	omnirpcClient omniClient.RPCClient
	handler       metrics.Handler
	submittter    submitter.TransactionSubmitter
	signer        signer.Signer
	apiServer     *stipapi.Server
	// screener is used to screen addresses using the screener-api client stub.
	screener client.ScreenerClient
}

// NewSTIPRelayer creates a new STIPRelayer with the provided context and configuration.
func NewSTIPRelayer(ctx context.Context,
	cfg stipconfig.Config,
	handler metrics.Handler,
	omniRPCClient omniClient.RPCClient,
	store db.STIPDB,
) (*STIPRelayer, error) {
	sg, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not get signer: %w", err)
	}
	sm := submitter.NewTransactionSubmitter(handler, sg, omniRPCClient, store.SubmitterDB(), &cfg.SubmitterConfig)

	apiServer, err := stipapi.NewStipAPI(ctx, cfg, handler)
	if err != nil {
		return nil, fmt.Errorf("could not get api server: %w", err)
	}

	screener, _ := client.NewNoOpClient()

	if cfg.ScreenerAPIUrl != "" {
		screener, err = client.NewClient(handler, cfg.ScreenerAPIUrl)
		if err != nil {
			return nil, fmt.Errorf("could not create screener client: %w", err)
		}
	}

	return &STIPRelayer{
		cfg:           cfg,
		db:            store,
		handler:       handler,
		omnirpcClient: omniRPCClient,
		submittter:    sm,
		signer:        sg,
		apiServer:     apiServer,
		screener:      screener,
	}, nil
}

// QueryResult represents the result of a Dune query.
type QueryResult struct {
	ExecutionID        string    `json:"execution_id"`
	QueryID            int       `json:"query_id"`
	State              string    `json:"state"`
	SubmittedAt        time.Time `json:"submitted_at"`
	ExpiresAt          time.Time `json:"expires_at"`
	ExecutionStartedAt time.Time `json:"execution_started_at"`
	ExecutionEndedAt   time.Time `json:"execution_ended_at"`
	Result             Result    `json:"result"`
}

// Result represents the data structure for the result of a query execution.
type Result struct {
	Rows     []Row    `json:"rows"`
	Metadata Metadata `json:"metadata"`
}

// Row represents a single row of the result of a query execution.
type Row struct {
	Address    string     `json:"address"`
	Amount     float64    `json:"amount"`
	AmountUsd  float64    `json:"amount_usd"`
	ArbPrice   float64    `json:"arb_price"`
	BlockTime  CustomTime `json:"block_time"`
	Direction  string     `json:"direction"`
	Hash       string     `json:"hash"`
	Module     string     `json:"module"`
	Token      string     `json:"token"`
	TokenPrice float64    `json:"token_price"`
}

// Metadata represents the metadata of a query execution result.
type Metadata struct {
	ColumnNames         []string `json:"column_names"`
	ResultSetBytes      int      `json:"result_set_bytes"`
	TotalRowCount       int      `json:"total_row_count"`
	DatapointCount      int      `json:"datapoint_count"`
	PendingTimeMillis   int      `json:"pending_time_millis"`
	ExecutionTimeMillis int      `json:"execution_time_millis"`
}

// CustomTime is a custom time type for handling specific time format in JSON unmarshalling.
type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05.000 MST"

// UnmarshalJSON overrides the default JSON unmarshaling for CustomTime to handle specific time format.
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}
	t, err := time.Parse(ctLayout, s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}
	ct.Time = t
	return nil
}

// Run starts the STIPRelayer service by initiating various goroutines.
func (s *STIPRelayer) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// Start the submitter goroutine
	g.Go(func() error {
		return s.StartSubmitter(ctx)
	})

	g.Go(func() error {
		err := s.apiServer.Run(ctx)
		if err != nil {
			return fmt.Errorf("could not start api server: %w", err)
		}
		return nil
	})

	err := s.ProcessExecutionResults(ctx)
	if err != nil {
		return fmt.Errorf("error processing execution results: %w", err)
	}

	// Start the ticker goroutine for requesting and storing execution results
	g.Go(func() error {
		return s.RequestAndStoreResults(ctx)
	})

	// Start the goroutine for querying, rebating/relaying, and updating results
	g.Go(func() error {
		return s.QueryRebateAndUpdate(ctx)
	})

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not run: %w", err) // handle the error from goroutines
	}

	return nil
}

// StartSubmitter handles the initialization of the submitter.
func (s *STIPRelayer) StartSubmitter(ctx context.Context) error {
	err := s.submittter.Start(ctx)
	if err != nil {
		fmt.Printf("could not start submitter: %v", err)
		// TODO: Will this force a panic in the Run() function?
		return fmt.Errorf("could not start submitter: %w", err) // panic in case submitter cannot start
	}
	return nil
}

// RequestAndStoreResults handles the continuous request of new execution results and storing them in the database.
func (s *STIPRelayer) RequestAndStoreResults(ctx context.Context) error {
	// TODO: If undefined, what do? Need a default, otherwise, panic
	ticker := time.NewTicker(s.cfg.DuneInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			//nolint: wrapcheck
			return ctx.Err() // exit if context is canceled
		case <-ticker.C:
			err := s.ProcessExecutionResults(ctx)
			if err != nil {
				// Log the error and decide whether to continue based on the error
				fmt.Printf("Error processing execution results: %v", err)
				// Optionally, you can return the error to stop the goroutine
				// return err
			}
		}
	}
}

// ProcessExecutionResults encapsulates the logic for requesting and storing execution results.
func (s *STIPRelayer) ProcessExecutionResults(parentCtx context.Context) (err error) {
	fmt.Println("Starting execution logic")

	ctx, span := s.handler.Tracer().Start(parentCtx, "ProcessExecutionResults")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	executionID, err := s.ExecuteDuneQuery(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute Dune query: %w", err)
	}
	span.SetAttributes(attribute.String("execution_id", executionID))

	// TODO: remove if exponentialBackoff.InitialInterval waits 30 seconds?
	// time.Sleep(30 * time.Second) // Consider replacing this with a more robust solution
	var getResultsJSONResult QueryResult
	operation := func() error {
		jsonResult, err := s.GetExecutionResults(ctx, executionID)
		if err != nil {
			return fmt.Errorf("failed to get execution results: %w", err)
		}

		if jsonResult.State != "QUERY_STATE_COMPLETED" {
			// query state is not completed, so return an error to retry
			return fmt.Errorf("query state is not completed")
		}
		getResultsJSONResult = *jsonResult
		return nil
	}

	// Create a new exponential backoff policy
	expBackOff := backoff.NewExponentialBackOff()
	expBackOff.InitialInterval = 30 * time.Second
	expBackOff.MaxElapsedTime = 300 * time.Second

	// Retry the operation with the backoff policy
	err = backoff.Retry(operation, expBackOff)
	if err != nil {
		return fmt.Errorf("failed to get execution results after retries: %w", err)
	}

	var rowsAfterStartDate []Row
	var firstResultTime time.Time
	for _, row := range getResultsJSONResult.Result.Rows {
		// TODO: Will this panic if StartDate not set?
		if row.BlockTime.After(s.cfg.StartDate) {
			rowsAfterStartDate = append(rowsAfterStartDate, row)
		}
		if firstResultTime.IsZero() || row.BlockTime.Before(firstResultTime) {
			firstResultTime = row.BlockTime.Time
		}
	}
	span.SetAttributes(
		attribute.Int("number_of_rows", len(rowsAfterStartDate)),
		attribute.String("first_result_time", firstResultTime.String()),
	)
	fmt.Println("Number of rows after start date:", len(rowsAfterStartDate))

	// Convert each Row to a STIPTransactions and store them in the database
	return s.StoreResultsInDatabase(ctx, rowsAfterStartDate, getResultsJSONResult.ExecutionID)
}

// StoreResultsInDatabase handles the storage of results in the database.
func (s *STIPRelayer) StoreResultsInDatabase(ctx context.Context, rows []Row, executionID string) error {
	stipTransactions := make([]db.STIPTransactions, len(rows))
	for i, row := range rows {
		stipTransactions[i] = db.STIPTransactions{
			Address:     row.Address,
			Amount:      row.Amount,
			AmountUSD:   row.AmountUsd,
			ArbPrice:    row.ArbPrice,
			BlockTime:   row.BlockTime.Time,
			Direction:   row.Direction,
			ExecutionID: executionID,
			Hash:        row.Hash,
			Module:      row.Module,
			Token:       row.Token,
			TokenPrice:  row.TokenPrice,
			Rebated:     false,
		}
	}

	if len(stipTransactions) > 0 {
		if err := s.db.InsertNewStipTransactions(ctx, stipTransactions); err != nil {
			return fmt.Errorf("error inserting new STIP transactions: %w", err)
		}
	}

	return nil
}

// QueryRebateAndUpdate handles the querying for new, non-relayed/rebated results, rebates/relays them, and updates the result row.
func (s *STIPRelayer) QueryRebateAndUpdate(ctx context.Context) error {
	// TODO: If undefined, what do? Need a default, otherwise, panic
	ticker := time.NewTicker(s.cfg.RebateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			//nolint: wrapcheck
			return ctx.Err() // exit if context is canceled
		case <-ticker.C:
			if err := s.RelayAndRebateTransactions(ctx); err != nil {
				// Log the error and decide whether to continue based on the error
				fmt.Printf("Error relaying and rebating transactions: %v", err)
				// Optionally, you can return the error to stop the goroutine
				// return err
			}
		}
	}
}

// RelayAndRebateTransactions encapsulates the logic for querying, rebating/relaying, and updating results.
func (s *STIPRelayer) RelayAndRebateTransactions(ctx context.Context) error {
	// Define the rate limit (e.g., 5 transactions per second)
	// You can adjust r (rate per second) and b (burst size) according to your specific requirements
	// TODO: Consider making these values configurable.
	r := rate.Limit(2)
	b := 1
	limiter := rate.NewLimiter(r, b)

	// Query DB to get all STIPs that need to be relayed
	stipTransactionsNotRebated, err := s.db.GetSTIPTransactionsNotRebated(ctx)
	if err != nil {
		return fmt.Errorf("error getting STIP transactions not rebated: %w", err)
	}
	if len(stipTransactionsNotRebated) == 0 {
		fmt.Println("No STIP transactions found that have not been rebated.")
		return nil
	}
	fmt.Println("Found", len(stipTransactionsNotRebated), "STIP transactions that have not been rebated.")

	// Relay and rebate transactions with rate limiting
	for _, transaction := range stipTransactionsNotRebated {
		// Wait for the limiter to allow another event
		if err := limiter.Wait(ctx); err != nil {
			fmt.Printf("Error waiting for rate limiter: %v", err)
			// Handle the error (e.g., break the loop or return the error)
			return fmt.Errorf("error waiting for rate limiter: %w", err)
		}

		// Submit and rebate the transaction
		if err := s.SubmitAndRebateTransaction(ctx, transaction); err != nil {
			// Log the error and continue processing the rest of the transactions
			fmt.Printf("Error relaying and rebating transaction: %v", err)
			// Optionally, you can return the error to stop processing further transactions
			// return err
		}
	}

	return nil
}

const stipRuleset = "stip"

// SubmitAndRebateTransaction handles the relaying and rebating of a single transaction.
func (s *STIPRelayer) SubmitAndRebateTransaction(ctx context.Context, transaction *db.STIPTransactions) error {
	// Check if the address is blocked
	blocked, err := s.screener.ScreenAddress(ctx, stipRuleset, transaction.Address)
	if err != nil {
		return fmt.Errorf("could not screen address: %w", err)
	}

	if blocked {
		err = s.db.UpdateSTIPTransactionDoNotProcess(ctx, transaction.Hash)
		if err != nil {
			return fmt.Errorf("could not update STIP transaction as do not process: %w", err)
		}
		return fmt.Errorf("address is blocked: %s", transaction.Address)
	}

	// Calculate the transfer amount based on transaction details
	// This function encapsulates the logic for determining the transfer amount
	// You can define it elsewhere and call it here
	transferAmount, err := s.CalculateTransferAmount(ctx, transaction)
	if err != nil {
		err := s.db.UpdateSTIPTransactionDoNotProcess(ctx, transaction.Hash)
		if err != nil {
			return fmt.Errorf("could not update STIP transaction as do not process: %w", err)
		}
		return fmt.Errorf("could not calculate transfer amount: %w", err)
	}

	// Setup for submitting the transaction
	chainID := s.cfg.ArbChainID
	arbAddress := s.cfg.ArbAddress
	backendClient, err := s.omnirpcClient.GetClient(ctx, big.NewInt(int64(chainID)))
	if err != nil {
		return fmt.Errorf("could not get client: %w", err)
	}

	// Submit the transaction
	nonceSubmitted, err := s.submittter.SubmitTransaction(ctx, big.NewInt(int64(chainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		erc20, err := ierc20.NewIERC20(common.HexToAddress(arbAddress), backendClient)
		if err != nil {
			return nil, fmt.Errorf("could not get erc20: %w", err)
		}

		// Use the calculated transfer amount in the actual transfer
		transferTx, err := erc20.Transfer(transactor, common.HexToAddress(transaction.Address), transferAmount)
		if err != nil {
			return nil, fmt.Errorf("could not transfer: %w", err)
		}

		return transferTx, nil
	})

	if err != nil {
		return fmt.Errorf("could not submit transfer: %w", err)
	}

	// Update the database to mark the transaction as rebated
	err = s.db.UpdateSTIPTransactionRebated(ctx, transaction.Hash, nonceSubmitted, transferAmount.String())
	if err != nil {
		return fmt.Errorf("could not update STIP transaction as rebated: %w", err)
	}

	return nil
}

// CalculateTransferAmount determines the amount to transfer based on the transaction.
func (s *STIPRelayer) CalculateTransferAmount(ctx context.Context, transaction *db.STIPTransactions) (*big.Int, error) {
	var toChainID int
	switch transaction.Direction {
	case "ARB":
		toChainID = 42161
	case "ETH":
		toChainID = 1
	case "AVAX":
		toChainID = 43114
	}

	moduleConfig, ok := s.cfg.FeesAndRebates[toChainID][transaction.Module]
	if !ok {
		return nil, fmt.Errorf("module configuration not found for module %s", transaction.Module)
	}

	tokenConfig, ok := moduleConfig[transaction.Token]
	if !ok {
		return nil, fmt.Errorf("token configuration not found for token %s", transaction.Token)
	}

	// Convert values to big.Float for precision during calculations
	amountUSD := new(big.Float).SetFloat64(transaction.AmountUSD)
	rebateBPS := new(big.Float).SetFloat64(tokenConfig.RebateBps)

	// Calculate rebate in USD (amountUSD * rebateBPS / 10000)
	// Divide rebateBPS by 10000 to get the actual rebate rate
	rebateRate := new(big.Float).Quo(rebateBPS, big.NewFloat(10000))
	rebateUSD := new(big.Float).Mul(amountUSD, rebateRate)

	// Convert arbPrice to big.Float
	arbPrice := new(big.Float).SetFloat64(transaction.ArbPrice)

	// Calculate the amount of ARB to transfer (rebateUSD / arbPrice)
	transferAmountFloat := new(big.Float).Quo(rebateUSD, arbPrice)

	// Convert the transfer amount to big.Int (assuming we want to truncate to the integer value)
	// Multiply by 10^18 to get the value in wei (like params.Ether does)
	transferAmountFloatWei := new(big.Float).Mul(transferAmountFloat, big.NewFloat(1e18))
	transferAmount, _ := transferAmountFloatWei.Int(nil) // Truncate fractional part
	// Check if transferAmount is greater than configured max ARB (MaxAmount * 10^18 wei)
	limit := big.NewInt(s.cfg.ARBMaxTransfer)
	limit = limit.Mul(limit, big.NewInt(1e18)) // Convert to wei
	if transferAmount.Cmp(limit) > 0 {
		return nil, fmt.Errorf("transfer amount exceeds the limit of %d ARB", s.cfg.ARBMaxTransfer)
	}
	// Check if transferAmount is lower than configured min ARB (MinAmount * 10^18 wei)
	minAmountFloat := new(big.Float).SetFloat64(s.cfg.ARBMinTransfer)
	minAmountFloatWei := new(big.Float).Mul(minAmountFloat, big.NewFloat(1e18))
	minAmount, _ := minAmountFloatWei.Int(nil) // Truncate fractional part
	if transferAmount.Cmp(minAmount) < 0 {
		return nil, fmt.Errorf("transfer amount is lower than the minimum of %f ARB", s.cfg.ARBMinTransfer)
	}
	// If you need to round to the nearest integer instead of truncating, use the following:
	// transferAmount := new(big.Int)
	// transferAmountFloat.Int(transferAmount) // Round to the nearest integer

	// Finally, apply the rebate cap
	var err error
	transferAmount, err = s.applyRebateCap(ctx, transaction, transferAmount)
	if err != nil {
		return nil, fmt.Errorf("could not apply rebate cap: %w", err)
	}

	return transferAmount, nil
}

func (s *STIPRelayer) applyRebateCap(ctx context.Context, transaction *db.STIPTransactions, amount *big.Int) (*big.Int, error) {
	totalArbRebated, err := s.db.GetTotalArbRebated(ctx, transaction.Address)
	if err != nil {
		return nil, fmt.Errorf("could not get total ARB rebated: %w", err)
	}
	rebateCap := new(big.Int).Mul(big.NewInt(s.cfg.GetArbCapPerAddress()), big.NewInt(1e18)) // Convert to wei
	remainingAmount := new(big.Int).Sub(rebateCap, totalArbRebated)

	if remainingAmount.Cmp(big.NewInt(0)) <= 0 {
		return nil, fmt.Errorf("address has reached the ARB rebate cap: %s", rebateCap.String())
	} else if amount.Cmp(remainingAmount) >= 0 {
		return remainingAmount, nil
	}
	return amount, nil
}
