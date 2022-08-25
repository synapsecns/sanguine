package config

import (
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/agents/types"
	"strings"
)

// chainTypeList contains a list of stringified chain types for convenience.
var chainTypeList []string

func init() {
	for _, chainType := range types.AllChainTypes() {
		chainTypeList = append(chainTypeList, chainType.String())
	}

	ErrInvalidChainType = fmt.Errorf("chain type must be one of (%s)", strings.Join(chainTypeList, ","))
}

// ErrInvalidChainType indicates chain type was invalid.
var ErrInvalidChainType error

// ErrInvalidDomainID indicates domain id is invalid.
var ErrInvalidDomainID = errors.New("domain ID invalid")

// ErrRequiredField indicates a required field was left blank.
var ErrRequiredField = errors.New("field is required")
