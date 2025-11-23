package ui

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

// newChromedpContext creates a new chromedp context with a timeout and no-sandbox option.
func NewChromedpContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("disable-dev-shm-usage", true),
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
func GetPageTitle(ctx context.Context, url string) (string, error) {

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
func TakeScreenshot(ctx context.Context, url string) ([]byte, error) {

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
