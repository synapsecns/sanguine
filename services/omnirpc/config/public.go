package config

import (
	// for embedding the config.
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/jftuga/ellipsis"
	"strconv"
)

//go:embed extraRpcs.json
var extraRpcs []byte

// GetPublicRPCConfig gets the rpc map. This should be done at startup time.
func GetPublicRPCConfig() (c Config, err error) {
	c, err = parseConfig(extraRpcs)
	if err != nil {
		return c, fmt.Errorf("could not parse config: %w", err)
	}

	return c, nil
}

// parseConfig parses a chain map from a json payload.
func parseConfig(rawData []byte) (c Config, err error) {
	c = Config{
		Chains: make(map[uint32]ChainConfig),
	}

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

		// skip nil
		if len(rpcArr) == 0 {
			return nil
		}

		// public rpcs always use
		c.Chains[uint32(chainID)] = ChainConfig{
			RPCs:   rpcArr,
			Checks: 1,
		}

		return nil
	})

	if err != nil {
		return c, fmt.Errorf("could not parse rpc map: %w", err)
	}

	return c, nil
}
