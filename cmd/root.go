package cmd

import (
	"fmt"

	"github.com/lithammer/dedent"
	"github.com/spenserblack/boop/pkg/boop"
	"github.com/spf13/cobra"
)

// executable defines if the created file should be an executable.
var executable bool

// rootCmd is the root command for the CLI.
var rootCmd = &cobra.Command{
	Use:   "boop [flags] PATH...",
	Short: "Create or update files and directories",
	Long: dedent.Dedent(`
		Create or update files and directories

		If the path ends in a trailing slash, it will be created as a directory.
		Otherwise, it will be created as a file.

		For deeply nested paths, it will create each directory that is required.
	`),
	Example: "boop path/to/new/dir/ path/to/new/file",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stderr := cmd.ErrOrStderr()
		boop := boop.New()
		boop.Executable(executable)

		for _, path := range args {
			if err := boop.Boop(path); err != nil {
				fmt.Fprintln(stderr, err)
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&executable, "executable", "x", false, "Make the file(s) executable (Unix-only)")
}
