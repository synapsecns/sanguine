package updater_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/agents/updater"
	"github.com/synapsecns/sanguine/core/config"
)

func (u UpdaterSuite) TestUpdater() {
	u.T().Skip("could not update")
	testConfig := config.Config{
		Domains: map[string]config.DomainConfig{
			"test": u.domainClient.Config(),
		},
		DBPath: filet.TmpDir(u.T(), ""),
	}
	_, err := updater.NewUpdater(u.GetTestContext(), testConfig)
	Nil(u.T(), err)
}
