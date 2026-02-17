package ai

import (
	"context"
	"fmt"
	"strings"

	"orion/internal/ai/local"
	"orion/internal/config"
	env "orion/internal/context"
	"orion/internal/plan"
	"orion/models"
	"orion/providers"
)

// Main entry point for converting NL to a Command Plan
func InferPlan(ctx context.Context, input string, cfg config.Config) (*plan.CommandPlan, error) {
	provider, err := selectProvider(cfg)
	if err != nil {
		return nil, err
	}

	// Gather context (CWD, tools, project type)
	currentEnv := env.Get("")

	prompt := buildPlanPrompt(input, currentEnv)
	output, err := provider.Complete(ctx, prompt)
	if err != nil {
		return nil, err
	}

	return plan.ParseStrict([]byte(output))
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
	case "local":
		return local.Local{BinPath: cfg.LocalAIPath}, nil
	default:
		return nil, fmt.Errorf("unsupported AI provider: %s", cfg.AIProvider)
	}
}

func buildPlanPrompt(input string, currentEnv env.Context) string {
	tools := strings.Join(currentEnv.Tools, ", ")
	project := string(currentEnv.ProjectType)

	const schema = `
{
  "intent": "string",
  "summary": "string",
  "cwd": "path",
  "commands": [{"cmd":"string", "risk":"low|medium|high"}],
  "questions": ["string"]
}`

	return fmt.Sprintf(`You are a macOS command planner.
Goal: Return a JSON execution plan for the user request.

CONTEXT:
- OS: macOS
- CWD: %s
- Project Type: %s
- Available Tools: %s

SCHEMA: %s

RULES:
1. Return ONLY valid JSON. No markdown blocking.
2. If the user request is vague, populate "questions".
3. If specific, list sequence of commands.
4. "cwd" should be absolute path or empty.
5. Risk: "low" (read/safe), "medium" (install/create), "high" (delete/sudo).
6. Prefer using available tools.

USER REQUEST: %s
`, currentEnv.Cwd, project, tools, schema, input)
}

// Deprecated: existing logic kept for backward compatibility if needed,
// but we should move to InferPlan entirely.
func InferIntent(ctx context.Context, input string, cfg config.Config) (models.Intent, error) {
	// ... (legacy implementation or map new plan to old intent)
	return models.Intent{}, fmt.Errorf("InferIntent is deprecated, use InferPlan")
}
