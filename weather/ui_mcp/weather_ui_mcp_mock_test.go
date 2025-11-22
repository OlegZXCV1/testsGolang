package uimcp

import (
	"bytes"
	"context"
	"fmt"
	"image/gif"
	"image/png"
	"os"
	"testing"

	"github.com/google/generative-ai-go/genai"
)

type mockGenerativeModel struct {
	GenerateContentFunc func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error)
}

func (m *mockGenerativeModel) GenerateContent(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
	return m.GenerateContentFunc(ctx, parts...)
}

func TestWeatherUIScreenshotMock(t *testing.T) {
	model := &mockGenerativeModel{
		GenerateContentFunc: func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
			return &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{
					{
						Content: &genai.Content{
							Parts: []genai.Part{
								genai.Text("A description of the weather UI"),
							},
						},
					},
				},
			}, nil
		},
	}

	file, err := os.Open("testdata/weather_ui.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}

	prompt := []genai.Part{
		genai.ImageData("png", buf.Bytes()),
		genai.Text("Describe this image of a weather website."),
	}

	resp, err := model.GenerateContent(context.Background(), prompt...)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Candidates) == 0 {
		t.Fatal("no candidates returned")
	}

	description := resp.Candidates[0].Content.Parts[0]

	if description != genai.Text("A description of the weather UI") {
		t.Errorf("expected description to be 'A description of the weather UI', got %q", description)
	}
}

func TestWeatherUIScreenshotMockError(t *testing.T) {
	model := &mockGenerativeModel{
		GenerateContentFunc: func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
			return nil, fmt.Errorf("an error occurred")
		},
	}

	file, err := os.Open("testdata/weather_ui.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}

	prompt := []genai.Part{
		genai.ImageData("png", buf.Bytes()),
		genai.Text("Describe this image of a weather website."),
	}

	_, err = model.GenerateContent(context.Background(), prompt...)
	if err == nil {
		t.Fatal("expected an error to be returned")
	}
}

func TestWeatherUIScreenshotEmptyImage(t *testing.T) {
	model := &mockGenerativeModel{
		GenerateContentFunc: func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
			return nil, fmt.Errorf("image is empty")
		},
	}

	prompt := []genai.Part{
		genai.ImageData("png", []byte{}),
		genai.Text("Describe this image."),
	}

	_, err := model.GenerateContent(context.Background(), prompt...)
	if err == nil {
		t.Fatal("expected an error for empty image")
	}
}

func TestWeatherUIScreenshotCorruptedImage(t *testing.T) {
	model := &mockGenerativeModel{
		GenerateContentFunc: func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
			return nil, fmt.Errorf("image is corrupted")
		},
	}

	prompt := []genai.Part{
		genai.ImageData("png", []byte("corrupted image data")),
		genai.Text("Describe this image."),
	}

	_, err := model.GenerateContent(context.Background(), prompt...)
	if err == nil {
		t.Fatal("expected an error for corrupted image")
	}
}

func TestWeatherUIScreenshotGIF(t *testing.T) {
	model := &mockGenerativeModel{
		GenerateContentFunc: func(ctx context.Context, parts ...genai.Part) (*genai.GenerateContentResponse, error) {
			return &genai.GenerateContentResponse{
				Candidates: []*genai.Candidate{
					{
						Content: &genai.Content{
							Parts: []genai.Part{genai.Text("A GIF image")},
						},
					},
				},
			}, nil
		},
	}

	file, err := os.Open("testdata/weather_ui.gif")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, err := gif.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if err := gif.Encode(&buf, img, nil); err != nil {
		t.Fatal(err)
	}

	prompt := []genai.Part{
		genai.ImageData("gif", buf.Bytes()),
		genai.Text("Describe this image."),
	}

	resp, err := model.GenerateContent(context.Background(), prompt...)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Candidates[0].Content.Parts[0] != genai.Text("A GIF image") {
		t.Error("unexpected response for GIF image")
	}
}