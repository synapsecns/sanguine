package internal_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/tools/abigen/internal"
	"os"
	"os/exec"
	"testing"
)

func TestCreateRunFile(t *testing.T) {
	runFile, err := internal.CreateRunFile("0.8.13")
	Nil(t, err)

	//nolint: gosec
	cmd := exec.Command("bash", "-n", runFile.Name())
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		Nil(t, err)
	}
}

func (a *AbiSuite) TestCompileSolidity() {
	if os.Getenv("CI") != "" {
		a.T().Skip("doesn't work on ci")
	}

	vals, err := internal.CompileSolidity("0.8.4", a.exampleFilePath, 1)
	Nil(a.T(), err)

	Len(a.T(), vals, 1)
	for _, value := range vals {
		Equal(a.T(), value.Info.CompilerVersion, "0.8.4")
		Equal(a.T(), value.Info.LanguageVersion, "0.8.4")
	}
}
