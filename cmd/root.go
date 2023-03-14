/*
 */
package cmd

import (
	"fmt"
	"os"

	"github.com/senzing/senzing-tools/constant"
	"github.com/senzing/senzing-tools/envar"
	"github.com/senzing/senzing-tools/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfiguration string = ""
)

var (
	buildVersion   string = "0.1.0"
	buildIteration string = "0"
)

var RootCmd = &cobra.Command{
	Use:   "senzing-tools",
	Short: "Tools to help use the Senzing API",
	Long: `
Welcome to senzing-tools!
	`,
	PreRun: func(cobraCommand *cobra.Command, args []string) {
		fmt.Println(">>>>> senzing-tools.PreRun")
		cobraCommand.SetVersionTemplate(constant.VersionTemplate)
	},
	Version: helper.MakeVersion(buildVersion, buildIteration),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Since init() is always invoked, define persistent command line parameters
// that apply to all senzing-tool subcommands.
func init() {
	fmt.Println(">>>>> senzing-tools.init()")
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().String(constant.Configuration, defaultConfiguration, fmt.Sprintf("Path to configuration file [%s]", envar.Configuration))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println(">>>>> senzing-tools.initConfig()")
	viper.AutomaticEnv()

	configuration := RootCmd.Flags().Lookup(constant.Configuration).Value.String()
	if configuration != "" { // Use configuration file specified as a command line option.
		viper.SetConfigFile(configuration)
	} else { // Search for a configuration file.

		// Determine home directory.

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Specify configuration file name.

		viper.SetConfigName("senzing-tools")
		viper.SetConfigType("yaml")

		// Define search path order.

		viper.AddConfigPath(home + "/.senzing-tools")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/senzing-tools")
	}

	// If a config file is found, read it in.

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
