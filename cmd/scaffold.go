package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tuuturu/hevd/cmd/scaffold"
)

// scaffoldCmd represents the scaffold command
var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Generate a test suite example",
	Args:  cobra.ExactArgs(1),
	RunE:  scaffold.RunE(),
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)
}
