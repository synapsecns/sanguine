package stip_relayer_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/synapsecns/sanguine/services/stip_relayer"
)

func TestExecuteDuneQuery(t *testing.T) {
	resp, err := stip_relayer.ExecuteDuneQuery()
	if err != nil {
		t.Fatalf("Failed to execute Dune query: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if len(body) == 0 {
		t.Error("Expected non-empty response body, got empty")
	}

	fmt.Println(string(body))
}

func TestGetExecutionResults(t *testing.T) {
	resp, err := stip_relayer.ExecuteDuneQuery()
	if err != nil {
		t.Fatalf("Failed to execute Dune query: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	execution_id, ok := result["execution_id"]
	if !ok {
		t.Fatal("No execution_id found in the response")
	}

	time.Sleep(20000 * time.Millisecond)

	resp, err = stip_relayer.GetExecutionResults(execution_id)
	if err != nil {
		t.Fatalf("Failed to get execution results: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	getResultsBody, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(getResultsBody))

}

func (c *STIPRelayerSuite) TestStartRelayer() {
	c.stipRelayer.Run()
}
