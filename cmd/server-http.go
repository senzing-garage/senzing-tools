package cmd

import "github.com/senzing-garage/serve-http/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
