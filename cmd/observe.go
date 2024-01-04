package cmd

import "github.com/senzing-garage/observe/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
