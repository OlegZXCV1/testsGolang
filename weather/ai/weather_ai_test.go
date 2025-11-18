package ai

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"google.golang.org/genai"
	"google.golang.org/api/option"
)

// TestWeatherHaiku uses the Gemini API to generate a haiku about the weather.
func TestWeatherHaiku(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		t.Skip("GEMINI_API_KEY not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := http.Get("https://wttr.in/?format=j1")
	if err != nil {
		t.Fatalf("failed to get weather: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	prompt := fmt.Sprintf("Write a haiku about this weather: %s", string(body))

	genResp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		t.Fatal(err)
	}

	if len(genResp.Candidates) == 0 {
		t.Fatal("no candidates returned")
	}

	haiku := genResp.Candidates[0].Content.Parts[0]

	fmt.Println(haiku)
}
