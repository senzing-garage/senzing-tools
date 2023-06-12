package cmdhelper

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/senzing/senzing-tools/envar"
	"github.com/senzing/senzing-tools/help"
	"github.com/senzing/senzing-tools/option"
	"github.com/spf13/cobra"
)

var contextVariables = ContextVariables{
	Bools: []ContextBool{
		{
			Default: false,
			Envar:   envar.EnableSwaggerUi,
			Help:    help.EnableSwaggerUi,
			Option:  option.EnableSwaggerUi,
		},
	},
	Ints: []ContextInt{
		{
			Default: 0,
			Envar:   envar.EngineLogLevel,
			Help:    help.EngineLogLevel,
			Option:  option.EngineLogLevel,
		},
	},
	Strings: []ContextString{
		{
			Default: "",
			Envar:   envar.Configuration,
			Help:    help.Configuration,
			Option:  option.Configuration,
		},
	},
	StringSlices: []ContextStringSlice{
		{
			Default: []string{},
			Envar:   envar.XtermArguments,
			Help:    help.XtermArguments,
			Option:  option.XtermArguments,
		},
	},
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
	var contextVariables = ContextVariables{
		Strings: []ContextString{
			{
				Default: "",
				Envar:   "MY_VARIABLE",
				Help:    "Description of my variable [%s]",
				Option:  "my-variable",
			},
		},
	}
	Init(cobraCommand, contextVariables)
	// Output:
}

func ExamplePreRun() {
	cobraCommand := &cobra.Command{
		Use:   "example-use",
		Short: "example-short",
		Long:  `example-long`,
	}
	var contextVariables = ContextVariables{
		Strings: []ContextString{
			{
				Default: "",
				Envar:   "MY_VARIABLE",
				Help:    "Description of my variable [%s]",
				Option:  "my-variable",
			},
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
