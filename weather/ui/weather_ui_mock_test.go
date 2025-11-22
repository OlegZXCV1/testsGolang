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
	if title != expectedTitle {
		t.Errorf("expected title %q, got %q", expectedTitle, title)
	}
}

func TestWeatherUICheckH1(t *testing.T) {
	server := newMockServer("<html><head><title>Weather report</title></head><body><h1>Big Weather</h1></body></html>", http.StatusOK)
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

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
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

func TestWeatherUICheckCSSSelector(t *testing.T) {
	server := newMockServer("<html><body><div class='weather'>sunny</div></body></html>", http.StatusOK)
	defer server.Close()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var text string
	err := chromedp.Run(ctx,
		chromedp.Navigate(server.URL),
		chromedp.Text(".weather", &text),
	)
	if err != nil {
		t.Fatal(err)
	}

	if text != "sunny" {
		t.Errorf("expected text to be 'sunny', got '%s'", text)
	}
}

func TestWeatherUICheckAttribute(t *testing.T) {
	server := newMockServer("<html><body><a href='/weather'>weather</a></body></html>", http.StatusOK)
	defer server.Close()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var href string
	err := chromedp.Run(ctx,
		chromedp.Navigate(server.URL),
		chromedp.AttributeValue("a", "href", &href, nil),
	)
	if err != nil {
		t.Fatal(err)
	}

	if href != "/weather" {
		t.Errorf("expected href to be '/weather', got '%s'", href)
	}
}
