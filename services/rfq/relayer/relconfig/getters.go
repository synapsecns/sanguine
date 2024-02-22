package relconfig

import (
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
)

// DefaultChainConfig is the default chain config.
var DefaultChainConfig = ChainConfig{
	DeadlineBufferSeconds: 600,
	OriginGasEstimate:     160000,
	DestGasEstimate:       100000,
	MinGasToken:           "100000000000000000", // 1 ETH
	QuotePct:              100,
	QuoteOffsetBps:        0,
	FixedFeeMultiplier:    1,
}

// getChainConfigValue gets the value of a field from ChainConfig.
// It returns the value from Chains[chainID] if non-zero,
// else from BaseChainConfig if non-zero,
// else from DefaultChainConfig.
func (c Config) getChainConfigValue(chainID int, fieldName string) (interface{}, error) {
	chainConfig, ok := c.Chains[chainID]
	if ok {
		value, err := getFieldValue(chainConfig, fieldName)
		if err != nil {
			return nil, err
		}
		if isNonZero(value) {
			return value, nil
		}
	}

	baseValue, err := getFieldValue(c.BaseChainConfig, fieldName)
	if err != nil {
		return nil, err
	}
	if isNonZero(baseValue) {
		return baseValue, nil
	}

	defaultValue, err := getFieldValue(DefaultChainConfig, fieldName)
	if err != nil {
		return nil, err
	}
	return defaultValue, nil
}

func getFieldValue(obj interface{}, fieldName string) (interface{}, error) {
	if !isChainConfigField(fieldName) {
		return nil, fmt.Errorf("invalid chain config field: %s", fieldName)
	}
	val := reflect.ValueOf(obj)
	fieldVal := val.FieldByName(fieldName)

	if !fieldVal.IsValid() {
		return nil, fmt.Errorf("invalid field: %s", fieldName)
	}

	return fieldVal.Interface(), nil
}

var chainConfigFields map[string]bool

// isChainConfigField returns true if the field is a valid field in ChainConfig.
// We use this to avoid unsafe-reflect-by-name.
func isChainConfigField(fieldName string) bool {
	// Cache the fields for faster lookups.
	if chainConfigFields == nil {
		chainConfigFields = make(map[string]bool)
		for i := 0; i < reflect.TypeOf(ChainConfig{}).NumField(); i++ {
			chainConfigFields[reflect.TypeOf(ChainConfig{}).Field(i).Name] = true
		}
	}

	_, ok := chainConfigFields[fieldName]
	return ok
}

func isNonZero(value interface{}) bool {
	return reflect.ValueOf(value).Interface() != reflect.Zero(reflect.TypeOf(value)).Interface()
}

// GetRFQAddress returns the RFQ address for the given chainID.
func (c Config) GetRFQAddress(chainID int) (value string, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "RFQAddress")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(string)
	if !ok {
		return value, fmt.Errorf("failed to cast RFQAddress to string")
	}
	return value, nil
}

// GetCCTPAddress returns the RFQ address for the given chainID.
func (c Config) GetCCTPAddress(chainID int) (value string, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "CCTPAddress")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(string)
	if !ok {
		return value, fmt.Errorf("failed to cast CCTPAddress to string")
	}
	return value, nil
}

// GetConfirmations returns the Confirmations for the given chainID.
func (c Config) GetConfirmations(chainID int) (value uint64, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "Confirmations")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(uint64)
	if !ok {
		return value, fmt.Errorf("failed to cast Confirmations to int")
	}
	return value, nil
}

// GetNativeToken returns the NativeToken for the given chainID.
func (c Config) GetNativeToken(chainID int) (value string, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "NativeToken")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(string)
	if !ok {
		return value, fmt.Errorf("failed to cast NativeToken to string")
	}
	return value, nil
}

// GetDeadlineBuffer returns the DeadlineBuffer for the given chainID.
func (c Config) GetDeadlineBuffer(chainID int) (seconds time.Duration, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "DeadlineBufferSeconds")
	if err != nil {
		return seconds, err
	}

	value, ok := rawValue.(int)
	if !ok {
		return seconds, fmt.Errorf("failed to cast DeadlineBufferSeconds to int")
	}
	seconds = time.Duration(value) * time.Second
	return seconds, nil
}

// GetOriginGasEstimate returns the OriginGasEstimate for the given chainID.
func (c Config) GetOriginGasEstimate(chainID int) (value int, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "OriginGasEstimate")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(int)
	if !ok {
		return value, fmt.Errorf("failed to cast OriginGasEstimate to int")
	}
	return value, nil
}

// GetDestGasEstimate returns the DestGasEstimate for the given chainID.
func (c Config) GetDestGasEstimate(chainID int) (value int, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "DestGasEstimate")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(int)
	if !ok {
		return value, fmt.Errorf("failed to cast DestGasEstimate to int")
	}
	return value, nil
}

// GetL1FeeChainID returns the L1FeeChainID for the given chainID.
func (c Config) GetL1FeeChainID(chainID int) (value uint32, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "L1FeeChainID")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(uint32)
	if !ok {
		return value, fmt.Errorf("failed to cast L1FeeChainID to int")
	}
	return value, nil
}

// GetL1FeeOriginGasEstimate returns the L1FeeOriginGasEstimate for the given chainID.
func (c Config) GetL1FeeOriginGasEstimate(chainID int) (value int, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "L1FeeOriginGasEstimate")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(int)
	if !ok {
		return value, fmt.Errorf("failed to cast L1FeeOriginGasEstimate to int")
	}
	return value, nil
}

// GetL1FeeDestGasEstimate returns the L1FeeDestGasEstimate for the given chainID.
func (c Config) GetL1FeeDestGasEstimate(chainID int) (value int, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "L1FeeDestGasEstimate")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(int)
	if !ok {
		return value, fmt.Errorf("failed to cast L1FeeDestGasEstimate to int")
	}
	return value, nil
}

// GetMinGasToken returns the MinGasToken for the given chainID.
func (c Config) GetMinGasToken(chainID int) (value *big.Int, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "MinGasToken")
	if err != nil {
		return value, err
	}

	strValue, ok := rawValue.(string)
	if !ok {
		return value, fmt.Errorf("failed to cast MinGasToken to int")
	}

	value, ok = new(big.Int).SetString(strValue, 10)
	if !ok {
		return value, fmt.Errorf("failed to cast MinGasToken to bigint")
	}
	return value, nil
}

// GetQuotePct returns the QuotePct for the given chainID.
func (c Config) GetQuotePct(chainID int) (value float64, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "QuotePct")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(float64)
	if !ok {
		return value, fmt.Errorf("failed to cast QuotePct to int")
	}
	return value, nil
}

// GetQuoteOffsetBps returns the QuoteOffsetBps for the given chainID.
func (c Config) GetQuoteOffsetBps(chainID int) (value float64, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "QuoteOffsetBps")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(float64)
	if !ok {
		return value, fmt.Errorf("failed to cast QuoteOffsetBps to int")
	}
	if value <= 0 {
		value = DefaultChainConfig.QuoteOffsetBps
	}
	return value, nil
}

// GetFixedFeeMultiplier returns the FixedFeeMultiplier for the given chainID.
func (c Config) GetFixedFeeMultiplier(chainID int) (value float64, err error) {
	rawValue, err := c.getChainConfigValue(chainID, "FixedFeeMultiplier")
	if err != nil {
		return value, err
	}

	value, ok := rawValue.(float64)
	if !ok {
		return value, fmt.Errorf("failed to cast FixedFeeMultiplier to int")
	}
	if value <= 0 {
		value = DefaultChainConfig.FixedFeeMultiplier
	}
	return value, nil
}

// GetL1FeeParams returns the L1 fee params for the given chain.
func (c Config) GetL1FeeParams(chainID uint32, origin bool) (uint32, int, bool) {
	var gasEstimate int
	var err error
	if origin {
		gasEstimate, err = c.GetL1FeeOriginGasEstimate(int(chainID))
		if err != nil {
			return 0, 0, false
		}
	} else {
		gasEstimate, err = c.GetL1FeeDestGasEstimate(int(chainID))
		if err != nil {
			return 0, 0, false
		}
	}

	l1FeeChainID, err := c.GetL1FeeChainID(int(chainID))
	if err != nil || l1FeeChainID <= 0 || gasEstimate <= 0 {
		return 0, 0, false
	}
	return l1FeeChainID, gasEstimate, true
}

// GetChains returns the chains config.
func (c Config) GetChains() map[int]ChainConfig {
	return c.Chains
}

// GetOmniRPCURL returns the OmniRPCURL.
func (c Config) GetOmniRPCURL() string {
	return c.OmniRPCURL
}

// GetRfqAPIURL returns the RFQ API URL.
func (c Config) GetRfqAPIURL() string {
	return c.RfqAPIURL
}

// GetDatabase returns the database config.
func (c Config) GetDatabase() DatabaseConfig {
	return c.Database
}

// GetSigner returns the signer config.
func (c Config) GetSigner() config.SignerConfig {
	return c.Signer
}

// GetFeePricer returns the fee pricer config.
func (c Config) GetFeePricer() FeePricerConfig {
	return c.FeePricer
}

const defaultHTTPTimeoutMs = 1000

// GetHTTPTimeout returns the HTTP timeout.
func (c Config) GetHTTPTimeout() time.Duration {
	feePricerCfg := c.GetFeePricer()
	timeoutMs := defaultHTTPTimeoutMs
	if feePricerCfg.HTTPTimeoutMs > 0 {
		timeoutMs = feePricerCfg.HTTPTimeoutMs
	}
	return time.Duration(timeoutMs) * time.Millisecond
}

func (c Config) getTokenConfigByAddr(chainID int, tokenAddr string) (cfg TokenConfig, name string, err error) {
	chainConfig, ok := c.Chains[chainID]
	if !ok {
		return cfg, name, fmt.Errorf("no chain config for chain %d", chainID)
	}
	for tokenName, tokenConfig := range chainConfig.Tokens {
		if common.HexToAddress(tokenConfig.Address).Hex() == common.HexToAddress(tokenAddr).Hex() {
			return tokenConfig, tokenName, nil
		}
	}
	return cfg, name, fmt.Errorf("no token config for chain %d and address %s", chainID, tokenAddr)
}

// GetRebalanceMethod returns the rebalance method for the given chain and token address.
func (c Config) GetRebalanceMethod(chainID int, tokenAddr string) (method RebalanceMethod, err error) {
	tokenConfig, tokenName, err := c.getTokenConfigByAddr(chainID, tokenAddr)
	if err != nil {
		return 0, err
	}
	for cid, chainCfg := range c.Chains {
		tokenCfg, ok := chainCfg.Tokens[tokenName]
		if ok {
			if tokenConfig.RebalanceMethod != tokenCfg.RebalanceMethod {
				return RebalanceMethodNone, fmt.Errorf("rebalance method mismatch for token %s on chains %d and %d", tokenName, chainID, cid)
			}
		}
	}
	switch tokenConfig.RebalanceMethod {
	case "cctp":
		return RebalanceMethodCCTP, nil
	case "native":
		return RebalanceMethodNative, nil
	}
	return RebalanceMethodNone, nil
}

// GetRebalanceMethods returns all rebalance methods present in the config.
func (c Config) GetRebalanceMethods() (methods map[RebalanceMethod]bool, err error) {
	methods = make(map[RebalanceMethod]bool)
	for chainID, chainCfg := range c.Chains {
		for _, tokenCfg := range chainCfg.Tokens {
			method, err := c.GetRebalanceMethod(chainID, tokenCfg.Address)
			if err != nil {
				return nil, err
			}
			if method != RebalanceMethodNone {
				methods[method] = true
			}
		}
	}
	return methods, nil
}

// GetMaintenanceBalancePct returns the maintenance balance percentage for the given chain and token address.
func (c Config) GetMaintenanceBalancePct(chainID int, tokenAddr string) (float64, error) {
	tokenConfig, _, err := c.getTokenConfigByAddr(chainID, tokenAddr)
	if err != nil {
		return 0, err
	}
	if tokenConfig.MaintenanceBalancePct <= 0 {
		return 0, fmt.Errorf("maintenance balance pct must be positive: %f", tokenConfig.MaintenanceBalancePct)
	}
	return tokenConfig.MaintenanceBalancePct, nil
}

// GetInitialBalancePct returns the initial balance percentage for the given chain and token address.
func (c Config) GetInitialBalancePct(chainID int, tokenAddr string) (float64, error) {
	tokenConfig, _, err := c.getTokenConfigByAddr(chainID, tokenAddr)
	if err != nil {
		return 0, err
	}
	if tokenConfig.InitialBalancePct <= 0 {
		return 0, fmt.Errorf("initial balance pct must be positive: %f", tokenConfig.InitialBalancePct)
	}
	return tokenConfig.InitialBalancePct, nil
}

// GetTokenID returns the tokenID for the given chain and address.
func (c Config) GetTokenID(chain int, addr string) (string, error) {
	chainConfig, ok := c.Chains[chain]
	if !ok {
		return "", fmt.Errorf("no chain config for chain %d", chain)
	}
	for tokenID, tokenConfig := range chainConfig.Tokens {
		if tokenConfig.Address == addr {
			return tokenID, nil
		}
	}
	return "", fmt.Errorf("no tokenID found for chain %d and address %s", chain, addr)
}

// GetQuotableTokens returns the quotable tokens for the given token.
func (c Config) GetQuotableTokens(token string) ([]string, error) {
	tokens, ok := c.QuotableTokens[token]
	if !ok {
		return nil, fmt.Errorf("no quotable tokens for token %s", token)
	}
	return tokens, nil
}

// GetTokenDecimals returns the token decimals for the given chain and token.
func (c Config) GetTokenDecimals(chainID uint32, token string) (uint8, error) {
	chainConfig, ok := c.Chains[int(chainID)]
	if !ok {
		return 0, fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	for tokenName, tokenConfig := range chainConfig.Tokens {
		if token == tokenName {
			return tokenConfig.Decimals, nil
		}
	}
	return 0, fmt.Errorf("could not get token decimals for chain %d and token %s", chainID, token)
}

// GetTokens returns the tokens for the given chain.
func (c Config) GetTokens(chainID uint32) (map[string]TokenConfig, error) {
	chainConfig, ok := c.Chains[int(chainID)]
	if !ok {
		return nil, fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	return chainConfig.Tokens, nil
}

// GetTokenName returns the token name for the given chain and address.
func (c Config) GetTokenName(chain uint32, addr string) (string, error) {
	chainConfig, ok := c.Chains[int(chain)]
	if !ok {
		return "", fmt.Errorf("no chain config for chain %d", chain)
	}
	for tokenName, tokenConfig := range chainConfig.Tokens {
		if common.HexToAddress(tokenConfig.Address).Hex() == common.HexToAddress(addr).Hex() {
			return tokenName, nil
		}
	}
	return "", fmt.Errorf("no tokenName found for chain %d and address %s", chain, addr)
}

const defaultMinQuoteAmount = 0

// GetMinQuoteAmount returns the quote amount for the given chain and address.
// Note that this getter returns the value in native token decimals.
func (c Config) GetMinQuoteAmount(chainID int, addr common.Address) *big.Int {
	chainCfg, ok := c.Chains[chainID]
	if !ok {
		return big.NewInt(defaultMinQuoteAmount)
	}

	var tokenCfg *TokenConfig
	for _, cfg := range chainCfg.Tokens {
		if common.HexToAddress(cfg.Address).Hex() == addr.Hex() {
			cfgCopy := cfg
			tokenCfg = &cfgCopy
			break
		}
	}
	if tokenCfg == nil {
		return big.NewInt(defaultMinQuoteAmount)
	}
	quoteAmountFlt, ok := new(big.Float).SetString(tokenCfg.MinQuoteAmount)
	if !ok {
		return big.NewInt(defaultMinQuoteAmount)
	}
	if quoteAmountFlt.Cmp(big.NewFloat(0)) <= 0 {
		return big.NewInt(defaultMinQuoteAmount)
	}

	// Scale the minQuoteAmount by the token decimals.
	denomDecimalsFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenCfg.Decimals)), nil)
	quoteAmountScaled, _ := new(big.Float).Mul(quoteAmountFlt, new(big.Float).SetInt(denomDecimalsFactor)).Int(nil)
	return quoteAmountScaled
}

const defaultDBSelectorIntervalSeconds = 1

// GetDBSelectorInterval returns the interval for the DB selector.
func (c Config) GetDBSelectorInterval() time.Duration {
	interval := c.DBSelectorIntervalSeconds
	if interval <= 0 {
		return defaultDBSelectorIntervalSeconds
	}
	return time.Duration(interval) * time.Second
}

const defaultRebalanceIntervalSeconds = 30

// GetRebalanceInterval returns the interval for rebalancing.
func (c Config) GetRebalanceInterval() time.Duration {
	interval := c.RebalanceInterval
	if interval == 0 {
		interval = defaultRebalanceIntervalSeconds
	}
	return interval
}
