package ai

import (
	"strings"
	"testing"

	env "orion/internal/context"
)

func TestBuildPlanPrompt(t *testing.T) {
	// Mock context
	mockEnv := env.Context{
		Cwd:         "/Users/test/project",
		ProjectType: env.ProjectNode,
		Tools:       []string{"git", "npm", "node"},
	}

	input := "initialize project"
	prompt := buildPlanPrompt(input, mockEnv)

	// Verify key context elements are present
	checks := []string{
		"CWD: /Users/test/project",
		"Project Type: node",
		"Available Tools: git, npm, node",
		"USER REQUEST: initialize project",
	}

	for _, check := range checks {
		if !strings.Contains(prompt, check) {
			t.Errorf("Prompt missing expected context: %s", check)
		}
	}
}
