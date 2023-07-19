package cmdhelper

import (
	"fmt"
	"testing"

	"github.com/senzing/go-common/option"
	"github.com/senzing/go-common/option/optiontype"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var contextVariables = []option.ContextVariable{
	option.Configuration,
	option.EngineLogLevel,
	option.EnableSwaggerUi,
	option.XtermArguments,
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestInit(test *testing.T) {
	cobraCommand := &cobra.Command{
		Use:   "test-use",
		Short: "test-short",
		Long:  `test-long`,
	}
	Init(cobraCommand, contextVariables)
}

func TestOsLookupEnvBool(test *testing.T) {
	assert.True(test, option.OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
}

func TestOsLookupEnvInt(test *testing.T) {
	assert.Equal(test, 10, option.OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
}

func TestOsLookupEnvString(test *testing.T) {
	assert.Equal(test, "default", option.OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
}

func TestPreRun(test *testing.T) {
	cobraCommand := &cobra.Command{
		Use:   "test-use",
		Short: "test-short",
		Long:  `test-long`,
	}
	Init(cobraCommand, contextVariables)
	PreRun(cobraCommand, []string{}, "test-cmd", contextVariables)
}

func TestVersion(test *testing.T) {
	assert.Equal(test, "1.2.3-4", Version("1.2.3", "4"))
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleInit() {
	cobraCommand := &cobra.Command{
		Use:   "example-use",
		Short: "example-short",
		Long:  `example-long`,
	}
	var contextVariables = []option.ContextVariable{
		{
			Default: "",
			Envar:   "MY_VARIABLE",
			Help:    "Description of my variable [%s]",
			Arg:     "my-variable",
			Type:    optiontype.String,
		},
	}
	Init(cobraCommand, contextVariables)
	// Output:
}

func ExampleOsLookupEnvBool() {
	fmt.Println(option.OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
	// Output: true
}

func ExampleOsLookupEnvInt() {
	fmt.Println(option.OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
	// Output: 10
}

func ExampleOsLookupEnvString() {
	fmt.Println(option.OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
	// Output: default
}

func ExamplePreRun() {
	cobraCommand := &cobra.Command{
		Use:   "example-use",
		Short: "example-short",
		Long:  `example-long`,
	}
	var contextVariables = []option.ContextVariable{
		{
			Default: "",
			Envar:   "MY_VARIABLE",
			Help:    "Description of my variable [%s]",
			Arg:     "my-variable",
			Type:    optiontype.String,
		},
	}
	Init(cobraCommand, contextVariables)
	PreRun(cobraCommand, []string{}, "example-cmd", contextVariables)
	// Output:
}

func ExampleVersion() {
	result := Version("1.2.3", "4")
	fmt.Println(result)
	// Output: 1.2.3-4
}
