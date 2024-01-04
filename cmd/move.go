package cmd

import "github.com/senzing-garage/move/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
