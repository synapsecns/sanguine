package toml_test

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/toml"
)

// WorkingConfig is a config that will work.
type WorkingConfig struct {
	SubConfig *WorkingSubConfig `toml:"SubConfig"`
}

// Encode enodes the sub config.
func (w *WorkingConfig) Encode() (string, error) {
	//nolint: wrapcheck
	return toml.Encode(w)
}

var _ toml.Encodable = &WorkingConfig{}

// WorkingSubConfig is a config with that uses the marshaller method and will work
// this will work because we wrap it in the working subconfig burntToml.
type WorkingSubConfig struct {
	// OperatorKeyFile is a path to the operator key this node will use
	IsaConfig bool `toml:"IsAConfig"`
	// Name is the password used to decrypt the keyfile
	Name string `toml:"OperatorKeyPassword"`
}

// MarshallText uses the marshall text method from marshaller
// Important: w must be a pointer.
func (w *WorkingSubConfig) MarshalText() (text []byte, err error) {
	//nolint: wrapcheck
	return toml.MarshalTextPtr(w)
}

var _ encoding.TextMarshaler = &WorkingSubConfig{}

// BrokenConfig is a config that will not work b/c of a missing MarshallTextPtr encode on a pointer config.
type BrokenConfig struct {
	SubConfig *BrokenSubConfig `toml:"SubConfig"`
}

// Encode encodes the broken config.
func (b *BrokenConfig) Encode() (string, error) {
	//nolint: wrapcheck
	return toml.Encode(b)
}

var _ toml.Encodable = &BrokenConfig{}

// BrokenConfig is a broken config since it is passed by pointer but does not implement marshall text
// this will return nothing.
type BrokenSubConfig struct {
	// OperatorKeyFile is a path to the operator key this node will use
	IsaConfig bool `burntToml:"IsAConfig"`
	// Name is the password used to decrypt the keyfile
	Name string `burntToml:"OperatorKeyPassword"`
}

// generateWorkingConfig is a helper method to generate a working config with random data.
func generateWorkingConfig() *WorkingConfig {
	return &WorkingConfig{
		SubConfig: &WorkingSubConfig{
			IsaConfig: gofakeit.Bool(),
			Name:      gofakeit.Name(),
		}}
}

// generateBrokenCofnig generates a broken config is a helper method to generate a broken config with random data.
func generateBrokenConfig() *BrokenConfig {
	return &BrokenConfig{SubConfig: &BrokenSubConfig{
		IsaConfig: gofakeit.Bool(),
		Name:      gofakeit.Name(),
	}}
}

// ExampleTestMarshallerImplementation shows how to use the marshaller function
// this is useful for when you have a config and a sub config that uses a pointer.
func ExampleMarshalTextPtr() {
	// if we use a struct that doesn't use marshaller text, we see no output
	brokenConfig := generateBrokenConfig()
	fmt.Println(brokenConfig.Encode())

	// otherwise, we see output
	workingConfig := generateWorkingConfig()
	fmt.Println(workingConfig.Encode())
}

func (t TomlSuite) TestExample() {
	NotPanics(t.T(), func() {
		ExampleMarshalTextPtr()
	})
}
