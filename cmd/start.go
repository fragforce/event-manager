package cmd

import (
	"github.com/fragforce/event-manager/pkg/webserver"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:    "start",
	Short:  "Start manager and services",
	Long:   `Starts the event manager service manager`,
	PreRun: setVerbose,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("starting webserver")
		webserver.Start(*config)
		//log.Debug("starting sub-services")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")
}
