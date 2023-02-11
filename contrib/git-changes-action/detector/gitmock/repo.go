// Package gitmock provides utilities for testing git repos w/ mock data
package gitmock

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// Repo is a test repo.
type Repo struct {
	repo         *git.Repository
	dir          string
	changedFiles []string
	tree         tree.Tree
	tb           testing.TB
}

// NewTestRepo creates a new test repo.
func NewTestRepo(tb testing.TB, r *git.Repository, dir string) (*Repo, error) {
	tb.Helper()

	testTree := tree.NewTree()
	// create a tree with the following structure:
	// we'll use this to add some paths
	err := tree.AddDirectoryPaths(testTree, dir, dir)
	assert.Nil(tb, err, "should not return an error")

	return &Repo{
		repo: r,
		dir:  dir,
		tb:   tb,
		tree: testTree,
	}, nil
}

// trimFiles trims files from a list of paths.
func (t *Repo) trimFiles(paths []string) (res []string) {
	t.tb.Helper()

	for _, path := range paths {
		stats, err := os.Stat(filepath.Join(t.dir, path))
		assert.Nil(t.tb, err)

		if stats.IsDir() {
			res = append(res, path)
		}
	}
	return res
}

// AddRandomFiles adds random files to the git repo
// addedFiles are returned relative to the repo.
func (t *Repo) AddRandomFiles(fileCount int) (addedFiles []string) {
	if fileCount > 0 {
		defer t.Commit()
	}

	dirPaths := t.trimFiles(t.tree.AllPaths())

	for i := 0; i < fileCount; i++ {
		newFile := filepath.Join(t.dir, gofakeit.RandomString(dirPaths), fmt.Sprintf("%s.dontignoreme", gofakeit.Word()))
		//nolint: gosec
		testFile, err := os.Create(newFile)
		assert.Nil(t.tb, err, "should not return an error")

		_, err = testFile.Write(gofakeit.ImageJpeg(20, 20))
		assert.Nil(t.tb, err, "should not return an error")

		addedFiles = append(addedFiles, strings.TrimPrefix(newFile, t.dir))
	}
	t.changedFiles = append(t.changedFiles, addedFiles...)
	return addedFiles
}

// Commit commits all changed files to the repo.
func (t *Repo) Commit() (hash string) {
	wt, err := t.repo.Worktree()
	assert.Nil(t.tb, err, "should be able to load work tree")

	err = wt.AddGlob(".")
	assert.Nil(t.tb, err, "should be able to add all files")

	var commitHash plumbing.Hash
	commitHash, err = wt.Commit("test Commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			When:  time.Now(),
		},
	})
	assert.Nil(t.tb, err, "should be able to Commit")

	return commitHash.String()
}
