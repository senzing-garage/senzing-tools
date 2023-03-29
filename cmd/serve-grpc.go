package cmd

import "github.com/senzing/serve-grpc/cmd"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
