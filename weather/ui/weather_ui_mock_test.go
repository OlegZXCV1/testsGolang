package ui

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func newMockServer(response string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	}))
}

func TestWeatherUINavigationMock(t *testing.T) {
	server := newMockServer("<html><head><title>Weather report</title></head><body><h1>Weather</h1></body></html>", http.StatusOK)
	defer server.Close()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var title string
	err := chromedp.Run(ctx,
		chromedp.Navigate(server.URL),
		chromedp.Title(&title),
	)
	if err != nil {
		t.Fatal(err)
	}

	expectedTitle := "Weather report"
	if title != expectedTitle {
		t.Errorf("expected title %q, got %q", expectedTitle, title)
	}
}

func TestWeatherUICheckH1(t *testing.T) {
	server := newMockServer("<html><head><title>Weather report</title></head><body><h1>Big Weather</h1></body></html>", http.StatusOK)
	defer server.Close()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var h1 string
	err := chromedp.Run(ctx,
		chromedp.Navigate(server.URL),
		chromedp.Text("h1", &h1),
	)
	if err != nil {
		t.Fatal(err)
	}

	expectedH1 := "Big Weather"
	if h1 != expectedH1 {
		t.Errorf("expected h1 %q, got %q", expectedH1, h1)
	}
}

func TestWeatherUICheckInput(t *testing.T) {
	server := newMockServer("<html><body><input type='text' value='london'></body></html>", http.StatusOK)
	defer server.Close()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var val string
	err := chromedp.Run(ctx,
		chromedp.Navigate(server.URL),
		chromedp.Value("input", &val),
	)
	if err != nil {
		t.Fatal(err)
	}

	if val != "london" {
		t.Errorf("expected input value to be 'london', got '%s'", val)
	}
}