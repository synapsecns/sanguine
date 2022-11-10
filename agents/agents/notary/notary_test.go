package notary_test

import (
	"math/big"
	"time"

	"github.com/Flaque/filet"
	awsTime "github.com/aws/smithy-go/time"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

func (u NotarySuite) TestNotaryE2E() {
	testConfig := config.Config{
		Domains: map[string]config.DomainConfig{
			"test": u.domainClient.Config(),
		},
		Signer: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.wallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
	}
	ud, err := notary.NewNotary(u.GetTestContext(), testConfig)
	Nil(u.T(), err)

	auth := u.testBackend.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	desinationDomain := gofakeit.Uint32()
	tx, err := u.originContract.Dispatch(auth.TransactOpts, desinationDomain, [32]byte{}, gofakeit.Uint32(), encodedTips, []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = ud.Start(u.GetTestContext())
	}()

	u.Eventually(func() bool {
		// TODO (joe): Figure out why attestationContract points to old version and fix this test after the GlobalRegistry changes
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)
		latestNonce, err := u.attestationContract.GetLatestNonce(&bind.CallOpts{Context: u.GetTestContext()}, u.domainClient.Config().DomainID, desinationDomain, u.signer.Address())
		Nil(u.T(), err)

		return latestNonce != 0
	})
}
