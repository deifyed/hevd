package test

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/tuuturu/hevd/pkg/apis/hevd.tuuturu.org/v1alpha1"
)

var (
	failedTestStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))
	passedTestStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00"))
	verboseStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#504945"))
)

type logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}

type Runner struct {
	testCases []v1alpha1.Case
	Log       logger
	Verbose   bool
	FailFast  bool
}

func (r *Runner) Push(t []v1alpha1.Case) {
	r.testCases = append(r.testCases, t...)
}

func (r *Runner) Run() (bool, error) {
	var ok bool = true

	for _, t := range r.testCases {
		err := t.Run()
		if err == nil {
			fmt.Printf("[ %s ] %s\n", passedTestStyle.Render("PASS"), t.Name())
		} else {
			fmt.Printf("[ %s ] %s\n", failedTestStyle.Render("FAIL"), t.Name())

			if r.Verbose {
				fmt.Printf("\t%s\n", verboseStyle.Render(err.Error()))
			}

			if r.FailFast {
				return false, nil
			}

			ok = false
		}
	}

	return ok, nil
}
