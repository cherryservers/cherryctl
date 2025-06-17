package docs

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   `docs <destination>`,
		Short: "Generate a local version of the CLI documentation.",
		Long:  "Generates a local version of the CLI documentation in the specified directory. Each command gets a markdown file.",
		Example: `  # Generate documentation in the ./docs directory:
  cherryctl docs ./docs`,

		DisableFlagsInUseLine: true,
		Args:                  cobra.ExactValidArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			dest := args[0]

			return doc.GenMarkdownTree(cmd.Parent(), dest)
		},
	}
}
