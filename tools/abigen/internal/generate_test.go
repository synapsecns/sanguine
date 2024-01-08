package internal_test

import (
	"encoding/json"
	"errors"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/debug"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/tools/abigen/internal"
)

func TestCreateRunFile(t *testing.T) {
	runFile, err := internal.CreateRunFile("0.8.17")
	Nil(t, err)

	//nolint: gosec
	cmd := exec.Command("bash", "-n", runFile.Name())
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		Nil(t, err)
	}
}

func (a *AbiSuite) TestCompileSolidityImplicitEVM() {
	vals, err := internal.CompileSolidity("0.8.4", a.exampleFilePath, 1, nil)
	Nil(a.T(), err)

	Len(a.T(), vals, 1)
	for _, value := range vals {
		Equal(a.T(), value.Info.CompilerVersion, "0.8.4")
		Equal(a.T(), value.Info.LanguageVersion, "0.8.4")
	}
}

func (a *AbiSuite) TestCompileSolidityExplicitEVM() {
	// default would be shnghai
	const testEvmVersion = "istanbul"
	vals, err := internal.CompileSolidity("0.8.20", a.exampleFilePath, 1, core.PtrTo(testEvmVersion))
	Nil(a.T(), err)

	Len(a.T(), vals, 1)
	for _, value := range vals {
		Equal(a.T(), value.Info.CompilerVersion, "0.8.20")
		Equal(a.T(), value.Info.LanguageVersion, "0.8.20")

		var metadata debug.ContractMetadata
		err = json.Unmarshal([]byte(value.Info.Metadata), &metadata)
		a.Require().NoError(err)

		if metadata.Settings.EvmVersion != testEvmVersion {
			a.T().Errorf("expected %s to be %s", metadata.Language, testEvmVersion)
		}
	}
}

func TestFilePathsAreEqual(t *testing.T) {
	tests := []struct {
		file1 string
		file2 string
		want  bool
		err   error
	}{
		{"path/to/file1.txt", "path/to/file2.txt", false, nil},
		{"path/to/file1.txt", "path/to/file1.txt", true, nil},
		{"path/to/file2.txt", "path/to/file2.txt", true, nil},
		{"path/to/file1.txt", "", false, filepath.ErrBadPattern},
		{"", "path/to/file2.txt", false, filepath.ErrBadPattern},
		{"nonexistent/file.txt", "path/to/file.txt", false, nil},
	}

	for _, tt := range tests {
		got, err := internal.FilePathsAreEqual(tt.file1, tt.file2)

		if got != tt.want {
			t.Errorf("filePathsAreEqual(%v, %v) got %v, want %v", tt.file1, tt.file2, got, tt.want)
		}

		if err != nil && !errors.Is(err, tt.err) {
			t.Errorf("filePathsAreEqual(%v, %v) error got %v, want %v", tt.file1, tt.file2, err, tt.err)
		}
	}
}
