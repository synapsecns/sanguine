package detector

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/go-github/v41/github"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/actionscore"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
	"github.com/synapsecns/sanguine/core"
	"golang.org/x/mod/modfile"
	"os"
	"path"
	"strings"
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
// nolint: cyclop, gocognit
func getChangeTreeFromGit(repoPath string, ghContext *actionscore.Context, head, base string) (tree.Tree, error) {
	// open the repository
	repository, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("could not open repository %s: %w", repoPath, err)
	}

	head, err = getHead(repository, ghContext, head)
	if err != nil {
		return nil, fmt.Errorf("could not get head: %w", err)
	}
	head = getShortName(head)

	base = getShortName(getBase(ghContext, base))

	_, err = hex.DecodeString(base)
	isBaseSha := err == nil
	isBaseSameAsHead := base == head

	var baseSha string

	// If base is commit SHA we will do comparison against the referenced commit
	// Or if base references same branch it was pushed to, we will do comparison against the previously pushed commit
	//nolint: nestif
	if isBaseSha || isBaseSameAsHead {
		if !isBaseSha {
			var ok bool
			baseSha, head, ok, _ = tryGetPushEvent()
			isBaseSha = true

			if !ok {
				baseSha, head, err = getHeadBase(repository)
				if err != nil {
					// TODO: we might need to add error handling here for last commit on a branch, see: https://github.com/dorny/paths-filter/blob/master/src/main.ts#L141
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

	// nolint: nestif
	if !isBaseSha {
		res, err := convertToSha(repository, base)
		if err != nil {
			return nil, fmt.Errorf("could not convert base to sha: %w", err)
		}

		baseSha = res.String()
	}

	// create the change tree
	changeTree := tree.NewTree()

	// get each commit object (before and after)
	baseObject, err := repository.CommitObject(plumbing.NewHash(baseSha))
	if err != nil {
		return nil, fmt.Errorf("could not get commit object for base %s: %w", baseSha, err)
	}

	headHash, err := convertToSha(repository, head)
	if err != nil {
		return nil, fmt.Errorf("could not convert head to sha: %w", err)
	}

	headObject, err := repository.CommitObject(*headHash)
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

func convertToSha(repository *git.Repository, ref string) (res *plumbing.Hash, err error) {
	// already a sha
	_, err = hex.DecodeString(ref)
	if err == nil {
		return core.PtrTo(plumbing.NewHash(ref)), nil
	}

	refs, err := repository.References()
	if err != nil {
		return nil, fmt.Errorf("could not get references: %w", err)
	}

	remotes, err := repository.Remotes()
	if err != nil {
		return nil, fmt.Errorf("could not get remotes: %w", err)
	}

	err = refs.ForEach(func(reference *plumbing.Reference) error {
		if reference.Name().String() == ref {
			res = core.PtrTo(reference.Hash())
		}

		for _, remote := range remotes {
			refName := plumbing.NewRemoteReferenceName(remote.Config().Name, ref)
			if refName == reference.Name() {
				res = core.PtrTo(reference.Hash())
			}
		}

		return nil
	})

	if err != nil {
		return nil, errors.New("could not iterate through references")
	}

	if res != nil {
		return res, nil
	}

	return nil, fmt.Errorf("could not convert reference %s to %T", ref, core.PtrTo(plumbing.NewHash("")))
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

func tryGetPushEvent() (base, head string, ok bool, err error) {
	f, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return "", "", false, fmt.Errorf("could not open event path: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()

	var gpe github.PushEvent

	if err := json.NewDecoder(f).Decode(&gpe); err != nil {
		return "gpe", "", false, fmt.Errorf("could not decode event: %w", err)
	}

	return gpe.GetBefore(), gpe.GetAfter(), true, nil
}

// note: we don't  handle the case of no previous commit, this will error.
func getHeadBase(repo *git.Repository) (head string, base string, err error) {
	co, err := repo.Head()
	if err != nil {
		return "", "", fmt.Errorf("could not get head: %w", err)
	}

	citer, err := repo.Log(&git.LogOptions{From: co.Hash()})
	if err != nil {
		return "", "", fmt.Errorf("could not get logs: %w", err)
	}

	_, err = citer.Next()
	if err != nil {
		return "", "", fmt.Errorf("could not pass head: %w", err)
	}
	lastCommit, err := citer.Next()
	if err != nil {
		return "", "", fmt.Errorf("could not get last commit: %w", err)
	}

	return co.Hash().String(), lastCommit.Hash.String(), nil
}

// getHead gets the head of the current branch.
// it attempts to mirror the logic of  https://github.com/dorny/paths-filter/blob/0ef5f0d812dc7b631d69e07d2491d70fcebc25c8/src/main.ts#L104
func getHead(repo *git.Repository, ghContext *actionscore.Context, head string) (string, error) {
	if head != "" {
		return head, nil
	}

	if ghContext.Ref != "" {
		return ghContext.Ref, nil
	}

	gitHead, err := repo.Head()
	if err != nil {
		// TODO: there's some other logic here: https://github.com/dorny/paths-filter/blob/4067d885736b84de7c414f582ac45897079b0a78/src/git.ts#L174
		// we might want to build in
		return "", fmt.Errorf("could not get head: %w", err)
	}

	return gitHead.Name().Short(), nil
}

func getBase(ghContext *actionscore.Context, base string) string {
	if base != "" {
		return base
	}

	return ghContext.Payload.Repository.GetDefaultBranch()
}

// emulates https://github.com/dorny/paths-filter/blob/master/src/git.ts#L185
func getShortName(ref string) string {
	const heads = "refs/heads/"
	const tags = "refs/tags/"

	if strings.HasPrefix(ref, heads) {
		return strings.TrimPrefix(ref, heads)
	}

	if strings.HasPrefix(ref, tags) {
		return strings.TrimPrefix(ref, tags)
	}

	return ref
}
