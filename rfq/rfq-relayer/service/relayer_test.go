package service_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service"
)

func (t *RelayerSuite) TestNewRelayer() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()

		// create the test omnirpc backend
		relayer, err := service.NewRelayer(ctx, t.config, testDB, t.metrics)
		Nil(t.T(), err)
		NotNil(t.T(), relayer)
	})
}

// TODO
// e2e test for a successful bridge, relay, prove
// test for a unsuccessful relay
// test for a unsuccessful prove
// test for nuked balance/balance issues
