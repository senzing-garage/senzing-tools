/*
 */
package cmd

import (
	"fmt"
	"os"

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

func makeVersion(version string, iteration string) string {
	result := ""
	if buildIteration == "0" {
		result = version
	} else {
		result = fmt.Sprintf("%s-%s", buildVersion, buildIteration)
	}
	return result
}

var RootCmd = &cobra.Command{
	Use:   "senzing-tools",
	Short: "Tools to help use the Senzing API",
	Long: `
Welcome to senzing-tools!
	`,
	PreRun: func(cobraCommand *cobra.Command, args []string) {
		fmt.Println(">>>>> senzing-tools.PreRun")
		versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
		cobraCommand.SetVersionTemplate(versionTemplate)
	},
	Version: makeVersion(buildVersion, buildIteration),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println(">>>>> senzing-tools.init()")
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().String("configuration", defaultConfiguration, "Path to configuration file [SENZING_TOOLS_CONFIGURATION]")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println(">>>>> senzing-tools.initConfig()")
	viper.AutomaticEnv()
	viper.SetDefault("configuration", defaultConfiguration)
	viper.BindPFlag("configuration", RootCmd.Flags().Lookup("configuration"))

	if viper.GetString("configuration") != "" {
		viper.SetConfigFile(viper.GetString("configuration"))
	} else {

		// Find home directory.

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Configuration file: senzing-tools.yaml

		viper.SetConfigName("senzing-tools")
		viper.SetConfigType("yaml")

		// Search path order.

		viper.AddConfigPath(home + "/.senzing-tools")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/senzing-tools")
	}

	// If a config file is found, read it in.

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
