package stip_relayer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/stip_relayer/db"
	"github.com/synapsecns/sanguine/services/stip_relayer/stipconfig"
	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"
)

// Check Dune Query
// Store in database

// Call database
// Submit transactions for corresponding rebate

var DuneAPIKey = os.Getenv("DUNE_API_KEY")

func ExecuteDuneQuery() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.dune.com/api/v1/query/3345214/execute", bytes.NewBuffer([]byte(`{"performance": "medium"}`)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Dune-API-Key", DuneAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("EXECUTING DUNE QUERY")
	return resp, nil
}

func GetExecutionResults(execution_id string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.dune.com/api/v1/execution/"+execution_id+"/results", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Dune-API-Key", DuneAPIKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println("GETTING EXECUTION RESULTS")

	return resp, nil
}

// QuoterAPIServer is a struct that holds the configuration, database connection, gin engine, RPC client, metrics handler, and fast bridge contracts.
// It is used to initialize and run the API server.
type STIPRelayer struct {
	cfg           stipconfig.Config
	db            db.STIPDB
	omnirpcClient omniClient.RPCClient
	handler       metrics.Handler
	submittter    submitter.TransactionSubmitter
	signer        signer.Signer
}

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
	return &STIPRelayer{
		cfg:           cfg,
		db:            store,
		handler:       handler,
		omnirpcClient: omniRPCClient,
		submittter:    sm,
		signer:        sg,
	}, nil
}

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

type Result struct {
	Rows     []Row    `json:"rows"`
	Metadata Metadata `json:"metadata"`
}
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

type Metadata struct {
	ColumnNames         []string `json:"column_names"`
	ResultSetBytes      int      `json:"result_set_bytes"`
	TotalRowCount       int      `json:"total_row_count"`
	DatapointCount      int      `json:"datapoint_count"`
	PendingTimeMillis   int      `json:"pending_time_millis"`
	ExecutionTimeMillis int      `json:"execution_time_millis"`
}

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05.000 MST"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}
	t, err := time.Parse(ctLayout, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (s *STIPRelayer) Run(ctx context.Context) error {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // ensure cancel is called to clean up resources

	g, ctx := errgroup.WithContext(ctx)

	// Start the submitter goroutine
	g.Go(func() error {
		return s.StartSubmitter(ctx)
	})

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
		return err // handle the error from goroutines
	}

	return nil
}

// startSubmitter handles the initialization of the submitter
func (s *STIPRelayer) StartSubmitter(ctx context.Context) error {
	err := s.submittter.Start(ctx)
	if err != nil {
		fmt.Printf("could not start submitter: %v", err)
		// TODO: Will this force a panic in the Run() function?
		return err // panic in case submitter cannot start
	}
	return nil
}

// requestAndStoreResults handles the continuous request of new execution results and storing them in the database
func (s *STIPRelayer) RequestAndStoreResults(ctx context.Context) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err() // exit if context is cancelled
		case <-ticker.C:
			if err := s.ProcessExecutionResults(ctx); err != nil {
				// Log the error and decide whether to continue based on the error
				fmt.Printf("Error processing execution results: %v", err)
				// Optionally, you can return the error to stop the goroutine
				// return err
			}
		}
	}
}

// processExecutionResults encapsulates the logic for requesting and storing execution results
func (s *STIPRelayer) ProcessExecutionResults(ctx context.Context) error {
	resp, err := ExecuteDuneQuery()
	if err != nil {
		return fmt.Errorf("failed to execute Dune query: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	executionID, ok := result["execution_id"]
	if !ok {
		return fmt.Errorf("no execution_id found in response")
	}

	// Implement a more robust solution for waiting, such as polling with a timeout.
	time.Sleep(20 * time.Second) // Consider replacing this with a more robust solution

	executionResults, err := GetExecutionResults(executionID)
	if err != nil {
		return fmt.Errorf("failed to get execution results: %v", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	getResultsBody, err := ioutil.ReadAll(executionResults.Body)
	var jsonResult QueryResult
	err = json.Unmarshal(getResultsBody, &jsonResult)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	fmt.Println(jsonResult.Result.Rows)
	fmt.Println("Number of rows:", len(jsonResult.Result.Rows))

	// Convert each Row to a STIPTransactions and store them in the database
	return s.StoreResultsInDatabase(ctx, jsonResult.Result.Rows, jsonResult.ExecutionID)
}

// storeResultsInDatabase handles the storage of results in the database
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
			return fmt.Errorf("error inserting new STIP transactions: %v", err)
		}
	}

	// Optionally, confirm that the insert occurred
	stipTransactionsNotRebated, err := s.db.GetSTIPTransactionsNotRebated(ctx)
	if err != nil {
		return fmt.Errorf("error getting STIP transactions not rebated: %v", err)
	}
	if len(stipTransactionsNotRebated) == 0 {
		fmt.Println("No STIP transactions found that have not been rebated.")
	} else {
		fmt.Println("Found", len(stipTransactionsNotRebated), "STIP transactions that have not been rebated.")
	}

	return nil
}

// queryRebateAndUpdate handles the querying for new, non-relayed/rebated results, rebates/relays them, and updates the result row
func (s *STIPRelayer) QueryRebateAndUpdate(ctx context.Context) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err() // exit if context is cancelled
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

// relayAndRebateTransactions encapsulates the logic for querying, rebating/relaying, and updating results
func (s *STIPRelayer) RelayAndRebateTransactions(ctx context.Context) error {
	// Define the rate limit (e.g., 5 transactions per second)
	// You can adjust r (rate per second) and b (burst size) according to your specific requirements
	r := rate.Limit(5)
	b := 1
	limiter := rate.NewLimiter(r, b)

	// Query DB to get all STIPs that need to be relayed
	stipTransactionsNotRebated, err := s.db.GetSTIPTransactionsNotRebated(ctx)
	if err != nil {
		return fmt.Errorf("error getting STIP transactions not rebated: %v", err)
	}
	if len(stipTransactionsNotRebated) == 0 {
		fmt.Println("No STIP transactions found that have not been rebated.")
		return nil
	} else {
		fmt.Println("Found", len(stipTransactionsNotRebated), "STIP transactions that have not been rebated.")
	}

	// Relay and rebate transactions with rate limiting
	for _, transaction := range stipTransactionsNotRebated {
		// Wait for the limiter to allow another event
		if err := limiter.Wait(ctx); err != nil {
			fmt.Printf("Error waiting for rate limiter: %v", err)
			// Handle the error (e.g., break the loop or return the error)
			return err
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

// submitAndRebateTransaction handles the relaying and rebating of a single transaction
func (s *STIPRelayer) SubmitAndRebateTransaction(ctx context.Context, transaction *db.STIPTransactions) error {
	// Calculate the transfer amount based on transaction details
	// This function encapsulates the logic for determining the transfer amount
	// You can define it elsewhere and call it here
	transferAmount, err := s.CalculateTransferAmount(transaction)
	if err != nil {
		return fmt.Errorf("could not calculate transfer amount: %w", err)
	}

	// Setup for submitting the transaction
	chainId := s.cfg.ArbChainID
	arbAddress := s.cfg.ArbAddress
	backendClient, err := s.omnirpcClient.GetClient(ctx, big.NewInt(int64(chainId)))
	if err != nil {
		return fmt.Errorf("could not get client: %w", err)
	}

	// Submit the transaction
	nonceSubmitted, err := s.submittter.SubmitTransaction(ctx, big.NewInt(int64(chainId)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
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
	err = s.db.UpdateSTIPTransactionRebated(ctx, transaction.Hash, nonceSubmitted)
	if err != nil {
		return fmt.Errorf("could not update STIP transaction as rebated: %w", err)
	}

	return nil
}

// calculateTransferAmount determines the amount to transfer based on the transaction
// This is a placeholder function. Implement the actual calculation logic based on your requirements.
func (s *STIPRelayer) CalculateTransferAmount(transaction *db.STIPTransactions) (*big.Int, error) {
	// Pseudocode:
	// 1. Retrieve the necessary fields from the transaction.
	// 2. Apply any calculations as defined in the config.
	// 3. Return the result as a *big.Int representing the amount to transfer.

	// Example calculation (replace with actual logic):
	return big.NewInt(params.Ether), nil
}
