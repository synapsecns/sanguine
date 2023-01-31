# Bundle

This is a tool for bundling a package into a single source file. Most of the source code is copied from there, but this attempts to fix the bugs around shadowing in a rather hacky way (by aliasing all imports). See [here](https://pkg.go.dev/golang.org/x/tools/cmd/bundle) for the original code
