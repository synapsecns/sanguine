package kmsmock_test

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	. "github.com/stretchr/testify/assert"
	kmsHelper "github.com/synapsecns/sanguine/ethergo/signer/signer/awssigner/kmsmock"
	"sync"
)

// TestConfigSwap tests that configs are correctly swapped. Functional tests in the original repository
// (see: https://github.com/nsmithuk/local-kms/tree/master/tests) cover everything else so this serves as a sanity
// check that configs are correctly swapped by starting two services.
func (k *KMSSuite) TestConfigSwap() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			testConfig := kmsHelper.NewMockKMS(k.GetTestContext(), k.T())
			k.testCreation(testConfig)
		}()
	}
	wg.Wait()
}

// testCreation tests key creation.
func (k *KMSSuite) testCreation(testConfig *kmsHelper.MockKMSService) {
	// create a key
	key, err := testConfig.Client().CreateKey(k.GetTestContext(), &kms.CreateKeyInput{
		CustomerMasterKeySpec: types.CustomerMasterKeySpecEccSecgP256k1,
		Description:           aws.String("this is a test key"),
		KeyUsage:              types.KeyUsageTypeSignVerify,
		MultiRegion:           aws.Bool(false),
	})
	Nil(k.T(), err)
	NotNil(k.T(), key)
}
