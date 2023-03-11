package run

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/tuuturu/hevd/pkg/apis/hevd.tuuturu.org/v1alpha1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func validatePath(fs *afero.Afero, target string) error {
	exists, err := fs.Exists(target)
	if err != nil {
		return fmt.Errorf("checking if target exists: %w", err)
	}

	if !exists {
		return fmt.Errorf("target does not exist: %s", target)
	}

	return nil
}

func acquireTestType(result []byte) (string, error) {
	var test struct {
		Type string `json:"type"`
	}

	err := yaml.Unmarshal(result, &test)
	if err != nil {
		return "", fmt.Errorf("unmarshalling test type: %w", err)
	}

	return test.Type, nil
}

func mapToGenericCase[T v1alpha1.Case](list []T) []v1alpha1.Case {
	var result []v1alpha1.Case

	for _, item := range list {
		result = append(result, item)
	}

	return result
}
