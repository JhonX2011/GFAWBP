package config

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
)

const (
	_defaultBaseDir = "/configs/latest"
	_baseDirEnvVar  = "CONFIG_DIR"
)

func Read(profileName string) ([]byte, error) {
	dir := _defaultBaseDir
	if d := os.Getenv(_baseDirEnvVar); d != "" {
		dir = d
	}

	_, content, err := read(profileName, dir)

	return content, err
}

var supportedExtensions = []string{"json", "yaml", "properties"} //nolint:gochecknoglobals

func read(profile string, basePath string) (string, []byte, error) {
	files, err := os.ReadDir(basePath)
	if err != nil {
		return "", nil, err
	}

	for _, file := range files {
		name, extension, found := strings.Cut(file.Name(), ".")
		if !found {
			continue
		}

		if !slices.Contains(supportedExtensions, extension) {
			continue
		}

		if strings.EqualFold(name, profile) {
			content, err := os.ReadFile(path.Join(basePath, file.Name()))
			if err != nil {
				return "", nil, err
			}

			return extension, content, nil
		}
	}

	return "", nil, fmt.Errorf("config: no such configuration profile: %s", profile)
}
