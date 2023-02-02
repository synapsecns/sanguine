// Package tree implements a tree data structure for representing a directory structure.
//
// It provides methods for adding paths, checking for the presence of a path, generating
// a string representation of the tree, getting all paths in the tree, and retrieving a random path.
//
// It leverages the github.com/xlab/treeprint package to generate the string representation.
// The implementation is concurrent-safe through the use of a sync.RWMutex.
package tree

import (
	"github.com/xlab/treeprint"
	"strings"
	"sync"
)

// Tree is a tree data structure for representing a directory structure.
type Tree interface {
	// Add adds a path to the tree
	Add(paths ...string)
	// HasPath checks if any of the paths in the tree contain the given path
	HasPath(path string) bool
	// ToString gets a string representation of the tree
	ToString() string
	// AllPaths gets all paths in the tree. This will include paths that might not have been explicitly added
	// e.g. if you Add() a/b/c/d.go, AllPaths() will return a/b/c/d.go, a/b/c a/b/ and a/
	AllPaths() []string
}

type treeImpl struct {
	nodes []Node
	mux   sync.RWMutex
}

// NewTree creates a new tree.
func NewTree() Tree {
	return &treeImpl{}
}

// Node is a node in the tree.
type Node struct {
	Name     string `json:"name"`
	Children []Node `json:"children,omitempty"`
}

// Add adds a path to the tree and all of it's component paths to the tree.
func (t *treeImpl) Add(paths ...string) {
	t.mux.Lock()
	defer t.mux.Unlock()

	for _, path := range paths {
		path = strings.TrimPrefix(path, "./")

		t.nodes = addToTree(t.nodes, strings.Split(path, "/"))
	}
}

// addToTree adds a path to the tree. It takes the current root node and an array of names (the names of nodes in the path).
// If a node with a given name already exists in the tree, it is used. If not, a new node with that name is added.
// Returns the updated tree root.
func addToTree(root []Node, names []string) []Node {
	if len(names) > 0 {
		var i int
		for i = 0; i < len(root); i++ {
			if root[i].Name == names[0] { // already in tree
				break
			}
		}
		if i == len(root) {
			root = append(root, Node{Name: names[0]})
		}
		root[i].Children = addToTree(root[i].Children, names[1:])
	}
	return root
}

func (t *treeImpl) HasPath(path string) bool {
	t.mux.RLock()
	defer t.mux.RUnlock()
	// trim a prefix slash if it exists
	path = strings.TrimPrefix(path, "./")
	path = strings.TrimPrefix(path, "/")
	for i := 0; i < len(t.nodes); i++ {
		if has(t.nodes[i], path) {
			return true
		}
	}
	return false
}

// has checks wether a path is in the tree recursively
// path is the path to check and nodePath is any previosus paths that belong to the node
// nodePath is passed in as an array so that it can be left blank, there should never be more than 1 node path.
func has(node Node, path string, nodePath ...string) bool {
	newNodePath := append(nodePath, node.Name)
	newNodePathStr := strings.Join(newNodePath, "/")

	if !strings.HasPrefix(path, newNodePathStr) {
		return false
	}

	if newNodePathStr == path {
		return true
	}
	for i := 0; i < len(node.Children); i++ {
		if has(node.Children[i], path, newNodePath...) {
			return true
		}
	}
	return false
}

func (t *treeImpl) AllPaths() []string {
	t.mux.RLock()
	defer t.mux.RUnlock()

	var allPaths []string
	for i := 0; i < len(t.nodes); i++ {
		allPaths = append(allPaths, allPathsFromNode(t.nodes[i])...)
	}
	return allPaths
}

func allPathsFromNode(node Node, nodePaths ...string) []string {
	var allPaths []string
	newNodePaths := append(nodePaths, node.Name)
	newNodePathsStr := strings.Join(newNodePaths, "/")
	allPaths = append(allPaths, newNodePathsStr)
	for i := 0; i < len(node.Children); i++ {
		allPaths = append(allPaths, allPathsFromNode(node.Children[i], newNodePaths...)...)
	}
	return allPaths
}

func (t *treeImpl) ToString() string {
	t.mux.RLock()
	defer t.mux.RUnlock()
	printer := treeprint.New()
	for _, node := range t.nodes {
		printer = addPrinterNode(printer, node)
	}
	return printer.String()
}

// addPrinterNode adds a node to the printer tree.
func addPrinterNode(tree treeprint.Tree, node Node) treeprint.Tree {
	newNode := tree.AddBranch(node.Name)
	for _, child := range node.Children {
		addPrinterNode(newNode, child)
	}
	return tree
}
