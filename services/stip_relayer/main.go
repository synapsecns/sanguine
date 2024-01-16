package stip_relayer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/synapsecns/sanguine/core/metrics"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/stip_relayer/db"
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

	return resp, nil
}

func main() {
	ExecuteDuneQuery()
}

// QuoterAPIServer is a struct that holds the configuration, database connection, gin engine, RPC client, metrics handler, and fast bridge contracts.
// It is used to initialize and run the API server.
type STIPRelayer struct {
	cfg           config.Config
	db            db.STIPDB
	omnirpcClient omniClient.RPCClient
	handler       metrics.Handler
}

func NewSTIPRelayer(ctx context.Context,
	cfg config.Config,
	handler metrics.Handler,
	omniRPCClient omniClient.RPCClient,
	store db.STIPDB,
) (*STIPRelayer, error) {
	return &STIPRelayer{
		cfg:           cfg,
		db:            store,
		handler:       handler,
		omnirpcClient: omniRPCClient,
	}, nil
}

func (s STIPRelayer) Run() error {
	ticker := time.NewTicker(5 * time.Minute)
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
				fmt.Println(executionResults)

				// Store executionResults in DB
				// This part is left as a comment because it depends on the specific implementation of your DB
				// s.db.StoreExecutionResults(executionResults)
			}
		}
	}()

	err := <-errChan
	if err != nil {
		return err
	}

	// Relayer event loop will live here

	return nil

	// Call ExecuteDuneQuery, wait 20 seconds
	// Call GetExecutionResults
	// Store Execution Results in DB

	// Query DB to get all STIPs that need to be relayed
	// Submit transactions for corresponding rebate
	// Once in submitter, assume we do not need to submit again
	// Update DB to reflect that STIP rebate has been submitted
}
