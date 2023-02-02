package detector

import (
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
)

func GetChangeTree(repoPath string, commitHash string) (tree.Tree, error) {
	return getChangeTree(repoPath, commitHash)
}

func GetDependencyDag(repoPath string) (map[string][]string, error) {
	return getDependencyGraph(repoPath)
}
