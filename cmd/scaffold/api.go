package scaffold

import (
	_ "embed"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed templates/http.yaml
var httpExample string

func RunE() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		testType := args[0]

		var output string

		switch testType {
		case "http":
			output = httpExample
		default:
			return fmt.Errorf("unknown test type: %s", testType)
		}

		io.Copy(cmd.OutOrStdout(), strings.NewReader(output))

		return nil
	}
}
