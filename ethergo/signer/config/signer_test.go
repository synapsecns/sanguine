package config_test

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"os"
	"path/filepath"
)

func (c ConfigSuite) TestInvalidType() {
	testWallet, err := wallet.FromRandom()
	Nil(c.T(), err)

	testConfig := config.SignerConfig{
		Type: gofakeit.LoremIpsumWord(),
		File: filet.TmpFile(c.T(), "", testWallet.PrivateKeyHex()).Name(),
	}

	ok, err := testConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrUnsupportedSignerType)
}

func (c ConfigSuite) TestInvalidFile() {
	testConfig := config.SignerConfig{
		Type: config.FileType.String(),
		File: filepath.Join(os.TempDir(), "idontexist"),
	}

	ok, err := testConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	Error(c.T(), err)

	// now test one with an existing file, but garbage data
	testConfig.File = filet.TmpFile(c.T(), "", gofakeit.Sentence(10)).Name()

	ok, err = testConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	Error(c.T(), err)
}

func (c ConfigSuite) TestValidFileSigner() {
	testWallet, err := wallet.FromRandom()
	Nil(c.T(), err)

	testConfig := config.SignerConfig{
		Type: config.FileType.String(),
		File: filet.TmpFile(c.T(), "", testWallet.PrivateKeyHex()).Name(),
	}

	ok, err := testConfig.IsValid(c.GetTestContext())
	True(c.T(), ok)
	Nil(c.T(), err)
}

func (c ConfigSuite) TestMarshallUnmarshallAWSConfig() {
	awsConfig := config.AWSConfig{
		Region:       gofakeit.Name(),
		AccessKey:    gofakeit.Name(),
		AccessSecret: gofakeit.Name(),
		KeyID:        gofakeit.Name(),
	}

	encodedConfig, err := awsConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", string(encodedConfig))

	resCfg, err := config.DecodeAWSConfig(file.Name())
	Nil(c.T(), err)
	Equal(c.T(), awsConfig, resCfg)
}

func (c ConfigSuite) TestGCPConfigMarshallUnmarshall() {
	gcpConfig := config.GCPConfig{
		KeyName:        gofakeit.Name(),
		CredentialFile: fmt.Sprintf("%s/%s.%s", os.TempDir(), gofakeit.Name(), gofakeit.FileExtension()),
	}

	encodedConfig, err := gcpConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", string(encodedConfig))

	decodedConfig, err := config.DecodeGCPConfig(file.Name())
	Nil(c.T(), err)

	Equal(c.T(), gcpConfig, decodedConfig)
}
