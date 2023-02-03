package detector_test

import (
	"fmt"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/gitmock"
	"os"
	"path/filepath"
)

func (d *DetectorSuite) TestChangedModules() {
	// store the headref for integrity checking after tests
	headRef, err := d.sourceRepo.repo.Head()
	Nil(d.T(), err, "could not get source repo head ref")

	testRepo, err := gitmock.NewTestRepo(d.T(), d.sourceRepo.repo, d.sourceRepo.dir)
	Nil(d.T(), err, "should not return an error")

	_, err = os.Create(filepath.Join(d.sourceRepo.dir, "lib", "newfile.go"))
	Nil(d.T(), err, "should not return an error")

	testRepo.Commit()

	ct, err := detector.GetChangeTree(d.GetTestContext(), d.sourceRepo.dir, "", headRef.Hash().String(), "", "main")
	Nil(d.T(), err, "should not return an error")

	withDeps, err := detector.DetectChangedModules(d.sourceRepo.dir, ct, true)
	Nil(d.T(), err, "should not return an error")

	withoutDeps, err := detector.DetectChangedModules(d.sourceRepo.dir, ct, false)
	Nil(d.T(), err, "should not return an error")

	False(d.T(), withoutDeps["./cmd/app1"])
	False(d.T(), withoutDeps["./cmd/app2"])
	False(d.T(), withoutDeps["./cmd/app3"])
	True(d.T(), withoutDeps["./lib"])

	True(d.T(), withDeps["./cmd/app1"])
	True(d.T(), withDeps["./cmd/app2"])
	False(d.T(), withDeps["./cmd/app3"])
	True(d.T(), withDeps["./lib"])
}

func (d *DetectorSuite) TestGetDependencyDag() {
	d.T().Skip()
	deps, err := detector.GetDependencyDag(d.sourceRepo.dir)
	Nil(d.T(), err, "should not return an error")

	Equal(d.T(), deps["./cmd/app1"], []string{"./lib"})
	Equal(d.T(), deps["./cmd/app2"], []string{"./lib"})
}

func (d *DetectorSuite) TestChangeTree() {
	testRepo, err := gitmock.NewTestRepo(d.T(), d.sourceRepo.repo, d.sourceRepo.dir)
	Nil(d.T(), err, "should not return an error")

	prevHash, err := d.sourceRepo.repo.Head()
	Nil(d.T(), err, "should not return an error")

	addedFiles := testRepo.AddRandomFiles(5)

	changeTree, err := detector.GetChangeTree(d.GetTestContext(), d.sourceRepo.dir, "", prevHash.Hash().String(), "", "main")
	Nil(d.T(), err, "should not empty change tree")

	for _, file := range addedFiles {
		if !changeTree.HasPath(file) {
			fmt.Println("hi")
		}
		True(d.T(), changeTree.HasPath(file), "could not find added file in change tree", file)
	}
}
