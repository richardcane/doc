package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/docker/libcompose/project/options"
	"github.com/docker/libcompose/utils"
	"github.com/richardcane/doc/actions"
	"github.com/richardcane/doc/flags"
	"github.com/spf13/cobra"
)

var upPreRun = func(cmd *cobra.Command, args []string) {
	var err error

	if loadedProject == nil {
		loadedProject, err = actions.LoadProjectFromPaths(flags.FilePaths)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var upRun = func(cmd *cobra.Command, args []string) {
	configuredServices := loadedProject.ServiceConfigs.Keys()

	for _, arg := range args {
		serviceExists := utils.Contains(configuredServices, arg)

		if !serviceExists {
			fmt.Printf("No such service \"%v\"\n\n", arg)
			log.Fatalf("Services available in %v: \n - %v",
				strings.Join(loadedProject.Files, ", "),
				strings.Join(configuredServices, "\n - "))
		}
	}

	loadedProject.Up(context.Background(), options.Up{}, args...)
}
