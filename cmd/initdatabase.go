package cmd

import "github.com/senzing/initdatabase/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
