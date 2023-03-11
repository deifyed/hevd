package http

import (
	"fmt"
	"net/http"
)

type Suite struct {
	Cases []Case `json:"cases"`
}

type Case struct {
	name string

	With   CaseWith   `json:"with"`
	Expect CaseExpect `json:"expect"`
}

func (c Case) Name() string {
	return c.name
}

func (c Case) Run() error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", c.With.URL, nil)
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

type CaseWith struct {
	URL string `json:"url"`
}

type CaseExpect struct {
	StatusCode int `json:"statusCode"`
}
