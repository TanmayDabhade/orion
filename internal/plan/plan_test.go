package plan

import (
	"testing"
)

func TestParseStrict(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name: "Valid Plan",
			input: `{
				"intent": "setup_nextjs",
				"summary": "Create app",
				"cwd": "/tmp",
				"commands": [
					{"cmd": "echo hello", "risk": "low"}
				],
				"questions": []
			}`,
			wantErr: false,
		},
		{
			name: "Unknown Field",
			input: `{
				"intent": "bad",
				"extra_field": "should fail",
				"commands": []
			}`,
			wantErr: true,
		},
		{
			name: "Missing Intent",
			input: `{
				"commands": [{"cmd":"ls", "risk":"low"}]
			}`,
			wantErr: true,
		},
		{
			name: "Invalid Risk",
			input: `{
				"intent": "test",
				"commands": [{"cmd":"ls", "risk":"critical"}]
			}`,
			wantErr: true,
		},
		{
			name: "Too Many Commands",
			input: `{
				"intent": "flood",
				"commands": [
					{"cmd":"1", "risk":"low"}, {"cmd":"2", "risk":"low"},
					{"cmd":"3", "risk":"low"}, {"cmd":"4", "risk":"low"},
					{"cmd":"5", "risk":"low"}, {"cmd":"6", "risk":"low"},
					{"cmd":"7", "risk":"low"}, {"cmd":"8", "risk":"low"},
					{"cmd":"9", "risk":"low"}
				]
			}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseStrict([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStrict() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSafety(t *testing.T) {
	tests := []struct {
		name          string
		cmd           Command
		expectBlocked bool
		expectHigh    bool
	}{
		{
			name:          "Safe Command",
			cmd:           Command{Cmd: "echo hello", Risk: RiskLow},
			expectBlocked: false,
			expectHigh:    false,
		},
		{
			name:          "Blocked Command (rm -rf /)",
			cmd:           Command{Cmd: "rm -rf /", Risk: RiskLow},
			expectBlocked: true,
			expectHigh:    false, // Blocked before upgrade
		},
		{
			name:          "Blocked Command (System)",
			cmd:           Command{Cmd: "touch /System/file", Risk: RiskLow},
			expectBlocked: true,
			expectHigh:    false,
		},
		{
			name:          "High Risk Escalation (sudo)",
			cmd:           Command{Cmd: "sudo apt update", Risk: RiskLow},
			expectBlocked: false,
			expectHigh:    true,
		},
		{
			name:          "High Risk Escalation (chmod recursive)",
			cmd:           Command{Cmd: "chmod -R 777 .", Risk: RiskLow},
			expectBlocked: false,
			expectHigh:    true,
		},
		{
			name:          "Sensitive Path (~/.ssh)",
			cmd:           Command{Cmd: "cat ~/.ssh/id_rsa", Risk: RiskLow},
			expectBlocked: false,
			expectHigh:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plan := &CommandPlan{
				Intent:   "test",
				Commands: []Command{tt.cmd},
			}
			err := ValidateSafety(plan)

			if tt.expectBlocked {
				if err == nil {
					t.Errorf("Expected block for command %s, but got nil", tt.cmd.Cmd)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if tt.expectHigh && plan.Commands[0].Risk != RiskHigh {
				t.Errorf("Expected risk upgrade to high for %s, got %s", tt.cmd.Cmd, plan.Commands[0].Risk)
			}
		})
	}
}
