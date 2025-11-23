package ui_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather/pkg/weather/ui"
)

func newMockServer(response string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, response)
	}))
}

func TestGetPageTitle(t *testing.T) {
	server := newMockServer(`<html><head><title>Test Title</title></head></html>`, http.StatusOK)
	defer server.Close()

	title, err := ui.GetPageTitle(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if title != "Test Title" {
		t.Errorf("expected title 'Test Title', got '%s'", title)
	}
}

func TestTakeScreenshot(t *testing.T) {
	server := newMockServer(`<html><body><h1>Hello</h1></body></html>`, http.StatusOK)
	defer server.Close()

	buf, err := ui.TakeScreenshot(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(buf) == 0 {
		t.Error("expected screenshot buffer to not be empty")
	}
}
