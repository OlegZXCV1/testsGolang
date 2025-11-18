package api

import (
	"net/http"
	"testing"
)

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
