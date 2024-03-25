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
  if typeOfDependency == "module" {
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
      allPackagesInModule := packagesPerModule[module.Path]

      for _, packageInModule := range allPackagesInModule {
        renamedPackage, hasPackage := dependencyNames.Get(packageInModule) 
        if hasPackage {
          for dep, _ := range dependencies[renamedPackage] {
            dep = strings.TrimPrefix(dep, `"`)
            dep = strings.TrimSuffix(dep, `"`)

            renamedDep, hasDep := dependencyNames.GetInverse(dep)
            if hasDep {
              err = depGraph.DependOn(packageInModule, renamedDep) 
              if err != nil {
                fmt.Println("THERE IS AN ERROR", err, packageInModule, renamedDep)
              }
            }
          }
        }

        for dep := range depGraph.Dependencies(packageInModule) {
          moduleDeps[packageInModule] = append(moduleDeps[packageInModule], dep)
        }
      }
    }
  }

	return moduleDeps, packagesPerModule, nil
}

func extractGoFileNames(pwd string, currentModule string, currentPackage string, goFiles map[string][]string) {
  ls, err := os.ReadDir(pwd)
  if err != nil {
  }

  for _, entry := range ls {
    if entry.IsDir() {
      extractGoFileNames(pwd + "/" + entry.Name(), currentModule + "/" + entry.Name(), entry.Name(), goFiles)
    } else if strings.Contains(entry.Name(), ".go") {
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
}

// makeDepMaps makes a dependency map and a bidirectional map of dep<->module.
func makeDepMaps(repoPath string, uses []*modfile.Use, typeOfDependency string) (dependencies map[string]map[string]struct{}, dependencyNames *bimap.BiMap[string, string], packagesPerModule map[string][]string, err error) {
	// map of module->dependencies + replaces
  // map of packages -> dependencies
	dependencies = make(map[string]map[string]struct{})
	// bidirectional map of module->module name
  // bidirectional map of package->package name, relative to public names.
	dependencyNames = bimap.NewBiMap[string, string]()
  // map of module->packages
  packagesPerModule = make(map[string][]string)

	// iterate through each module in the go.work file
	// create a list of dependencies for each module based on modules or packages
	// and module names or package names
		//nolint: gose

    if typeOfDependency == "module" {
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
      }


      for _, module := range uses {
        extractedGoFileNames[module.Path[1:]] = make(map[string][]string)

        modContents, err := os.ReadFile(filepath.Join(repoPath, module.Path, "go.mod"))
        if err != nil {
          return dependencies, dependencyNames, packagesPerModule, fmt.Errorf("failed to read module file %s: %w", module.Path, err)
        }

        parsedModFile, err := modfile.Parse(module.Path, modContents, nil)
        if err != nil {
          return dependencies, dependencyNames, packagesPerModule, fmt.Errorf("failed to parse module file %s: %w", module.Path, err)
        }


        extractGoFileNames(pwd + module.Path[1:], "", module.Path[2:], extractedGoFileNames[module.Path[1:]])

        for packageName, _ := range extractedGoFileNames[module.Path[1:]] {
          var relativePackageName string
          if strings.HasSuffix(module.Path[1:], packageName) {
            relativePackageName = module.Path[1:] 
          } else {
            relativePackageName = module.Path[1:] + packageName
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

      for _, module := range uses {
        for packageInModule, files := range extractedGoFileNames[module.Path[1:]] {
          publicPackageName, _ := dependencyNames.Get(module.Path[1:] + packageInModule)

          dependencies[publicPackageName] = make(map[string]struct{})
          for _, file := range files {
            fset := token.NewFileSet()
            f, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)

            if err != nil {
            }

            for _, s := range f.Imports {
              _, hasDep := dependencyNames.GetInverse(s.Path.Value[1:len(s.Path.Value)-1])

              if hasDep {
                dependencies[publicPackageName][s.Path.Value] = struct{}{} 
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
