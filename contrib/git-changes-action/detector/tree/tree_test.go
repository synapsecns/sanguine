package tree_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/git-changes-action/detector/tree"
	"testing"
)

func TestTree(t *testing.T) {
	testPaths := []string{
		"a/b/c/d.go",
		"a/b/g",
		"a/d",
	}

	testTree := tree.NewTree()
	testTree.Add(testPaths...)

	// add again to make sure it doesn't add duplicates
	testTree.Add(testPaths...)

	for _, path := range testPaths {
		if !testTree.HasPath(path) {
			t.Errorf("tree should have path %s", path)
		}
	}

	if testTree.HasPath("a/b/c/d/e.go") {
		t.Errorf("tree should not have path a/b/c/d/e.go")
	}

	if !testTree.HasPath("a/b/c") {
		t.Errorf("tree should have path a/b/c")
	}

	expectedOutput := `.
└── a
    ├── b
    │   ├── c
    │   │   └── d.go
    │   └── g
    └── d
`

	Equal(t, expectedOutput, testTree.ToString())
	Equal(t, 6, len(testTree.AllPaths()))

	for _, path := range testPaths {
		Contains(t, testTree.AllPaths(), path)
	}
}
