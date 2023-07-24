/*
 */
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/senzing/go-cmdhelping/cmdhelper"
	"github.com/senzing/go-cmdhelping/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Short string = "Tools to help use the Senzing API"
	Use   string = "senzing-tools"
	Long  string = `
Welcome to senzing-tools!
The value of [command] may also be specified by the
SENZING_TOOLS_COMMAND environment variable.
For more information, visit https://github.com/Senzing/senzing-tools
    `
	defaultConfiguration string = ""
)

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// Since init() is always invoked, define persistent command line parameters
// that apply to all senzing-tool subcommands.
func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().String(option.Configuration.Arg, option.Configuration.Default.(string), fmt.Sprintf(option.Configuration.Help, option.Configuration.Envar))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv()

	configuration := RootCmd.Flags().Lookup(option.Configuration.Arg).Value.String()
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
		fmt.Fprintf(os.Stderr, "Applying configuration file: %s\n", viper.ConfigFileUsed())
	}
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// Handle choice of command set by environment variable.

	command, isSet := os.LookupEnv("SENZING_TOOLS_COMMAND")
	if isSet {
		if (len(os.Args) == 1) || (strings.HasPrefix(os.Args[1], "-")) {
			newArgs := strings.Split(command, " ")
			for index, arg := range os.Args {
				if index > 0 {
					newArgs = append(newArgs, arg)
				}
			}
			fmt.Fprintf(os.Stderr, "Using SENZING_TOOLS_COMMAND value of '%s' resulting in command: `senzing-tools %s'\n", command, strings.Join(newArgs, " "))
			RootCmd.SetArgs(newArgs)
		}
	}

	// Execute command.

	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Used in construction of cobra.Command
func PreRun(cobraCommand *cobra.Command, args []string) {
	cobraCommand.SetVersionTemplate(`{{printf "%s: %s - version %s\n" .Name .Short .Version}}`)
}

// Used in construction of cobra.Command
func Version() string {
	return cmdhelper.Version(githubVersion, githubIteration)
}

// ----------------------------------------------------------------------------
// Command
// ----------------------------------------------------------------------------

// RootCmd represents the command.
var RootCmd = &cobra.Command{
	Use:     Use,
	Short:   Short,
	Long:    Long,
	PreRun:  PreRun,
	Version: Version(),
}
