package config

import (
	"bytes"
	"fmt"
	"github.com/synapsecns/sanguine/core/internal/assets"
	"os"
	"path/filepath"
)

// logoFileName is the name of the logo file.
const logoFileName = "logo.svg"

// nameSuffix is the name suffix to use after the app name.
// this is useful for testing/versioning in a future version of this app.
var nameSuffix string

// GetLogoPath fetches the logo path from the config and writes it to the global config if it doesn't exist
// TODO this should be more generalized to work across static assets when we have more.
// returns an empty string if not present.
//
//nolint:nestif
func GetLogoPath() (_ string, err error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", fmt.Errorf("could not get config dir: %w", err)
	}
	logoPath := filepath.Join(configDir, logoFileName)

	shouldWrite := true
	var fileHandle *os.File

	if _, err := os.Stat(logoPath); os.IsNotExist(err) {
		//nolint: gosec
		fileHandle, err = os.Create(logoPath)

		if err != nil {
			return "", fmt.Errorf("could not get logo: %w", err)
		}
	} else {
		//nolint: gosec
		fileHandle, err = os.Open(logoPath)

		if err != nil {
			return "", fmt.Errorf("could not get logo: %w", err)
		}

		data, _ := os.ReadFile(fileHandle.Name())
		if bytes.Equal(data, assets.Logo) {
			shouldWrite = false
		}
	}

	// otherwise create the file and return it
	if shouldWrite {
		_, err = fileHandle.Write(assets.Logo)
		if err != nil {
			return "", fmt.Errorf("could not write file: %w", err)
		}
	}
	_ = fileHandle.Close()

	return fileHandle.Name(), nil
}
