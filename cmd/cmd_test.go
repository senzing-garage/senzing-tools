package cmd_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/senzing-garage/senzing-tools/cmd"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func Test_Execute(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "--help"}

	cmd.Execute()
}

func Test_Execute_completion(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "completion"}

	cmd.Execute()
}

func Test_Execute_docs(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "docs"}

	cmd.Execute()
}

func Test_Execute_help(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "--help"}

	cmd.Execute()
}

func Test_PreRun(test *testing.T) {
	_ = test
	args := []string{"command-name", "--help"}
	cmd.PreRun(cmd.RootCmd, args)
}

func Test_RootCmd(test *testing.T) {
	_ = test
	err := cmd.RootCmd.Execute()
	// err = RootCmd.RunE(RootCmd, []string{})
	// require.NoError(test, err)
	require.NoError(test, err)
}

func Test_completionCmd(test *testing.T) {
	_ = test
	err := cmd.CompletionCmd.Execute()
	require.NoError(test, err)
	err = cmd.CompletionCmd.RunE(cmd.CompletionCmd, []string{})
	require.NoError(test, err)
}

func Test_docsCmd(test *testing.T) {
	_ = test
	err := cmd.DocsCmd.Execute()
	require.NoError(test, err)
	err = cmd.DocsCmd.RunE(cmd.DocsCmd, []string{})
	require.NoError(test, err)
}

// Must be run near last.
func Test_Execute_SenzingToolsCommand(test *testing.T) {
	_ = test
	test.Setenv("SENZING_TOOLS_COMMAND", "explain --help")

	os.Args = []string{"command-name"}

	cmd.Execute()
}

func Test_docsAction_badDir(test *testing.T) {
	var buffer bytes.Buffer

	badDir := "/tmp/no/directory/exists"
	err := cmd.DocsAction(&buffer, badDir)
	require.Error(test, err)
}
