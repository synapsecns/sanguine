package origin_scanner_test

import (
	"time"

	"github.com/Flaque/filet"
	awsTime "github.com/aws/smithy-go/time"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/origin_scanner"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

func (o OriginScannerSuite) TestOriginScannerE2E() {
	testConfig := config.Config{
		Domains: map[string]config.DomainConfig{
			"test0": o.domainClients[0].Config(),
			"test1": o.domainClients[1].Config(),
			"test2": o.domainClients[2].Config(),
			"test3": o.domainClients[3].Config(),
			"test4": o.domainClients[4].Config(),
			"test5": o.domainClients[5].Config(),
			"test6": o.domainClients[6].Config(),
			"test7": o.domainClients[7].Config(),
			"test8": o.domainClients[8].Config(),
			"test":  o.domainClients[9].Config(),
		},
		Signer: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(o.T(), "", o.wallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(o.T(), ""),
			ConnString: filet.TmpDir(o.T(), ""),
		},
	}
	originScanner, err := origin_scanner.NewOriginScanner(o.GetTestContext(), testConfig, 5*time.Second)
	Nil(o.T(), err)

	/*auth := o.testBackend.GetTxContext(o.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(o.T(), err)

	tx, err := o.originContract.Dispatch(auth.TransactOpts, gofakeit.Uint32(), [32]byte{}, gofakeit.Uint32(), encodedTips, []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(o.T(), err)
	o.testBackend.WaitForConfirmation(o.GetTestContext(), tx)*/

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = originScanner.Start(o.GetTestContext())
	}()

	o.Eventually(func() bool {
		_ = awsTime.SleepWithContext(o.GetTestContext(), time.Second*10)
		return true
	})
}
