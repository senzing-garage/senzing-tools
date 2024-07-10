package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test public functions
// ----------------------------------------------------------------------------

/*
 * The unit tests in this file simulate command line invocation.
 */

func Test_MakeVersion_Iteration(test *testing.T) {
	expected := "1.2.3-1"
	actual := MakeVersion("1.2.3", "1")
	require.Equal(test, expected, actual)
}

func Test_MakeVersion_NoIteration(test *testing.T) {
	expected := "1.2.3"
	actual := MakeVersion(expected, "0")
	require.Equal(test, expected, actual)
}
