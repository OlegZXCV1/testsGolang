package ai

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GenerativeModel is an interface that matches the `genai.GenerativeModel`'s relevant methods.
type GenerativeModel interface {
	GenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error)
}

// HaikuGenerator holds the model for generating haikus.
type HaikuGenerator struct {
	Model GenerativeModel
}

// GenerateHaiku uses the provided model to generate a haiku.
func (hg *HaikuGenerator) GenerateHaiku(prompt string) (string, error) {
	resp, err := hg.Model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content generated")
	}

	if text, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
		return string(text), nil
	}

	return "", fmt.Errorf("unexpected response format")
}

// NewHaikuGenerator creates a new HaikuGenerator with a real Gemini client.
func NewHaikuGenerator(ctx context.Context, apiKey, modelName string) (*HaikuGenerator, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %w", err)
	}
	model := client.GenerativeModel(modelName)
	return &HaikuGenerator{Model: model}, nil
}