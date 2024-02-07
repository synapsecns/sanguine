package config

import (
	kms "cloud.google.com/go/kms/apiv1"
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/awssigner"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/gcpsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
)

// SignerConfig contains a signer config. Currently this config
// only supports local based signers due to a lack of isomorphic types
// when we parse yaml.
type SignerConfig struct {
	// Type is the driver used for the signer
	Type string
	// File is the file used for the key.
	File string
}

// IsValid determines if the config is valid.
func (s SignerConfig) IsValid(_ context.Context) (ok bool, err error) {
	if !strings.EqualFold(s.Type, FileType.String()) {
		return false, fmt.Errorf("%w: %s. must be one of: %s", ErrUnsupportedSignerType, s.Type, allSignerTypesList())
	}

	// TODO: we'll need to switch validity here based on type once we have more then one supported configuration type
	// alternatively, we could try to use an awsconfig type file, but this makes the virtual box setup more tedious. A third option is a json blob
	_, err = wallet.FromKeyFile(s.File)
	if err != nil {
		return false, fmt.Errorf("file %s invalid: %w", s.File, err)
	}

	return true, nil
}

// ErrUnsupportedSignerType indicates the signer type being used is unsupported.
var ErrUnsupportedSignerType = errors.New("unsupported signer type")

// SignerType is the signer type
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=SignerType -linecomment
type SignerType int

const (
	// FileType is a file-based signer.
	FileType SignerType = iota + 1 // File
	// AWSType is an aws kms based signer.
	AWSType // AWS
	// GCPType is a gcp cloud based signer.
	GCPType // GCP
)

// AllSignerTypes is a list of all contract types. Since we use stringer and this is a testing library, instead
// of manually copying all these out we pull the names out of stringer. In order to make sure stringer is updated, we panic on
// any method called where the index is higher than the stringer array length.
var AllSignerTypes []SignerType

// set all contact types.
func init() {
	for i := 0; i < len(_SignerType_index); i++ {
		contractType := SignerType(i)
		AllSignerTypes = append(AllSignerTypes, contractType)
	}
}

// allSignerTypesList prints a list of all signer types. This is useful for returning errors.
func allSignerTypesList() string {
	var res []string
	for _, signerType := range AllSignerTypes {
		res = append(res, signerType.String())
	}

	return strings.Join(res, ",")
}

// SignerFromConfig creates a new signer from a signer config.
// TODO: this needs to be moved to some kind of common package.
// in the old code configs were split into responsible packages. Maybe something like that works here?
func SignerFromConfig(ctx context.Context, config SignerConfig) (signer.Signer, error) {
	switch config.Type {
	case FileType.String():
		wall, err := wallet.FromKeyFile(core.ExpandOrReturnPath(config.File))
		if err != nil {
			return nil, fmt.Errorf("could not add signer: %w", err)
		}

		res := localsigner.NewSigner(wall.PrivateKey())

		return res, nil
	case AWSType.String():
		awsConfig, err := DecodeAWSConfig(config.File)
		if err != nil {
			return nil, fmt.Errorf("could not decode aws config: %w", err)
		}
		res, err := awssigner.NewKmsSigner(ctx, awsConfig.Region, awsConfig.AccessKey, awsConfig.AccessSecret, awsConfig.KeyID)
		if err != nil {
			return nil, fmt.Errorf("could not decode aws config: %w", err)
		}
		return res, nil
	case GCPType.String():
		gcpConfig, err := DecodeGCPConfig(config.File)
		if err != nil {
			return nil, fmt.Errorf("could not decode gcp config: %w", err)
		}

		return makeGCPSigner(ctx, gcpConfig)
	default:
		return nil, fmt.Errorf("could not create signer: %w", ErrUnsupportedSignerType)
	}
}

func makeGCPSigner(ctx context.Context, gcpConfig GCPConfig) (signer.Signer, error) {
	var options []option.ClientOption
	if gcpConfig.CredentialFile != "" {
		options = append(options, option.WithCredentialsFile(gcpConfig.CredentialFile))
	}

	if gcpConfig.Endpoint != "" {
		options = append(options, option.WithEndpoint(gcpConfig.Endpoint))
	}

	keyClient, err := kms.NewKeyManagementClient(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("could not create key client: %w", err)
	}

	res, err := gcpsigner.NewManagedKey(ctx, keyClient, gcpConfig.KeyName)
	if err != nil {
		return nil, fmt.Errorf("could not create managed key: %w", err)
	}

	return res, nil
}

// GCPConfig is the config for a GCP signer.
type GCPConfig struct {
	// KeyName is the name of the key to use.
	KeyName string `yaml:"key_name"`
	// CredentialFile is the path to the credentials file.
	// note: this is not recommended for production use.
	// workload identity federation is recommended.
	CredentialFile string `yaml:"credential_file"`
	// Endpoint is the endpoint to use. This is useful for testing.
	Endpoint string `yaml:"endpoint"`
}

// Encode encodes the config to yaml.
func (a GCPConfig) Encode() ([]byte, error) {
	output, err := yaml.Marshal(&a)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(a), 20), err)
	}
	return output, nil
}

// DecodeGCPConfig decodes the config from a file.
func DecodeGCPConfig(filePath string) (cfg GCPConfig, err error) {
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return GCPConfig{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return GCPConfig{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return cfg, nil
}

// AWSConfig is the config for an AWS signer.
// this should match the schema of the file passed in.
type AWSConfig struct {
	// Region is the region the signer is in.
	Region string `yaml:"region"`
	// AccessKey is the access key for the signer.
	AccessKey string `yaml:"access_key"`
	// AccessSecret is the access secret for the signer.
	AccessSecret string `yaml:"access_secret"`
	// KeyID is the key id for the signer.
	KeyID string `yaml:"key_id"`
}

// Encode encodes the config to yaml.
func (a AWSConfig) Encode() ([]byte, error) {
	output, err := yaml.Marshal(&a)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(a), 20), err)
	}
	return output, nil
}

// DecodeAWSConfig decodes the config from a file.
func DecodeAWSConfig(filePath string) (cfg AWSConfig, err error) {
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return AWSConfig{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return AWSConfig{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return cfg, nil
}
