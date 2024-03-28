package packagedetector

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kendru/darwin/go/depgraph"
	"github.com/vishalkuo/bimap"
	"go/parser"
	"go/token"
	"golang.org/x/mod/modfile"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// nolint: cyclop
func getPackageDependencyGrap(repoPath string) (moduleDeps map[string][]string, packagesPerModule map[string][]string, err error) {
	moduleDeps = make(map[string][]string)
	// parse the go.work file
	goWorkPath := path.Join(repoPath, "go.work")

	if !common.FileExist(goWorkPath) {
		return nil, nil, fmt.Errorf("go.work file not found in %s", repoPath)
	}

	//nolint: gosec
	workFile, err := os.ReadFile(goWorkPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read go.work file: %w", err)
	}

	parsedWorkFile, err := modfile.ParseWork(goWorkPath, workFile, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse go.work file: %w", err)
	}

	// map of package->dependencies
	var dependencies map[string]map[string]struct{}

	// iterate through each module in the go.work file
	// create a list of dependencies for each package 
	// and generate a list of packages per module 
  // nolint: gocognit
	dependencies, packagesPerModule, err = makePackageDepMaps(repoPath, parsedWorkFile.Use)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create dependency maps: %w", err)
	}

	depGraph := depgraph.New()

	for _, module := range parsedWorkFile.Use {
		for _, relativePackageName := range packagesPerModule[module.Path] {
			for relativePackageDependencyName := range dependencies[relativePackageName] {
				err = depGraph.DependOn(relativePackageName, relativePackageDependencyName)
				// Circular dependencies are fine as long as both packages are in the same module
				if err != nil && !(strings.Contains(relativePackageDependencyName, module.Path) && strings.Contains(relativePackageName, module.Path)) {
					return nil, nil, fmt.Errorf("failed to add dependency %s -> %s: %w", relativePackageName, relativePackageDependencyName, err)
				}
			}
		}
	}

	for _, module := range parsedWorkFile.Use {
		for _, relativePackageName := range packagesPerModule[module.Path] {
			for dep := range depGraph.Dependencies(relativePackageName) {
				moduleDeps[relativePackageName] = append(moduleDeps[relativePackageName], dep)
			}
		}
	}
	return moduleDeps, packagesPerModule, nil
}

func extractGoFileNames(pwd string, currentPackage string, goFiles map[string][]string) (err error) {
	searchNext := make(map[string]string)
	_, packageDir := path.Split(currentPackage)
	searchNext[pwd] = packageDir

	for len(searchNext) > 0 {
		discovered := make(map[string]string)
		for path, dirName := range searchNext {
			err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() && !(path == filePath) {
					discovered[filePath] = info.Name()
					return filepath.SkipDir
				} else if strings.HasSuffix(info.Name(), ".go") {
					goFiles["/"+dirName] = append(goFiles["/"+dirName], filePath)
				}

				return nil
			})

			if err != nil {
				return fmt.Errorf("failed to walk path: %w", err)
			}
		}
		searchNext = discovered
	}
	return nil
}

// nolint: gocognit, cyclop
func makePackageDepMaps(repoPath string, uses []*modfile.Use) (dependencies map[string]map[string]struct{}, packagesPerModule map[string][]string, err error) {
	// map of packages -> dependencies
	dependencies = make(map[string]map[string]struct{})

	// bidirectional map of package->package name
	// Maps relative to public names, used to filer out all external libraries/packages.
	dependencyNames := bimap.NewBiMap[string, string]()

	// map of module->packages
	packagesPerModule = make(map[string][]string)

	// map module->package->goFiles
	// Maps each module to all packages and each package to its go files.
	extractedGoFileNames := make(map[string]map[string][]string)

	pwd, err := os.Getwd()
	if err != nil {
		return dependencies, packagesPerModule, fmt.Errorf("failed to read current directory: %w", err)
	}
	// iterate through each module in the go.work file
	// 1. Extract all go files filepaths for each package.
	// 2. Create a map where key is module, value is an array with all packages in the module
	// 3. Map public name to relative name for each package (used to filter external library/package imports)
	for _, module := range uses {
		// nolint: gosec
		modContents, err := os.ReadFile(filepath.Join(repoPath, module.Path, "go.mod"))
		if err != nil {
			return dependencies, packagesPerModule, fmt.Errorf("failed to read module file %s: %w", module.Path, err)
		}

		parsedModFile, err := modfile.Parse(module.Path, modContents, nil)
		if err != nil {
			return dependencies, packagesPerModule, fmt.Errorf("failed to parse module file %s: %w", module.Path, err)
		}

		extractedGoFileNames[module.Path] = make(map[string][]string)
		err = extractGoFileNames(pwd+module.Path[1:], module.Path[1:], extractedGoFileNames[module.Path])
		if err != nil {
			return dependencies, packagesPerModule, fmt.Errorf("failed to extract go files for module %s: %w", module.Path, err)
		}

		for packageName := range extractedGoFileNames[module.Path] {
			var relativePackageName string
			if strings.HasSuffix(module.Path, packageName) {
				relativePackageName = module.Path
			} else {
				relativePackageName = module.Path + packageName
			}

			var publicPackageName string
			if strings.HasSuffix(parsedModFile.Module.Mod.Path, packageName) {
				publicPackageName = parsedModFile.Module.Mod.Path
			} else {
				publicPackageName = parsedModFile.Module.Mod.Path + packageName
			}

			packagesPerModule[module.Path] = append(packagesPerModule[module.Path], relativePackageName)
			dependencyNames.Insert(relativePackageName, publicPackageName)
		}
	}

	// iterate through each module in the go.work file
	// For every package in the module
	// using the filepaths extracted on the previous loop, parse files and extract imports.
	// Ignore any external library/package imports
	for _, module := range uses {
		for packageInModule, files := range extractedGoFileNames[module.Path] {
			var relativePackageName string
			if strings.HasSuffix(module.Path, packageInModule) {
				relativePackageName = module.Path
			} else {
				relativePackageName = module.Path + packageInModule
			}

			dependencies[relativePackageName] = make(map[string]struct{})
			for _, file := range files {
				fset := token.NewFileSet()
				f, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)

				if err != nil {
					return dependencies, packagesPerModule, fmt.Errorf("failed to parse go file %s in package %s: %w", file, relativePackageName, err)
				}

				for _, s := range f.Imports {
					// s.Path.Value contains double quotation marks that must be removed before indexing dependencyNames
					renamedDep, hasDep := dependencyNames.GetInverse(s.Path.Value[1 : len(s.Path.Value)-1])

					if hasDep && (relativePackageName != renamedDep) {
						dependencies[relativePackageName][renamedDep] = struct{}{}
					}
				}
			}
		}
	}
	return dependencies, packagesPerModule, nil
}
