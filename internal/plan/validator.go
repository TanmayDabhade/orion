package plan

import (
	"fmt"
	"strings"
)

// ValidateSafety checks the plan for forbidden commands and ensures risk levels aren't underestimated.
func ValidateSafety(plan *CommandPlan) error {
	for i := range plan.Commands {
		cmd := &plan.Commands[i]

		isHighRisk, err := checkCommandSafety(cmd.Cmd)
		if err != nil {
			return fmt.Errorf("command %d blocked: %w", i, err)
		}

		if isHighRisk {
			// Upgrade risk to high if logic detects it, regardless of AI output
			cmd.Risk = RiskHigh
		}
	}
	return nil
}

// checkCommandSafety returns true if the command requires manual confirmation (high risk),
// and an error if the command is strictly forbidden.
func checkCommandSafety(cmd string) (bool, error) {
	cmd = strings.TrimSpace(cmd)

	// Hard blocks (Forbidden)
	forbiddenPatterns := []string{
		"rm -rf /",
		"rm -fr /",
		":(){ :|:& };:", // fork bomb
		"> /dev/sda",    // disk wipe attempt (simplistic)
		"mkfs",
	}

	for _, pattern := range forbiddenPatterns {
		if strings.Contains(cmd, pattern) {
			return false, fmt.Errorf("unsafe command detected: %s", pattern)
		}
	}

	// Protect System integrity
	if strings.HasPrefix(cmd, "/System") || strings.Contains(cmd, " /System") {
		return false, fmt.Errorf("modification of /System is forbidden")
	}

	if strings.Contains(cmd, "csrutil disable") {
		return false, fmt.Errorf("disabling SIP is forbidden")
	}

	// High Risk (Requires Confirmation)
	highRiskKeywords := []string{
		"sudo",
		"chmod -R",
		"chown -R",
		"mv /*",
		"rm /*",
	}

	for _, keyword := range highRiskKeywords {
		if strings.HasPrefix(cmd, keyword) || strings.Contains(cmd, " "+keyword) {
			return true, nil
		}
	}

	// Sensitive paths
	sensitivePaths := []string{
		"~/.ssh",
		"/etc",
		"/Library",
		"/var",
	}

	for _, path := range sensitivePaths {
		if strings.Contains(cmd, path) {
			return true, nil
		}
	}

	return false, nil
}
