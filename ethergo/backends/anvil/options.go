package anvil

import (
	"errors"
	"fmt"
	"github.com/ImVexed/fasturl"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ory/dockertest/v3/docker"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/exp/slices"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func _() {
	// make sure we don't panic
	NewAnvilOptionBuilder()
}

const defaultMnemonic = "tag volcano eight thank tide danger coast health above argue embrace heavy"

// NewAnvilOptionBuilder creates a new option builder.
func NewAnvilOptionBuilder() *OptionBuilder {
	optionsBuilder := &OptionBuilder{}
	// set default non-anvil values
	optionsBuilder.SetMaxWaitTime(time.Minute * 2)
	optionsBuilder.SetRestartPolicy(Autoremove)
	optionsBuilder.SetExpirySeconds(600)

	optionsBuilder.SetAccounts(10)
	optionsBuilder.SetBlockTime(0)
	optionsBuilder.SetBalance(10000)
	optionsBuilder.SetGasLimit(gasLimit * 5)

	optionsBuilder.SetDerivationPath(accounts.DefaultBaseDerivationPath)
	// note: ordinarily we strongly discourage using panics, but the setup should be deterministic.
	// to prevent unintended behavior, we have a boot time static assertion check above to make sure this will not panic
	err := optionsBuilder.SetHardfork(Latest)
	if err != nil {
		panic(err)
	}
	err = optionsBuilder.SetOrder(Fifo)
	if err != nil {
		panic(err)
	}

	optionsBuilder.SetBaseFee(1)
	optionsBuilder.SetStepsTracing(true)
	optionsBuilder.SetTimestamp(0)
	optionsBuilder.SetAllowOrigin("*")
	optionsBuilder.SetHost("0.0.0.0")
	err = optionsBuilder.SetMnemonic(defaultMnemonic)
	if err != nil {
		panic(err)
	}

	return optionsBuilder
}

// tagName is the tag name used for anvil options building.
const tagName = "anvil"

type evmOptions struct {
	ForkURL                   string `anvil:"fork-url"`
	BlockNumber               uint64 `anvil:"fork-block-number"`
	ForkRetryBackoff          uint   `anvil:"fork-retry-backoff"`
	Retries                   uint16 `anvil:"retries"`
	Timeout                   uint64 `anvil:"timeout"`
	ComputeUnitsPerSecond     uint64 `anvil:"compute-units-per-second"`
	NoRateLimitUnitsPerSecond bool   `anvil:"no-rate-limit-units-per-second"`
	NoStorageCaching          bool   `anvil:"no-storage-caching"`
}

type generalOptions struct {
	Accounts       uint8                   `anvil:"accounts"`
	BlockTime      uint64                  `anvil:"block-time"`
	Balance        uint64                  `anvil:"balance"`
	DerivationPath accounts.DerivationPath `anvil:"derivation-path"`
	Hardfork       Hardfork                `anvil:"hardfork"`
	Order          Order                   `anvil:"order"`
	StepsTracing   bool                    `anvil:"steps-tracing"`
	Silent         bool                    `anvil:"silent"`
	Timestamp      uint64                  `anvil:"timestamp"`
	Mnemonic       string                  `anvil:"mnemonic"`
	NoMining       bool                    `anvil:"no-mining"`
}

type executorOptions struct {
	BaseFee            uint64 `anvil:"base-fee"`
	BlockBaseFeePerGas uint64 `anvil:"block-base-fee-per-gas"`
	ChainID            uint64 `anvil:"chain-id"`
	CodeSizeLimit      uint64 `anvil:"code-size-limit"`
	GasLimit           uint64 `anvil:"gas-limit"`
	GasPrice           uint64 `anvil:"gas-price"`
}

type serverOptions struct {
	AllowOrigin  string `anvil:"allow-origin"`
	disableCors  bool   `anvil:"disable-cors"`
	PruneHistory bool   `anvil:"prune-history"`
	Host         string `anvil:"host"`
}

// nonAnvilOptions options are options that are not part of the anvil spec, but are used by the anvil backend.
// it's important that these do not include the anvil annotation.
type nonAnvilOptions struct {
	// enableOtterscan specifies wether or not to start a container for otterscan
	enableOtterscan bool
	// maxWait is the maximum time to wait for the server to start
	maxWait time.Duration
	// expirySeconds is the number of seconds to wait before expiring a request
	// set to 0 to disable
	expirySeconds uint
	// autoremove is wether the container should be deleted after the run
	autoremove bool
	// restartPolicy restarts a policy
	restartPolicy *docker.RestartPolicy
	// processOptions are options for the process logger
	processOptions []processlog.StdStreamLogArgsOption
}

// OptionBuilder is a builder for anvil options.
type OptionBuilder struct {
	generalOptions
	evmOptions
	executorOptions
	serverOptions
	nonAnvilOptions
}

// ========= General Options =========

// OtterscanEnabled sets whether or not to enable otterscan.
// Note: for this to work correctly, the underlying node we fork from must also support the ots namespace.
func (o *OptionBuilder) OtterscanEnabled(enabled bool) {
	o.enableOtterscan = enabled
}

// SetAccounts sets the number of Accounts to create (defaults to 10).
func (o *OptionBuilder) SetAccounts(accountCount uint8) {
	o.Accounts = accountCount
}

// GetAccounts gets the number of accounts to use.
func (o *OptionBuilder) GetAccounts() uint8 {
	return o.Accounts
}

// SetBlockTime sets the block time (defaults to 1 second).
// if block time is 0 or less than a second, no mining is used.
func (o *OptionBuilder) SetBlockTime(blockTime time.Duration) {
	o.BlockTime = uint64(blockTime / time.Second)
}

// SetBalance sets the Balance of each account in ether (defaults to 1000 ether).
func (o *OptionBuilder) SetBalance(balance uint64) {
	o.Balance = balance
}

// SetNoMining sets whether or not to mine blocks.
// this is automatically set and must be disabled for other options to be sent.
// if disabled, this can cause concurrency issues since only one account can be impersonated at a time.
func (o *OptionBuilder) SetNoMining(noMining bool) {
	o.NoMining = noMining
}

// SetDerivationPath sets the derivation path to use for the Accounts.
// defaults to m/44'/60'/0'/0.
func (o *OptionBuilder) SetDerivationPath(derivationPath accounts.DerivationPath) {
	o.DerivationPath = derivationPath
}

// GetDerivationPath returns the derivation path to use for the Accounts.
func (o *OptionBuilder) GetDerivationPath() accounts.DerivationPath {
	return o.DerivationPath
}

// SetHardfork sets the Hardfork to use for the chain.
func (o *OptionBuilder) SetHardfork(hardfork Hardfork) error {
	if !slices.Contains(allHardforks, hardfork) {
		return fmt.Errorf("invalid Hardfork: %s", hardfork.String())
	}
	o.Hardfork = hardfork
	return nil
}

// GetHardfork returns the Hardfork to use for the chain.
func (o *OptionBuilder) GetHardfork() Hardfork {
	return o.Hardfork
}

// GetMnemonic returns the mnemonic to use for the chain.
func (o *OptionBuilder) GetMnemonic() string {
	return o.Mnemonic
}

// SetInit sets the genesis file to use for the chain.
// Warning: this option is not currently supported.
func (o *OptionBuilder) SetInit(_ string) error {
	return fmt.Errorf("unsupported option (not yet implemented): %s", getFunctionName(o.SetInit))
}

// SetMnemonic sets the mnemonic to use for the chain.
func (o *OptionBuilder) SetMnemonic(mnemonic string) error {
	o.Mnemonic = mnemonic
	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("invalid mnemonic: %s", mnemonic)
	}

	return nil
}

// SetOrder sets the transaction Order to use for the chain.
// Defaults to Fees.
func (o *OptionBuilder) SetOrder(order Order) error {
	if !slices.Contains(allOrders, order) {
		return fmt.Errorf("invalid Order: %s", order.String())
	}
	o.Order = order
	return nil
}

// SetStepsTracing enables or disables steps tracing.
// Defaults to enabled.
func (o *OptionBuilder) SetStepsTracing(enabled bool) {
	o.StepsTracing = enabled
}

// SetIPC sets the IPC path to use for the chain.
// This option is not currently supported.
func (o *OptionBuilder) SetIPC(path string) error {
	return fmt.Errorf("unsupported option (not yet implemented): %s", getFunctionName(o.SetIPC))
}

// SetSilent enables or disables Silent boot.
func (o *OptionBuilder) SetSilent(silent bool) {
	o.Silent = silent
}

// SetTimestamp sets the Timestamp of the genesis block. This option is ignored if set to 0.
func (o *OptionBuilder) SetTimestamp(timestamp uint64) {
	o.Timestamp = timestamp
}

// ========= EVM Options =========

// SetForkURL sets the fork URL to use for the chain.
func (o *OptionBuilder) SetForkURL(url string) error {
	_, err := fasturl.ParseURL(url)
	if err != nil {
		return fmt.Errorf("invalid fork url: %s", url)
	}
	o.ForkURL = url
	return nil
}

// SetForkBlockNumber sets the fork block number to use for the chain.
// only valid if SetForkURL is also called.
func (o *OptionBuilder) SetForkBlockNumber(block uint64) {
	o.BlockNumber = block
}

// SetForkRetryBackoff sets the fork retry backoff to use for the chain.
func (o *OptionBuilder) SetForkRetryBackoff(backoff uint) {
	o.ForkRetryBackoff = backoff
}

// SetRetries sets the number of retries to use for forked rpc requests.
func (o *OptionBuilder) SetRetries(retries uint16) {
	o.Retries = retries
}

// SetTimeout sets the timeout to use for forked rpc requests in milliseconds.
func (o *OptionBuilder) SetTimeout(timeout time.Duration) {
	o.Timeout = uint64(timeout / time.Millisecond)
}

// SetComputeUnitsPerSecond sets the CUPS per second, see: https://docs.alchemy.com/reference/compute-units
func (o *OptionBuilder) SetComputeUnitsPerSecond(units uint64) {
	o.ComputeUnitsPerSecond = units
}

// SetRateLimitUnitsPerSecondEnabled enables or disables rate limiting of CUPS.
func (o *OptionBuilder) SetRateLimitUnitsPerSecondEnabled(enabled bool) {
	o.NoRateLimitUnitsPerSecond = enabled
}

// SetStorageCachingEnabled enables or disables storage caching.
func (o *OptionBuilder) SetStorageCachingEnabled(enabled bool) {
	o.NoStorageCaching = enabled
}

// SetBaseFee sets the base fee for the chain.
func (o *OptionBuilder) SetBaseFee(baseFee uint64) {
	o.BaseFee = baseFee
}

// SetBlockBaseFeePerGas sets the block base fee per gas for the chain.
func (o *OptionBuilder) SetBlockBaseFeePerGas(baseFee uint64) {
	o.BlockBaseFeePerGas = baseFee
}

// SetChainID sets the chain id to use in foundry.
func (o *OptionBuilder) SetChainID(chainID uint64) {
	o.ChainID = chainID
}

// SetCodeSizeLimit sets the code size limit for the chain.
func (o *OptionBuilder) SetCodeSizeLimit(codeSizeLimit uint64) {
	o.CodeSizeLimit = codeSizeLimit
}

// SetGasLimit sets the gas limit per block for the chain.
func (o *OptionBuilder) SetGasLimit(gasLimit uint64) {
	o.GasLimit = gasLimit
}

// SetGasPrice sets the gas price for the chain.
func (o *OptionBuilder) SetGasPrice(gasPrice uint64) {
	o.GasPrice = gasPrice
}

// SetAllowOrigin sets the allow origin for the chain.
// defaults to "*".
func (o *OptionBuilder) SetAllowOrigin(allowOrigin string) {
	o.AllowOrigin = allowOrigin
}

// SetDisableCors disables cors for the chain.
func (o *OptionBuilder) SetDisableCors(disableCors bool) {
	o.disableCors = disableCors
}

// SetPruneHistory enables or disables pruning of history.
// defaults to false.
func (o *OptionBuilder) SetPruneHistory(pruneHistory bool) {
	o.PruneHistory = pruneHistory
}

// SetHost sets the host to listen on.
func (o *OptionBuilder) SetHost(host string) {
	o.Host = host
}

// SetMaxWaitTime sets the max wait time for the docker container(s) to start.
func (o *OptionBuilder) SetMaxWaitTime(maxWait time.Duration) {
	o.maxWait = maxWait
}

// SetExpirySeconds sets the expiry seconds for the docker containers(s) to be removed.
func (o *OptionBuilder) SetExpirySeconds(expirySeconds uint) {
	o.expirySeconds = expirySeconds
}

// SetProcessLogOptions sets the process log options for the docker container(s).
func (o *OptionBuilder) SetProcessLogOptions(customOpts ...processlog.StdStreamLogArgsOption) {
	o.processOptions = customOpts
}

// SetRestartPolicy sets the restart policy for the docker container(s).
func (o *OptionBuilder) SetRestartPolicy(restartPolicy RestartPolicy) {
	alwaysRestart := docker.AlwaysRestart()
	failure := docker.RestartOnFailure(10)
	unlessStopped := docker.RestartUnlessStopped()
	neverRestart := docker.NeverRestart()

	o.autoremove = false

	switch restartPolicy {
	case Restart:
		o.restartPolicy = &alwaysRestart
	case Failure:
		o.restartPolicy = &failure
	case UnlessStopped:
		o.restartPolicy = &unlessStopped
	case No:
		o.restartPolicy = &neverRestart
	case Autoremove:
		o.restartPolicy = nil
		o.autoremove = true
	}
}

// RestartPolicy defines the restart policy for the docker container.
type RestartPolicy string

const (
	// Restart defines the restart policy for the docker container.
	Restart RestartPolicy = "always"
	// Failure defines the restart policy for the docker container.
	Failure = "on-failure"
	// UnlessStopped defines the restart policy for the docker container.
	UnlessStopped = "unless-stopped"
	// No defines the restart policy for the docker container.
	No = "no"
	// Autoremove removes the container when it's finished.
	Autoremove = "autoremove"
)

// Build converts the option builder into a list of command line parameters.
// for use in anvil.
func (o *OptionBuilder) Build() (args []string, err error) {
	v := reflect.ValueOf(*o)

	fields, err := getFields(v, make(map[string]string))
	if err != nil {
		return []string{}, fmt.Errorf("error getting fields: %w", err)
	}

	for key, value := range fields {
		args = append(args, fmt.Sprintf("--%s %s", key, value))
	}
	return args, nil
}

func getFields(v reflect.Value, options map[string]string) (_ map[string]string, err error) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		value := v.Field(i)

		// iterate over all available tags in the sub field, we'll ignore parent fields
		if field.Type.Kind() == reflect.Struct {
			options, err = getFields(value, options)
			if err != nil {
				return map[string]string{}, fmt.Errorf("error getting fields: %w", err)
			}
		}

		tag := field.Tag.Get(tagName)
		if tag == "" {
			continue
		}

		// check if the field has the default value
		if fieldIsEmpty(value) {
			continue
		}

		castField, err := valueToString(value)
		if err != nil {
			return map[string]string{}, fmt.Errorf("error casting field of type %v: %w", value.Kind(), err)
		}
		options[tag] = castField
	}
	return options, nil
}

// nolint: cyclop
func fieldIsEmpty(v reflect.Value) bool {
	//nolint: exhaustive
	switch v.Kind() {
	case reflect.Bool:
		if !v.Bool() {
			return true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() == 0 {
			return true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if v.Uint() == 0 {
			return true
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() == 0.0 {
			return true
		}
	case reflect.String:
		if v.String() == "" {
			return true
		}
	case reflect.Slice, reflect.Map:
		if v.IsNil() || v.Len() == 0 {
			return true
		}
	case reflect.Pointer:
		return v.IsNil()
	case reflect.Struct:
		if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
			return true
		}
	}
	return false
}

// valueToString converts a reflect.Value to a string.
// bools are cast to empty.
// nolint: cyclop
func valueToString(v reflect.Value) (string, error) {
	switch v.Type().Name() {
	case reflect.TypeOf(Frontier).Name():
		//nolint: forcetypeassert
		return v.Interface().(Hardfork).String(), nil
	case reflect.TypeOf(Fees).Name():
		//nolint: forcetypeassert
		return v.Interface().(Order).String(), nil
	}

	//nolint: exhaustive
	switch v.Kind() {
	// custom types first
	case reflect.TypeOf(accounts.DerivationPath{}).Kind():
		//nolint: forcetypeassert
		derivationPath := v.Interface().(accounts.DerivationPath)
		if len(derivationPath) == 0 {
			return "", errors.New("derivation path is empty")
		}

		// Remove the last item from the slice (since this is used as a base derivation path rather than a fully qualified account name)
		derivationPath = derivationPath[:len(derivationPath)-1]

		return fmt.Sprintf("\"%s\"", derivationPath.String()), nil
	case reflect.Bool:
		return "", nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	case reflect.String:
		return fmt.Sprintf("\"%s\"", v.String()), nil
	default:
		return "", fmt.Errorf("could not convert %v to string", v.Kind())
	}
}

func getFunctionName(temp interface{}) string {
	strs := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
	return strs[len(strs)-1]
}
