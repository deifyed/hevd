package http

import (
	"fmt"
	"net/http"
	"strings"
)

type Case struct {
	Title string `json:"name"`

	With   CaseWith   `json:"with"`
	Expect CaseExpect `json:"expect"`
}

func (c Case) Run() error {
	client := &http.Client{}

	req, err := http.NewRequest(c.With.method(), c.With.URL, nil)
	if err != nil {
		return fmt.Errorf("error while creating request: %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error while sending request: %s", err)
	}

	if resp.StatusCode != c.Expect.StatusCode {
		return fmt.Errorf("expected status code %d, got %d", c.Expect.StatusCode, resp.StatusCode)
	}

	return nil
}

func (c Case) Name() string {
	return fmt.Sprintf("%s %s %s", c.With.method(), c.With.URL, c.Title)
}

type CaseWith struct {
	URL    string `json:"url"`
	Method string `json:"method"`
}

func (c CaseWith) method() string {
	if c.Method == "" {
		return "GET"
	}

	return strings.ToUpper(c.Method)
}

type CaseExpect struct {
	StatusCode int `json:"statusCode"`
}
