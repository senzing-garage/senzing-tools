package cmd

import "github.com/senzing/serve-http/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
