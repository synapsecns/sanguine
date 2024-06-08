package internal

import (
	"bytes"
	"fmt"
	"github.com/markbates/pkger"
	"github.com/thoas/go-funk"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// CopyModule copies a module path to a destination.
func CopyModule(toCopy, dest, packageName string) error {
	// walk through the dir, see: https://github.com/markbates/pkger/blob/09e9684b656b/examples/app/main.go#L29
	info, err := pkger.Info(toCopy)
	if err != nil {
		return fmt.Errorf("could not resolve %s", toCopy)
	}

	// get the go files to copy
	goFiles := append(info.GoFiles, info.TestGoFiles...)

	err = pkger.Walk(toCopy, func(filePath string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error while walking: %w", err)
		}
		// if it's not a go file, skip it
		if !funk.ContainsString(goFiles, info.Name()) {
			return nil
		}

		return copyGoFile(filePath, packageName, dest, info)
	})

	if err != nil {
		return fmt.Errorf("error while copying: %w", err)
	}
	return nil
}

// copyGoFile copies a go file using the package info.
func copyGoFile(filePath, packageName, dest string, info fs.FileInfo) error {
	fileContents, err := getUpdatedFileContents(filePath, packageName)
	if err != nil {
		return fmt.Errorf("could not get updated file contents: %w", err)
	}

	newFile := fmt.Sprintf("%s/%s", dest, getFileName(info.Name()))
	//nolint: gosec
	f, err := os.Create(newFile)
	if err != nil {
		return fmt.Errorf("could not open file")
	}

	// write the contents to the file
	_, err = f.Write(fileContents)
	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("could not close file: %w", err)
	}

	return nil
}

// CopyFile copies a single go file. This will not bring dependencies.
func CopyFile(fileToCopy, dest, packageName string) error {
	// first things first, pkger operates on go modules, so we need to trim
	modulePath := path.Dir(fileToCopy)
	fileName := path.Base(fileToCopy)

	// make sure the last element is a file
	if filepath.Ext(fileName) != ".go" {
		return fmt.Errorf("must specify a .go file after module, got %s", filepath.Ext(fileName))
	}

	err := pkger.Walk(modulePath, func(filePath string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error while walking: %w", err)
		}

		// only copy the target file
		if info.Name() != fileName {
			return nil
		}

		return copyGoFile(filePath, packageName, dest, info)
	})

	if err != nil {
		return fmt.Errorf("error while copying: %w", err)
	}

	return nil
}

// getFileName gets the new file name. Gen is added here before the .go in the case of non tests
// and before _test.go in the case of tests.
func getFileName(originalName string) string {
	suffix := filepath.Ext(originalName)
	noExtensionName := strings.TrimSuffix(originalName, suffix)

	const testSuffix = "_test"

	// if it's a test strip it from the original name and add it to the suffix
	testIndex := strings.LastIndex(noExtensionName, testSuffix)
	if testIndex != -1 {
		noExtensionName = noExtensionName[:testIndex] + strings.Replace(noExtensionName[testIndex:], testSuffix, "", 1)
		suffix = testSuffix + suffix
	}

	return noExtensionName + "_gen" + suffix
}

// getUpdatedFileContents rewrites adds the generation header and rewrites the package name.
func getUpdatedFileContents(path, newPackageName string) (fileContents []byte, err error) {
	file, err := pkger.Open(path)
	if err != nil {
		return fileContents, fmt.Errorf("could not open file at %s: %w", path, err)
	}

	fileContents, err = io.ReadAll(file)
	if err != nil {
		return fileContents, fmt.Errorf("could not read file %s: %w", fileContents, err)
	}

	// prepend the header to the file
	fileContents = append([]byte(makeGeneratedHeader(path)+"\n\n"), fileContents...)

	// rename the package by modifying the ast
	fset := token.NewFileSet()

	fileAst, err := parser.ParseFile(fset, filepath.Base(path), fileContents, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("could not parse ast. This could indicate an invalid source file: %w", err)
	}

	newAst := astutil.Apply(fileAst, nil, func(cursor *astutil.Cursor) bool {
		if ident, ok := cursor.Node().(*ast.Ident); ok {
			cursor.Replace(&ast.Ident{
				NamePos: ident.NamePos,
				Name:    newPackageName,
				Obj:     ident.Obj,
			})
			return false
		}
		return true
	})

	fileBuffer := bytes.NewBuffer([]byte{})
	err = printer.Fprint(fileBuffer, fset, newAst)
	if err != nil {
		return nil, fmt.Errorf("could not write resulting ast: %w", err)
	}

	// TODO: use golangci-lint
	formatted, err := format.Source(fileBuffer.Bytes())
	if err != nil {
		return nil, fmt.Errorf("could not format: %w", err)
	}

	return formatted, nil
}

// makeGenerated header makes the code generation header
// note: this must conform to https://github.com/golangci/golangci-lint/blob/1fb67fe448da8a3fb525ecef28decceb23b42d7a/pkg/result/processors/autogenerated_exclude.go#L76
// to bypass linters.
func makeGeneratedHeader(origin string) string {
	return fmt.Sprintf("// Code copied from %s for testing by synapse modulecopier DO NOT EDIT.\"", origin)
}
