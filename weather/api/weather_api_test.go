package api

import (
	"net/http"
	"testing"
)

// TestGetWeather tests the weather API by making a GET request to wttr.in
// and checking for a 200 OK status.
func TestGetWeather(t *testing.T) {
	resp, err := http.Get("https://wttr.in/")
	if err != nil {
		t.Fatalf("failed to get weather: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
