package run

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tuuturu/hevd/pkg/http"
	"github.com/tuuturu/hevd/pkg/test"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func RunE(log logger, opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]

		err := validatePath(opts.FileSystem, targetPath)
		if err != nil {
			return fmt.Errorf("validating path: %w", err)
		}

		result, err := opts.FileSystem.ReadFile(targetPath)
		if err != nil {
			return fmt.Errorf("reading target: %w", err)
		}

		testType, err := acquireTestType(result)
		if err != nil {
			return fmt.Errorf("acquiring test type: %w", err)
		}

		runner := test.Runner{Log: log, Verbose: opts.Verbose}

		switch testType {
		case "http":
			var suite http.Suite

			err = yaml.Unmarshal(result, &suite)
			if err != nil {
				return fmt.Errorf("unmarshalling http test: %w", err)
			}

			runner.Push(mapToGenericCase(suite.Cases))
		default:
			return fmt.Errorf("unknown test type: %s", testType)
		}

		ok, err := runner.Run()
		if err != nil {
			return fmt.Errorf("running test: %w", err)
		}

		if !ok {
			return errors.New("one or more tests failed")
		}

		return nil
	}
}
