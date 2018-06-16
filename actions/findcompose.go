package actions

import (
	"errors"
	"os"

	"github.com/richardcane/doc/docker"
)

// FindComposeFiles Returns array of compose files from given paths
func FindComposeFiles(filepaths []string) (docker.ComposeFileArray, error) {
	var defaults = []string{"docker-compose.yml", "docker-compose.yaml"}
	var targets docker.ComposeFileArray
	var target docker.ComposeFile
	var err error

	if len(filepaths) > 0 {
		for _, filepath := range filepaths {
			if target, err = os.Stat(filepath); err != nil {
				return nil, errors.New("Failed to find the compose file: " + filepath)
			}
			targets = append(targets, target)
		}
		return targets, nil
	}

	for _, path := range defaults {
		if target, err = os.Stat(path); err == nil {
			targets = append(targets, target)
		}
	}

	if len(targets) == 0 {
		return nil, errors.New("Failed to find compose file")
	}
	return targets, nil
}
