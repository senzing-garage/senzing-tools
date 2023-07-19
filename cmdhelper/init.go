package cmdhelper

import (
	"fmt"

	"github.com/senzing/go-common/option"
	"github.com/senzing/go-common/option/optiontype"
	"github.com/spf13/cobra"
)

func Init(cobraCommand *cobra.Command, contextVariables []option.ContextVariable) {

	for _, contextVariable := range contextVariables {
		SetCobraFlag(cobraCommand, contextVariable)
	}
}

func SetCobraFlag(cobraCommand *cobra.Command, contextVariable option.ContextVariable) {
	switch contextVariable.Type {
	case optiontype.Bool:
		cobraCommand.Flags().Bool(contextVariable.Arg, contextVariable.Default.(bool), fmt.Sprintf(contextVariable.Help, contextVariable.Envar))
	case optiontype.Int:
		cobraCommand.Flags().Int(contextVariable.Arg, contextVariable.Default.(int), fmt.Sprintf(contextVariable.Help, contextVariable.Envar))
	case optiontype.String:
		cobraCommand.Flags().String(contextVariable.Arg, contextVariable.Default.(string), fmt.Sprintf(contextVariable.Help, contextVariable.Envar))
	case optiontype.StringSlice:
		cobraCommand.Flags().StringSlice(contextVariable.Arg, contextVariable.Default.([]string), fmt.Sprintf(contextVariable.Help, contextVariable.Envar))
	}
}
