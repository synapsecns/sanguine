package notary_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/agents/notary"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/db/datastore/sql"
	"github.com/synapsecns/sanguine/core/types"
	"math/big"
)

func (u NotarySuite) TestNotaryE2E() {
	u.T().Skip("todo")

	testConfig := config.Config{
		Domains: map[string]config.DomainConfig{
			"test": u.domainClient.Config(),
		},
		Signer: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.wallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       sql.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
	}
	ud, err := notary.NewNotary(u.GetTestContext(), testConfig)
	Nil(u.T(), err)

	auth := u.testBackend.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	tx, err := u.homeContract.Dispatch(auth.TransactOpts, gofakeit.Uint32(), [32]byte{}, gofakeit.Uint32(), encodedTips, []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = ud.Start(u.GetTestContext())
	}()

	u.Eventually(func() bool {
		latestNonce, err := u.attestationContract.LatestNonce(&bind.CallOpts{Context: u.GetTestContext()}, u.domainClient.Config().DomainID)
		Nil(u.T(), err)

		return latestNonce != 0
	})
}
