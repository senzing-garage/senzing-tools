/*
 */
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	configurationFile string
	buildVersion      string = "0.1.0"
	buildIteration    string = "0"
)

var RootCmd = &cobra.Command{
	Use:   "senzing-tools",
	Short: "Tools to help use the Senzing API",
	Long: `
Welcome to senzing-tools!
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
