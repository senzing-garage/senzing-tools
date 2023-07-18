package cmdhelper

import (
	"fmt"

	"github.com/senzing/senzing-tools/cmdhelper/optiontype"
	"github.com/spf13/cobra"
)

// func InitBools(cobraCommand *cobra.Command, contextBools []ContextBool) {
// 	for _, contextBool := range contextBools {
// 		cobraCommand.Flags().Bool(contextBool.Option, contextBool.Default, fmt.Sprintf(contextBool.Help, contextBool.Envar))
// 	}
// }

// func InitInts(cobraCommand *cobra.Command, contextInts []ContextInt) {
// 	for _, contextInt := range contextInts {
// 		cobraCommand.Flags().Int(contextInt.Option, contextInt.Default, fmt.Sprintf(contextInt.Help, contextInt.Envar))
// 	}
// }

// func InitStrings(cobraCommand *cobra.Command, contextStrings []ContextString) {
// 	for _, contextString := range contextStrings {
// 		cobraCommand.Flags().String(contextString.Option, contextString.Default, fmt.Sprintf(contextString.Help, contextString.Envar))
// 	}
// }

// func InitStringSlices(cobraCommand *cobra.Command, contextStringSlices []ContextStringSlice) {
// 	for _, contextStringSlice := range contextStringSlices {
// 		cobraCommand.Flags().StringSlice(contextStringSlice.Option, contextStringSlice.Default, fmt.Sprintf(contextStringSlice.Help, contextStringSlice.Envar))
// 	}
// }

func Init(cobraCommand *cobra.Command, contextVariables []ContextVariable) {
	// if contextVariables.Bools != nil {
	// 	InitBools(cobraCommand, contextVariables.Bools)
	// }
	// if contextVariables.Ints != nil {
	// 	InitInts(cobraCommand, contextVariables.Ints)
	// }
	// if contextVariables.Strings != nil {
	// 	InitStrings(cobraCommand, contextVariables.Strings)
	// }
	// if contextVariables.StringSlices != nil {
	// 	InitStringSlices(cobraCommand, contextVariables.StringSlices)
	// }

	for _, contextVariable := range contextVariables {
		switch contextVariable.Type {
		case optiontype.Bool:
			cobraCommand.Flags().Bool(contextVariable.Option, contextVariable.Default.(bool), fmt.Sprintf(contextVariable.Help, contextVariable.Envar))
		case optiontype.Int:
			cobraCommand.Flags().Int(contextVariable.Option, contextVariable.Default.(int), fmt.Sprintf(contextVariable.Help, contextVariable.Envar))
		case optiontype.String:
			cobraCommand.Flags().String(contextVariable.Option, contextVariable.Default.(string), fmt.Sprintf(contextVariable.Help, contextVariable.Envar))
		case optiontype.StringSlice:
			cobraCommand.Flags().StringSlice(contextVariable.Option, contextVariable.Default.([]string), fmt.Sprintf(contextVariable.Help, contextVariable.Envar))
		}

	}
}
