package ui

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

// newChromedpContext creates a new chromedp context with a timeout and no-sandbox option.
func newChromedpContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancelCtx := chromedp.NewContext(allocCtx)
	ctx, cancelTimeout := context.WithTimeout(ctx, timeout)

	return ctx, func() {
		cancelTimeout()
		cancelCtx()
		cancelAlloc()
	}

}

// GetPageTitle navigates to a URL and returns the page title.
func GetPageTitle(url string) (string, error) {
	ctx, cancel := newChromedpContext(15 * time.Second)
	defer cancel()

	var title string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Title(&title),
	)
	if err != nil {
		return "", fmt.Errorf("could not get page title: %w", err)
	}
	return title, nil
}

// TakeScreenshot navigates to a URL and takes a full screenshot.
func TakeScreenshot(url string) ([]byte, error) {
	ctx, cancel := newChromedpContext(20 * time.Second)
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.FullScreenshot(&buf, 90),
	)
	if err != nil {
		return nil, fmt.Errorf("could not take screenshot: %w", err)
	}
	return buf, nil
}
