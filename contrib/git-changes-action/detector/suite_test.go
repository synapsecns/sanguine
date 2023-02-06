package detector_test

import (
	"github.com/Flaque/filet"
	"github.com/go-git/go-git/v5"
	copier "github.com/otiai10/copy"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"io"
	"os"
	"strings"
	"testing"
)

// DetectorSuite defines the basic test suite.
type DetectorSuite struct {
	*testsuite.TestSuite
	// sourceRepo stores the source repository for the test suite.
	sourceRepo TestRepo
	// unsafeSourceRepo stores the source repository for the test suite.
	// it is called unsafe since it should not be used by tests
	unsafeSourceRepo TestRepo
	// unsafeHeadRef stores the head ref for the unsafe source repository.
	// this makes sure it hasn't been modified by a test
	unsafeHeadRef string
}

// TestRepo is a test repository.
type TestRepo struct {
	repo *git.Repository
	dir  string
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *DetectorSuite {
	tb.Helper()
	return &DetectorSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (d *DetectorSuite) SetupSuite() {
	d.TestSuite.SetupSuite()

	d.unsafeSourceRepo.dir = filet.TmpDir(d.T(), "")
	var err error
	// clone the repo
	d.unsafeSourceRepo.repo, err = git.PlainClone(d.unsafeSourceRepo.dir, false, &git.CloneOptions{
		URL:      "https://github.com/xmlking/go-workspace",
		Progress: nullWriter{},
		Depth:    1,
	})
	Nil(d.T(), err, "could not clone source repo")

	// store the headref for integrity checking after tests
	headRef, err := d.unsafeSourceRepo.repo.Head()
	Nil(d.T(), err, "could not get source repo head ref")

	d.unsafeHeadRef = headRef.Hash().String()
}

func (d *DetectorSuite) SetupTest() {
	d.TestSuite.SetupTest()
	d.sourceRepo.dir = filet.TmpDir(d.T(), "")

	// copy the source repo to a new directory for testing
	// this saves us from having to clone the repo for every test
	err := copier.Copy(d.unsafeSourceRepo.dir, d.sourceRepo.dir)
	Nil(d.T(), err, "could not copy source repo")

	d.sourceRepo.repo, err = git.PlainOpen(d.sourceRepo.dir)
	Nil(d.T(), err, "could not open source repo")

	// unset all github env vars
	for _, osVar := range os.Environ() {
		splitVar := strings.Split(osVar, "=")
		key := splitVar[0]

		if strings.HasPrefix(key, "GITHUB_") {
			d.T().Setenv(key, "")
		}
	}
}

func (d *DetectorSuite) TearDownTest() {
	d.TestSuite.TearDownTest()

	// remove the test repo
	err := os.RemoveAll(d.sourceRepo.dir)
	if err != nil {
		d.T().Fatal(err)
	}

	// make sure there are no changes to the unsafe repo
	workTree, err := d.unsafeSourceRepo.repo.Worktree()
	Nil(d.T(), err, "could not get source repo working dir")

	// make sure nothings changed by seeing if any file is dirty
	status, err := workTree.Status()
	Nil(d.T(), err, "could not get source repo status")
	if !status.IsClean() {
		d.T().Fatal("unsafeSourceRepo repo has been modified, please use sourceRepo instead")
	}

	// make sure the head ref hasn't changed
	headRef, err := d.unsafeSourceRepo.repo.Head()
	Nil(d.T(), err, "could not get source repo head ref")
	if headRef.Hash().String() != d.unsafeHeadRef {
		d.T().Fatal("unsafeSourceRepo repo has been modified, please use sourceRepo instead")
	}
}

func (d *DetectorSuite) TearDownSuite() {
	err := os.RemoveAll(d.unsafeSourceRepo.dir)
	if err != nil {
		d.T().Fatal(err)
	}
}

func TestDetectorSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}

// nullWriter writes no progress.
type nullWriter struct{}

// Write implements io.Writer and does nothing.
func (nullWriter) Write(p []byte) (n int, err error) {
	// Do Nothing
	return 0, nil
}

var _ io.Writer = nullWriter{}
