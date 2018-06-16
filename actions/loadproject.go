package actions

import (
	compose "github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/richardcane/doc/docker"

	"github.com/docker/libcompose/project"
)

// LoadProjectFromPaths Loads project from given filepaths
func LoadProjectFromPaths(filepaths []string) (*project.Project, error) {
	var targets docker.ComposeFileArray
	var proj project.APIProject
	var err error

	if targets, err = FindComposeFiles(filepaths); err != nil {
		return nil, err
	}

	if proj, err = newProject(targets); err != nil {
		return nil, err
	}

	return proj.(*project.Project), nil
}

func newProject(targets docker.ComposeFileArray) (*project.Project, error) {
	var proj project.APIProject
	var err error

	if proj, err = compose.NewProject(
		&ctx.Context{
			Context: project.Context{
				ComposeFiles: targets.GetNames(),
			},
		}, nil); err != nil {
		return nil, err
	}

	return proj.(*project.Project), nil
}
