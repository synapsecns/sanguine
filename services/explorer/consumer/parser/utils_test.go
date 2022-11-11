package parser_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenYaml(t *testing.T) {
	pwd, _ := os.Getwd()
	path := pwd + filepath.Clean("/static/tokenIDToCoinGeckoID.yaml")
	parsedYaml, err := parser.OpenYaml(path)
	Nil(t, err)
	NotNil(t, parsedYaml)
}
