package run

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/tuuturu/hevd/pkg/apis/hevd.tuuturu.org/v1alpha1"
	"github.com/tuuturu/hevd/pkg/http"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func RunE(log logger, fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		log.Debugf("run args: %v", args)
		targetPath := args[0]

		exists, err := fs.Exists(targetPath)
		if err != nil {
			return fmt.Errorf("checking if target exists: %w", err)
		}

		if !exists {
			return fmt.Errorf("target does not exist: %s", targetPath)
		}

		result, err := fs.ReadFile(targetPath)
		if err != nil {
			return fmt.Errorf("reading target: %w", err)
		}

		var test v1alpha1.Test

		err = yaml.Unmarshal(result, &test)
		if err != nil {
			return fmt.Errorf("unmarshalling target: %w", err)
		}

		tests := make([]v1alpha1.Case, 0)

		switch test.Type {
		case "http":
			log.Debugf("test type: %s", test.Type)

			var suite http.Suite
			err = yaml.Unmarshal(result, &suite)
			if err != nil {
				return fmt.Errorf("unmarshalling http test: %w", err)
			}

			for _, c := range suite.Cases {
				tests = append(tests, &c)
			}

			break
		default:
			return fmt.Errorf("unknown test type: %s", test.Type)
		}

		for _, t := range tests {
			log.Debugf("running test: %s", t.Name())
			err = t.Run()
			if err != nil {
				return fmt.Errorf("running test: %w", err)
			}
		}

		return nil
	}
}
