package config

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
	"gopkg.in/yaml.v2"
)

// ScribeConfig is used to configure a scribe for the Executor.
type ScribeConfig struct {
	// Type is the type of scribe. This can be either "embedded" or "remote".
	Type ScribeType `yaml:"type"`

	// EmbeddedDBConfig is the database configuration for an embedded scribe.
	EmbeddedDBConfig scribeConfig.DBConfig `yaml:"embedded_db_config,omitempty"`
	// EmbeddedScribeConfig is the config for the embedded scribe.
	EmbeddedScribeConfig scribeConfig.Config `yaml:"embedded_scribe_config,omitempty"`

	// Port is the port to listen on for the remote scribe.
	Port uint `yaml:"port,omitempty"`
	// URL is the URL to connect to for the remote scribe.
	URL string `yaml:"url,omitempty"`
}

// Encode encodes the config into a yaml byte slice.
func (c ScribeConfig) Encode() ([]byte, error) {
	output, err := yaml.Marshal(&c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(c), 20), err)
	}
	return output, nil
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *ScribeConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.EmbeddedScribeConfig.IsValid(ctx); !ok {
		return false, fmt.Errorf("embedded scribe config is invalid: %w", err)
	}

	return true, nil
}

// ScribeType is the type of scribe.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ScribeType -linecomment
type ScribeType int

// UnmarshalYAML unmarshals a ScribeType from a yaml string.
func (i *ScribeType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var yamlObj interface{}
	err := unmarshal(&yamlObj)
	if err != nil {
		panic(err)
	}

	stringifiedType, ok := yamlObj.(string)
	if !ok {
		return fmt.Errorf("could not unmarshal scribe type. Expected %T got %T", stringifiedType, yamlObj)
	}

	for _, scribeType := range allScribeTypes {
		if scribeType.String() == stringifiedType {
			i = &scribeType
			return nil
		}
	}

	return fmt.Errorf("could not unmarshal scribe type. Expected one of %v got %v", allScribeTypes, stringifiedType)

	return nil
}

func (i ScribeType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

const (
	// RemoteScribeType is a remote scribe instance. This is a scribe
	// instance setup as a seperate service and ran independently.
	RemoteScribeType ScribeType = iota + 1 // remote

	// EmbeddedScribeType is an embedded scribe instance. This is a scribe
	// in the same process spun up by the agent.
	EmbeddedScribeType // embedded
)

var allScribeTypes = []ScribeType{
	EmbeddedScribeType,
	RemoteScribeType,
}

// make sure allScribeTypes are set
func init() {
	if len(allScribeTypes) != len(_ScribeType_index)-1 {
		panic("please update allScribeTypes before running test again")
	}
}

var _ yaml.Unmarshaler = (*ScribeType)(nil)
