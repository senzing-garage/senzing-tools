/*
 */
package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion for the command",
	Long: `To load completions, run:
source < (senzing-tools completion)

To load completions automaticallon on login, add this line to your .bashrc file:
source < (senzing-tools completion)
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = cmd
		_ = args
		return completionAction(os.Stdout)
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}

func completionAction(out io.Writer) error {
	return RootCmd.GenBashCompletion(out)
}
