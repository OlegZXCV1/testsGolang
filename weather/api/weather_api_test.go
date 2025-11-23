package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"weather/pkg/weather/api"
)

func newMockServer(response string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, response)
	}))
}

func TestGetWeather(t *testing.T) {
	server := newMockServer("Weather for London", http.StatusOK)
	defer server.Close()

	body, err := api.GetWeather(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(body, "Weather for London") {
		t.Errorf("expected body to contain 'Weather for London', got '%s'", body)
	}
}

func TestGetWeather_NotFound(t *testing.T) {
	server := newMockServer("Not Found", http.StatusNotFound)
	defer server.Close()

	_, err := api.GetWeather(server.URL)
	if err == nil {
		t.Fatal("expected an error, but got none")
	}

	expectedError := "expected status code 200, got 404"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error to contain '%s', got '%s'", expectedError, err.Error())
	}
}
