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
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/stip_relayer/db"
	"github.com/synapsecns/sanguine/services/stip_relayer/stipconfig"
	"golang.org/x/sync/errgroup"
)

// Check Dune Query
// Store in database

// Call database
// Submit transactions for corresponding rebate

// Dune API key dmcGJqYYuq36viagnjoBMTMuwM4wjzqf

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

func main() {
	ExecuteDuneQuery()
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
	// TODO: Context fix
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := s.submittter.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not start submitter: %w", err)
		}
		return nil
	})

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	errChan := make(chan error)

	go func() {
		for {
			select {
			case <-ticker.C:
				resp, err := ExecuteDuneQuery()
				if err != nil {
					errChan <- fmt.Errorf("Failed to execute Dune query: %v", err)
					return
				}

				body, err := ioutil.ReadAll(resp.Body)

				if err != nil {
					errChan <- fmt.Errorf("Failed to read response body: %v", err)
					return
				}

				var result map[string]string
				err = json.Unmarshal(body, &result)
				if err != nil {
					errChan <- fmt.Errorf("Failed to unmarshal response body: %v", err)
					return
				}

				execution_id, ok := result["execution_id"]
				if !ok {
					errChan <- fmt.Errorf("No execution_id found in response")
					return
				}

				time.Sleep(20 * time.Second)

				executionResults, err := GetExecutionResults(execution_id)
				if err != nil {
					errChan <- err
					return
				}
				// fmt.Println(executionResults)

				if err != nil {
					fmt.Errorf("Failed to get execution results: %v", err)
				}

				if resp.StatusCode != 200 {
					fmt.Errorf("Expected status code 200, got %d", resp.StatusCode)
				}

				getResultsBody, err := ioutil.ReadAll(executionResults.Body)
				var jsonResult QueryResult
				err = json.Unmarshal(getResultsBody, &jsonResult)
				if err != nil {
					// handle error, e.g., print it or log it
					fmt.Println("Error unmarshalling JSON:", err)
					return
				}
				fmt.Println(jsonResult.Result.Rows)
				fmt.Println("Number of rows:", len(jsonResult.Result.Rows))

				// Store executionResults in DB
				// This part is left as a comment because it depends on the specific implementation of your DB

				// Convert each Row to a STIPTransactions
				stipTransactions := make([]db.STIPTransactions, len(jsonResult.Result.Rows))
				for i, row := range jsonResult.Result.Rows {
					stipTransactions[i] = db.STIPTransactions{
						Address:     row.Address,
						Amount:      row.Amount,
						AmountUSD:   row.AmountUsd,
						ArbPrice:    row.ArbPrice,
						BlockTime:   row.BlockTime.Time,
						Direction:   row.Direction,
						ExecutionID: jsonResult.ExecutionID,
						Hash:        row.Hash,
						Module:      "Module",
						Token:       row.Token,
						TokenPrice:  row.TokenPrice,
						Rebated:     false,
					}
				}

				// Now you can pass stipTransactions to InsertNewStipTransactions
				if len(stipTransactions) > 0 {
					err = s.db.InsertNewStipTransactions(context.Background(), stipTransactions)
				}
				if err != nil {
					errChan <- err
					fmt.Println("Error inserting new STIP transactions:", err)
					return
				}

				// Confirm that the insert occurred
				stipTransactionsNotRebated, err := s.db.GetSTIPTransactionsNotRebated(context.Background())
				if err != nil {
					errChan <- err
					fmt.Println("Error getting STIP transactions not rebated:", err)
					return
				}
				if len(stipTransactionsNotRebated) == 0 {
					fmt.Println("No STIP transactions found that have not been rebated.")
				} else {
					fmt.Println("Found", len(stipTransactionsNotRebated), "STIP transactions that have not been rebated.")
				}

			}
		}
	}()

	go func() {
		// Query DB to get all STIPs that need to be relayed
		for {
			time.Sleep(5 * time.Second)
			fmt.Println("Relayer submitter event loop")
			stipTransactionsNotRebated, err := s.db.GetSTIPTransactionsNotRebated(context.Background())
			if err != nil {
				fmt.Println("Error getting STIP transactions not rebated:", err)
				continue
			}
			if len(stipTransactionsNotRebated) == 0 {
				fmt.Println("No STIP transactions found that have not been rebated.")
			} else {
				fmt.Println("Found", len(stipTransactionsNotRebated), "STIP transactions that have not been rebated.")
			}

			chainId := s.cfg.ArbChainID
			arbAddress := s.cfg.ArbAddress
			backendClient, err := s.omnirpcClient.GetClient(context.Background(), big.NewInt(int64(chainId)))

			_, err = s.submittter.SubmitTransaction(context.Background(), big.NewInt(int64(chainId)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
				erc20, err := ierc20.NewIERC20(common.HexToAddress(arbAddress), backendClient)
				if err != nil {
					return nil, fmt.Errorf("could not get erc20: %w", err)
				}

				transferTx, err := erc20.Transfer(transactor, s.signer.Address(), big.NewInt(0))
				if err != nil {
					return nil, fmt.Errorf("could not transfer: %w", err)
				}

				return transferTx, nil
			})

			if err != nil {
				fmt.Println("could not submit transfer:", err)
			}
		}
	}()

	err := <-errChan
	if err != nil {
		return err
	}

	// Relayer event loop will live here

	return nil
	// Submit transactions for corresponding rebate
	// Once in submitter, assume we do not need to submit again
	// Update DB to reflect that STIP rebate has been submitted
}
