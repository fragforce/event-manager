package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	log     = logrus.New()
	verbose string

	cfgDir  = "."
	cfgFile = "config.yml"
	config  = viper.New()

	version  = `v0.4.0`
	asciiArt = `                      __             __
    ____  ____ ______/ /_____  _____/ /__________  ____
   / __ \/ __ '/ ___/ //_/ _ \/ ___/ __/ ___/ __ \/ __ \
  / /_/ / /_/ / /  / ,< /  __/ /  / /_/ /  / /_/ / / / /
 / .___/\__,_/_/  /_/|_|\___/_/   \__/_/   \____/_/ /_/
/_/`
)

var (
	rootCmd = &cobra.Command{
		Use:    "event-manager",
		Short:  "event management software",
		Long:   `A software written to manage entire streaming events from fragforce`,
		PreRun: setVerbose,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", cfgFile, "config file name (default is 'config.yml'")
	rootCmd.PersistentFlags().StringVarP(&cfgDir, "dir", "d", cfgDir, "config directory path (default is '.)")
	rootCmd.PersistentFlags().StringVarP(&verbose, "verbosity", "v", logrus.InfoLevel.String(), "Log level (debug, info, warn, error")

	log.Info("starting event manager")
	// If a config file is found, read it in.
	config.AddConfigPath(cfgDir)
	config.SetConfigType("yaml")
	config.SetConfigName(cfgFile)

	err := config.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	log.Info("Using config file:", config.ConfigFileUsed())
}

func setVerbose(cmd *cobra.Command, args []string) {
	switch verbose {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.DebugLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.DebugLevel)
	}
}
