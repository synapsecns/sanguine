package parser_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"path/filepath"
	"testing"
)

func TestOpenYaml(t *testing.T) {
	path, _ := filepath.Abs("../../static/tokenIDToCoinGeckoID.yaml")
	parsedYaml, err := parser.OpenYaml(path)
	Nil(t, err)
	NotNil(t, parsedYaml)
}
