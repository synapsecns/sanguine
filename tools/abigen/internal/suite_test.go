package internal_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"os"
	"path/filepath"
	"testing"
)

// AbiSuite is the abigen suite.
type AbiSuite struct {
	*testutils.TestSuite
	exampleFilePath string
}

func NewAbiSuite(tb testing.TB) *AbiSuite {
	tb.Helper()

	return &AbiSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (a *AbiSuite) SetupTest() {
	a.TestSuite.SetupTest()

	tempDir := filet.TmpDir(a.T(), "")
	tempFile, err := os.Create(filepath.Join(tempDir, "testFile"))
	Nil(a.T(), err)

	_ = tempFile.Close()
	a.exampleFilePath = tempFile.Name()
	tempFile.Close()
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
