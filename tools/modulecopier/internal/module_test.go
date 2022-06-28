package internal_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/tools/modulecopier/internal"
	"os"
	"path"
	// required for copy test.
	_ "github.com/ethereum/go-ethereum/common"
)

// fileCheck is the file to check for. We check for .mailmap in ethereum because
// 1) we don't use it
// 2) it's relatively uncommon depiste being a git feature (https://git-scm.com/docs/git-check-mailmap)
// 3) it hasn't changed in 4 years.
// if you're seeing this test break, make sure this file wasn't deleted from ethereum.
const fileCheck = ".mailmap"
const ethModule = "github.com/ethereum/go-ethereum"

// TestGetEthModulePath tests a fetch of the ethereum module path.
func (s GeneratorSuite) TestGetEthModulePath() {
	ethModulePath, err := internal.GetModulePath(ethModule)
	Nil(s.T(), err)

	if _, err := os.Stat(path.Join(ethModulePath, fileCheck)); os.IsNotExist(err) {
		s.T().Errorf("expected to find module path for %s%s, did not find any. Used eth module path %s", ethModule, fileCheck, ethModulePath)
	}
}
