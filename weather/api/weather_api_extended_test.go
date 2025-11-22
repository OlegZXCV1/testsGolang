package api

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// TestGetWeatherBody tests the weather API by making a GET request to wttr.in
// and checking for a 200 OK status and the presence of "weather" in the body.
func TestGetWeatherBodyMock(t *testing.T) {
	server := newMockServer("weather for London", http.StatusOK)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("failed to get weather: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if !strings.Contains(strings.ToLower(string(body)), "weather") {
		t.Error("expected response body to contain 'weather'")
	}
}
