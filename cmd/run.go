package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tuuturu/hevd/cmd/run"
)

// runCmd represents the run command
var runCmdOpts = run.Options{
	FileSystem: fs,
	Verbose:    false,
}

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run a test suite",
	Example: `  hevd run ./suite.yaml`,
	Args:    cobra.ExactArgs(1),
	RunE:    run.RunE(log, &runCmdOpts),
}

func init() {
	log.SetLevel(logrus.DebugLevel)

	runCmd.Flags().BoolVarP(&runCmdOpts.Verbose, "verbose", "v", false, "Verbose output from failed tests")

	rootCmd.AddCommand(runCmd)
}
