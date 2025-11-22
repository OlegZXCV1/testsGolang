package ai

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/generative-ai-go/genai"
)

type mockGenerativeModel struct {
	GenerateContentFunc func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error)
}

func (m *mockGenerativeModel) GenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	return m.GenerateContentFunc(ctx, parts...)
}

func TestWeatherHaikuMock(t *testing.T) {
	model := &mockGenerativeModel{
		GenerateContentFunc: func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
			return &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{
					{
						Content: &genai.Content{
							Parts: []genai.Part{
								genai.Text("A haiku about weather"),
							},
						},
					},
				},
			}, nil
		},
	}

	resp, err := model.GenerateContent(context.Background(), genai.Text("Write a haiku about this weather: {}"))
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Candidates) == 0 {
		t.Fatal("no candidates returned")
	}

	haiku := resp.Candidates[0].Content.Parts[0]

	if haiku != genai.Text("A haiku about weather") {
		t.Errorf("expected haiku to be 'A haiku about weather', got %q", haiku)
	}
}

func TestWeatherHaikuMockNoCandidates(t *testing.T) {
	model := &mockGenerativeModel{
		GenerateContentFunc: func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
			return &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{},
			}, nil
		},
	}

	resp, err := model.GenerateContent(context.Background(), genai.Text("Write a haiku about this weather: {}"))
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Candidates) != 0 {
		t.Fatal("expected no candidates to be returned")
	}
}

func TestWeatherHaikuMockError(t *testing.T) {
	model := &mockGenerativeModel{
		GenerateContentFunc: func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
			return nil, fmt.Errorf("an error occurred")
		},
	}

	_, err := model.GenerateContent(context.Background(), genai.Text("Write a haiku about this weather: {}"))
	if err == nil {
		t.Fatal("expected an error to be returned")
	}
}

func TestWeatherHaikuEmptyPrompt(t *testing.T) {
	model := &mockGenerativeModel{
		GenerateContentFunc: func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
			return &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{
					{
						Content: &genai.Content{
							Parts: []genai.Part{
								genai.Text("A haiku about something"),
							},
						},
					},
				},
			}, nil
		},
	}

	resp, err := model.GenerateContent(context.Background(), genai.Text(""))
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Candidates) == 0 {
		t.Fatal("no candidates returned")
	}

	haiku := resp.Candidates[0].Content.Parts[0]

	if haiku != genai.Text("A haiku about something") {
		t.Errorf("expected haiku to be 'A haiku about something', got %q", haiku)
	}
}
