package cmd

import (
	"os"

	"github.com/docker/libcompose/project"

	"github.com/richardcane/doc/flags"

	"github.com/spf13/cobra"
)

var loadedProject *project.Project

func init() {
	doc.PersistentFlags().StringSliceVarP(&flags.FilePaths,
		"file",
		"f",
		[]string{},
		"Docker-compose file to load (default is docker-compose.yml)")

	doc.AddCommand(up)
	doc.AddCommand(down)
}

var doc = &cobra.Command{
	Use: "doc",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var up = &cobra.Command{
	Use:    "up",
	Short:  "Start/build one or more docker containers from docker-compose",
	PreRun: upPreRun,
	Run:    upRun,
}

var down = &cobra.Command{
	Use:   "down",
	Short: "Kill and remove one or more running containers.",
	Run:   downRun,
}

func Execute() {
	if err := doc.Execute(); err != nil {
		os.Exit(1)
	}
}
