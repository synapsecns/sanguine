package detector

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/go-github/v37/github"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
	"golang.org/x/mod/modfile"
	"os"
	"path"
)

// DetectChangedModules is the change detector client.
// nolint: cyclop
func DetectChangedModules(repoPath string, ct tree.Tree, includeDeps bool) (modules map[string]bool, err error) {
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

// getChangeTreeFromGit returns a tree of all the files that have changed between the current commit and the commit with the given hash.
// nolint: cyclop
func getChangeTreeFromGit(repoPath string, head, base string) (tree.Tree, error) {
	fmt.Println("og")
	fmt.Println(head)
	fmt.Println(base)
	// open the repository
	repository, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("could not open repository %s: %w", repoPath, err)
	}

	_, err = hex.DecodeString(base)
	isBaseSha := err == nil
	isBaseSameAsHead := base == head

	var baseSha string

	// If base is commit SHA we will do comparison against the referenced commit
	// Or if base references same branch it was pushed to, we will do comparison against the previously pushed commit
	//nolint: nestif
	if isBaseSha || isBaseSameAsHead {
		baseSha = base
		if isBaseSha {
			fmt.Println("base sha")
			var ok bool
			baseSha, ok, _ = tryGetPushEvent()

			if !ok {
				fmt.Println("could not get push event")
				baseSha, err = getLastCommitHash(repository)
				if err != nil {
					return nil, fmt.Errorf("could not get last commit hash: %w", err)
				}
			}
		}
	}

	if head == "" {
		rawHead, err := repository.Head()
		if err != nil {
			return nil, fmt.Errorf("could not get HEAD: %w", err)
		}

		head = rawHead.String()
	}

	if !isBaseSha {
		refs, err := repository.References()
		if err != nil {
			return nil, fmt.Errorf("could not get references: %w", err)
		}

		remotes, err := repository.Remotes()
		if err != nil {
			return nil, fmt.Errorf("could not get remotes: %w", err)
		}

		err = refs.ForEach(func(reference *plumbing.Reference) error {
			if reference.Name().String() == base {
				baseSha = reference.Hash().String()
			}

			for _, remote := range remotes {
				refName := plumbing.NewRemoteReferenceName(remote.Config().Name, base)
				if refName == reference.Name() {
					baseSha = reference.Hash().String()
				}

			}

			return nil
		})

		if err != nil {
			return nil, fmt.Errorf("could not iterate through references: %w", err)
		}
	}

	// create the change tree
	changeTree := tree.NewTree()

	// get each commit object (before and after)
	baseObject, err := repository.CommitObject(plumbing.NewHash(baseSha))
	if err != nil {
		return nil, fmt.Errorf("could not get commit object for base %s: %w", baseSha, err)
	}

	headObject, err := repository.CommitObject(plumbing.NewHash(head))
	if err != nil {
		return nil, fmt.Errorf("could not get commit object for head %s: %w", head, err)
	}

	diff, err := fastDiff(baseObject, headObject)
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

func tryGetPushEvent() (lastSha string, ok bool, err error) {
	f, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return "", false, fmt.Errorf("could not open event path: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()

	var gpe github.PushEvent

	if err := json.NewDecoder(f).Decode(&gpe); err != nil {
		return "gpe", false, fmt.Errorf("could not decode event: %w", err)
	}

	return gpe.GetBefore(), true, nil
}

func getDefaultBranch() (defaultBranch string, err error) {
	f, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return "", fmt.Errorf("could not open event path: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()

	var gpe github.PushEvent

	if err := json.NewDecoder(f).Decode(&gpe); err != nil {
		return "gpe", fmt.Errorf("could not decode event: %w", err)
	}

	return gpe.Repo.GetDefaultBranch(), nil
}

// note: we don't  handle the case of no previous commit, this will error.
func getLastCommitHash(repo *git.Repository) (string, error) {
	co, err := repo.CommitObjects()
	if err != nil {
		return "", fmt.Errorf("could not get head: %w", err)
	}

	lastCommit, err := co.Next()
	if err != nil {
		return "", fmt.Errorf("could not get last commit: %w", err)
	}

	return lastCommit.Hash.String(), nil
}
