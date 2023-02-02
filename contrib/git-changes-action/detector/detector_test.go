package detector_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/git-changest-action/detector"
	"github.com/synapsecns/sanguine/contrib/git-changest-action/detector/gitmock"
)

func (d *DetectorSuite) TestChangedModules() {
	detector.DetectChangedModules("/Users/jake/sanguine", "665a3b1d014c0f5482c3cf6393868f70438ad3f8")
}

func (d *DetectorSuite) TestGetDependencyDag() {
	d.T().Skip()
	deps, err := detector.GetDependencyDag(d.sourceRepo.dir)
	Nil(d.T(), err, "should not return an error")

	Equal(d.T(), deps["./cmd/app1"], []string{"./lib"})
	Equal(d.T(), deps["./cmd/app2"], []string{"./lib"})
}

func (d *DetectorSuite) TestChangeTree() {
	d.T().Skip()
	testRepo, err := gitmock.NewTestRepo(d.T(), d.sourceRepo.repo, d.sourceRepo.dir)
	Nil(d.T(), err, "should not return an error")

	prevHash, err := d.sourceRepo.repo.Head()
	Nil(d.T(), err, "should not return an error")

	addedFiles := testRepo.AddRandomFiles(5)

	changeTree, err := detector.GetChangeTree(d.sourceRepo.dir, prevHash.Hash().String())
	Nil(d.T(), err, "should not empty change tree")

	for _, file := range addedFiles {
		True(d.T(), changeTree.HasPath(file))
	}
}
