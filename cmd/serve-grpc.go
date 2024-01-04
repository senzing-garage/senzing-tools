package cmd

import "github.com/senzing-garage/serve-grpc/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
