package gcpsigner_test

import (
	"cloud.google.com/go/kms/apiv1/kmspb"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/gcpsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/gcpsigner/gcpmock"
)

func (g *GCPSignerSuite) TestSign() {
	mc, err := gcpmock.NewMockClient()
	Nil(g.T(), err, "should not return an error")

	mk, err := gcpsigner.NewManagedKey(g.GetTestContext(), mc, "test")
	Nil(g.T(), err, "should not return an error")

	mk.Address()
	res, err := mc.GetPublicKey(g.GetTestContext(), &kmspb.GetPublicKeyRequest{
		Name: "test",
	})

	Nil(g.T(), err, "should not return an error")

	_ = res
}
