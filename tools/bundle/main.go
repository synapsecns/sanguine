// we skip linting this file because it is largely copied from the standard library and mostly a patch until the bug is fixed in bundle
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Bundle creates a single-source-file version of a source package
// suitable for inclusion in a particular target package.
//
// Usage:
//
//	bundle [-o file] [-dst path] [-pkg name] [-prefix p] [-import old=new] [-tags build_constraints] <src>
//
// The src argument specifies the import path of the package to bundle.
// The bundling of a directory of source files into a single source file
// necessarily imposes a number of constraints.
// The package being bundled must not use cgo; must not use conditional
// file compilation, whether with build tags or system-specific file names
// like code_amd64.go; must not depend on any special comments, which
// may not be preserved; must not use any assembly sources;
// must not use renaming imports; and must not use reflection-based APIs
// that depend on the specific names of types or struct fields.
//
// By default, bundle writes the bundled code to standard output.
// If the -o argument is given, bundle writes to the named file
// and also includes a “//go:generate” comment giving the exact
// command line used, for regenerating the file with “go generate.”
//
// Bundle customizes its output for inclusion in a particular package, the destination package.
// By default bundle assumes the destination is the package in the current directory,
// but the destination package can be specified explicitly using the -dst option,
// which takes an import path as its argument.
// If the source package imports the destination package, bundle will remove
// those imports and rewrite any references to use direct references to the
// corresponding symbols.
// Bundle also must write a package declaration in the output and must
// choose a name to use in that declaration.
// If the -pkg option is given, bundle uses that name.
// Otherwise, the name of the destination package is used.
// Build constraints for the generated file can be specified using the -tags option.
//
// To avoid collisions, bundle inserts a prefix at the beginning of
// every package-level const, func, type, and var identifier in src's code,
// updating references accordingly. The default prefix is the package name
// of the source package followed by an underscore. The -prefix option
// specifies an alternate prefix.
//
// Occasionally it is necessary to rewrite imports during the bundling
// process. The -import option, which may be repeated, specifies that
// an import of "old" should be rewritten to import "new" instead.
//
// # Example
//
// Bundle archive/zip for inclusion in cmd/dist:
//
//	cd $GOROOT/src/cmd/dist
//	bundle -o zip.go archive/zip
//
// Bundle golang.org/x/net/http2 for inclusion in net/http,
// prefixing all identifiers by "http2" instead of "http2_", and
// including a "!nethttpomithttp2" build constraint:
//
//	cd $GOROOT/src/net/http
//	bundle -o h2_bundle.go -prefix http2 -tags '!nethttpomithttp2' golang.org/x/net/http2
//
// Update the http2 bundle in net/http:
//
//	go generate net/http
//
// Update all bundles in the standard library:
//
//	go generate -run bundle std
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/printer"
	"go/token"
	"go/types"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/packages"
)

var (
	outputFile = flag.String("o", "", "write output to `file` (default standard output)")
	dstPath    = flag.String("dst", ".", "set destination import `path`")
	pkgName    = flag.String("pkg", "", "set destination package `name`")
	prefix     = flag.String("prefix", "&_", "set bundled identifier prefix to `p` (default is \"&_\", where & stands for the original name)")
	buildTags  = flag.String("tags", "", "the build constraints to be inserted into the generated file")

	importMap = map[string]string{}
)

func init() {
	flag.Var(flagFunc(addImportMap), "import", "rewrite import using `map`, of form old=new (can be repeated)")
}

func addImportMap(s string) {
	if strings.Count(s, "=") != 1 {
		log.Fatal("-import argument must be of the form old=new")
	}
	i := strings.Index(s, "=")
	old, new := s[:i], s[i+1:]
	if old == "" || new == "" {
		log.Fatal("-import argument must be of the form old=new; old and new must be non-empty")
	}
	importMap[old] = new
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: bundle [options] <src>\n")
	flag.PrintDefaults()
}

func main() {
	log.SetPrefix("bundle: ")
	log.SetFlags(0)

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		usage()
		os.Exit(2)
	}

	cfg := &packages.Config{Mode: packages.NeedName}
	pkgs, err := packages.Load(cfg, *dstPath)
	if err != nil {
		log.Fatalf("cannot load destination package: %v", err)
	}
	if packages.PrintErrors(pkgs) > 0 || len(pkgs) != 1 {
		log.Fatalf("failed to load destination package")
	}
	if *pkgName == "" {
		*pkgName = pkgs[0].Name
	}

	code, err := bundle(args[0], pkgs[0].PkgPath, *pkgName, *prefix, *buildTags)
	if err != nil {
		log.Fatal(err)
	}
	if *outputFile != "" {
		err := ioutil.WriteFile(*outputFile, code, 0666)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := os.Stdout.Write(code)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// isStandardImportPath is copied from cmd/go in the standard library.
func isStandardImportPath(path string) bool {
	i := strings.Index(path, "/")
	if i < 0 {
		i = len(path)
	}
	elem := path[:i]
	return !strings.Contains(elem, ".")
}

var testingOnlyPackagesConfig *packages.Config

func bundle(src, dst, dstpkg, prefix, buildTags string) ([]byte, error) {
	// Load the initial package.
	cfg := &packages.Config{}
	if testingOnlyPackagesConfig != nil {
		*cfg = *testingOnlyPackagesConfig
	} else {
		// Bypass default vendor mode, as we need a package not available in the
		// std module vendor folder.
		var environVars []string
		// list of go keys to copy
		validGoKeys := []string{"GOPATH", "GOROOT", "GO111MODULE", "GOPRIVATE"}
		for _, val := range os.Environ() {
			key := strings.Split(val, "=")[0]

			if slices.Contains(validGoKeys, key) || !strings.HasPrefix(val, "GO") {
				environVars = append(environVars, val)
			}
		}
		cfg.Env = append(environVars)
	}
	cfg.Mode = packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedImports | packages.NeedDeps
	pkgs, err := packages.Load(cfg, src)
	if err != nil {
		return nil, err
	}
	if packages.PrintErrors(pkgs) > 0 || len(pkgs) != 1 {
		return nil, fmt.Errorf("failed to load source package")
	}
	pkg := pkgs[0]

	if strings.Contains(prefix, "&") {
		prefix = strings.Replace(prefix, "&", pkg.Syntax[0].Name.Name, -1)
	}

	objsToUpdate := make(map[types.Object]bool)
	var rename func(from types.Object)
	rename = func(from types.Object) {
		if !objsToUpdate[from] {
			objsToUpdate[from] = true

			// Renaming a type that is used as an embedded field
			// requires renaming the field too. e.g.
			// 	type T int // if we rename this to U..
			// 	var s struct {T}
			// 	print(s.T) // ...this must change too
			if _, ok := from.(*types.TypeName); ok {
				for id, obj := range pkg.TypesInfo.Uses {
					if obj == from {
						if field := pkg.TypesInfo.Defs[id]; field != nil {
							rename(field)
						}
					}
				}
			}
		}
	}

	// Rename each package-level object.
	scope := pkg.Types.Scope()
	for _, name := range scope.Names() {
		rename(scope.Lookup(name))
	}

	var out bytes.Buffer

	// Concatenate package comments from all files...
	for _, f := range pkg.Syntax {
		if doc := f.Doc.Text(); strings.TrimSpace(doc) != "" {
			for _, line := range strings.Split(doc, "\n") {
				fmt.Fprintf(&out, "// %s\n", line)
			}
		}
	}
	// ...but don't let them become the actual package comment.
	fmt.Fprintln(&out)

	fmt.Fprintf(&out, "package %s\n\n", dstpkg)

	// BUG(adonovan,shurcooL): bundle may generate incorrect code
	// due to shadowing between identifiers and imported package names.
	//
	// The generated code will either fail to compile or
	// (unlikely) compile successfully but have different behavior
	// than the original package. The risk of this happening is higher
	// when the original package has renamed imports (they're typically
	// renamed in order to resolve a shadow inside that particular .go file).

	// TODO(adonovan,shurcooL):
	// - detect shadowing issues, and either return error or resolve them
	// - preserve comments from the original import declarations.

	// pkgStd and pkgExt are sets of printed import specs. This is done
	// to deduplicate instances of the same import name and path.
	var pkgStd = make(map[string]bool)
	var pkgExt = make(map[string]bool)
	// renamedImportsFile keeps track of all new aliases
	// alias->ogimport->newimport
	var renamedImportsAlias = make(map[string]map[string]string)
	for _, f := range pkg.Syntax {
		// create a standard alias for every import in the file
		aliasPrefix := getAliasPrefix(pkg, f)
		renamedImportsAlias[aliasPrefix] = make(map[string]string)

		for _, imp := range f.Imports {
			path, err := strconv.Unquote(imp.Path.Value)
			if err != nil {
				log.Fatalf("invalid import path string: %v", err) // Shouldn't happen here since packages.Load succeeded.
			}
			if path == dst {
				continue
			}

			ogPath := path

			if newPath, ok := importMap[path]; ok {
				path = newPath
			}

			var name string
			if imp.Name != nil {
				name = imp.Name.Name
			}

			importName := pkg.Imports[ogPath].Types.Name()

			refName := name
			if name == "" {
				refName = importName
			}

			name = aliasPrefix + importName + name
			spec := fmt.Sprintf("%s %q", name, path)

			renamedImportsAlias[aliasPrefix][refName] = name

			if isStandardImportPath(path) {
				pkgStd[spec] = true
			} else {
				pkgExt[spec] = true
			}
		}
	}

	// Print a single declaration that imports all necessary packages.
	fmt.Fprintln(&out, "import (")
	for p := range pkgStd {
		fmt.Fprintf(&out, "\t%s\n", p)
	}
	if len(pkgExt) > 0 {
		fmt.Fprintln(&out)
	}
	for p := range pkgExt {
		fmt.Fprintf(&out, "\t%s\n", p)
	}
	fmt.Fprint(&out, ")\n\n")

	// Modify and print each file.
	for _, f := range pkg.Syntax {
		// Update renamed identifiers.
		for id, obj := range pkg.TypesInfo.Defs {
			if objsToUpdate[obj] {
				id.Name = prefix + obj.Name()
			}
		}
		for id, obj := range pkg.TypesInfo.Uses {
			if objsToUpdate[obj] {
				id.Name = prefix + obj.Name()
			}
		}

		// For each qualified identifier that refers to the
		// destination package, remove the qualifier.
		// The "@@@." strings are removed in postprocessing.
		ast.Inspect(f, func(n ast.Node) bool {
			if sel, ok := n.(*ast.SelectorExpr); ok {
				if id, ok := sel.X.(*ast.Ident); ok {
					if obj, ok := pkg.TypesInfo.Uses[id].(*types.PkgName); ok {
						if obj.Imported().Path() == dst {
							id.Name = "@@@"
						}
					}
				}
			}
			return true
		})

		// For each reference to an import, replace it with the new alias
		aliasedImports := renamedImportsAlias[getAliasPrefix(pkg, f)]
		// convert to a slice for quick lookup
		importKeys := maps.Keys(aliasedImports)

		// manage a stack to do ancestor checks, see: https://stackoverflow.com/a/66810485
		var stack []ast.Node
		ast.Inspect(f, func(n ast.Node) bool {
		OUTER:
			switch x := n.(type) {
			case *ast.Ident:
				if slices.Contains(importKeys, x.Name) {
					// don't rename struct vars
					for _, item := range stack {
						// check if any ancestor is a struct
						if _, ok := item.(*ast.StructType); ok {
							if x.Obj != nil {
								// if our current object is a field don't rename it
								_, isField := x.Obj.Decl.(*ast.Field)
								if isField {
									break OUTER
								}
							}
						}
					}

					// check the parent node to make sure it's not a struct before replacing
					parent := stack[len(stack)-1]
					switch px := parent.(type) {
					case *ast.SelectorExpr:
						if sel, ok := px.X.(*ast.Ident); ok {
							if sel.Obj != nil && sel.Obj.Decl != nil {
								// check if its a field, if it is we break to the outer
								// since these dont' need to be rewritten
								_, isField := sel.Obj.Decl.(*ast.Field)
								if isField {
									break OUTER
								}

								assign, isAssign := sel.Obj.Decl.(*ast.AssignStmt)
								if isAssign {
									// check the right side of the assignment to see if its a field
									for _, item := range assign.Rhs {
										// TODO: there's probably more edge cases to handle around this but none that come up in google
										tae, isTypeAssert := item.(*ast.TypeAssertExpr)
										if isTypeAssert {
											// check the ident used in the type assertion
											taIdent, isTaIdent := tae.X.(*ast.Ident)
											if isTaIdent {
												// check the object for a field assertion
												if taIdent.Obj != nil && taIdent.Obj.Decl != nil {
													_, isField := taIdent.Obj.Decl.(*ast.Field)
													if isField {
														break OUTER
													}
												}
											}
										}
									}
								}
							}
						}
					}

					// don't rename fields
					if x.Obj != nil && x.Obj.Decl != nil {
						_, isField := x.Obj.Decl.(*ast.Field)
						if isField {
							break OUTER
						}
					}

					// skip vars
					x.Name = aliasedImports[x.Name]
				}
			case *ast.FuncType:
				if x.Results != nil {
					for _, res := range x.Results.List {
						if len(res.Names) == 0 {
							if ident, ok := res.Type.(*ast.Ident); ok && slices.Contains(importKeys, ident.Name) {
								ident.Name = aliasedImports[ident.Name]
							}
						}
					}
				}
			}

			// Manage the stack. Inspect calls a function like this:
			//   f(node)
			//   for each child {
			//      f(child) // and recursively for child's children
			//   }
			//   f(nil)
			if n == nil {
				// Done with node's children. Pop.
				stack = stack[:len(stack)-1]
			} else {
				// Push the current node for children.
				stack = append(stack, n)
			}

			return true
		})

		last := f.Package
		if len(f.Imports) > 0 {
			imp := f.Imports[len(f.Imports)-1]
			last = imp.End()
			if imp.Comment != nil {
				if e := imp.Comment.End(); e > last {
					last = e
				}
			}
		}

		// Pretty-print package-level declarations.
		// but no package or import declarations.
		var buf bytes.Buffer
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.GenDecl); ok && decl.Tok == token.IMPORT {
				continue
			}

			beg, end := sourceRange(decl)

			printComments(&out, f.Comments, last, beg)

			buf.Reset()
			format.Node(&buf, pkg.Fset, &printer.CommentedNode{Node: decl, Comments: f.Comments})
			// Remove each "@@@." in the output.
			// TODO(adonovan): not hygienic.
			out.Write(bytes.Replace(buf.Bytes(), []byte("@@@."), nil, -1))

			last = printSameLineComment(&out, f.Comments, pkg.Fset, end)

			out.WriteString("\n\n")
		}

		printLastComments(&out, f.Comments, last)
	}

	result, err := imports.Process("", out.Bytes(), &imports.Options{})
	if err != nil {
		return nil, fmt.Errorf("error processing imports: %v", err)
	}

	if buildTags != "" {
		result = append([]byte(fmt.Sprintf("//go:build %s\n", buildTags)), result...)
		result = append([]byte(fmt.Sprintf("// +build %s\n\n", buildTags)), result...)
	}

	result = append([]byte("//nolint\n"), result...)
	result = append([]byte("// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.\n"), result...)
	if *outputFile != "" && buildTags == "" {
		// skip this for now
		//fmt.Fprintf(&out, "//go:generate bundle %s\n", strings.Join(quoteArgs(os.Args[1:]), " "))
	} else {
		result = append([]byte(fmt.Sprintf("//   $ bundle %s\n", strings.Join(os.Args[1:], " "))), result...)
	}
	result = append([]byte("\n"), result...)

	// Now format the entire thing.
	result, err = format.Source(result)
	if err != nil {
		log.Fatalf("formatting failed: %v", err)
	}

	return result, nil
}

// sourceRange returns the [beg, end) interval of source code
// belonging to decl (incl. associated comments).
func sourceRange(decl ast.Decl) (beg, end token.Pos) {
	beg = decl.Pos()
	end = decl.End()

	var doc, com *ast.CommentGroup

	switch d := decl.(type) {
	case *ast.GenDecl:
		doc = d.Doc
		if len(d.Specs) > 0 {
			switch spec := d.Specs[len(d.Specs)-1].(type) {
			case *ast.ValueSpec:
				com = spec.Comment
			case *ast.TypeSpec:
				com = spec.Comment
			}
		}
	case *ast.FuncDecl:
		doc = d.Doc
	}

	if doc != nil {
		beg = doc.Pos()
	}
	if com != nil && com.End() > end {
		end = com.End()
	}

	return beg, end
}

func printComments(out *bytes.Buffer, comments []*ast.CommentGroup, pos, end token.Pos) {
	for _, cg := range comments {
		if pos <= cg.Pos() && cg.Pos() < end {
			for _, c := range cg.List {
				fmt.Fprintln(out, c.Text)
			}
			fmt.Fprintln(out)
		}
	}
}

const infinity = 1 << 30

func printLastComments(out *bytes.Buffer, comments []*ast.CommentGroup, pos token.Pos) {
	printComments(out, comments, pos, infinity)
}

func printSameLineComment(out *bytes.Buffer, comments []*ast.CommentGroup, fset *token.FileSet, pos token.Pos) token.Pos {
	tf := fset.File(pos)
	for _, cg := range comments {
		if pos <= cg.Pos() && tf.Line(cg.Pos()) == tf.Line(pos) {
			for _, c := range cg.List {
				fmt.Fprintln(out, c.Text)
			}
			return cg.End()
		}
	}
	return pos
}

func quoteArgs(ss []string) []string {
	// From go help generate:
	//
	// > The arguments to the directive are space-separated tokens or
	// > double-quoted strings passed to the generator as individual
	// > arguments when it is run.
	//
	// > Quoted strings use Go syntax and are evaluated before execution; a
	// > quoted string appears as a single argument to the generator.
	//
	var qs []string
	for _, s := range ss {
		if s == "" || containsSpace(s) {
			s = strconv.Quote(s)
		}
		qs = append(qs, s)
	}
	return qs
}

func containsSpace(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

type flagFunc func(string)

func (f flagFunc) Set(s string) error {
	f(s)
	return nil
}

func (f flagFunc) String() string { return "" }

// getAliasPrefix gets a unique import alias name for a package.
func getAliasPrefix(packages *packages.Package, file *ast.File) string {
	filename := path.Base(packages.Fset.Position(file.Package).Filename)
	extension := filepath.Ext(filename)

	return filename[0:len(filename)-len(extension)] + "_"
}
