package config

func init() {
	// init sets default name suffix to the original suffix for testing
	defaultNameSuffix = nameSuffix
}

// defaultNameSuffix is the original name suffix
// this is lower case as to be immutable.
var defaultNameSuffix string

// GetDefaultNameSuffix returns the original name suffix, so it can be reset after each test.
func GetDefaultNameSuffix() string {
	return defaultNameSuffix
}

// SetNameSuffix sets the name suffix, this should be reset to GetDefaultNameSuffix after each test.
func SetNameSuffix(newNameSuffix string) {
	nameSuffix = newNameSuffix
}

// defaultSkipNetworkedChecks is whether or not to skip the networks.
var defaultSkipNetworkedChecks bool

// GetDefaultSkipNetworkedChecks returns the opriginal skip networked check setting, so it can be reset after
// each test.
func GetDefaultSkipNetworkedChecks() bool {
	return defaultSkipNetworkedChecks
}
