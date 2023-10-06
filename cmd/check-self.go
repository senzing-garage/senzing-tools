package cmd

import "github.com/senzing/check-self/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
