// Package cmd contains a generator for copying files exported files from geth
// in order to use private fields. The resulting files should not be modified directly
// but if there are new methods you need exported, generators, etc that can be done in other files
// that will now have access to the private fields. These generated files should only be used for testing
//
// TODO: look into implementing a tag for tests in order to make sure nothing in testutils/ is used in a production build
// we haven't done this yet because of the poor ux in an ide as far as having to add a `-tag`.
package cmd
