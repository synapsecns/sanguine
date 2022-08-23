package internal_test

import (
	"bytes"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/tools/modulecopier/internal"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
)

// TestCopyModule runs some sanity checks on the copy process.
func (s GeneratorSuite) TestCopyModule() {
	newPackageName := gofakeit.Word()

	destDir := filet.TmpDir(s.T(), "")
	err := internal.CopyModule("github.com/ethereum/go-ethereum/accounts/abi/bind/backends", destDir, newPackageName)
	Nil(s.T(), err)

	// run some sanity checks on the resulting dir. This is by no means complete, but this is a testutil
	err = filepath.WalkDir(destDir, func(path string, d fs.DirEntry, err error) error {
		Nil(s.T(), err)
		// skip the tld
		if d.IsDir() {
			return nil
		}

		// make sure file is not empty
		//nolint: staticcheck
		fileInfo, err := d.Info()
		Nil(s.T(), err)

		NotZero(s.T(), fileInfo.Size())

		s.validateGoFile(path, newPackageName)

		return nil
	})
	Nil(s.T(), err)
}

func (s GeneratorSuite) TestCopyFile() {
	newPackageName := gofakeit.Word()
	destDir := filet.TmpDir(s.T(), "")
	err := internal.CopyFile("github.com/ethereum/go-ethereum/ethclient/signer.go", destDir, newPackageName)
	Nil(s.T(), err)

	path := filepath.Join(destDir, "signer_gen.go")

	s.validateGoFile(path, newPackageName)
}

// validateGoFile validates that the file was correctly copied with the correct prefix.
func (s GeneratorSuite) validateGoFile(path, packageName string) {
	//nolint: gosec
	src, err := os.ReadFile(path)
	Nil(s.T(), err)

	True(s.T(), bytes.Contains(src, []byte("DO NOT EDIT")))

	fset := token.NewFileSet()

	// verify package name was correctly changed
	ast, err := parser.ParseFile(fset, filepath.Base(path), src, parser.PackageClauseOnly)
	Nil(s.T(), err)

	realPackageName := ast.Name.Name
	Equal(s.T(), realPackageName, packageName)
}
