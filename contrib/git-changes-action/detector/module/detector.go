package moduledetector

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
	"golang.org/x/mod/modfile"
	"os"
	"path"
)

// DetectChangedModules is the change detector client.
// Will flag modules as changed if any module in their dependency tree has been modified.
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
