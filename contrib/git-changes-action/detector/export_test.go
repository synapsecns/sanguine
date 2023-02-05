package detector

import (
	"github.com/go-git/go-git/v5"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/actionscore"
)

func GetDependencyDag(repoPath string) (map[string][]string, error) {
	return getDependencyGraph(repoPath)
}

func GetHead(repo *git.Repository, ghContext *actionscore.Context, head string) (string, error) {
	return getHead(repo, ghContext, head)
}
