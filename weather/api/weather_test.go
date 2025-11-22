package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bool64/godogx/allure"
	"github.com/cucumber/godog"
)

var resp *http.Response
var server *httptest.Server

func iAmAUser() error {
	return nil
}

func iRequestTheWeatherFor(city string) error {
	var err error
	if city == "NonExistentCity" {
		server = newMockServer("Not Found", http.StatusNotFound)
	} else {
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, fmt.Sprintf("Weather for %s", city))
		}))
	}
	resp, err = http.Get(server.URL)
	if err != nil {
		return err
	}
	return nil
}

func theResponseShouldContain(text string) error {
	defer server.Close()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if !strings.Contains(string(body), text) {
		return fmt.Errorf("expected response to contain '%s', but it didn't", text)
	}
	return nil
}

func theResponseShouldHaveStatusCode(statusCode int) error {
	defer server.Close()
	if resp.StatusCode != statusCode {
		return fmt.Errorf("expected status code %d, but got %d", statusCode, resp.StatusCode)
	}
	return nil
}

func theResponseHeaderShouldBe(headerName, headerValue string) error {
	if resp.Header.Get(headerName) != headerValue {
		return fmt.Errorf("expected header '%s' to be '%s', but got '%s'", headerName, headerValue, resp.Header.Get(headerName))
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I am a user$`, iAmAUser)
	ctx.Step(`^I request the weather for "([^"]*)"$`, iRequestTheWeatherFor)
	ctx.Step(`^the response should contain "([^"]*)"$`, theResponseShouldContain)
	ctx.Step(`^the response should have status code (\d+)$`, theResponseShouldHaveStatusCode)
	ctx.Step(`^the response header "([^"]*)" should be "([^"]*)"$`, theResponseHeaderShouldBe)
}

func TestMain(m *testing.M) {
	allure.RegisterFormatter()
	godog.TestSuite{
		Name:                "godog",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "allure",
			Paths:  []string{"features"},
		},
	}.Run()
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Important!
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
