package detector

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kendru/darwin/go/depgraph"
	"github.com/vishalkuo/bimap"
	"golang.org/x/mod/modfile"
)

// getDependencyGraph returns a dependency graph of all the modules in the go.work file that refer to other modules in the go.work file
// returns a map of module (./my_module)->(./my_module_dependency1,./my_module_dependency2).
// nolint: cyclop
func getDependencyGraph(repoPath string, typeOfDependency string) (moduleDeps map[string][]string, packagesPerModule map[string][]string, err error) {
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

	// map of module->dependencies + replaces
	var dependencies map[string]map[string]struct{}
	// bidirectional map of module->module name or dependency->dependency
	var dependencyNames *bimap.BiMap[string, string]

  // iterate through each module in the go.work file
	// create a list of dependencies for each module
	// and module names
	dependencies, dependencyNames, packagesPerModule, err = makeDepMaps(repoPath, parsedWorkFile.Use, typeOfDependency)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create dependency maps: %w", err)
	}

	depGraph := depgraph.New()
	// build the dependency graph
  if typeOfDependency == "modules" {
    for _, module := range parsedWorkFile.Use {
      for dep, _ := range dependencies[module.Path] {
        // check if the full package name (e.g. github.com/myorg/myrepo/mymodule) is in the list of modules. If it is, add it as a dependency after renaming
        renamedDep, hasDep := dependencyNames.GetInverse(dep)
        if hasDep {
          err = depGraph.DependOn(module.Path, renamedDep)
          if err != nil {
            return nil, nil, fmt.Errorf("failed to add dependency %s -> %s: %w", module.Path, dep, err)
          }
        }

        if isRelativeDep(dep) {
          // if the dependency is relative, add it as a dependency
          err = depGraph.DependOn(module.Path, dep)
          if err != nil {
            return nil, nil, fmt.Errorf("failed to add dependency %s -> %s: %w", module.Path, dep, err)
          }
        }
      }
    }

    for _, module := range parsedWorkFile.Use {
      for dep := range depGraph.Dependencies(module.Path) {
        moduleDeps[module.Path] = append(moduleDeps[module.Path], dep)
      }
    }
  }

  if typeOfDependency == "packages" {
    for _, module := range parsedWorkFile.Use {
      for _, relativePackageName := range packagesPerModule[module.Path] {
          for relativePackageDependencyName, _ := range dependencies[relativePackageName] {
            err = depGraph.DependOn(relativePackageName, relativePackageDependencyName) 
            // Circular dependencies are fine as long as both packages are in the same module
            if err != nil && !(strings.Contains(relativePackageDependencyName, module.Path) && strings.Contains(relativePackageName, module.Path)){
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
  }

	return moduleDeps, packagesPerModule, nil
}

func extractGoFileNames(pwd string, currentModule string, currentPackage string, goFiles map[string][]string) (err error) {
  ls, err := os.ReadDir(pwd)
  if err != nil {
    return err
  }

  for _, entry := range ls {
    if entry.IsDir() {
      extractGoFileNames(pwd + "/" + entry.Name(), currentModule + "/" + entry.Name(), entry.Name(), goFiles)
    } else if strings.HasSuffix(entry.Name(), ".go") {
      fileName := pwd + "/" + entry.Name()
      var packageName string
      if currentModule == "" {
        packageName = "/" + currentPackage
      } else {
        packageName = currentModule 
      }
      goFiles[packageName] = append(goFiles[packageName], fileName)
    }
  }

  return nil
}

// makeDepMaps makes a dependency map and a bidirectional map of dep<->module.
func makeDepMaps(repoPath string, uses []*modfile.Use, typeOfDependency string) (dependencies map[string]map[string]struct{}, dependencyNames *bimap.BiMap[string, string], packagesPerModule map[string][]string, err error) {
  // Can be either:
	// map of module->depndencies
  // map of packages -> dependencies
  // depends on typeOfDependency
	dependencies = make(map[string]map[string]struct{})

	// bidirectional map of module->module name
  // bidirectional map of package->package name
  // Maps relative to public names, used to filer out all external libraries/packages.
	dependencyNames = bimap.NewBiMap[string, string]()

  // map of module->packages
  packagesPerModule = make(map[string][]string)

	// iterate through each module in the go.work file
	// 1. Create a list of dependencies for each module 
  // 2. Map public names to private names for each module
		//nolint: gose
    if typeOfDependency == "modules" {
    	for _, module := range uses {
        modContents, err := os.ReadFile(filepath.Join(repoPath, module.Path, "go.mod"))
        if err != nil {
          return dependencies, dependencyNames, packagesPerModule, fmt.Errorf("failed to read module file %s: %w", module.Path, err)
        }

        parsedModFile, err := modfile.Parse(module.Path, modContents, nil)
        if err != nil {
          return dependencies, dependencyNames, packagesPerModule, fmt.Errorf("failed to parse module file %s: %w", module.Path, err)
        }

        dependencyNames.Insert(module.Path, parsedModFile.Module.Mod.Path)
        dependencies[module.Path] = make(map[string]struct{})
      
        // include all requires and replaces, as they are dependencies
        for _, require := range parsedModFile.Require {
          dependencies[module.Path][convertRelPath(repoPath, module.Path, require.Mod.Path)] = struct{}{}
        }
        for _, require := range parsedModFile.Replace {
          dependencies[module.Path][convertRelPath(repoPath,module.Path, require.New.Path)] = struct{}{} 
        }
      }
    }


    if typeOfDependency == "packages" {
      extractedGoFileNames := make(map[string]map[string][]string)

      pwd, err := os.Getwd()
      if err != nil {
        return dependencies, dependencyNames, packagesPerModule, fmt.Errorf("Failed to read current directory: %w", err)
      }


	// iterate through each module in the go.work file
  // 1. Extract all go files filepaths for each package.
  // 2. Create a map where key is module, value is an array with all packages in the module
  // 3. Map public name to relative name for each package (used to filter external library/package imports)
      for _, module := range uses {
        extractedGoFileNames[module.Path] = make(map[string][]string)

        modContents, err := os.ReadFile(filepath.Join(repoPath, module.Path, "go.mod"))
        if err != nil {
          return dependencies, dependencyNames, packagesPerModule, fmt.Errorf("failed to read module file %s: %w", module.Path, err)
        }

        parsedModFile, err := modfile.Parse(module.Path, modContents, nil)
        if err != nil {
          return dependencies, dependencyNames, packagesPerModule, fmt.Errorf("failed to parse module file %s: %w", module.Path, err)
        }

        // module.Path = ./moduleName
        err = extractGoFileNames(pwd + module.Path[1:], "", module.Path[2:], extractedGoFileNames[module.Path])
        if err != nil {
          return dependencies, dependencyNames, packagesPerModule, fmt.Errorf("failed to extract go files for module %s: %w", module.Path, err)
        }

        for packageName, _ := range extractedGoFileNames[module.Path] {
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
              return dependencies, dependencyNames, packagesPerModule, fmt.Errorf("failed to parse go file %s in package %s: %w", file, relativePackageName, err)
            }

            for _, s := range f.Imports {
              // s.Path.Value contains double quotation marks that must be removed before indexing dependencyNames 
              renamedDep, hasDep := dependencyNames.GetInverse(s.Path.Value[1:len(s.Path.Value)-1])

              if hasDep && (relativePackageName != renamedDep){
                dependencies[relativePackageName][renamedDep] = struct{}{} 
              }
            }
          }
        }
      }
    }

	return dependencies, dependencyNames, packagesPerModule, nil
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
