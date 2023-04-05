package config_test

import (
	"bytes"
	"fmt"
	"github.com/synapsecns/sanguine/core/internal/assets"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
)

// testSuffix allows for easy bulk deleting in the case this isn't automatically taken care of.
const testSuffix = "_test"

func (c ConfigSuite) TestGetLogo() {
	c.generateSafeConfigDir()

	// first fetch will create the path, second should use it
	const iterations = 2
	for i := 0; i < iterations; i++ {
		NotPanics(c.T(), func() {
			logoPath, _ := config.GetLogoPath()
			//nolint:gosec
			logoContents, err := os.ReadFile(logoPath)
			Nil(c.T(), err)
			True(c.T(), bytes.Equal(logoContents, assets.Logo))
		})
	}
}

// generateSafeConfigDir gets a safe (deletable after test) config dir for testing
// this will be deleted after the test.
func (c ConfigSuite) generateSafeConfigDir() {
	c.T().Helper()
	// use a different dir each time for testing
	gofakeit.Seed(time.Now().UnixNano() * int64(c.GetTestID()))

	// get the original config dir so we don't acidentally delete this.
	// we do a compare and fail the test if these are the same
	ogConfigDir, err := config.GetConfigDir()
	Nil(c.T(), err)

	config.SetNameSuffix(fmt.Sprintf("%s_%s", testSuffix, gofakeit.AppName()))
	newConfigDir, err := config.GetConfigDir()
	Nil(c.T(), err)

	NotEqual(c.T(), ogConfigDir, newConfigDir)

	c.DeferAfterTest(func() {
		err := os.RemoveAll(newConfigDir)
		if err != nil {
			c.T().Logf("could not remove path: %v, should do manually", err)
		}
	})
}
