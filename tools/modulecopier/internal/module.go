package internal

import (
	"fmt"
	"github.com/markbates/pkger"
	"github.com/pkg/errors"
	"golang.org/x/mod/modfile"
	"os"
	"path"
	"path/filepath"
)

const modFileName = "go.mod"

// GetModulePath gets the module path for a dependency
// for example, for ethereum, dependencyName would be github.com/ethereum/go-ethereum
// note: we keep this in place w/ packager so tests can determine valid resolution
// pkger is safe to use directly with the tests in place.
func GetModulePath(dependencyName string) (modPath string, err error) {
	modFile, err := getModfile()
	if err != nil {
		return "", err
	}

	// make sure the module is not a replace which we don't have functionality for yet
	if _, err := hasUnsupportedDirective(modFile, dependencyName); err != nil {
		return "", fmt.Errorf("module has unupoorted directive: %w", err)
	}

	var resolvedModule *modfile.Require
	for _, mod := range modFile.Require {
		// this is our module
		if mod.Mod.Path == dependencyName {
			resolvedModule = mod
		}
	}

	if resolvedModule == nil {
		return "", fmt.Errorf("could not find module at %s in go.mod", dependencyName)
	}

	// now we use pkger to resolve the module name. If we could've done this the whole time, why didn't we?
	// a) we need the module included in the go.mod so we don't have to run go mod tidy after generation.
	//   pkger is go module aware, but it's user friendliness comes at a cost. It'll try to import
	//   things that aren't in the modules file
	// b) pkger will not handle replaces: see the above check
	depModFile, err := pkger.Open(fmt.Sprintf("%s/:go.mod", dependencyName))
	if err != nil {
		return "", fmt.Errorf("pkger could not resolve go.mod file: %w", err)
	}
	resolvedModFile := path.Join(depModFile.Info().Dir, modFileName)

	//nolint: gosec
	depModFileContents, err := os.ReadFile(resolvedModFile)
	if err != nil {
		return "", fmt.Errorf("could not read resolved module file at %s: %w", depModFile.Path().String(), err)
	}

	// parse the resolved module file
	parsedFile, err := modfile.Parse(depModFile.Path().String(), depModFileContents, nil)
	if err != nil {
		return "", fmt.Errorf("could not read mod file: %w", err)
	}

	if parsedFile.Module.Mod.Path != depModFile.Info().Module.Path {
		return "", fmt.Errorf("incorrect module resolved at path %s, expected: %s got %s", depModFile.Path().String(),
			parsedFile.Module.Mod.String(),
			resolvedModule.Mod.String())
	}

	return depModFile.Info().Dir, nil
}

// hasUnsupportedDirective checks if the module is either a replace or exclude which are not currently supported
// note: there's no reason they can't be. We just don't use them at all yet.
func hasUnsupportedDirective(modFile *modfile.File, dependencyName string) (ok bool, err error) {
	for _, mod := range modFile.Replace {
		if mod.Old.Path == dependencyName {
			return true, errors.New("replaced modules are not currently supported")
		}
	}

	for _, mod := range modFile.Exclude {
		if mod.Mod.Path == dependencyName {
			return true, errors.New("excluded modules are not currently supported")
		}
	}
	return false, nil
}

// findModPath recursively searches parent directories for the module path.
// Throws an error if it hits a breakpoint (either due to permissions or getting to repo root).
func findModPath() (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get current path: %w", err)
	}

	for {
		exists := true

		prospectiveFile := filepath.Join(currentPath, modFileName)

		if _, err := os.Stat(prospectiveFile); os.IsNotExist(err) {
			exists = false
		}

		if !exists {
			lastPath := currentPath
			currentPath = filepath.Dir(currentPath)

			if lastPath == currentPath {
				return "", errors.New("could not find go.mod file")
			}

			continue
		}

		return prospectiveFile, nil
	}
}

// getModFile gets the module file from the root of the repo. It returns an error if the module cannot be found.
func getModfile() (*modfile.File, error) {
	modFile, err := findModPath()
	if err != nil {
		return nil, fmt.Errorf("could not get modfile: %w", err)
	}

	// read the file
	//nolint: gosec
	modContents, err := os.ReadFile(modFile)
	if err != nil {
		return nil, fmt.Errorf("could not read modfile: %w", err)
	}

	parsedFile, err := modfile.Parse(modFile, modContents, nil)
	if err != nil {
		return nil, fmt.Errorf("could not parse mod file")
	}

	return parsedFile, nil
}
