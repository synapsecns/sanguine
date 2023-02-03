package detector

import (
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
)

func GetChangeTreeFromGit(repoPath string, commitHash string) (tree.Tree, error) {
	return getChangeTreeFromGit(repoPath, commitHash)
}

func GetDependencyDag(repoPath string) (map[string][]string, error) {
	return getDependencyGraph(repoPath)
}
