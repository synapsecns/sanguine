package moduledetector

func GetDependencyDag(repoPath string) (map[string][]string, error) {
	return getDependencyGraph(repoPath)
}
