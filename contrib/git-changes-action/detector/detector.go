package detector

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/synapsecns/sanguine/contrib/git-changest-action/detector/tree"
	"golang.org/x/mod/modfile"
	"os"
	"path"
)

// DetectChangedModules is the change detector client.
func DetectChangedModules(repoPath, fromHash string) (modules map[string]bool, err error) {
	modules = make(map[string]bool)

	goWorkPath := path.Join(repoPath, "go.work")

	if !common.FileExist(goWorkPath) {
		return nil, fmt.Errorf("go.work file not found in %s", repoPath)
	}

	workFile, err := os.ReadFile(goWorkPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read go.work file: %w", err)
	}

	parsedWorkFile, err := modfile.ParseWork(goWorkPath, workFile, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go.work file: %w", err)
	}

	ct, err := getChangeTree(repoPath, fromHash)
	if err != nil {
		return nil, fmt.Errorf("could not get change tree: %w", err)
	}

	depGraph, err := getDependencyGraph(repoPath)
	if err != nil {
		return nil, fmt.Errorf("could not get dep graph: %w", err)
	}

	for _, module := range parsedWorkFile.Use {
		changed := false
		if ct.HasPath(module.Path) {
			changed = true
		}

		deps := depGraph[module.Path]
		for _, dep := range deps {
			if ct.HasPath(dep) {
				changed = true
			}
		}

		modules[module.Path] = changed
	}

	return modules, nil
}

// getChangeTree returns a tree of all the files that have changed between the current commit and the commit with the given hash.
func getChangeTree(repoPath string, fromHash string) (tree.Tree, error) {
	// create the change tree
	changeTree := tree.NewTree()

	// open the repository
	repository, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("could not open repository %s: %w", repoPath, err)
	}

	// get the head of the repository (git status)
	repoHead, err := repository.Head()
	if err != nil {
		return nil, fmt.Errorf("could not get head for repository %s: %w", repoPath, err)
	}

	// get each commit object (before and after)
	toCommitObject, err := repository.CommitObject(repoHead.Hash())
	if err != nil {
		return nil, fmt.Errorf("could not get commit object for hash %s: %w", fromHash, err)
	}

	fromCommitObject, err := repository.CommitObject(plumbing.NewHash(fromHash))
	if err != nil {
		return nil, fmt.Errorf("could not get commit object for hash %s: %w", fromHash, err)
	}

	// generate a diff
	patch, err := fromCommitObject.Patch(toCommitObject)
	if err != nil {
		return nil, fmt.Errorf("could not get patch for commit objects %s and %s: %w", fromHash, toCommitObject.Hash.String(), err)
	}

	// add diff items to the tree
	for _, filePatch := range patch.FilePatches() {
		fromFile, toFile := filePatch.Files()
		if fromFile != nil {
			changeTree.Add(fromFile.Path())
		}
		if toFile != nil {
			changeTree.Add(toFile.Path())
		}
	}

	return changeTree, nil
}
