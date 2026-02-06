package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Gemini struct {
	APIKey string
	Model  string
	Client *http.Client
}

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
}

type geminiContent struct {
	Parts []geminiPart `json:"parts"`
}

type geminiPart struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []geminiCandidate `json:"candidates"`
	Error      *geminiError      `json:"error,omitempty"`
}

type geminiCandidate struct {
	Content geminiContent `json:"content"`
}

type geminiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (g Gemini) Complete(ctx context.Context, prompt string) (string, error) {
	if g.APIKey == "" {
		return "", fmt.Errorf("gemini api key required")
	}
	model := g.Model
	if model == "" {
		model = "gemini-pro"
	}

	payload, err := json.Marshal(geminiRequest{
		Contents: []geminiContent{
			{
				Parts: []geminiPart{
					{Text: prompt},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, g.APIKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := g.Client
	if client == nil {
		client = &http.Client{Timeout: 30 * time.Second}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var out geminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}

	if out.Error != nil {
		return "", fmt.Errorf("gemini error %d: %s", out.Error.Code, out.Error.Message)
	}

	if len(out.Candidates) == 0 || len(out.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("gemini returned empty response")
	}

	return out.Candidates[0].Content.Parts[0].Text, nil
}

func (g Gemini) Health(ctx context.Context) error {
	// Simple health check by listing models or just checking if key is present
	if g.APIKey == "" {
		return fmt.Errorf("gemini api key missing")
	}
	// Ideally we'd make a lightweight call here, but for now we'll assume if key is present it's ok
	// to avoid wasting quota on every health check.
	// A real check would call https://generativelanguage.googleapis.com/v1beta/models?key=...
	return nil
}
