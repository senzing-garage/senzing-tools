package cmd

import "github.com/senzing/init-database/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
