package cmd

import "github.com/senzing/load/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
