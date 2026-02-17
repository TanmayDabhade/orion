package executor

import (
	"orion/internal/config"
	"orion/internal/plan"
	"testing"
)

func TestExecute_Echo(t *testing.T) {
	// Simple test to ensure Execute runs without panic and executes a basic command
	p := &plan.CommandPlan{
		Intent:  "test",
		Summary: "Test Plan",
		Commands: []plan.Command{
			{Cmd: "echo hello", Risk: plan.RiskLow},
		},
	}

	cfg := config.Default()
	// Mock LogsDir in a real scenario would be needed, but NewLogger uses config.LogsDir() directly
	// which uses os.UserHomeDir. This test generates a real log file in ~/.config/orion/logs
	// We might want to clean it up or mock config, but for now this integration test is fine.

	if err := Execute(p, cfg); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
}

func TestExecute_Fail(t *testing.T) {
	p := &plan.CommandPlan{
		Intent: "fail_test",
		Commands: []plan.Command{
			{Cmd: "false", Risk: plan.RiskLow}, // assumes 'false' command exists and returns non-zero
		},
	}

	cfg := config.Default()
	if err := Execute(p, cfg); err == nil {
		t.Fatal("Expected Execute to fail, but it succeeded")
	}
}
