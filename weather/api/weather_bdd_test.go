package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"weather/pkg/weather/api"
)

type apiFeature struct {
	server *httptest.Server
	resp   string
	err    error
}

func (f *apiFeature) iRequestTheWeatherFor(city string) error {
	f.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if city == "NonExistentCity" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Weather for %s", city)
	}))

	f.resp, f.err = api.GetWeather(f.server.URL)
	return nil
}

func (f *apiFeature) theResponseShouldContain(text string) error {
	defer f.server.Close()
	if f.err != nil {
		return f.err
	}
	if !strings.Contains(f.resp, text) {
		return fmt.Errorf("expected response to contain '%s', but it didn't", text)
	}
	return nil
}

func (f *apiFeature) theResponseShouldHaveStatusCode(statusCode int) error {
	defer f.server.Close()
	if f.err == nil {
		return fmt.Errorf("expected an error, but got none")
	}

	expectedError := fmt.Sprintf("got %d", statusCode)
	if !strings.Contains(f.err.Error(), expectedError) {
		return fmt.Errorf("expected error to contain status code %d, but it was: %w", statusCode, f.err)
	}

	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			f := &apiFeature{}
			s.Step(`^I am a user$`, func() error { return nil })
			s.Step(`^I request the weather for "([^"]*)"$`, f.iRequestTheWeatherFor)
			s.Step(`^the response should contain "([^"]*)"$`, f.theResponseShouldContain)
			s.Step(`^the response should have status code (\d+)$`, f.theResponseShouldHaveStatusCode)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}