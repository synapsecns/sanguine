// Package packagedetector implements a detector for detecting changes in a git repo.
// Dependencies are transitive, at the package level. Modules depend on packages, whose dep graph is also at package level. Module->package->package
package packagedetector
