package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

// CircleAPI is a wrapper for Circle's REST API..
type CircleAPI struct {
	client  *http.Client
	baseURL string
}

const circleAttestationURL = "https://iris-api-sandbox.circle.com/v1/attestations"

// NewCircleAPI creates a new CircleAPI.
func NewCircleAPI() CircleAPI {
	return CircleAPI{
		client:  &http.Client{},
		baseURL: circleAttestationURL,
	}
}

type circleAttestationResponse struct {
	Data struct {
		Attestation string `json:"attestation"`
		Status      string `json:"status"`
	} `json:"data"`
}

// Wrapper for GET /attestations/{txHash}.
func (c CircleAPI) GetAttestation(ctx context.Context, txHash common.Hash) (attestation []byte, err error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, txHash.String())
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		err = fmt.Errorf("could not create request: %w", err)
		return
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	//nolint:errcheck
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var attestationResp circleAttestationResponse
	err = json.Unmarshal(body, &attestationResp)
	if err != nil {
		err = fmt.Errorf("could not unmarshal body: %w", err)
		return
	}

	attestation, err = hex.DecodeString(attestationResp.Data.Attestation)
	if err != nil {
		err = fmt.Errorf("could not decode signature: %w", err)
		return
	}
	return
}
