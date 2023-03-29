package client_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/client"
	"testing"
)

func TestOptions(t *testing.T) {
	options := client.MakeOptions([]client.OptionsArgsOption{client.WithDefaultConfirmations(4)})
	Equal(t, options.Confirmations(), 4)

	options = client.MakeOptions([]client.OptionsArgsOption{})
	Equal(t, options.Confirmations(), 0)
}
