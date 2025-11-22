package uimcp

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// TestWeatherUIScreenshot uses the Gemini API to describe a screenshot of the weather UI.
func TestWeatherUIScreenshot(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		t.Skip("GEMINI_API_KEY not set")
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://wttr.in/`),
		chromedp.FullScreenshot(&buf, 90),
	)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("weather_ui.png", buf, 0644)
	if err != nil {
		t.Fatal(err)
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	prompt := []genai.Part{
		genai.ImageData("png", buf),
		genai.Text("Describe this image of a weather website."),
	}

	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Candidates) == 0 {
		t.Fatal("no candidates returned")
	}

	description := resp.Candidates[0].Content.Parts[0]

	fmt.Println(description)
}
