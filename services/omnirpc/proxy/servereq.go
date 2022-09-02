package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gin-gonic/gin"
	resty "github.com/go-resty/resty/v2"
	"github.com/hedzr/cmdr/tool"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"
)

// RPCRequest is a raw rpc request format.
type RPCRequest struct {
	ID     json.RawMessage   `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

func parseRPCPayload(body []byte) (request RPCRequest, err error) {
	rpcPayload := RPCRequest{}
	err = json.Unmarshal(body, &rpcPayload)
	if err != nil {
		return RPCRequest{}, errors.Wrap(err, "failed to parse json RPC payload")
	}

	return rpcPayload, nil
}

func isBlockNumConfirmable(arg json.RawMessage) bool {
	// nonConfirmableBlockNumArgs is a list of non numerical block args
	var nonConfirmableBlockNumArgs = []string{"latest", "pending"}

	return !slices.Contains(nonConfirmableBlockNumArgs, tool.StripQuotes(string(arg)))
}

// isFilterArgConfirmable checks if filter.filterCriteria is confirmable.
func isFilterArgConfirmable(arg json.RawMessage) (bool, error) {
	// cast latest block number to a big int for comparison
	latestBlockNumber := new(big.Int).SetInt64(rpc.LatestBlockNumber.Int64())

	filterCriteria := filters.FilterCriteria{}
	err := filterCriteria.UnmarshalJSON(arg)
	if err != nil {
		return false, fmt.Errorf("could not unmarshall filter: %w", err)
	}

	// Block filter requested, construct a single-shot filter
	if filterCriteria.BlockHash != nil {
		return true, nil
	}

	usesLatest := filterCriteria.FromBlock.Cmp(latestBlockNumber) == 0 || filterCriteria.ToBlock.Cmp(latestBlockNumber) == 0
	return !usesLatest, nil
}

// nolint: cyclop
func isConfirmable(body []byte) (bool, error) {
	payload, err := parseRPCPayload(body)
	if err != nil {
		return false, fmt.Errorf("could not parse payload: %w", err)
	}

	// TODO: handle batch methods
	// TODO: should we error on default?
	switch payload.Method {
	case "eth_getBlockByNumber":
		return isBlockNumConfirmable(payload.Params[0]), nil
	case "eth_blockNumber":
		return false, nil
	case "eth_syncing":
		return false, nil
	case "eth_getBlockTransactionCountByNumber":
		return isBlockNumConfirmable(payload.Params[0]), nil
	case "eth_getBalance":
		return isBlockNumConfirmable(payload.Params[1]), nil
	case "eth_getStorageAt":
		return isBlockNumConfirmable(payload.Params[2]), nil
	case "eth_getCode":
		return isBlockNumConfirmable(payload.Params[1]), nil
	case "eth_getTransactionCount":
		return isBlockNumConfirmable(payload.Params[1]), nil
	case "eth_getLogs":
		return isFilterArgConfirmable(payload.Params[0])
	case "eth_call":
		return isBlockNumConfirmable(payload.Params[1]), nil
	case "eth_gasPrice":
		return false, nil
	case "eth_maxPriorityFeePerGas":
		return false, nil
	case "eth_estimateGas":
		return false, nil
	// not confirmable because tx could be pending. We might want to handle w/ omnicast though
	case "eth_sendRawTransaction":
		return false, nil
	}
	return true, nil
}

func (r *RPCProxy) serveRPCReq(c *gin.Context, chainID int) {
	rpcList := r.rpcMap.ChainID(chainID)

	if len(rpcList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("no endpoint for chain %d", chainID),
		})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}

	for _, endpoint := range rpcList {
		endpointURL, err := url.Parse(endpoint)
		if err != nil {
			continue
		}

		// websockets not yet supported
		if !slices.Contains([]string{"http", "https"}, endpointURL.Scheme) {
			continue
		}

		client := resty.New()
		resp, err := client.R().
			SetContext(c).
			SetBody(body).
			Post(endpoint)

		if err != nil {
			// continue until we exhaust endpoints
			continue
		}

		if resp.StatusCode() < 200 || resp.StatusCode() > 400 {
			// error
			continue
		}

		c.Header("x-forwarded-from", endpoint)
		c.Data(http.StatusOK, gin.MIMEJSON, resp.Body())
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf("no rpc online for chain %d, attempted: %s", chainID, strings.Join(rpcList, ",")),
	})
}
