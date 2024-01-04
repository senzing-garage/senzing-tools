package cmd

import "github.com/senzing-garage/load/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
