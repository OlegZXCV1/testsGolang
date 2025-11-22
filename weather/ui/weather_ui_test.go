package ui

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

// TestWeatherUINavigation uses chromedp to navigate to wttr.in and checks the page title.
func TestWeatherUINavigation(t *testing.T) {
	server := newMockServer("<html><head><title>Weather report</title></head><body><h1>Weather</h1></body></html>", http.StatusOK)
	defer server.Close()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
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
	if !strings.Contains(title, expectedTitle) {
		t.Errorf("expected title to contain %q, got %q", expectedTitle, title)
	}
}
