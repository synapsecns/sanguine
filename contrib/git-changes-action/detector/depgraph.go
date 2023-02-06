package detector

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kendru/darwin/go/depgraph"
	"github.com/vishalkuo/bimap"
	"golang.org/x/mod/modfile"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// getDependencyGraph returns a dependency graph of all the modules in the go.work file that refer to other modules in the go.work file
// returns a map of module (./my_module)->(./my_module_dependency1,./my_module_dependency2).
// nolint: cyclop
func getDependencyGraph(repoPath string) (moduleDeps map[string][]string, err error) {
	moduleDeps = make(map[string][]string)
	// parse the go.work file
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

	// map of module->dependencies + replaces
	var dependencies map[string][]string
	// bidirectional map of module->module name
	var moduleNames *bimap.BiMap[string, string]

	// iterate through each module in the go.work file
	// create a list of dependencies for each module
	// and module names
	dependencies, moduleNames, err = makeDepMaps(repoPath, parsedWorkFile.Use)
	if err != nil {
		return nil, fmt.Errorf("failed to create dependency maps: %w", err)
	}

	depGraph := depgraph.New()
	// build the dependency graph
	for _, module := range parsedWorkFile.Use {
		moduleDeps := dependencies[module.Path]
		for _, dep := range moduleDeps {
			// check if the full package name (e.g. github.com/myorg/myrepo/mymodule) is in the list of modules. If it is, add it as a dependency after renaming
			renamedDep, hasDep := moduleNames.GetInverse(dep)
			if hasDep {
				err = depGraph.DependOn(module.Path, renamedDep)
				if err != nil {
					return nil, fmt.Errorf("failed to add dependency %s -> %s: %w", module.Path, dep, err)
				}
			}

			if isRelativeDep(dep) {
				// if the dependency is relative, add it as a dependency
				err = depGraph.DependOn(module.Path, dep)
				if err != nil {
					return nil, fmt.Errorf("failed to add dependency %s -> %s: %w", module.Path, dep, err)
				}
			}
		}
	}

	for _, module := range parsedWorkFile.Use {
		for dep := range depGraph.Dependencies(module.Path) {
			moduleDeps[module.Path] = append(moduleDeps[module.Path], dep)
		}
	}

	return moduleDeps, nil
}

// makeDepMaps makes a dependency map and a bidirectional map of dep<->module.
func makeDepMaps(repoPath string, uses []*modfile.Use) (dependencies map[string][]string, moduleNames *bimap.BiMap[string, string], err error) {
	// map of module->dependencies + replaces
	dependencies = make(map[string][]string)
	// bidirectional map of module->module name
	moduleNames = bimap.NewBiMap[string, string]()

	// iterate through each module in the go.work file
	// create a list of dependencies for each module
	// and module names
	for _, module := range uses {
		//nolint: gosec
		modContents, err := os.ReadFile(filepath.Join(repoPath, module.Path, "go.mod"))
		if err != nil {
			return dependencies, moduleNames, fmt.Errorf("failed to read module file %s: %w", module.Path, err)
		}

		parsedModFile, err := modfile.Parse(module.Path, modContents, nil)
		if err != nil {
			return dependencies, moduleNames, fmt.Errorf("failed to parse module file %s: %w", module.Path, err)
		}

		moduleNames.Insert(module.Path, parsedModFile.Module.Mod.Path)

		dependencies[module.Path] = make([]string, 0)
		// include all requires and replaces, as they are dependencies
		for _, require := range parsedModFile.Require {
			dependencies[module.Path] = append(dependencies[module.Path], convertRelPath(repoPath, module.Path, require.Mod.Path))
		}
		for _, require := range parsedModFile.Replace {
			dependencies[module.Path] = append(dependencies[module.Path], convertRelPath(repoPath, module.Path, require.New.Path))
		}
	}

	return dependencies, moduleNames, nil
}

// isRelativeDep returns true if the dependency is relative to the module (starts with ./ or ../).
func isRelativeDep(path string) bool {
	return strings.HasPrefix(path, "./") || strings.HasPrefix(path, "../")
}

// convertRelPath converts a path relative to a module to a path relative to the repository root.
// it does nothing if the path does not start with ./ or ../.
func convertRelPath(repoPath string, modulePath, dependency string) string {
	if !isRelativeDep(dependency) {
		return dependency
	}

	// repo/./module => repo/module
	fullModulePath := filepath.Join(repoPath, modulePath)
	// repo/module/../dependency => repo/dependency
	fullDependencyPath := filepath.Join(fullModulePath, dependency)
	// repo/dependency => dependency
	trimmedPath := strings.TrimPrefix(fullDependencyPath, repoPath)
	if len(trimmedPath) == 0 {
		return "."
	}

	return fmt.Sprintf(".%s", trimmedPath)
}
