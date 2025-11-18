package ui

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

// TestWeatherUINavigation uses chromedp to navigate to wttr.in and checks the page title.
func TestWeatherUINavigation(t *testing.T) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var title string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://wttr.in/`),
		chromedp.Title(&title),
	)
	if err != nil {
		t.Fatal(err)
	}

	expectedTitle := "wttr.in"
	if title != expectedTitle {
		t.Errorf("expected title %q, got %q", expectedTitle, title)
	}
}
