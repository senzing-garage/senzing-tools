package cmd

import "github.com/senzing-garage/check-self/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
