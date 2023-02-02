package detector

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
	"golang.org/x/mod/modfile"
	"os"
	"path"
)

// DetectChangedModules is the change detector client.
// nolint: cyclop
func DetectChangedModules(repoPath, fromHash string, includeDeps bool) (modules map[string]bool, err error) {
	modules = make(map[string]bool)

	goWorkPath := path.Join(repoPath, "go.work")

	if !common.FileExist(goWorkPath) {
		return nil, fmt.Errorf("go.work file not found in %s", repoPath)
	}

	//nolint: gosec
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

		if includeDeps {
			deps := depGraph[module.Path]
			for _, dep := range deps {
				if ct.HasPath(dep) {
					changed = true
				}
			}
		}

		modules[module.Path] = changed
	}

	return modules, nil
}

// getChangeTree returns a tree of all the files that have changed between the current commit and the commit with the given hash.
func getChangeTree(repoPath string, toHash string) (tree.Tree, error) {
	// open the repository
	repository, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("could not open repository %s: %w", repoPath, err)
	}

	_, err = hex.DecodeString(toHash)
	if err != nil {
		// this is a ref, convert it to a hash
		refs, err := repository.References()
		if err != nil {
			return nil, fmt.Errorf("could not get references for repository %s: %w", repoPath, err)
		}

		err = refs.ForEach(func(ref *plumbing.Reference) error {
			if ref.Name().String() == toHash {
				toHash = ref.Hash().String()
			}
			return nil
		})

		if err != nil {
			return nil, fmt.Errorf("could not get reference for repository %s: %w", repoPath, err)
		}
	}

	// create the change tree
	changeTree := tree.NewTree()

	// get the head of the repository (git status)
	repoHead, err := repository.Head()
	if err != nil {
		return nil, fmt.Errorf("could not get head for repository %s: %w", repoPath, err)
	}

	// get each commit object (before and after)
	toCommitObject, err := repository.CommitObject(repoHead.Hash())
	if err != nil {
		return nil, fmt.Errorf("could not get commit object for hash %s: %w", toHash, err)
	}

	fromCommitObject, err := repository.CommitObject(plumbing.NewHash(toHash))
	if err != nil {
		return nil, fmt.Errorf("could not get commit object for hash %s: %w", toHash, err)
	}

	diff, err := fastDiff(fromCommitObject, toCommitObject)
	if err != nil {
		return nil, fmt.Errorf("could not get diff: %w", err)
	}

	changeTree.Add(diff...)

	return changeTree, nil
}

// fastDiff is a faster way to get the diff between two commits.
// it returns a boolean rather than a full blob diff.
func fastDiff(from, to *object.Commit) (changedFiles []string, err error) {
	fromTree, err := from.Tree()
	if err != nil {
		return []string{}, fmt.Errorf("could not get tree for commit %s: %w", from.Hash.String(), err)
	}

	toTree, err := to.Tree()
	if err != nil {
		return []string{}, fmt.Errorf("could not get tree for commit %s: %w", to.Hash.String(), err)
	}

	changes, err := fromTree.Diff(toTree)
	if err != nil {
		return []string{}, fmt.Errorf("could not get diff for commit %s: %w", to.Hash.String(), err)
	}

	for _, change := range changes {
		if change.From.Name != "" {
			changedFiles = append(changedFiles, change.From.Name)
		}
		if change.To.Name != "" {
			changedFiles = append(changedFiles, change.To.Name)
		}
	}

	return changedFiles, nil
}
