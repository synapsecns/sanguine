package config_test

import (
	"github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/contrib/promexporter/config"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	// create an empty file
	tmpFile := filet.TmpFile(t, "", "")

	_, err := config.DecodeConfig(tmpFile.Name())
	assert.Nil(t, err)
}
