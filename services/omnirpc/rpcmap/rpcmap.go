package rpcmap

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/jftuga/ellipsis"
	backoffHelper "github.com/jpillora/backoff"
	"gitlab.com/1f320/x/duration"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// RPCMap maps [chainid]->list of rpcs for that chain.
// RPCMap is thread safe.
type RPCMap struct {
	rpcs map[int][]string
	mux  sync.RWMutex
}

// GetChainIDs gets all chainids.
func (r *RPCMap) GetChainIDs() (chainIDs []int) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	for key := range r.rpcs {
		chainIDs = append(chainIDs, key)
	}
	return chainIDs
}

// RawMap gets a copy of the underlying rpc map.
// this function makes a fully copy.
func (r *RPCMap) RawMap() (res map[int][]string) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	res = make(map[int][]string)

	for key, value := range r.rpcs {
		res[key] = value
	}
	return res
}

// ChainID gets all rpc urls for a given chainid.
func (r *RPCMap) ChainID(chainID int) []string {
	r.mux.RLock()
	defer r.mux.RUnlock()

	return r.rpcs[chainID]
}

// PutChainID overwrites the existing slice for the chain id.
func (r *RPCMap) PutChainID(chainID int, newSlice []string) {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.rpcs[chainID] = newSlice
}

// NewRPCMap returns an empty rpc map.
func NewRPCMap() *RPCMap {
	return &RPCMap{
		rpcs: make(map[int][]string),
		mux:  sync.RWMutex{},
	}
}

// NewRPCMapFromMap creates a new rpc map from a raw map.
func NewRPCMapFromMap(rawMap map[int][]string) *RPCMap {
	return &RPCMap{
		rpcs: rawMap,
		mux:  sync.RWMutex{},
	}
}

// PublicRPCMapURL is the url we pull the rpc list from.
// TODO: this has some rate-limits, they're relatively aggressive but something like gitcdn.xyz would be good here.
const PublicRPCMapURL = "https://raw.githubusercontent.com/DefiLlama/chainlist/master/constants/extraRpcs.json"

// GetPublicRPCMap gets the rpc map from defillama. This should be done at startup time.
// this will retry on a backoffHelper until context cancellation.
func GetPublicRPCMap(ctx context.Context) (m *RPCMap, err error) {
	backoff := &backoffHelper.Backoff{
		Factor: 1.3,
		Jitter: true,
		Min:    time.Second * 1,
		Max:    time.Second * 10,
	}

	var waitTime time.Duration
	httpClient := &http.Client{}
	for {
		select {
		case <-ctx.Done():
			return NewRPCMap(), fmt.Errorf("could not get rpc map: %w", ctx.Err())
		case <-time.After(waitTime):
			waitTime = backoff.Duration()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, PublicRPCMapURL, nil)
			if err != nil {
				return NewRPCMap(), fmt.Errorf("could not create request: %w", err)
			}

			resp, err := httpClient.Do(req)
			if err != nil {
				logger.Errorf("could not retrieve rpc list from %s, waiting %s before trying again (error: %v)", PublicRPCMapURL, duration.Format(waitTime), err)
				continue
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Errorf("could not read body from %s, waiting %s before trying again (error: %v)", PublicRPCMapURL, duration.Format(waitTime), err)
				continue
			}

			_ = resp.Body.Close()

			m, err = parseRPCMap(body)
			if err != nil {
				logger.Error(err)
				continue
			}

			return m, nil
		}
	}
}

// parseRPCMap parses a chain map from a json payload.
func parseRPCMap(rawData []byte) (m *RPCMap, err error) {
	m = NewRPCMap()

	// iterate over chain ids which are strings
	err = jsonparser.ObjectEach(rawData, func(key []byte, value []byte, dataType jsonparser.ValueType, _ int) error {
		if dataType != jsonparser.Object {
			return fmt.Errorf("expected %s got %s when parsing %s", jsonparser.Object, dataType, ellipsis.Shorten(string(rawData), 10))
		}

		chainID, err := strconv.Atoi(string(key))
		if err != nil {
			return fmt.Errorf("could not parse %s to int", key)
		}

		rawRPCList, dataType, _, err := jsonparser.Get(value, "rpcs")
		// skip this key
		if err != nil {
			logger.Debugf("key %s missing", key)
			//nolint: nilerr
			return nil
		}

		if dataType != jsonparser.Array {
			return fmt.Errorf("expected %s got %s when parsing %s", jsonparser.Array, dataType, ellipsis.Shorten(string(value), 10))
		}

		var rpcArr []string
		err = json.Unmarshal(rawRPCList, &rpcArr)
		if err != nil {
			return fmt.Errorf("could not unmarshal array: %w", err)
		}

		m.rpcs[chainID] = rpcArr

		return nil
	})

	if err != nil {
		return &RPCMap{}, fmt.Errorf("could not parse rpc map: %w", err)
	}

	return m, nil
}
