package cmd

import (
	"context"
	"doc/flags"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/project/options"
	"github.com/spf13/cobra"
)

var upRun = func(cmd *cobra.Command, args []string) {
	composeFiles := []string{"docker-compose.yml", "docker-compose.yaml"}

	var target os.FileInfo
	if len(flags.FilePath) > 0 {
		fileInfo, err := os.Stat(flags.FilePath)
		if os.IsNotExist(err) {
			log.Fatal("Failed to find the compose file: " + flags.FilePath)
		}
		target = fileInfo
	} else {
		for _, name := range composeFiles {
			if fileInfo, err := os.Stat(name); err == nil {
				target = fileInfo
			}
		}
	}
	if target == nil {
		log.Fatal("Failed to find compose file.")
	}

	if apiProj, err := docker.NewProject(&ctx.Context{
		Context: project.Context{
			ComposeFiles: []string{target.Name()},
		},
	}, nil); err != nil {
		log.Fatal(err)
	} else {
		var proj *project.Project
		proj = apiProj.(*project.Project)

		configuredServices := proj.ServiceConfigs.Keys()

		var serviceExists bool
		for _, arg := range args {
			serviceExists = func(requested string, configured []string) bool {
				for _, service := range configured {
					if requested == service {
						return true
					}
				}
				return false
			}(arg, configuredServices)

			if !serviceExists {
				fmt.Printf("No such service \"%v\"\n\n", arg)
				log.Fatalf("Services available in %v: \n - %v", target.Name(), strings.Join(configuredServices, "\n - "))
			}
		}

		apiProj.Up(context.Background(), options.Up{}, args...)
	}
}
