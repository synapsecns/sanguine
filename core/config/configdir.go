package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shibukawa/configdir"
	"path/filepath"
)

// GetConfigDir gets the config dir. We create a default file to mark the directory level.
// if the config dir can not be created, an error is returned.
func GetConfigDir() (string, error) {
	configDir := configdir.New(VendorName, AppName+nameSuffix)
	file := configDir.QueryFolderContainsFile(readmeName)

	if file != nil {
		return filepath.Dir(file.Path), nil
	}
	// create the config folder
	folders := configDir.QueryFolders(configdir.All)
	if len(folders) == 0 {
		return "", errors.New("could not create config folder")
	}

	folder := folders[0]

	// create the file handle
	fileHandle, err := folder.Create(readmeName)
	if err != nil {
		return "", fmt.Errorf("could not get logo: %w", err)
	}

	_, err = fileHandle.WriteString(readmeContents)
	if err != nil {
		return "", fmt.Errorf("could not write to file: %w", err)
	}

	_ = fileHandle.Close()

	return filepath.Dir(fileHandle.Name()), nil
}

// readmeName is the name of the readme file.
const readmeName = "README.md"

// readmeContents are the contents of the readme file.
var readmeContents = fmt.Sprintf("This is the config directory for the %s application.", AppName)
