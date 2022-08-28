package config

import (
	"fmt"
	"github.com/synapsecns/sanguine/serivces/omnirpc/rpcmap"
	"gopkg.in/yaml.v2"
	"os"
)

// UnmarshallConfig unmarshalls an rpc config from an input.
func UnmarshallConfig(input string) (*rpcmap.RPCMap, error) {
	var rawMap map[int][]string
	err := yaml.Unmarshal([]byte(input), &rawMap)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall rpc map: %w", err)
	}
	return rpcmap.NewRPCMapFromMap(rawMap), nil
}

// UnmarshallConfigFromFile gets a config from a file.
func UnmarshallConfigFromFile(file string) (*rpcmap.RPCMap, error) {
	//nolint: gosec
	contents, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	return UnmarshallConfig(string(contents))
}

// MarshallFromMap marshalls a config from an rpc map.
func MarshallFromMap(rpcMap *rpcmap.RPCMap) string {
	// errors are impossible here
	output, _ := yaml.Marshal(rpcMap.RawMap())

	return string(output)
}
