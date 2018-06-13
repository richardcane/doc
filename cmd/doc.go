package cmd

import (
	"doc/flags"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	doc.PersistentFlags().StringVarP(&flags.FilePath, "file", "f", "", "Docker-compose file to load (default is docker-compose.yml)")

	doc.AddCommand(up)
	doc.AddCommand(down)
}

var doc = &cobra.Command{
	Use: "doc",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if len(os.Args) > 1 && os.Args[1] != "help" {
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var up = &cobra.Command{
	Use:   "up",
	Short: "Start/build one or more docker containers from docker-compose",
	Run:   upRun,
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
