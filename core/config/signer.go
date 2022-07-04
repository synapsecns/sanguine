package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
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
		return false, fmt.Errorf("%w: %s. must be one of: %s", UnsupportedSignerType, s.Type, allSignerTypesList())
	}

	// TODO: we'll need to switch validity here based on type once we have more then one supported configuration type
	// alternatively, we could try to use an awsconfig type file, but this makes the virtual box setup more tedious. A third option is a json blob
	_, err = wallet.FromKeyFile(s.File)
	if err != nil {
		return false, fmt.Errorf("file %s invalid: %w", s.File, err)
	}

	return true, nil
}

// UnsupportedSignerType indicates the signer type being used is unsupported.
var UnsupportedSignerType = errors.New("unsupported signer type")

//go:generate go run golang.org/x/tools/cmd/stringer -type=SignerType -linecomment
type SignerType int

const (
	// FileType is a file-based signer.
	FileType SignerType = 0 // File
	// KMSType is a non-file based signer.
	KMSType SignerType = iota // KMS
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
