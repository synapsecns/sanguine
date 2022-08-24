package internal_test

import (
	"github.com/Flaque/filet"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// AbiSuite is the abigen suite.
type AbiSuite struct {
	*testsuite.TestSuite
	exampleFilePath string
}

func NewAbiSuite(tb testing.TB) *AbiSuite {
	tb.Helper()

	return &AbiSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (a *AbiSuite) SetupTest() {
	a.TestSuite.SetupTest()

	tempFile := filet.TmpFile(a.T(), "", testFileContents)

	_ = tempFile.Close()
	a.exampleFilePath = tempFile.Name()
}

func TestAbiSuite(t *testing.T) {
	suite.Run(t, NewAbiSuite(t))
}

const testFileContents = `// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.7.4 and less than 0.9.0
pragma solidity ^0.8.4;

contract HelloWorld {
    string public greet = "Hello World!";
}
`
