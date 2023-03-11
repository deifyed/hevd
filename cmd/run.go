package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tuuturu/hevd/cmd/run"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run a test suite",
	Example: `  hevd run ./suite.yaml`,
	Args:    cobra.ExactArgs(1),
	RunE:    run.RunE(log, fs),
}

func init() {
	log.SetLevel(logrus.DebugLevel)
	rootCmd.AddCommand(runCmd)
}
