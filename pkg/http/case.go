package http

import (
	"fmt"
	"io"
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

	body, err := c.With.body()
	if err != nil {
		return fmt.Errorf("error preparing body: %s", err)
	}

	req, err := http.NewRequest(c.With.method(), c.With.URL, body)
	if err != nil {
		return fmt.Errorf("error while creating request: %s", err)
	}

	for k, v := range c.With.Headers {
		req.Header.Set(k, v)
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
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func (c CaseWith) method() string {
	if c.Method == "" {
		return "GET"
	}

	return strings.ToUpper(c.Method)
}

func (c CaseWith) body() (io.Reader, error) {
	// TODO: ideally DX wise we should allow for creating the payload with yaml instead of as a string
	return strings.NewReader(c.Body), nil
}

type CaseExpect struct {
	StatusCode int `json:"statusCode"`
}
