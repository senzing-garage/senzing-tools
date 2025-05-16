package helper_test

import (
	"testing"

	"github.com/senzing-garage/senzing-tools/helper"
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
	actual := helper.MakeVersion("1.2.3", "1")
	require.Equal(test, expected, actual)
}

func Test_MakeVersion_NoIteration(test *testing.T) {
	expected := "1.2.3"
	actual := helper.MakeVersion(expected, "0")
	require.Equal(test, expected, actual)
}
