package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"orion/internal/config"
	"orion/models"
	"orion/providers"
)

type llmIntent struct {
	Action string                 `json:"action"`
	Args   map[string]interface{} `json:"args"`
	Risk   string                 `json:"risk"`
}

func InferIntent(ctx context.Context, input string, cfg config.Config) (models.Intent, error) {
	provider, err := selectProvider(cfg)
	if err != nil {
		return models.Intent{}, err
	}

	prompt := buildPrompt(input)
	output, err := provider.Complete(ctx, prompt)
	if err != nil {
		return models.Intent{}, err
	}

	return parseIntent(output)
}

func HealthCheck(ctx context.Context, cfg config.Config) error {
	provider, err := selectProvider(cfg)
	if err != nil {
		return err
	}
	return provider.Health(ctx)
}

func selectProvider(cfg config.Config) (providers.Provider, error) {
	switch strings.ToLower(strings.TrimSpace(cfg.AIProvider)) {
	case "", "ollama":
		return providers.Ollama{Endpoint: cfg.AIEndpoint, Model: cfg.AIModel}, nil
	case "gemini":
		return providers.Gemini{APIKey: cfg.AIKey, Model: cfg.AIModel}, nil
	default:
		return nil, fmt.Errorf("unsupported AI provider: %s", cfg.AIProvider)
	}
}

func buildPrompt(input string) string {
	return fmt.Sprintf(`You are an intent router. Respond with ONLY valid JSON.

Schema:
{"action":"open_url|open_app|search|run_shell","args":{"key":"value"},"risk":"low|medium|high"}

Rules:
- No extra keys.
- Use "open_url" with args.url
- Use "open_app" with args.app
- Use "search" with args.query
- Use "run_shell" with args.command
- Keep risk conservative.

CRITICAL:
- If the user asks for multiple steps (e.g. "Open Safari AND go to github.com"), use "run_shell".
- Combine commands using "open -a AppName URL" or "cmd1 && cmd2".
- Example: "Open Safari and search for cats" -> {"action":"run_shell", "args":{"command":"open -a Safari 'https://google.com/search?q=cats'"}, "risk":"medium"}
- Example: "Open report.pdf in Preview" -> {"action":"run_shell", "args":{"command":"open -a Preview report.pdf"}, "risk":"low"}
- If user attempts to open a known website/dashboard (e.g. "gemini api dashboard"), resolve the URL yourself and use "open_url".
- Example: "open gemini console" -> {"action":"open_url", "args":{"url":"https://aistudio.google.com"}, "risk":"low"}

User request: %s
`, input)
}

func parseIntent(raw string) (models.Intent, error) {
	raw = strings.TrimSpace(raw)
	var parsed llmIntent
	if err := json.Unmarshal([]byte(raw), &parsed); err != nil {
		return models.Intent{}, fmt.Errorf("invalid AI JSON: %w", err)
	}

	action := models.Action(strings.TrimSpace(parsed.Action))
	switch action {
	case models.ActionOpenURL, models.ActionOpenApp, models.ActionSearch, models.ActionRunShell:
	default:
		return models.Intent{}, fmt.Errorf("unsupported AI action: %s", parsed.Action)
	}

	args := make(map[string]string)
	for key, value := range parsed.Args {
		args[key] = fmt.Sprint(value)
	}

	intent := models.Intent{
		Action: action,
		Args:   args,
		Risk:   models.ParseRisk(parsed.Risk),
	}

	if err := validateArgs(intent); err != nil {
		return models.Intent{}, err
	}
	return intent, nil
}

func validateArgs(intent models.Intent) error {
	switch intent.Action {
	case models.ActionOpenURL:
		if intent.Args["url"] == "" {
			return fmt.Errorf("AI intent missing url")
		}
	case models.ActionOpenApp:
		if intent.Args["app"] == "" {
			return fmt.Errorf("AI intent missing app")
		}
	case models.ActionSearch:
		if intent.Args["query"] == "" {
			return fmt.Errorf("AI intent missing query")
		}
	case models.ActionRunShell:
		if intent.Args["command"] == "" {
			return fmt.Errorf("AI intent missing command")
		}
	}
	return nil
}
