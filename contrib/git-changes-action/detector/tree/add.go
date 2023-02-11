package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// AddDirectoryPaths is a helper function that adds all paths in a directory to the tree
// you can optionally pass in a prefix to trim from the paths.
func AddDirectoryPaths(tree Tree, dirPath, trimPrefix string) error {
	err := filepath.Walk(dirPath, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking path %s: %w", path, err)
		}

		if trimPrefix != "" {
			path = path[len(trimPrefix):]
		}

		// skip git files
		if strings.Contains(path, ".git") {
			return nil
		}

		tree.Add(path)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error adding directory paths: %w", err)
	}
	return nil
}
