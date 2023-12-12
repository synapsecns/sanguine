package config_test

import (
	"testing"

	"github.com/synapsecns/sanguine/rfq/rfq-relayer/config"

	"github.com/stretchr/testify/assert"
)

// TODO we can write better tests here.
func TestLoadConfig(t *testing.T) {
	config, err := config.DecodeConfig("../config.yaml") // Assuming the config.yaml is in the parent directory of the config package
	assert.NoError(t, err, "Failed to load config")

	// Test for chain 42
	chain42, exists42 := config.Chains[42]
	assert.True(t, exists42, "Chain 42 does not exist in config")

	assert.Equal(t, uint32(42), chain42.ChainID, "ChainID does not match for chain 42")
	assert.Equal(t, "http://localhost:8042", chain42.RpcUrl, "RpcUrl does not match for chain 42")
	assert.Equal(t, "0x6438CB36cb18520774EfC7A172410D8BBBe9a428", chain42.FastBridgeAddress, "FastBridgeContract does not match for chain 42")

	// Test for chain 43
	chain43, exists43 := config.Chains[43]
	assert.True(t, exists43, "Chain 43 does not exist in config")

	assert.Equal(t, uint32(43), chain43.ChainID, "ChainID does not match for chain 43")
	assert.Equal(t, "http://localhost:8043", chain43.RpcUrl, "RpcUrl does not match for chain 43")
	assert.Equal(t, "0x6438CB36cb18520774EfC7A172410D8BBBe9a428", chain43.FastBridgeAddress, "FastBridgeContract does not match for chain 43")

	// Test for database configuration
	assert.Equal(t, "sqlite", config.Database.Type, "Database type does not match")
	assert.Equal(t, "rfq_solver_bot.db", config.Database.DSN, "Database DSN does not match")

	assert.Equal(t, "0xa01474A3b29535A90Cbd05912af81a4fcc5276Ef", config.RelayerAddress, "RelayerAddress does not match")

	assert.Equal(t, "localhost:0000", config.OmnirpcURL, "OmniRPC does not match")
}
