package cmd

import "github.com/senzing/servegrpc/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
