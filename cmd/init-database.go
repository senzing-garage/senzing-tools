package cmd

import "github.com/senzing-garage/init-database/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
