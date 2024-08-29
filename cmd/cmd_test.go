package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

func Test_Execute(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "--help"}
	Execute()
}

func Test_Execute_completion(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "completion"}
	Execute()
}

func Test_Execute_docs(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "docs"}
	Execute()
}

func Test_Execute_help(test *testing.T) {
	_ = test
	os.Args = []string{"command-name", "--help"}
	Execute()
}

func Test_PreRun(test *testing.T) {
	_ = test
	args := []string{"command-name", "--help"}
	PreRun(RootCmd, args)
}

func Test_RootCmd(test *testing.T) {
	_ = test
	err := RootCmd.Execute()
	require.NoError(test, err)
	// err = RootCmd.RunE(RootCmd, []string{})
	// require.NoError(test, err)
}

func Test_completionCmd(test *testing.T) {
	_ = test
	err := completionCmd.Execute()
	require.NoError(test, err)
	err = completionCmd.RunE(completionCmd, []string{})
	require.NoError(test, err)
}

func Test_docsCmd(test *testing.T) {
	_ = test
	err := docsCmd.Execute()
	require.NoError(test, err)
	err = docsCmd.RunE(docsCmd, []string{})
	require.NoError(test, err)
}

// Must be run near last.
func Test_Execute_SenzingToolsCommand(test *testing.T) {
	_ = test
	test.Setenv("SENZING_TOOLS_COMMAND", "explain --help")
	os.Args = []string{"command-name"}
	Execute()
}

// ----------------------------------------------------------------------------
// Test private functions
// ----------------------------------------------------------------------------

func Test_completionAction(test *testing.T) {
	var buffer bytes.Buffer
	err := completionAction(&buffer)
	require.NoError(test, err)
}

func Test_docsAction_badDir(test *testing.T) {
	var buffer bytes.Buffer
	badDir := "/tmp/no/directory/exists"
	err := docsAction(&buffer, badDir)
	require.Error(test, err)
}
