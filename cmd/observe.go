package cmd

import "github.com/senzing/observe/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
