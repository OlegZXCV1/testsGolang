package ai_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/generative-ai-go/genai"
	"weather/pkg/weather/ai"
)

// mockGenerativeModel is a mock implementation of the GenerativeModel interface.
type mockGenerativeModel struct {
	resp *genai.GenerateContentResponse
	err  error
}

func (m *mockGenerativeModel) GenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	return m.resp, m.err
}

func TestGenerateHaiku(t *testing.T) {
	mockModel := &mockGenerativeModel{
		resp: &genai.GenerateContentResponse{
			Candidates: []*genai.Candidate{
				{
					Content: &genai.Content{
						Parts: []genai.Part{genai.Text("Test haiku")},
					},
				},
			},
		},
	}
	generator := &ai.HaikuGenerator{Model: mockModel}

	haiku, err := generator.GenerateHaiku("a prompt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if haiku != "Test haiku" {
		t.Errorf("expected haiku 'Test haiku', got '%s'", haiku)
	}
}

func TestGenerateHaiku_Error(t *testing.T) {
	mockModel := &mockGenerativeModel{
		err: fmt.Errorf("API error"),
	}
	generator := &ai.HaikuGenerator{Model: mockModel}

	_, err := generator.GenerateHaiku("a prompt")
	if err == nil {
		t.Fatal("expected an error, but got none")
	}
}

func TestGenerateHaiku_NoContent(t *testing.T) {
	mockModel := &mockGenerativeModel{
		resp: &genai.GenerateContentResponse{
			Candidates: []*genai.Candidate{},
		},
	}
	generator := &ai.HaikuGenerator{Model: mockModel}

	_, err := generator.GenerateHaiku("a prompt")
	if err == nil {
		t.Fatal("expected an error for no content, but got none")
	}
}
