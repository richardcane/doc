package docker

import "os"

// ComposeFile Represents a docker compose file
type ComposeFile os.FileInfo

// ComposeFileArray Array of ComposeFiles
type ComposeFileArray []ComposeFile

// GetNames Returns the names of all compose files as a string array
func (files ComposeFileArray) GetNames() (names []string) {
	for _, file := range files {
		names = append(names, file.Name())
	}
	return
}
