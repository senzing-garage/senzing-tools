package cmdhelper

import (
	"fmt"
	"os"
	"strings"

	"github.com/senzing/senzing-tools/constant"
	"github.com/senzing/senzing-tools/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// If a configuration file is present, load it.
func loadConfigurationFile(cobraCommand *cobra.Command, configName string) {
	configuration := ""
	configFlag := cobraCommand.Flags().Lookup(option.Configuration)
	if configFlag != nil {
		configuration = configFlag.Value.String()
	}
	if configuration != "" { // Use configuration file specified as a command line option.
		viper.SetConfigFile(configuration)
	} else { // Search for a configuration file.

		// Determine home directory.

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Specify configuration file name.

		viper.SetConfigName(configName)
		viper.SetConfigType("yaml")

		// Define search path order.

		viper.AddConfigPath(home + "/.senzing-tools")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/senzing-tools")
	}

	// If a config file is found, read it in.

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Applying configuration file:", viper.ConfigFileUsed())
	}
}

// Configure Viper with user-specified options.
func loadOptions(cobraCommand *cobra.Command, contextVariables []ContextVariable) {
	var err error = nil
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	if contextVariables != nil {
		for _, contextVar := range contextVariables {
			viper.SetDefault(contextVar.Option, contextVar.Default)
			err = viper.BindPFlag(contextVar.Option, cobraCommand.Flags().Lookup(contextVar.Option))
			if err != nil {
				panic(err)
			}
		}
	}

	// // Bools

	// if contextVariables.Bools != nil {
	// 	for _, contextVar := range contextVariables.Bools {
	// 		viper.SetDefault(contextVar.Option, contextVar.Default)
	// 		err = viper.BindPFlag(contextVar.Option, cobraCommand.Flags().Lookup(contextVar.Option))
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }

	// // Ints

	// if contextVariables.Ints != nil {
	// 	for _, contextVar := range contextVariables.Ints {
	// 		viper.SetDefault(contextVar.Option, contextVar.Default)
	// 		err = viper.BindPFlag(contextVar.Option, cobraCommand.Flags().Lookup(contextVar.Option))
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }

	// // Strings

	// if contextVariables.Strings != nil {
	// 	for _, contextVar := range contextVariables.Strings {
	// 		viper.SetDefault(contextVar.Option, contextVar.Default)
	// 		err = viper.BindPFlag(contextVar.Option, cobraCommand.Flags().Lookup(contextVar.Option))
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }

	// // StringSlice

	// if contextVariables.StringSlices != nil {
	// 	for _, contextVar := range contextVariables.StringSlices {
	// 		viper.SetDefault(contextVar.Option, contextVar.Default)
	// 		err = viper.BindPFlag(contextVar.Option, cobraCommand.Flags().Lookup(contextVar.Option))
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }
}

func PreRun(cobraCommand *cobra.Command, args []string, configName string, contextVariables []ContextVariable) {
	loadConfigurationFile(cobraCommand, configName)
	loadOptions(cobraCommand, contextVariables)
	cobraCommand.SetVersionTemplate(constant.VersionTemplate)
}
