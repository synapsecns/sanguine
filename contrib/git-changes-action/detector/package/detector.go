package packagedetector

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
	"golang.org/x/mod/modfile"
	"os"
	"path"
)

// DetectChangedModules is the change detector client.
// Modules will be flagged as changed if any of the packaged in their dependency tree has been modified.
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

	depGraph, packagesPerModule, err := getPackageDependencyGrap(repoPath)
	if err != nil {
		return nil, fmt.Errorf("could not get dep graph: %w", err)
	}

	for _, module := range parsedWorkFile.Use {
		changed := false

		if ct.HasPath(module.Path) {
			changed = true
		}

		if includeDeps && !changed {
			for _, packageName := range packagesPerModule[module.Path] {
				if isPackageChanged(packageName, ct, depGraph) {
					changed = true
					// If a package is flagged as changed
					// its not necessary to analyze the remaining packages,
					// module can be flagged as changed
					break
				}
			}
		}

		modules[module.Path] = changed
	}

	return modules, nil
}

func isPackageChanged(packageName string, ct tree.Tree, depGraph map[string][]string) bool {
	if ct.HasPath(packageName) {
		return true
	}

	for _, dep := range depGraph[packageName] {
		if ct.HasPath(dep) {
			return true
		}
	}
	return false
}
